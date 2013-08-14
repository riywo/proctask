package main

import (
    "os"
    "os/exec"
    "io"
    "path/filepath"
)

type engine struct {
    Config *config
    Cmd    *exec.Cmd
    Stdin  *os.File
    Stdout *os.File
    Stderr *os.File
}

func NewEngine(dir string) (*engine, error) {
    config, err := loadConfig(dir)
    if err != nil {
        log.Error("%s", err)
        return nil, err
    }

    cmd, err := buildCmd(dir, config)
    if err != nil {
        log.Error("%s", err)
        return nil, err
    }

    stdin, err := buildStdin(dir)
    if err != nil {
        log.Error("can't open stdin file")
        return nil, err
    }

    stdout, err := buildStdout(dir)
    if err != nil {
        log.Error("can't create stdout file")
        return nil, err
    }

    stderr, err := buildStderr(dir)
    if err != nil {
        log.Error("can't create stderr file")
        return nil, err
    }

    var engine engine
    engine.Config = config
    engine.Cmd    = cmd
    engine.Stdin  = stdin
    engine.Stdout = stdout
    engine.Stderr = stderr

    return &engine, nil
}

func (e *engine) Run() error {
    e.Cmd.Stdin  = e.Stdin

    stdout, err := e.Cmd.StdoutPipe()
    if err != nil { return err }
    go io.Copy(e.Stdout, stdout)

    stderr, err := e.Cmd.StderrPipe()
    if err != nil { return err }
    stderrTee := io.TeeReader(stderr, os.Stderr)
    go io.Copy(e.Stderr, stderrTee)

    log.Info("run cmd: %s, env: %s, dir: %s", e.Cmd.Args, e.Cmd.Env, e.Cmd.Dir)
    err = e.Cmd.Run()

    defer func() {
        if err := e.Stdout.Close(); err != nil {
            panic(err)
        }
        if err := e.Stderr.Close(); err != nil {
            panic(err)
        }
    }()

    return err
}

func loadConfig(dir string) (*config, error) {
    configPath := filepath.Join(dir, configFile)
    f, err := os.Open(configPath)
    if err != nil {
        log.Error("reading path: %s", configPath)
        return nil, err
    }
    defer f.Close()

    var config config
    if err = decodeConfig(f, &config); err != nil {
        log.Error("parse json path: %s", configPath)
        return nil, err
    }

    log.Info("load config path: %s", configPath)
    return &config, nil
}

func buildCmd(dir string, config *config) (*exec.Cmd, error) {
    cmd := exec.Command(config.Command[0], config.Command[1:]...)
    cmd.Env    = config.buildEnv()
    cmd.Dir    = dir
    return cmd, nil
}

func buildStdin(dir string) (*os.File, error) {
    stdinPath := filepath.Join(dir, stdinFile)
    stdin, err := os.Open(stdinPath)
    if err != nil {
        if os.IsNotExist(err) {
            log.Info("stdin file is not found. use null device")
            return nil, nil
        } else {
            return nil, err
        }
    }
    return stdin, err
}

func buildStdout(dir string) (*os.File, error) {
    stdoutPath := filepath.Join(dir, stdoutFile)
    stdout, err := os.Create(stdoutPath)
    if err != nil {
        return nil, err
    }
    return stdout, err
}

func buildStderr(dir string) (*os.File, error) {
    stderrPath := filepath.Join(dir, stderrFile)
    stderr, err := os.Create(stderrPath)
    if err != nil {
        return nil, err
    }
    return stderr, err
}

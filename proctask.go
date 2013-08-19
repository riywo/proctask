package main

import (
    "os"
    "path/filepath"
    "github.com/jcelliott/lumber"
)

const runFile    = "run"
const envFile    = "env"
const stdinFile  = "stdin"
const stdoutFile = "stdout"
const stderrFile = "stderr"

var log *lumber.ConsoleLogger

func main() {
    os.Exit(realMain())
}

func realMain() int {
    log = buildLogger()
    dir, err := workingDir()
    if err != nil {
        log.Error("%s", err)
        return 1
    }

    engine, err := NewEngine(dir)
    if err != nil {
        log.Error("%s", err)
        return 1
    }

    err = engine.Run()
    if err != nil {
        log.Error("Command failed: %s", err)
        return 1
    }

    log.Info("Command succeeded!")
    return 0
}

func buildLogger() *lumber.ConsoleLogger {
    return lumber.NewConsoleLogger(lumber.DEBUG)
}

func workingDir() (string, error) {
    var dir string
    var err error
    if len(os.Args) < 2 {
        dir, err = os.Getwd()
        if err != nil { return "", err }
    } else {
        dir = os.Args[1]
    }
    dir, err = filepath.Abs(dir)
    if err != nil { return "", err }
    return dir, nil
}

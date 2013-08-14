package main

import (
    "os"
    "github.com/jcelliott/lumber"
)

const configFile = "proctask.json"
const stdinFile  = "stdin"
const stdoutFile = "stdout"
const stderrFile = "stderr"

var log *lumber.ConsoleLogger

func main() {
    os.Exit(realMain())
}

func realMain() int {
    log = buildLogger()
    dir := workingDir()

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

func workingDir() string {
    var dir string
    if len(os.Args) < 2 {
        dir, _ = os.Getwd()
    } else {
        dir = os.Args[1]
    }
    return dir
}

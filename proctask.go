package main

import (
    "os"
    "os/exec"
    "io/ioutil"
    "encoding/json"
    "path/filepath"
)

type Proctask struct {
    Command  []string
    Env      []string
    Dir      string
}

func NewProctask() *Proctask {
    conf, err := ioutil.ReadFile("./proctask.json")
    if err != nil { panic(err) }
    var proctask Proctask
    err = json.Unmarshal(conf, &proctask)
    if err != nil { panic(err) }

    return &proctask
}

func main() {
    var proctask = *NewProctask()

    stdi, err := os.Open(filepath.Join(proctask.Dir, "stdin"))
    if err != nil { panic(err) }
    defer func() {
        if err := stdi.Close(); err != nil {
            panic(err)
        }
    }()

    stdo, err := os.Create(filepath.Join(proctask.Dir, "stdout"))
    if err != nil { panic(err) }
    defer func() {
        if err := stdo.Close(); err != nil {
            panic(err)
        }
    }()

    stde, err := os.Create(filepath.Join(proctask.Dir, "stderr"))
    if err != nil { panic(err) }
    defer func() {
        if err := stde.Close(); err != nil {
            panic(err)
        }
    }()

    cmd := exec.Command(proctask.Command[0], proctask.Command[1:]...)
    cmd.Env    = proctask.Env
    cmd.Dir    = proctask.Dir
    cmd.Stdin  = stdi
    cmd.Stdout = stdo
    cmd.Stderr = stde
    err = cmd.Run()
    if err != nil {
        panic(err)
    }
}

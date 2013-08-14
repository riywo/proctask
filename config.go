package main

import (
    "encoding/json"
    "io"
    "fmt"
)

type config struct {
    Command []string
    Env     map[string]string
}

func decodeConfig(r io.Reader, c *config) error {
    decoder:= json.NewDecoder(r)
    return decoder.Decode(c)
}

func (c *config) buildEnv() []string {
    env := make([]string, 0, len(c.Env))
    for key, val := range c.Env {
        env = append(env, fmt.Sprintf("%s=%s", key, val))
    }
    return env
}

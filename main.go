package main

import (
    "fmt"
    "os"
    "github.com/rbledsaw3/blog_aggregator/internal/config"
)

type state struct {
    cfg *config.Config
}

func main() {
    cfg, err := config.Read()
    if err != nil {
        fmt.Println("error reading config: %v", err)
        os.Exit(1)
    }

    programState := &state{
        cfg: &cfg,
    }

    cmds := commands{
        registeredCommands: make(map[string]func(*state, command) error),
    }
    cmds.register("login", handlerLogin)

    if len(os.Args) < 2 {
        fmt.Println("Usage: cli <command> [args...]")
        os.Exit(1)
    }

    cmdName := os.Args[1]
    cmdArgs := os.Args[2:]

    err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

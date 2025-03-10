package main

import (
    "database/sql"
    "log"
    "os"

    "github.com/rbledsaw3/blog_aggregator/internal/config"
    "github.com/rbledsaw3/blog_aggregator/internal/database"
    _ "github.com/lib/pq"
)

type state struct {
    db  *database.Queries
    cfg *config.Config
}

func main() {
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }

    db, err := sql.Open("postgres", cfg.DBURL)
    if err != nil {
        log.Fatalf("error connecting to database: %v", err)
    }
    defer db.Close()
    dbQueries := database.New(db)

    programState := &state {
        db: dbQueries,
        cfg: &cfg,
    }

    cmds := commands{
        registeredCommands: make(map[string]func(*state, command) error),
    }
    cmds.register("login", handlerLogin)
    cmds.register("register", handlerRegister)

    if len(os.Args) < 2 {
        log.Fatalf("Usage: %s <command> [args...]", os.Args[0])
        return
    }

    cmdName := os.Args[1]
    cmdArgs := os.Args[2:]

    err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
    if err != nil {
        log.Fatal(err)
    }
}

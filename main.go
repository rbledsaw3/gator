package main

import (
    "context"
    "database/sql"
    "fmt"
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
    cmds.register("reset", handlerReset)
    cmds.register("users", handlerGetUsers)
    cmds.register("agg", handlerAgg)
    cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
    cmds.register("feeds", handlerGetFeeds)
    cmds.register("follow", middlewareLoggedIn(handlerFollow))
    cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
    cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
    cmds.register("browse", middlewareLoggedIn(handlerBrowse))

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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
    return func(s *state, cmd command) error {
        user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
        if err != nil {
            return fmt.Errorf("error getting user: %w", err)
        }
        return handler(s, cmd, user)
    }
}

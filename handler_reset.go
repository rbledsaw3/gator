package main

import (
    "context"
    "fmt"
)

func handlerReset(s *state, cmd command) error {
    if len(cmd.Args) > 1 {
        return fmt.Errorf("usage: %v <name>", cmd.Name)
    }

    err := s.db.DeleteAllUsers(context.Background())
    if err != nil {
        return fmt.Errorf("error deleting all users: %w", err)
    }
    fmt.Println("Deleting all users...")
    err = s.db.DeleteAllFeeds(context.Background())
    if err != nil {
        return fmt.Errorf("error deleting all feeds: %w", err)
    }
    fmt.Println("Deleting all feeds...")
    fmt.Println("Database reset successfully!")
    return nil
}


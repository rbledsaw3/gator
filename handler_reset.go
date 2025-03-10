package main

import (
    "context"
    "fmt"
)

func handlerReset(s *state, cmd command) error {
    if len(cmd.Args) > 1 {
        return fmt.Errorf("usage: %v <name>", cmd.Name)
    }

    // call query DeleteAllUsers to delete all users
    err := s.db.DeleteAllUsers(context.Background())
    if err != nil {
        return fmt.Errorf("error deleting all users: %w", err)
    }
    fmt.Println("Database reset successfully!")
    return nil
}


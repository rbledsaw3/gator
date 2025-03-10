package main

import (
	"context"
    "fmt"
    "time"

    "github.com/rbledsaw3/blog_aggregator/internal/database"
    "github.com/google/uuid"
)

func handlerGetUsers(s *state, cmd command) error {
    // prints all users with a "(current)" next to the current user
    users, err := s.db.GetUsers(context.Background())
    if err != nil {
        return fmt.Errorf("error getting users: %w", err)
    }
    if len(users) == 0 {
        fmt.Println("No users found.")
        return nil
    }
    fmt.Println("Users:")
    for _, user := range users {
        fmt.Printf(" * %v", user.Name)
        if user.Name == s.cfg.CurrentUserName {
            fmt.Print(" (current)")
        }
        fmt.Println()
    }
    return nil
}

func handlerRegister(s *state, cmd command) error {
    if len(cmd.Args) < 1 {
        return fmt.Errorf("usage: %v <name>", cmd.Name)
    }

    name := cmd.Args[0]

    user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        Name:      name,
    })
    if err != nil {
        return fmt.Errorf("error creating new user: %w", err)
    }

    err = s.cfg.SetUser(name)
    if err != nil {
        return fmt.Errorf("error setting current user in config: %w", err)
    }

    fmt.Println("User created successfully!")
    printUser(user)
    return nil
}

func handlerLogin(s *state, cmd command) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("usage: %s <name>", cmd.Name)
    }
    name := cmd.Args[0]

    _, err := s.db.GetUser(context.Background(), name)
    if err != nil {
        return fmt.Errorf("couldn't find user: %w", err)
    }

    err = s.cfg.SetUser(name)
    if err != nil {
        return fmt.Errorf("couldn't set current user: %w", err)
    }

    fmt.Printf("User switched to %s\n", name)
    return nil
}

func printUser(user database.User) {
    fmt.Printf(" * ID:      %v\n", user.ID)
    fmt.Printf(" * Name:    %v\n", user.Name)
}

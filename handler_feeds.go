package main

import (
    "context"
    "fmt"
    "time"

    "github.com/rbledsaw3/blog_aggregator/internal/database"
    "github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
    user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
        return fmt.Errorf("error getting user: %w", err)
    }
    if len(cmd.Args) != 2 {
        return fmt.Errorf("usage: %v <feed_name> <feed_url>", cmd.Name)
    }
    name := cmd.Args[0]
    url := cmd.Args[1]

    feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        Name:      name,
        Url:       url,
        UserID:    user.ID,
    })
    if err != nil {
        return fmt.Errorf("error creating new feed: %w", err)
    }

    fmt.Println("Feed created successfully!")
    printFeed(feed)

    return nil
}

func handlerGetFeeds(s *state, cmd command) error {
    // Gets all feeds for all users
    feeds, err := s.db.GetAllFeeds(context.Background(), 100)
    if err != nil {
        return fmt.Errorf("error getting feeds: %w", err)
    }

    for _, feed := range feeds {
        user, err := s.db.GetUserById(context.Background(), feed.UserID)
        if err != nil {
            return fmt.Errorf("error getting user: %w", err)
        }
        fmt.Printf("User: %v\n", user.Name)
        printFeed(feed)
    }
    return nil
}

func printFeed(feed database.Feed) {
    fmt.Printf("ID: %v\n", feed.ID)
    fmt.Printf("Name: %v\n", feed.Name)
    fmt.Printf("URL: %v\n", feed.Url)
    fmt.Printf("UserID: %v\n", feed.UserID)
    fmt.Printf("CreatedAt: %v\n", feed.CreatedAt)
    fmt.Printf("UpdatedAt: %v\n", feed.UpdatedAt)
}

package main

import (
    "context"
    "fmt"
    "time"

    "github.com/rbledsaw3/blog_aggregator/internal/database"
    "github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
    if len(cmd.Args) < 1 {
        return fmt.Errorf("usage: %v <url>", cmd.Name)
    }
    feedURL := cmd.Args[0]
    feed, err := s.db.GetFeedByUrl(context.Background(), feedURL)
    if err != nil {
        return fmt.Errorf("error getting feed: %w", err)
    }

    feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
        ID: uuid.New(),
        UpdatedAt: time.Now(),
        CreatedAt: time.Now(),
        UserID: user.ID,
        FeedID: feed.ID,
    })

    if err != nil {
        return fmt.Errorf("error creating feed follow: %w", err)
    }

    fmt.Println("Feed followed successfully!")
    printFeedFollow(feed_follow.UserName, feed_follow.FeedName)
    return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
    feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
    if err != nil {
        return fmt.Errorf("error getting feed follows: %w", err)
    }

    if len(feed_follows) == 0 {
        fmt.Println("No feeds follows found for this user.")
        return nil
    }

    fmt.Printf("Feeds followed by %v:\n", user.Name)
    for _, ff := range feed_follows {
        fmt.Printf("  %s\n", ff.FeedName)
    }
    return nil
}


func printFeedFollow(username, feedname string) {
    fmt.Printf("* User:            %s\n", username)
    fmt.Printf("* Feed:            %s\n", feedname)
}

package main

import (
    "context"
    "fmt"
    "time"

    "github.com/rbledsaw3/blog_aggregator/internal/database"
    "github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
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

    feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        UserID:    user.ID,
        FeedID:    feed.ID,
    })
    if err != nil {
        return fmt.Errorf("error creating new feed follow: %w", err)
    }

    fmt.Println("Feed created successfully!")
    printFeed(feed, user)
    fmt.Println()
    fmt.Println("Feed followed successfully!")
    printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
    fmt.Println()
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
        printFeed(feed, user)
        fmt.Println("======================================")
    }
    return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("usage: %v <feed_url>", cmd.Name)
    }

    feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
    if err != nil {
        return fmt.Errorf("error getting feed: %w", err)
    }

    err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
        UserID: user.ID,
        FeedID: feed.ID,
    })
    if err != nil {
        return fmt.Errorf("error deleting feed follow: %w", err)
    }

    fmt.Printf("Feed %s unfollowed successfully!\n", feed.Name)
    return nil
}

func printFeed(feed database.Feed, user database.User) {
    fmt.Printf("* ID:          %s\n", feed.ID)
    fmt.Printf("* Created:     %v\n", feed.CreatedAt)
    fmt.Printf("* Updated:     %v\n", feed.UpdatedAt)
    fmt.Printf("* Name:        %s\n", feed.Name)
    fmt.Printf("* URL:         %s\n", feed.Url)
    fmt.Printf("* User:        %s\n", user.Name)
    fmt.Printf("* LastFetched: %v\n", feed.LastFetchedAt.Time)
}

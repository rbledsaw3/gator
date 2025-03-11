package main

import (
    "context"
    "fmt"
)

func handlerAgg(_ *state, cmd command) error {
    var feedURL string
    numArgs := len(cmd.Args)
    fmt.Printf("numArgs: %v\n", numArgs)
    fmt.Printf("cmd.Args: %v\n", cmd.Args)
    switch numArgs {
    case 1:
        feedURL = cmd.Args[0]
    case 0:
        feedURL = "https://www.wagslane.dev/index.xml"
    default:
        return fmt.Errorf("usage: %v <feedURL> (optional)", cmd.Name)
    }

    feed, err := fetchFeed(context.Background(), feedURL)
    if err != nil {
        return fmt.Errorf("couldn't fetch feed: %w", err)
    }
    fmt.Printf("Feed: %v\n", feed)
    return nil
}

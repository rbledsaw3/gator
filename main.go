package main

import (
    "fmt"
    "log"
    "github.com/rbledsaw3/blog_aggregator/internal/config"
)

func main() {
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("failed to read config: %v", err)
    }
    fmt.Printf("Read config: %+v\n", cfg)

    err = cfg.SetUser("lane")
    if err != nil {
        log.Fatalf("failed to set user: %v", err)
    }

    cfg, err = config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }
    fmt.Printf("Read config again: %+v\n", cfg)
}

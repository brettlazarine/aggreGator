package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/brettlazarine/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected a time duration argument for agg command")
	}
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	nextFeed.LastFetchedAt = sql.NullTime{Time: time.Now(), Valid: true}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        nextFeed.ID,
		UpdatedAt: nextFeed.LastFetchedAt.Time,
	})
	if err != nil {
		return err
	}

	fmt.Println("*** Feed updated ***")
	fmt.Printf("ID: %d\n", nextFeed.ID)
	fmt.Printf("LastFetchedAt: %v\n", nextFeed.LastFetchedAt.Time)

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("*** Title: %s\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Println("Title:", item.Title)
	}

	return nil
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/brettlazarine/aggreGator/internal/database"
	"github.com/google/uuid"
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

	for _, item := range feed.Channel.Item {
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: parseTime(item.PubDate),
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			return err
		}

		fmt.Printf("*** Post %v created ***\n", post.Title)
	}

	return nil
}

func parseTime(timeStr string) time.Time {
	layout := "Mon, 02 Jan 2006 15:04:05 MST"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return time.Time{}
	}
	return t
}

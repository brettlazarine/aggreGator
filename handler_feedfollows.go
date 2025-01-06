package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brettlazarine/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow requires a url -> %v <url>", cmd.Name)
	}
	url := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Printf("*** Feed: %v\n", feed.Name)
	fmt.Printf("*** User: %v\n", user.Name)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("following takes no arguments -> %v", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %v", err)
	}
	if len(following) == 0 {
		fmt.Println("*** No feeds followed")
		return nil
	}

	for _, ff := range following {
		fmt.Printf("*** %v\n", ff.FeedName)
	}

	return nil
}
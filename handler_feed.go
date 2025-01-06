package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brettlazarine/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUsername
	if currentUser == "" {
		return fmt.Errorf("no user logged in")
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("addfeed requires a name and url -> %v <name> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Println("feed created successfully")
	printFeed(feed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("*** ID: %v\n", feed.ID)
	fmt.Printf("*** Name: %v\n", feed.Name)
	fmt.Printf("*** URL: %v\n", feed.Url)
	fmt.Printf("*** User ID: %v\n", feed.UserID)
	fmt.Printf("*** Created At: %v\n", feed.CreatedAt)
	fmt.Printf("*** Updated At: %v\n", feed.UpdatedAt)
}

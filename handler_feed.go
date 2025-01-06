package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brettlazarine/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if user.Name == "" {
		return fmt.Errorf("no user logged in")
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("addfeed requires a name and url -> %v <name> <url>", cmd.Name)
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
	printFeed(feed, user)

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

	fmt.Println("feed follow created successfully")

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("*** ID: %v\n", feed.ID)
	fmt.Printf("*** Name: %v\n", feed.Name)
	fmt.Printf("*** URL: %v\n", feed.Url)
	fmt.Printf("*** User Name: %v\n", user.Name)
	fmt.Printf("*** User ID: %v\n", feed.UserID)
	fmt.Printf("*** Created At: %v\n", feed.CreatedAt)
	fmt.Printf("*** Updated At: %v\n", feed.UpdatedAt)
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}
	if len(feeds) == 0 {
		return fmt.Errorf("no feeds found")
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting user: %v", err)
		}

		printFeed(feed, user)
		fmt.Println()
	}

	return nil
}

package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/brettlazarine/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		arg, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = arg
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		printPost(post)
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Println("***")
	fmt.Printf("Title:%v\n", post.Title)
	fmt.Printf("URL:%v\n", post.Url)
	fmt.Printf("Description:%v\n", post.Description)
	fmt.Printf("Published At:%v\n", post.PublishedAt)
	fmt.Println("***")
}

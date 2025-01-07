package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brettlazarine/aggreGator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login requires a username -> %v <username>", cmd.Name)
	}

	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Println("user changed to", username)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("register requires a username -> %v <username>", cmd.Name)
	}

	username := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Println("user registered:")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf("*** ID: %v\n", user.ID)
	fmt.Printf("*** Name: %v\n", user.Name)
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %v", err)
	}
	printUsersList(s, users)

	return nil
}

func printUsersList(s *state, users []database.User) {
	for _, user := range users {
		if user.Name == s.cfg.CurrentUsername {
			fmt.Printf("*** %v (current)\n", user.Name)
		} else {
			fmt.Printf("*** %v\n", user.Name)
		}
	}
}

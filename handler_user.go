package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login requires a username")
	}

	username := cmd.Args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("user changed to", username)
	return nil
}

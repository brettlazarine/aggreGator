package main

import (
	"fmt"
	"os"

	"github.com/brettlazarine/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("usage: gator <command> [args]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{cmdName, cmdArgs})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

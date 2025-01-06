package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/brettlazarine/gator/internal/config"
	"github.com/brettlazarine/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to database: %v\n", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)

	if len(os.Args) < 2 {
		log.Fatalf("usage: cli <command> [args]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{cmdName, cmdArgs})
	if err != nil {
		log.Fatalf("error running command: %v\n", err)
	}

}

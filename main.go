package main

import (
	"fmt"

	"github.com/brettlazarine/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DB URL: %v, Username: %v\n", cfg.DbUrl, cfg.CurrentUsername)
	err = cfg.SetUser("admin")
	if err != nil {
		panic(err)
	}
	cfg, err = config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Printf("*** DB URL: %v, Username: %v\n", cfg.DbUrl, cfg.CurrentUsername)
}

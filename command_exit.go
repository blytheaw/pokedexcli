package main

import (
	"os"
)

func commandExit(config *config, params []string) error {
	os.Exit(0)

	return nil
}

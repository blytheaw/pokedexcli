package main

import (
	"os"
)

func commandExit(config *config) error {
	os.Exit(0)

	return nil
}

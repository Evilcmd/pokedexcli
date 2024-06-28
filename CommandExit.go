package main

import "os"

func exitCommand(cfg *config, _ string) error {
	os.Exit(0)
	return nil
}

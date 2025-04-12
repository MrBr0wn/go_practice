package main

import (
	"signatures/cmd"
	"signatures/cmd/keys"
	"signatures/cmd/signatures"
	"signatures/logger"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(err, "Initial setup failed")
	}
}

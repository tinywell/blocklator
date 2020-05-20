package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tinywell/blocklator/cmd/server"
)

var (
	rootCmd = &cobra.Command{
		Use:   "blocklator",
		Short: "blocklator is a tool to translate fabric block",
	}
)

func main() {
	rootCmd.AddCommand(server.ServerCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

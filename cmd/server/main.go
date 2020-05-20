package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tinywell/blocklator/api"
)

var (
	port       int
	production bool

	// ServerCmd server subcommand
	ServerCmd = &cobra.Command{
		Use:   "server",
		Short: "server is a backend web service use of gin",
		Run: func(cmd *cobra.Command, args []string) {
			Execute()
		},
	}
)

func init() {
	ServerCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "listen port")
	ServerCmd.PersistentFlags().BoolVarP(&production, "production", "m", false, "is in production mode")
}

// Execute execute server command
func Execute() {
	var mode = gin.DebugMode
	if production {
		mode = gin.ReleaseMode
	}

	r := api.CollectRouter(mode)

	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", port)); err != nil {
		panic(err)
	}
}

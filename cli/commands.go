package cli

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Limerio/calculator-api/lib/logger"
	"github.com/Limerio/calculator-api/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "calculator-api",
	Run: func(cmd *cobra.Command, args []string) {

		formatLog, err := cmd.Flags().GetString("format-log")
		if err != nil {
			slog.Error(err.Error())
		}

		if formatLog != "json" && formatLog != "text" {
			slog.Error("❌ Unknown log format")

			return
		}

		switch formatLog {
		case "json":
			slog.SetDefault(logger.NewJson())
		default:
			slog.SetDefault(logger.NewText())
		}

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			slog.Error(err.Error())
		}
		portEnv := os.Getenv("PORT")

		if len(portEnv) < 1 && len(port) < 1 {
			port = "8080"
		} else {
			if len(portEnv) > 1 && len(port) > 1 {
				fmt.Println("⚠️ Due to an priority order, the environment variable was chosen first")
			}

			if len(portEnv) > 1 {
				port = portEnv
			}
		}

		wait, err := cmd.Flags().GetDuration("graceful-timeout")
		if err != nil {
			slog.Error(err.Error())
		}

		server.Run(port, wait)
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringP("port", "p", "", "Port number for the web server")
	rootCmd.PersistentFlags().String("format-log", "text", "Formatting of logs send to the stdout and stderr")
	rootCmd.PersistentFlags().Duration("graceful-timeout", time.Second*15, "Time in seconds of graceful timeout")

	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

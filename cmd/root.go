package cmd

import (
	"fmt"
	slog "github.com/go-eden/slf4go"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [flags] (log file)", appName),
		Short: "go-login",
		Args:  cobra.ExactArgs(0),
		RunE:  help,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func help(cmd *cobra.Command, args []string) error {
	err := cmd.Help()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	InitLogging()

	slog.SetLevel(slog.InfoLevel)
}

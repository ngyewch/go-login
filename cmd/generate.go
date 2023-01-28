package cmd

import (
	generate2 "github.com/ngyewch/go-login/generate"
	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:   "generate [flags] (input directory) (output directory)",
		Short: "Generate login pages",
		Args:  cobra.ExactArgs(2),
		RunE:  generate,
	}
)

func generate(cmd *cobra.Command, args []string) error {
	inputDirectory := args[0]
	outputDirectory := args[1]

	g, err := generate2.NewGenerator(inputDirectory, outputDirectory)
	if err != nil {
		return err
	}
	err = g.Generate()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

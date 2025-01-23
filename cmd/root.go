package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/edulustosa/fconv/pkg/converter"
	"github.com/spf13/cobra"
)

var (
	outputExt string

	rootCmd = &cobra.Command{
		Use:   "fconv [file]",
		Short: "Converts a file from one format to another",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputPath := args[0]
			info, err := os.Stat(inputPath)
			if err != nil {
				return errors.New("could not find input file")
			}

			if info.IsDir() {
				if outputExt == "" {
					return errors.New("output extension must be provided in directories conversions")
				}

				converter.ConvertDir(inputPath, outputExt)
				return nil
			}

			var outputPath string
			if len(args) > 1 {
				outputPath = args[1]
			} else if outputExt != "" {
				filename := strings.Split(filepath.Base(inputPath), ".")[0]
				outputPath = fmt.Sprintf("%s.%s", filename, outputExt)
			}

			if outputPath == "" {
				return errors.New("output file extension or path is required")
			}

			if err := converter.ConvertFile(inputPath, outputPath); err != nil {
				return err
			}

			return nil
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputExt, "output", "o", "", "Output file extension")
}

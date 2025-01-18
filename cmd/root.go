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
			var outputPath string
			if len(args) > 1 {
				outputPath = args[1]
				outputExt = getFileExtension(outputPath)
			}

			if outputExt == "" {
				return errors.New("output file extension is required")
			}

			inputPath := args[0]
			inputExt := getFileExtension(inputPath)
			if inputExt == outputExt {
				return errors.New("input and output file extensions must be different")
			}

			conv, err := converter.GetConversion(inputExt, outputExt)
			if err != nil {
				return err
			}

			file, err := os.Open(inputPath)
			if err != nil {
				return fmt.Errorf("could not open file: %w", err)
			}
			defer file.Close()

			output, err := conv(file, inputExt)
			if err != nil {
				return fmt.Errorf("could not convert file: %w", err)
			}

			if outputPath == "" {
				filename := strings.Split(filepath.Base(inputPath), ".")[0]
				outputPath = fmt.Sprintf("%s.%s", filename, outputExt)
			}

			if err := os.WriteFile(outputPath, output, 0644); err != nil {
				return fmt.Errorf("could not write output file: %w", err)
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

func getFileExtension(path string) string {
	ext := filepath.Ext(path)
	return strings.TrimPrefix(strings.ToLower(ext), ".")
}

package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/edulustosa/fconv/pkg/converter"
	"github.com/spf13/cobra"
)

type status int

const (
	converting status = iota
	success
	failed
)

type model struct {
	spinner    spinner.Model
	status     status
	inputFile  string
	outputFile string
	err        error
}

type conversionMsg struct {
	err error
}

var (
	outputExt string

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42")).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)

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

			p := tea.NewProgram(initialModel(filepath.Base(inputPath), outputExt))

			go func() {
				output, err := conv(file, inputExt)
				if err != nil {
					p.Send(conversionMsg{err: fmt.Errorf("could not convert file: %w", err)})
					return
				}

				if outputPath == "" {
					filename := strings.Split(filepath.Base(inputPath), ".")[0]
					outputPath = fmt.Sprintf("%s.%s", filename, outputExt)
				}

				if err := os.WriteFile(outputPath, output, 0644); err != nil {
					p.Send(conversionMsg{err: fmt.Errorf("could not write output file: %w", err)})
					return
				}

				p.Send(conversionMsg{err: nil})
			}()

			if _, err := p.Run(); err != nil {
				return fmt.Errorf("could not start program: %v", err)
			}

			return nil
		},
	}
)

func initialModel(inputFile, outputFile string) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("adb5bd"))

	return model{
		spinner:    s,
		status:     converting,
		inputFile:  inputFile,
		outputFile: outputFile,
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) View() string {
	switch m.status {
	case converting:
		return fmt.Sprintf("%s Converting %s to %s", m.spinner.View(), m.inputFile, m.outputFile)
	case success:
		return successStyle.Render(fmt.Sprintf("✓ Successfully converted %s to %s\n", m.inputFile, m.outputFile))
	case failed:
		return errorStyle.Render(fmt.Sprintf("✗ Error: %v\n", m.err))
	default:
		return ""
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}

	case conversionMsg:
		if msg.err != nil {
			m.err = msg.err
			m.status = failed
			return m, tea.Quit
		}
		m.status = success
		return m, tea.Quit

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

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

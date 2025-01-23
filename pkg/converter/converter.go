package converter

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/edulustosa/fconv/pkg/documents"
	"github.com/edulustosa/fconv/pkg/images"
)

type (
	ConversionFunc func(image io.Reader, ext string) ([]byte, error)

	Conversions map[string][]string

	Decoders map[string]ConversionFunc
)

var validConversions = Conversions{
	"jpeg": {"png", "webp", "bmp", "tiff", "gif", "tif"},
	"jpg":  {"png", "webp", "bmp", "tiff", "gif", "tif"},
	"png":  {"jpeg", "jpg", "webp", "bmp", "tiff", "gif", "tif"},
	"webp": {"jpeg", "jpg", "png", "bmp", "tiff", "gif", "tif"},
	"bmp":  {"jpeg", "jpg", "png", "webp", "tiff", "gif", "tif"},
	"tiff": {"jpeg", "jpg", "png", "webp", "bmp", "gif"},
	"tif":  {"jpeg", "jpg", "png", "webp", "bmp", "gif"},
	"gif":  {"jpeg", "jpg", "png", "webp", "bmp", "tiff", "tif"},
	"csv":  {"xlsx", "json", "yaml", "yml"},
	"json": {"yaml", "yml"},
	"yaml": {"json"},
	"xml":  {"json", "yaml", "yml"},
}

var decoders = Decoders{
	"jpeg": images.ToJpeg,
	"jpg":  images.ToJpeg,
	"png":  images.ToPng,
	"webp": images.ToWebp,
	"bmp":  images.ToBmp,
	"tiff": images.ToTiff,
	"tif":  images.ToTiff,
	"gif":  images.ToGif,
	"xlsx": documents.ToXlsx,
	"json": documents.ToJson,
	"yaml": documents.ToYaml,
	"yml":  documents.ToYaml,
}

func getConversion(inputExt, outputExt string) (ConversionFunc, error) {
	conversionsSupported, ok := validConversions[inputExt]
	if !ok {
		return nil, errors.New("unsupported input file extension")
	}

	for _, conversion := range conversionsSupported {
		if conversion == outputExt {
			conversionFunc := decoders[outputExt]
			return conversionFunc, nil
		}
	}

	return nil, fmt.Errorf("conversion from %s to %s is not supported", inputExt, outputExt)
}

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

func ConvertFile(inputPath, outputPath string) error {
	inputExt := getFileExtension(inputPath)
	outputExt := getFileExtension(outputPath)

	conversionFunc, err := getConversion(inputExt, outputExt)
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
		output, err := conversionFunc(file, inputExt)
		if err != nil {
			p.Send(conversionMsg{err: fmt.Errorf("could not convert file: %w", err)})
			return
		}

		if err := os.WriteFile(outputPath, output, 0644); err != nil {
			p.Send(conversionMsg{err: fmt.Errorf("could not write output file: %w", err)})
			return
		}

		p.Send(conversionMsg{err: nil})
	}()

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("could not start program: %w", err)
	}

	return nil
}

func ConvertDir(inputDir, outputExt string) {
	var wg sync.WaitGroup

	filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		filename := strings.Split(filepath.Base(path), ".")[0]
		outputPath := fmt.Sprintf("%s.%s", filename, outputExt)

		wg.Add(1)
		go func() {
			defer wg.Done()

			ConvertFile(path, outputPath)
		}()

		return nil
	})

	wg.Wait()
}

var (
	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42")).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)
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

func getFileExtension(path string) string {
	ext := filepath.Ext(path)
	return strings.TrimPrefix(strings.ToLower(ext), ".")
}

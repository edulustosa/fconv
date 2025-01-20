package documents

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/tealeg/xlsx"
)

func ToXlsx(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "csv":
		return csvToXlsx(file)
	}

	return nil, fmt.Errorf("unsupported conversion: %s to xlsx", ext)
}

func csvToXlsx(file io.Reader) ([]byte, error) {
	csvReader := csv.NewReader(file)

	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("failed to add sheet: %w", err)
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		row := sheet.AddRow()
		for _, value := range record {
			cell := row.AddCell()
			cell.Value = value
		}
	}

	buff := new(bytes.Buffer)
	if err := xlsxFile.Write(buff); err != nil {
		return nil, fmt.Errorf("failed to write xlsx file: %w", err)
	}

	return buff.Bytes(), nil
}

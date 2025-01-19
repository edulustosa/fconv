package images

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"

	"golang.org/x/image/tiff"
)

func ToTiff(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToTiff(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func jpegToTiff(jpegFile io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode jpeg: %w", err)
	}

	tiffBuff := new(bytes.Buffer)
	if err := tiff.Encode(tiffBuff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to encode tiff: %w", err)
	}

	return tiffBuff.Bytes(), nil
}

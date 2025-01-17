package images

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
)

func ToPng(imageBytes io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToPng(imageBytes)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func jpegToPng(jpegBytes io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegBytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding JPEG: %w", err)
	}

	pngBuff := new(bytes.Buffer)
	if err := png.Encode(pngBuff, img); err != nil {
		return nil, fmt.Errorf("error encoding PNG: %w", err)
	}

	return pngBuff.Bytes(), nil
}

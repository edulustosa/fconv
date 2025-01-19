package images

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"

	"golang.org/x/image/bmp"
)

func ToBmp(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToBmp(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func jpegToBmp(jpegFile io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode jpeg: %w", err)
	}

	bmpBuff := new(bytes.Buffer)
	if err := bmp.Encode(bmpBuff, img); err != nil {
		return nil, fmt.Errorf("failed to encode bmp: %w", err)
	}

	return bmpBuff.Bytes(), nil
}

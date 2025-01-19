package images

import (
	"bytes"
	"fmt"
	"image/gif"
	"image/jpeg"
	"io"
)

func ToGif(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToGif(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func jpegToGif(jpegFile io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode jpeg: %w", err)
	}

	gifBuff := new(bytes.Buffer)
	if err := gif.Encode(gifBuff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to encode gif: %w", err)
	}

	return gifBuff.Bytes(), nil
}

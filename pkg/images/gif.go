package images

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func ToGif(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToGif(file)
	case "png":
		return pngToGif(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func imageToGif(img image.Image) ([]byte, error) {
	gifBuff := new(bytes.Buffer)
	if err := gif.Encode(gifBuff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to encode gif: %w", err)
	}

	return gifBuff.Bytes(), nil
}

func jpegToGif(jpegFile io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode jpeg: %w", err)
	}

	return imageToGif(img)
}

func pngToGif(pngFile io.Reader) ([]byte, error) {
	img, err := png.Decode(pngFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode png: %w", err)
	}

	return imageToGif(img)
}

package images

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/tiff"
)

func ToTiff(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToTiff(file)
	case "png":
		return pngToTiff(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func imageToTiff(img image.Image) ([]byte, error) {
	tiffBuff := new(bytes.Buffer)
	if err := tiff.Encode(tiffBuff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to encode tiff: %w", err)
	}

	return tiffBuff.Bytes(), nil
}

func jpegToTiff(jpegFile io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode jpeg: %w", err)
	}

	return imageToTiff(img)
}

func pngToTiff(pngFile io.Reader) ([]byte, error) {
	img, err := png.Decode(pngFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode png: %w", err)
	}

	return imageToTiff(img)
}

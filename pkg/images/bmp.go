package images

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/bmp"
)

func ToBmp(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToBmp(file)
	case "png":
		return pngToBmp(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func imageToBmp(img image.Image) ([]byte, error) {
	bmpBuff := new(bytes.Buffer)
	if err := bmp.Encode(bmpBuff, img); err != nil {
		return nil, fmt.Errorf("failed to encode bmp: %w", err)
	}

	return bmpBuff.Bytes(), nil
}

func jpegToBmp(jpegFile io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode jpeg: %w", err)
	}

	return imageToBmp(img)
}

func pngToBmp(pngFile io.Reader) ([]byte, error) {
	img, err := png.Decode(pngFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode png file: %w", err)
	}

	return imageToBmp(img)
}

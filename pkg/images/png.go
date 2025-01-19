package images

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

func ToPng(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg", "jpg":
		return jpegToPng(file)
	case "bmp":
		return bmpToPng(file)
	case "gif":
		return gifToPng(file)
	case "tiff":
		return tiffToPng(file)
	case "webp":
		return webpToPng(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func imageToPng(img image.Image) ([]byte, error) {
	pngBuff := new(bytes.Buffer)
	if err := png.Encode(pngBuff, img); err != nil {
		return nil, fmt.Errorf("error encoding PNG: %w", err)
	}

	return pngBuff.Bytes(), nil
}

func jpegToPng(jpegBytes io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegBytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding JPEG: %w", err)
	}

	return imageToPng(img)
}

func bmpToPng(bmpFile io.Reader) ([]byte, error) {
	img, err := bmp.Decode(bmpFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode bmp file: %w", err)
	}

	return imageToPng(img)
}

func gifToPng(gifFile io.Reader) ([]byte, error) {
	img, err := gif.Decode(gifFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode gif file: %w", err)
	}

	return imageToPng(img)
}

func tiffToPng(tiffFile io.Reader) ([]byte, error) {
	img, err := tiff.Decode(tiffFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tiff file: %w", err)
	}

	return imageToPng(img)
}

func webpToPng(webpFile io.Reader) ([]byte, error) {
	img, err := webp.Decode(webpFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode webp file: %w", err)
	}

	return imageToPng(img)
}

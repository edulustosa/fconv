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

func ToJpeg(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "png":
		return pngToJpeg(file)
	case "bmp":
		return bmpToJpeg(file)
	case "gif":
		return gifToJpeg(file)
	case "tiff":
		return tiffToJpeg(file)
	case "webp":
		return webpToJpeg(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func imageToJpeg(img image.Image) ([]byte, error) {
	jpegBuff := new(bytes.Buffer)
	if err := jpeg.Encode(jpegBuff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to convert jpeg: %w", err)
	}

	return jpegBuff.Bytes(), nil
}

func pngToJpeg(pngFile io.Reader) ([]byte, error) {
	img, err := png.Decode(pngFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode png file: %w", err)
	}

	return imageToJpeg(img)
}

func bmpToJpeg(bmpFile io.Reader) ([]byte, error) {
	img, err := bmp.Decode(bmpFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode bmp file: %w", err)
	}

	return imageToJpeg(img)
}

func gifToJpeg(gifFile io.Reader) ([]byte, error) {
	img, err := gif.Decode(gifFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode gif file: %w", err)
	}

	return imageToJpeg(img)
}

func tiffToJpeg(tiffFile io.Reader) ([]byte, error) {
	img, err := tiff.Decode(tiffFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tiff file: %w", err)
	}

	return imageToJpeg(img)
}

func webpToJpeg(webpFile io.Reader) ([]byte, error) {
	img, err := webp.Decode(webpFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode webp file: %w", err)
	}

	return imageToJpeg(img)
}

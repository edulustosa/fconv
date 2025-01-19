package images

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func ToWebp(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToWebp(file)
	case "png":
		return pngToWebp(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func imageToWebp(img image.Image) ([]byte, error) {
	opts, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		return nil, fmt.Errorf("error creating WebP encoder options: %w", err)
	}

	webpBuff := new(bytes.Buffer)
	if err := webp.Encode(webpBuff, img, opts); err != nil {
		return nil, fmt.Errorf("error encoding WebP: %w", err)
	}

	return webpBuff.Bytes(), nil
}

func jpegToWebp(jpegBytes io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegBytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding JPEG: %w", err)
	}

	return imageToWebp(img)
}

func pngToWebp(pngFile io.Reader) ([]byte, error) {
	img, err := png.Decode(pngFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode png file: %w", err)
	}

	return imageToWebp(img)
}

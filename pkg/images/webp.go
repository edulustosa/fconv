package images

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func ToWebp(image io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "jpeg":
		return jpegToWebp(image)
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

func jpegToWebp(jpegBytes io.Reader) ([]byte, error) {
	img, err := jpeg.Decode(jpegBytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding JPEG: %w", err)
	}

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

package images

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	kolesaWebp "github.com/kolesa-team/go-webp/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

type DecoderFunc func(io.Reader) (image.Image, error)

var decoders = map[string]DecoderFunc{
	"jpeg": jpeg.Decode,
	"jpg":  jpeg.Decode,
	"png":  png.Decode,
	"gif":  gif.Decode,
	"bmp":  bmp.Decode,
	"tiff": tiff.Decode,
	"webp": webp.Decode,
}

func getImage(file io.Reader, ext string) (image.Image, error) {
	decoder, ok := decoders[ext]
	if !ok {
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	img, err := decoder(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s file: %w", ext, err)
	}

	return img, nil
}

func ToJpeg(file io.Reader, ext string) ([]byte, error) {
	img, err := getImage(file, ext)
	if err != nil {
		return nil, err
	}

	jpegBuff := new(bytes.Buffer)
	if err := jpeg.Encode(jpegBuff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to convert jpeg: %w", err)
	}

	return jpegBuff.Bytes(), nil
}

func ToPng(file io.Reader, ext string) ([]byte, error) {
	img, err := getImage(file, ext)
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	if err := png.Encode(buff, img); err != nil {
		return nil, fmt.Errorf("failed to convert jpeg: %w", err)
	}

	return buff.Bytes(), nil
}

func ToWebp(file io.Reader, ext string) ([]byte, error) {
	img, err := getImage(file, ext)
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	if err := kolesaWebp.Encode(buff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to convert jpeg: %w", err)
	}

	return buff.Bytes(), nil
}

func ToGif(file io.Reader, ext string) ([]byte, error) {
	img, err := getImage(file, ext)
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	if err := gif.Encode(buff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to convert jpeg: %w", err)
	}

	return buff.Bytes(), nil
}

func ToBmp(file io.Reader, ext string) ([]byte, error) {
	img, err := getImage(file, ext)
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	if err := bmp.Encode(buff, img); err != nil {
		return nil, fmt.Errorf("failed to convert jpeg: %w", err)
	}

	return buff.Bytes(), nil
}

func ToTiff(file io.Reader, ext string) ([]byte, error) {
	img, err := getImage(file, ext)
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	if err := tiff.Encode(buff, img, nil); err != nil {
		return nil, fmt.Errorf("failed to convert jpeg: %w", err)
	}

	return buff.Bytes(), nil
}

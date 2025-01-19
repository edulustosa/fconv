package converter

import (
	"errors"
	"io"

	"github.com/edulustosa/fconv/pkg/images"
)

type (
	ConversionFunc func(image io.Reader, ext string) ([]byte, error)

	Conversions map[string]map[string]ConversionFunc
)

var validConversions = Conversions{
	"jpeg": {
		"png":  images.ToPng,
		"webp": images.ToWebp,
		"bmp":  images.ToBmp,
		"tiff": images.ToTiff,
		"gif":  images.ToGif,
	},
	"jpg": {
		"png":  images.ToPng,
		"webp": images.ToWebp,
		"bmp":  images.ToBmp,
		"tiff": images.ToTiff,
		"gif":  images.ToGif,
	},
	"png": {
		"jpeg": images.ToJpeg,
		"jpg":  images.ToJpeg,
		"webp": images.ToWebp,
		"bmp":  images.ToBmp,
		"tiff": images.ToTiff,
		"gif":  images.ToGif,
	},
}

func GetConversion(inputExt, outputExt string) (ConversionFunc, error) {
	conversionsSupported, ok := validConversions[inputExt]
	if !ok {
		return nil, errors.New("unsupported input file extension")
	}

	conversionFunc, ok := conversionsSupported[outputExt]
	if !ok {
		return nil, errors.New("unsupported output file extension")
	}

	return conversionFunc, nil
}

package converter

import (
	"errors"
	"io"

	"github.com/edulustosa/fconv/pkg/images"
)

type ConversionFunc func(image io.Reader, ext string) ([]byte, error)

var validConversions = map[string]map[string]ConversionFunc{
	"jpeg": {
		"png":  images.ToPng,
		"webp": images.ToWebp,
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

package converter

import (
	"errors"
	"fmt"
	"io"

	"github.com/edulustosa/fconv/pkg/documents"
	"github.com/edulustosa/fconv/pkg/images"
)

type (
	ConversionFunc func(image io.Reader, ext string) ([]byte, error)

	Conversions map[string][]string

	Decoders map[string]ConversionFunc
)

var validConversions = Conversions{
	"jpeg": {"png", "webp", "bmp", "tiff", "gif", "tif"},
	"jpg":  {"png", "webp", "bmp", "tiff", "gif", "tif"},
	"png":  {"jpeg", "jpg", "webp", "bmp", "tiff", "gif", "tif"},
	"webp": {"jpeg", "jpg", "png", "bmp", "tiff", "gif", "tif"},
	"bmp":  {"jpeg", "jpg", "png", "webp", "tiff", "gif", "tif"},
	"tiff": {"jpeg", "jpg", "png", "webp", "bmp", "gif"},
	"tif":  {"jpeg", "jpg", "png", "webp", "bmp", "gif"},
	"gif":  {"jpeg", "jpg", "png", "webp", "bmp", "tiff", "tif"},
	"csv":  {"xlsx", "json", "yaml", "yml"},
	"json": {"yaml", "yml"},
	"yaml": {"json"},
	"xml":  {"json", "yaml", "yml"},
}

var decoders = Decoders{
	"jpeg": images.ToJpeg,
	"jpg":  images.ToJpeg,
	"png":  images.ToPng,
	"webp": images.ToWebp,
	"bmp":  images.ToBmp,
	"tiff": images.ToTiff,
	"tif":  images.ToTiff,
	"gif":  images.ToGif,
	"xlsx": documents.ToXlsx,
	"json": documents.ToJson,
	"yaml": documents.ToYaml,
	"yml":  documents.ToYaml,
}

func GetConversion(inputExt, outputExt string) (ConversionFunc, error) {
	conversionsSupported, ok := validConversions[inputExt]
	if !ok {
		return nil, errors.New("unsupported input file extension")
	}

	for _, conversion := range conversionsSupported {
		if conversion == outputExt {
			conversionFunc := decoders[outputExt]
			return conversionFunc, nil
		}
	}

	return nil, fmt.Errorf("conversion from %s to %s is not supported", inputExt, outputExt)
}

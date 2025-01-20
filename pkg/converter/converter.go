package converter

import (
	"errors"
	"io"

	"github.com/edulustosa/fconv/pkg/documents"
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
		"tif":  images.ToTiff,
	},
	"jpg": {
		"png":  images.ToPng,
		"webp": images.ToWebp,
		"bmp":  images.ToBmp,
		"tiff": images.ToTiff,
		"gif":  images.ToGif,
		"tif":  images.ToTiff,
	},
	"png": {
		"jpeg": images.ToJpeg,
		"jpg":  images.ToJpeg,
		"webp": images.ToWebp,
		"bmp":  images.ToBmp,
		"tiff": images.ToTiff,
		"gif":  images.ToGif,
		"tif":  images.ToTiff,
	},
	"webp": {
		"jpeg": images.ToJpeg,
		"jpg":  images.ToJpeg,
		"png":  images.ToPng,
		"bmp":  images.ToBmp,
		"tiff": images.ToTiff,
		"gif":  images.ToGif,
		"tif":  images.ToTiff,
	},
	"bmp": {
		"jpeg": images.ToJpeg,
		"jpg":  images.ToJpeg,
		"png":  images.ToPng,
		"webp": images.ToWebp,
		"tiff": images.ToTiff,
		"gif":  images.ToGif,
		"tif":  images.ToTiff,
	},
	"tiff": {
		"jpeg": images.ToJpeg,
		"jpg":  images.ToJpeg,
		"png":  images.ToPng,
		"webp": images.ToWebp,
		"bmp":  images.ToBmp,
		"gif":  images.ToGif,
	},
	"tif": {
		"jpeg": images.ToJpeg,
		"jpg":  images.ToJpeg,
		"png":  images.ToPng,
		"webp": images.ToWebp,
		"bmp":  images.ToBmp,
		"gif":  images.ToGif,
		"tif":  images.ToTiff,
	},
	"gif": {
		"jpeg": images.ToJpeg,
		"jpg":  images.ToJpeg,
		"png":  images.ToPng,
		"webp": images.ToWebp,
		"bmp":  images.ToBmp,
		"tiff": images.ToTiff,
		"tif":  images.ToTiff,
	},
	"csv": {
		"xlsx": documents.ToXlsx,
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

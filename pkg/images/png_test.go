package images_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/edulustosa/fconv/pkg/images"
)

func TestJpegToPng(t *testing.T) {
	image, err := os.Open("./test_images/sample_jpeg.jpeg")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	pngImage, err := images.ToPng(image, "jpeg")
	if err != nil {
		t.Errorf("Failed to convert jpeg to png: %v", err)
	}

	if http.DetectContentType(pngImage) != "image/png" {
		t.Errorf("Failed to convert jpeg to png: invalid image type)")
	}

	err = os.WriteFile("./test_images/sample_jpeg_converted.png", pngImage, 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}
}

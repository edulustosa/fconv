package images_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/edulustosa/fconv/pkg/images"
)

func TestJpegToWebP(t *testing.T) {
	image, err := os.Open("./test_images/sample_jpeg.jpeg")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	webpImage, err := images.ToWebp(image, "jpeg")
	if err != nil {
		t.Errorf("Failed to convert jpeg to webp: %v", err)
	}

	if http.DetectContentType(webpImage) != "image/webp" {
		t.Errorf("Failed to convert jpeg to webp: invalid image type")
	}

	err = os.WriteFile("./test_images/sample_jpeg_converted.webp", webpImage, 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}
}

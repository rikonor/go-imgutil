package imgutil

import (
	"image"
	"image/color"
	"reflect"
	"testing"
)

func TestIterate(t *testing.T) {
	testImage := image.NewRGBA(image.Rect(0, 0, 1, 1))
	testColor := color.RGBA{
		R: 123, G: 123, B: 123, A: 255,
	}

	// Make image the test color
	Iterate(testImage, func(x int, y int) {
		testImage.Set(x, y, testColor)
	})

	// Verify image is the test color
	Iterate(testImage, func(x int, y int) {
		// All pixels should be of `testColor`
		if c := testImage.At(x, y); !reflect.DeepEqual(c, testColor) {
			t.Fatalf("test image has unexpected color: %+v", c)
		}
	})
}

package imgutil

import "image"

// Iterate takes an image and a function and applies the fn
// Notice: It's up to the given function to access the same image as the one
// given to Iterate
func Iterate(img image.Image, fn func(x int, y int)) {
	b := img.Bounds()

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			fn(x, y)
		}
	}
}

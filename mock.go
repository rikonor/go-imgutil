package imgutil

import (
	"image"
	"image/color"
)

// MockImage is a mockable image.Image
type MockImage struct {
	ColorModelFn func() color.Model
	BoundsFn     func() image.Rectangle
	AtFn         func(x, y int) color.Color
}

// ColorModel calls the underlying ColorModel method
func (img *MockImage) ColorModel() color.Model {
	return img.ColorModelFn()
}

// Bounds calls the underlying Bounds method
func (img *MockImage) Bounds() image.Rectangle {
	return img.BoundsFn()
}

// At calls the underlying At method
func (img *MockImage) At(x, y int) color.Color {
	return img.AtFn(x, y)
}

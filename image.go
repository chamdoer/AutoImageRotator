package autorot

import (
	"image"
	"image/color"
	"math"

	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/num-analysis/linalg/ludecomp"
)

// Rotate rotates an image around its center and returns
// the largest centered square cropping that does not go
// out of the rotated image's bounds.
//
// The angle is specified in clockwise radians.
//
// The outSize argument specifies the side length of the
// resulting image.
func Rotate(img image.Image, angle float64, outSize int) image.Image {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	width := float64(img.Bounds().Dx())
	height := float64(img.Bounds()
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
	height := float64(img.Bounds().Dy())
	axisBasis := &linalg.Matrix{
		Rows: 2,
		Cols: 2,
		Data: []float64{
			cos * width / 2, -sin * height / 2,
			sin * width / 2, cos * height / 2,
		},
	}

	inv := ludecomp.Decompose(axisBasis)
	var sideLength float64
	for rectFits(inv, sideLength+1) {
		sideLength++
	}

	scale := sideLength / float64(outSize)

	inImage := newRGBACache(img)
	newImage := image.NewRGBA(image.Rect(0, 0, int(outSize), int(outSize)))
	for x := 0; x < int(outSize); x++ {
		for y := 0; y < int(outSize); y++ {
			xOff :
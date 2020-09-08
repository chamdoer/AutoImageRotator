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
			xOff := scale*float64(x) - sideLength/2
			yOff := scale*float64(y) - sideLength/2
			newX := cos*xOff + sin*yOff + width/2
			newY := cos*yOff - sin*xOff + height/2
			newImage.SetRGBA(x, y, interpolate(inImage, newX, newY))
		}
	}

	return newImage
}

func rectFits(axisBasis *ludecomp.LU, sideLength float64) bool {
	for xScale := -1; xScale <= 1; xScale += 2 {
		for yScale := -1; yScale <= 1; yScale += 2 {
			corner := []float64{
				sideLength * float64(xScale) / 2,
				sideLength * float64(yScale) / 2,
			}
			solution := axisBasis.Solve(corner)
			if solution.MaxAbs() > 1 {
				return false
			}
		}
	}
	return true
}

func interpolate(img *rgbaCache, x, y float64) color.RGBA {
	x1 := int(x)
	x2 := int(x + 1)
	y1 := int(y)
	y2 := int(y + 1)
	amountX1 := float64(x2) - x
	amountY1 := float64(y2) - y
	clipRange(0, img.Width(), &x1, &x2)
	clipRange(0, img.Height(
package autorot

import (
	"image"
	"math"
	"testing"
)

func BenchmarkRotate(b *testing.B) {
	img := image.NewYCbCr(image.Rect(0, 0, 900, 713), image.YCbCrSubsampleRatio444)
	b
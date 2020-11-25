package autorot

import (
	"image"
	"math"
	"testing"
)

func BenchmarkRotate(b *testing.B) {
	img := image.N
package autorot

import (
	"image"
	"image/color"
	"math"

	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/num-analysis/linalg/ludecomp"
)

// Rotate rotates an image around its center and returns
/
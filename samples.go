
package autorot

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/unixpickle/anynet/anyff"
	"github.com/unixpickle/anynet/anysgd"
	"github.com/unixpickle/anyvec/anyvec32"
)

// A SampleList is an anyff.SampleList of image samples.
//
// The samples are rotated by random angles.
//
// It is designed to work with data downloaded via
// https://github.com/unixpickle/imagenet.
type SampleList struct {
	Paths     []string
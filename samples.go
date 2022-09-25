
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
	ImageSize int
}

// ReadSampleList walks the directory and creates a sample
// for each of the images (with a random rotation).
func ReadSampleList(imageSize int, dir string) (*SampleList, error) {
	res := &SampleList{ImageSize: imageSize}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ext := filepath.Ext(path)
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
			res.Paths = append(res.Paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

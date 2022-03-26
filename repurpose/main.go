// Command repurpose converts an ImageNet classifier to an
// autorot network.
package main

import (
	"flag"

	"github.com/unixpickle/anydiff"
	"github.com/unixpickle/anynet"
	"github.com/unixpickle/anyvec/anyvec32"
	"github.com/unixpickle/autorot"
	"github.com/unixpickle/essentials"
	"github.com/unixpickle/imagenet"
	"github.com/unixpickle/serializer"
)

func main() {
	var inFile string
	var outFile string
	var removeLayers int
	var rightAngles bool
	var confidence bool

	flag.StringVar(&inFile, "in", "", "imagenet classifier path")
	flag.StringVar(&outFile, "out", "", "output network path")
	flag.IntVar(&removeLayers, "remove", 2, "number of layers to remove")
	flag.BoolVar(&
// Command post_train produces an *imagenet.Classifier for
// a neural network.
// As part of doing this, it converts batch normalization
// layers into affine transforms.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/unixpickle/anynet/anyconv"
	"github.com/unixpickle/anynet/anyff"
	"github.com/unixpickle/anynet/anysgd"
	"github.com/unixpickle/autorot"
	"github.com/unixpickle/essentials"
	"github.com/unixpickle/serializer"
)

func main() {
	var imgDir string
	var inNet string
	var outNet string

	var batchSize int
	var sampleCount int

	flag.StringVar(&imgDir, "samples", "", "sample directory")
	flag.StringVar(&inNet, "in", "", "input network")
	flag.StringVar(&outNet, "out", "", "output network")
	flag.IntVar(&batchSize, "batch", 8, "evaluation batch size")
	flag.IntVar(&sampleCount, "total", 512, "total samples for BatchNorm replacement")

	flag.Parse()

	if imgDir == "" || inNet == "" || outNet == "" {
		fmt.Fprintln(os.Stderr, "Required flags: -in, -out, and -samples")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Println("Loading network...")
	var net *autorot.Net
	if err := serializer.LoadAny(inNet, &net); err != nil {
		essentials.Die("Failed to read network:", err)
	}

	log.Println("Loading samples...")
	samples, err := autorot.ReadSampleList(net.InputSize, imgDir)
	if err != nil {
		essentials.Die("Failed to read 
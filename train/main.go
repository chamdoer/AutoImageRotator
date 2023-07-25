// Command train trains a network to find the correct
// rotation for images.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/unixpickle/anynet/anyff"
	"github.com/unixpickle/anynet/anysgd"
	"github.com/unixpickle/autorot"
	"github.com/unixpickle/essentials"
	"github.com/unixpickle/rip"
	"github.com/unixpickle/serializer"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var netFile string
	var dataDir string
	var stepSize float64
	var batchSize int
	flag.StringVar(&netFile, "net", "", "network file")
	flag.StringVar(&dataDir, "data", "", "image directory")
	flag.Float64Var(&stepSize, "step", 0.001, "SGD step size")
	flag.IntVar(&batchSize, "batch", 12, "SGD batch size")
	flag.Parse()

	if netFile == "" || dataDir == "" {
		essentials.Die("Required flags: -net and -data. See -help for more.")
	}

	log.Println("Loading network...")

	var net *autorot.Net
	if err := serializer.LoadAny(netFile, &net); err != nil {
		essentials.Die("Load network failed:", err)
	}

	log.Println("Loading samples...")

	samples, err := autorot.ReadSampleList(net.InputSize, dataDir)
	if err != nil {
		essentials.Die("Load data failed:", err)
	}

	log.Println("Training...")

	t := &anyff.Trainer{
		Net:     net.Net,
		Cost:    net,
		Params:  net.Net.Param
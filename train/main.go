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
	var batchSize in
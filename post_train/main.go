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

	var batchSize 
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
	"github.com/unix
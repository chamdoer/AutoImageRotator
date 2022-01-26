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

func ma
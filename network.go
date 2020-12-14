package autorot

import (
	"image"
	"math"

	"github.com/unixpickle/anydiff"
	"github.com/unixpickle/anynet"
	"github.com/unixpickle/anyvec"
	"github.com/unixpickle/anyvec/anyvec32"
	"github.com/unixpickle/serializer"

	_ "github.com/unixpickle/anynet/anyconv"
)

// OutputType specifies the output format and loss
// function for a network.
type OutputType int

const (
	RawAngle OutputType = iota
	RightAngles
	ConfidenceAngle
)

func init() {
	var n Net
	serializer.Regist
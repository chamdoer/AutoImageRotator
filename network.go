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
	serializer.RegisterTypedDeserializer(n.SerializerType(), DeserializeNet)
}

// A Net is a neural net that predicts angles from images.
type Net struct {
	// Side length of input images.
	InputSize int

	OutputType OutputType
	Net        anynet.Net
}

// DeserializeNet deserializes a Net.
func DeserializeNet(d []byte) (*Net, error) {
	var res Net
	err := serializer.DeserializeAny(d, &res.InputSize, &res.OutputType, &res.Net)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Evaluate generates a prediction for an image.
//
// The confidence measures how accurate the angle is
// likely to be.
// It should range between 0 and 1.
// Some output types do not yield a confidence measure.
func (n *Net) Evaluate(img image.Image) (angle, confidence float64) {
	if img.Bounds().Dx() != img.Bounds().Dy() ||
		img.Bounds().Dx() != n.InputSize {
		// Hack to crop the center square.
		img = Rotate(img, 0, n.InputSize)
	}
	inTensor := netInputTensor(img)
	inConst := anydiff.NewConst(anyvec32.MakeVectorData(inTensor))
	out := n.Net.Apply(inConst, 1).Output()
	switch n.OutputType {
	case RawAngle:
		return float64(anyvec.Sum(out).(float32)), 0
	case RightAngles:
		angles, probs := rightAngleMaxes(out)
		return float64(anyvec.Sum(angles).(float32)),
			float64(anyvec.Sum(probs).(float32))
	case ConfidenceAngle:
		angle := float64(anyvec.Sum(out.Slice(0, 1)).
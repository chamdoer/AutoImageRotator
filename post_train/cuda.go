// +build cuda

package main

import (
	"github.com/unixpickle/anyvec/anyvec32"
	"github.com/unixpickle/cudavec"
)

func init() {
	handle, er
// +build cuda

package main

import (
	"github.com/unixpickle/anyvec/anyvec32"
	"github.com/unixpickle/cudavec"
)

func init() {
	handle, err := cudavec.NewHandleDefault()
	if err != nil {
		panic(err)
	}
	anyvec32.
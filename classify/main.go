package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/unixpickle/autorot"
	"github.com/unixpickle/essentials"
	"github.com/unixpickle/serializer"
)

func main() {
	var dirPath string
	var netPath string
	flag.StringVar(&dirPath, "dir", "", "image directory")
	flag.StringVar(&netPath, "net", "", "network path")
	flag.Parse()
	if dirPath == "" || netPath == "" {
		essentials.Die("Required flags: -net and -dir. See -help for more.")
	}

	var net *autorot.Net
	if err := serializer.LoadAny(netPath, &net); err != nil {
		essentials.Die("Load network failed:", err)
	}

	outWriter := csv.NewWriter(os.Stdout)
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ext := strings.ToLower(filepath.Ex
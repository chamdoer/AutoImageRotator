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

	var net *autoro
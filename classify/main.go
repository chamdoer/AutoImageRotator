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
	"github.com/unixpi
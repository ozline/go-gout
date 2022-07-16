package gout

import (
	"io"
	"os"
)

const (
	debugCode = "debug"
)

var (
	goutMode = debugCode

	DefaultWriter io.Writer = os.Stdout
)

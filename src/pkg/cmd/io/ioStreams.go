package io

import (
	"github.com/mattn/go-colorable"
	"io"
	"os"
)

type Streams struct {
	Out io.Writer
	ErrOut io.Writer
}

func GetIOStreams() *Streams {
	io := &Streams{
		Out: colorable.NewColorable(os.Stdout),
		ErrOut: colorable.NewColorable(os.Stderr),
	}
	return io
}
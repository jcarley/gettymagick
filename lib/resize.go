package lib

import (
	"fmt"
	"os"

	bimg "gopkg.in/h2non/bimg.v1"
)

type Options struct {
	Compression int
	Destination string
	Height      int
	Quality     int
	Source      string
	Width       int
}

func ResizeImage(options Options) error {
	buffer, err := bimg.Read(options.Source)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	o := bimg.Options{
		Compression: options.Compression,
		Embed:       true,
		Height:      options.Height,
		Quality:     options.Quality,
		Width:       options.Width,
	}

	img := bimg.NewImage(buffer)

	newImage, err := img.Process(o)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write(options.Destination, newImage)

	buffer = buffer[:0]
	buffer = nil

	return nil
}

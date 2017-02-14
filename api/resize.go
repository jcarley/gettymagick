package api

import (
	"fmt"
	"os"

	bimg "gopkg.in/h2non/bimg.v1"
)

type Options struct {
	Source      string
	Destination string
	Width       int
	Height      int
	Quality     int
}

func ResizeImage(options Options) error {
	buffer, err := bimg.Read(options.Source)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	o := bimg.Options{
		Width:   options.Width,
		Height:  options.Height,
		Quality: options.Quality,
		Embed:   true,
	}
	newImage, err := bimg.NewImage(buffer).Process(o)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write(options.Destination, newImage)

	newImage = nil
	buffer = nil

	return nil
}

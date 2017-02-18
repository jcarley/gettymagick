package lib

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jcarley/gettymagick/helpers/jsonutil"

	bimg "gopkg.in/h2non/bimg.v1"
)

func ResizeImage(options Options) error {

	optionsJson, _ := jsonutil.EncodeJSONToString(&options)
	msg := fmt.Sprintf("Resizing %s", options.Source)
	defer Benchmark(time.Now(), msg)

	log.Printf("%s", optionsJson)

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

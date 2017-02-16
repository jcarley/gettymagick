package command

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/jcarley/gettymagick/lib"
	"github.com/mitchellh/cli"
)

type ResizeCommand struct {
	Ui          cli.Ui
	ResizeImage func(options lib.Options) error
}

func (this *ResizeCommand) Run(args []string) int {

	if len(args) < 2 {
		this.Ui.Info(this.Help())
		return 1
	}

	var sourceFile string
	var destinationFile string
	var sizeArg string
	var qualityArg string
	var compressionArg string

	flags := flag.NewFlagSet("resize", flag.ContinueOnError)
	flags.StringVar(&sizeArg, "size", "x600", "the size to resize the image to")
	flags.StringVar(&qualityArg, "quality", "80", "the quality to make the image")
	flags.StringVar(&compressionArg, "compression", "1", "the compression to apply to this image")

	if err := flags.Parse(args); err != nil {
		this.Ui.Error(err.Error())
		return 1
	}

	remainArgs := flags.Args()
	if len(remainArgs) < 2 {
		log.Printf("Gettymagick was called with to few args\n")
		this.Ui.Info(this.Help())
		return 1
	}

	sourceFile = remainArgs[0]
	destinationFile = remainArgs[1]

	dimensions := strings.Split(sizeArg, "x")
	if len(dimensions) < 2 {
		log.Printf("An invalid dimension was supplied for '%s'\n", sourceFile)
		this.Ui.Info(this.Help())
		return 1
	}

	widthValue := dimensions[0]
	heightValue := dimensions[1]

	width, err := strconv.Atoi(widthValue)
	if err != nil {
		numErr := err.(*strconv.NumError)

		// the string was empty, maintaining aspect ratio with height
		if numErr.Err == strconv.ErrSyntax && numErr.Num == "" {
			width = 0
		} else {
			// something else went wrong
			log.Printf("invalid width supplied for '%s'\n", sourceFile)
			return 1
		}

	}

	height, err := strconv.Atoi(heightValue)
	if err != nil {
		numErr := err.(*strconv.NumError)

		// the string was empty, maintaining aspect ratio with width
		if numErr.Err == strconv.ErrSyntax && numErr.Num == "" {
			height = 0
		} else {
			// something else went wrong
			log.Printf("invalid height supplied for '%s'\n", sourceFile)
			return 1
		}
	}

	quality, err := strconv.Atoi(qualityArg)
	if err != nil {
		log.Printf("invalid quality supplied for '%s', default to 60%\n", sourceFile)
		quality = 60
	}

	compression, err := strconv.Atoi(compressionArg)
	if err != nil {
		log.Printf("invalid compression supplied for '%s', default to 1\n", sourceFile)
		compression = 1
	}

	options := lib.Options{
		Compression: compression,
		Destination: destinationFile,
		Height:      height,
		Quality:     quality,
		Source:      sourceFile,
		Width:       width,
	}

	err = this.ResizeImage(options)
	if err != nil {
		log.Printf("An error occured for '%s': %s\n", sourceFile, err.Error())
		return 1
	}

	this.Ui.Info("Success")
	return 0
}

func (this *ResizeCommand) Help() string {
	helpText := `
Usage: gettymagick resize -size=WxH -quality=# -compression=$ <source> <destination>

	Resizes an image.
`
	return strings.TrimSpace(helpText)
}

func (this *ResizeCommand) Synopsis() string {
	return "Resizes an image."
}

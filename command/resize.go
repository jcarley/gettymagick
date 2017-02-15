package command

import (
	"flag"
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
		this.Ui.Info(this.Help())
		return 1
	}

	sourceFile = remainArgs[0]
	destinationFile = remainArgs[1]

	dimensions := strings.Split(sizeArg, "x")
	if len(dimensions) < 2 {
		this.Ui.Info(this.Help())
		return 1
	}

	widthValue := dimensions[0]
	heightValue := dimensions[1]
	width, err := strconv.Atoi(widthValue)
	if err != nil {
		this.Ui.Error("invalid width supplied")
		return 1
	}

	height, err_ := strconv.Atoi(heightValue)
	if err != nil {
		this.Ui.Error("invalid height supplied")
		return 1
	}

	quality, err := strconv.Atoi(qualityArg)
	if err != nil {
		this.Ui.Warn("invalid quality supplied, defaulting to 60%")
		quality = 60
	}

	compression, err := strconv.Atoi(compressionArg)
	if err != nil {
		this.Ui.Warn("invalid compression supplied, defaulting to 1")
		compression = 1
	}

	options := lib.Options{
		Source:      sourceFile,
		Destination: destinationFile,
		Quality:     quality,
		Width:       width,
		Height:      height,
	}

	err := this.ResizeImage(options)
	if err != nil {
		this.Ui.Error(err.Error())
		return 1
	}

	this.Ui.Info("Success")
	return 0
}

func (this *ResizeCommand) Help() string {
	helpText := `
Usage: gettymagick resize -size=WxH -quality=# <source> <destination>

	Resizes an image.
`
	return strings.TrimSpace(helpText)
}

func (this *ResizeCommand) Synopsis() string {
	return "Resizes an image."
}

package command

import (
	"flag"
	"strconv"
	"strings"

	"github.com/jcarley/gettymagick/api"
	"github.com/mitchellh/cli"
)

type ResizeCommand struct {
	Ui          cli.Ui
	ResizeImage func(options api.Options) error
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

	flags := flag.NewFlagSet("resize", flag.ContinueOnError)
	flags.StringVar(&sizeArg, "size", "x600", "the size to resize thumbnails to")
	flags.StringVar(&qualityArg, "quality", "80", "the quality to make the thumbnails")
	if err := flags.Parse(args); err != nil {
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
	width, _ := strconv.Atoi(widthValue)
	height, _ := strconv.Atoi(heightValue)

	quality, _ := strconv.Atoi(qualityArg)

	options := api.Options{
		Source:      sourceFile,
		Destination: destinationFile,
		Quality:     quality,
		Width:       width,
		Height:      height,
	}

	err := this.ResizeImage(options)
	if err != nil {
		this.Ui.Info("Failed")
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

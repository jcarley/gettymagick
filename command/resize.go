package command

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	bimg "gopkg.in/h2non/bimg.v1"
)

type ResizeCommand struct {
	Ui cli.Ui
}

func (this *ResizeCommand) Run(args []string) int {

	if len(args) < 3 {
		this.Ui.Info(this.Help())
		return 1
	}

	sourceFile := args[0]
	destinationFile := args[1]
	size := args[2]

	sizeValue := strings.Split(size, "=")

	err := this.resizeImage(sourceFile, destinationFile, sizeValue[1])
	if err != nil {
		this.Ui.Info("Failed")
		return 1
	}

	this.Ui.Info("Success")
	return 0
}

func (this *ResizeCommand) Help() string {
	helpText := `
Usage: gettymagick resize <source> <destination> size=WxH

	Resizes an image.
`
	return strings.TrimSpace(helpText)
}

func (this *ResizeCommand) Synopsis() string {
	return "Resizes an image."
}

func (this *ResizeCommand) resizeImage(sourceFile, destinationFile, sizeValue string) error {
	buffer, err := bimg.Read(sourceFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	widthValue := strings.Split(sizeValue, "x")[0]
	heightValue := strings.Split(sizeValue, "x")[1]

	width, _ := strconv.Atoi(widthValue)
	height, _ := strconv.Atoi(heightValue)

	newImage, err := bimg.NewImage(buffer).Resize(width, height)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write(destinationFile, newImage)

	newImage = nil
	buffer = nil

	return nil
}

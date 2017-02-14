package command

import (
	"os"
	"testing"

	. "github.com/jcarley/gettymagick/testing"

	"github.com/jcarley/gettymagick/api"
	"github.com/mitchellh/cli"
)

func TestRun(t *testing.T) {

	var expectedOptions api.Options

	cmd := ResizeCommand{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
		ResizeImage: func(options api.Options) error {
			expectedOptions = options
			return nil
		},
	}

	sourceFile := "/tmp/masters/source.jpg"
	destFile := "/tmp/thumbnail/source.jpg"

	args := []string{
		"-quality",
		"60",
		"-size",
		"400x600",
		sourceFile,
		destFile,
	}

	cmd.Run(args)

	AssertTrue(t, expectedOptions.Source == sourceFile, "Source should equal masters/source.jpg")
	AssertTrue(t, expectedOptions.Destination == destFile, "Destination should equal thumbnail/source.jpg")
	AssertTrue(t, expectedOptions.Quality == 60, "Quality should equal 60")
	AssertTrue(t, expectedOptions.Width == 400, "Width should equal 400")
	AssertTrue(t, expectedOptions.Height == 600, "Width should equal 600")
}

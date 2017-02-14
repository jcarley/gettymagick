package command

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

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

	Convey("Options were set from args", t, func() {
		So(expectedOptions.Source, ShouldEqual, sourceFile)
		So(expectedOptions.Destination, ShouldEqual, destFile)
		So(expectedOptions.Quality, ShouldEqual, 60)
		So(expectedOptions.Width, ShouldEqual, 400)
		So(expectedOptions.Height, ShouldEqual, 600)
	})

}

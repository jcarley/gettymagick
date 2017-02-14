package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jcarley/gettymagick/api"
	"github.com/jcarley/gettymagick/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Serf commands.
var Commands map[string]cli.CommandFactory

func init() {
	prefixedUi := &cli.PrefixedUi{
		InfoPrefix: "==> ",
		Ui:         &cli.BasicUi{Writer: os.Stdout},
	}

	ui := &cli.ColoredUi{
		InfoColor:  cli.UiColorGreen,
		ErrorColor: cli.UiColorRed,
		Ui:         prefixedUi,
	}

	Commands = map[string]cli.CommandFactory{

		"resize": func() (cli.Command, error) {
			return &command.ResizeCommand{
				Ui:          ui,
				ResizeImage: api.ResizeImage,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                ui,
			}, nil
		},
	}

}

// makeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// interrupt received.
func makeShutdownCh() <-chan struct{} {
	resultCh := make(chan struct{})

	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()

	return resultCh
}

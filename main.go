package main

import (
	"io"
	"log"
	"os"
	"runtime"

	"github.com/mitchellh/cli"
)

func main() {
	procs := runtime.NumCPU() - 2
	if procs <= 0 {
		procs = 1
	}
	runtime.GOMAXPROCS(procs)
	os.Exit(realMain())
}

func realMain() int {

	// setup application logging
	log.SetFlags(log.LstdFlags)

	err := os.Mkdir("logs", 0755)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("An error occured creating the log directory: %q\n", err.Error())
			return 1
		}
	}

	logFile, err := os.OpenFile("logs/gettymagick.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Fatalln("Permission denied.")
		} else {
			log.Fatalf("Failed to open log file: %q\n", err.Error())
		}
		return 1
	}

	multiWriter := io.MultiWriter(UiWriter, logFile)
	log.SetOutput(multiWriter)

	return wrappedMain()
}

func wrappedMain() int {

	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("gettymagick"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		log.Printf("Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}

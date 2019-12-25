package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
	appConfig                 = config{}
	configFileLocation        string
	defaultOutputFileLocation = "output.json"
	showVersionAndExit        = false
	silentMode                = false
)

func showVersion() bool {
	fmt.Println("Name:        ", appName)
	fmt.Println("Manufacturer:", appMaintainer)
	fmt.Println("Version:     ", appVersion)
	fmt.Println("GoVersion:   ", runtime.Version())
	fmt.Println("Os:          ", runtime.GOOS)
	fmt.Println("Arch:        ", runtime.GOARCH)

	return true
}

func initCMDFlags() {
	flag.BoolVar(&showVersionAndExit, "version", false, "print version and exit")
	flag.BoolVar(&silentMode, "silent", false, "do not print anything except errors")
	flag.StringVar(&configFileLocation, "config", "", "path to configuration file, overrides values passed directly as arguments")
	flag.Var(&appConfig.FeedURLs, "feed", "URL(s) of RSS feed(s) to read")
	flag.StringVar(&appConfig.OutputFilePath, "output", "", "path to output json file")
}

func processCMDFlags() bool {
	flag.Parse()

	if showVersionAndExit {
		return showVersion()
	}
	err := tryLoadConfig()
	if err != nil {
		println("failed to load configuration file: ", err.Error())
	}

	if len(appConfig.FeedURLs) == 0 {
		println("0 feed URLs passed: nothing to do")
		println("Usage:")
		flag.PrintDefaults()

		os.Exit(0)
	}

	if appConfig.OutputFilePath == "" {
		fmt.Printf("output file was not passed, using default one: %q\n", defaultOutputFileLocation)
		appConfig.OutputFilePath = defaultOutputFileLocation
	}

	return false
}

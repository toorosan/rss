package main

import (
	"fmt"
	"os"

	"github.com/toorosan/rss/pkg/rssreader"
)

const (
	// appName defines application name
	appName = "RSS Reader"
	// appName defines application version
	appVersion = "0.0.1"
	// appManufacturer defines application maintainer information
	appMaintainer = "https://github.com/toorosan"
)

func main() {
	fmt.Printf("%q initiated\n", appName)
	initCMDFlags()
	if processCMDFlags() {
		return
	}
	if runApp() {
		defer os.Exit(1)
	}
	fmt.Printf("%q is stopped\n", appName)
}

func runApp() bool {
	results, err := rssreader.Parse(appConfig.FeedURLs)
	toFile := make([]rssItem, len(results))
	for i := range results {
		toFile[i].Title = results[i].Title
		toFile[i].Source = results[i].Source
		toFile[i].SourceURL = results[i].SourceURL
		toFile[i].Link = results[i].Link
		toFile[i].PubishDate = results[i].PubishDate
		toFile[i].Description = results[i].Description

		println(results[i].String())
	}
	if err != nil {
		println("following errors occured:\n", err.Error())
	}

	return err != nil
}

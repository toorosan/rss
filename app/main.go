package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// prepare marshalable objects and print RSS parse results output.
	resultsConvertable := make([]rssItem, len(results))
	for i := range results {
		resultsConvertable[i].Title = results[i].Title
		resultsConvertable[i].Source = results[i].Source
		resultsConvertable[i].SourceURL = results[i].SourceURL
		resultsConvertable[i].Link = results[i].Link
		resultsConvertable[i].PubishDate = results[i].PubishDate
		resultsConvertable[i].Description = results[i].Description
		if !silentMode {
			println(results[i].String())
		}
	}
	// print parsing errors if any.
	if err != nil {
		println("following parsing errors occured:\n", err.Error())
	}
	// prepare marshaled data from results.
	toFile, err := json.Marshal(resultsConvertable)
	if err != nil {
		panic(fmt.Errorf("failed to marshal results to json: %v", err))
	}

	// create file with permissions to read-write for owner and read for others.
	err = ioutil.WriteFile(appConfig.OutputFilePath, toFile, 0644)
	if err != nil {
		println(fmt.Errorf("failed to write json results to output file: %v", err))
	}

	return err != nil
}

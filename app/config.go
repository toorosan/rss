package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type config struct {
	FeedURLs       arrayOfStrings `json:"feedURLs"`
	OutputFilePath string         `json:"outputFilePath"`
}

// arrayOfStrings describes array of string values, used in command line arguments to pass multiple URLs.
type arrayOfStrings []string

func (aos *arrayOfStrings) String() string {
	return strings.Join(*aos, ",")
}

func (aos *arrayOfStrings) Set(value string) error {
	*aos = append(*aos, value)

	return nil
}

func tryLoadConfig() error {
	if configFileLocation == "" {
		return nil
	}
	data, err := ioutil.ReadFile(configFileLocation)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &appConfig)
}

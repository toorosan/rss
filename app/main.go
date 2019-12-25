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
	if RssDemo() {
		defer os.Exit(1)
	}
	fmt.Printf("%q is stopped\n", appName)
}

func RssDemo() bool {
	urls := []string{"https://www.feedforall.com/sample.xml", "https://www.feedforall.com/sample-feed.xml"}
	results, err := rssreader.Parse(urls)
	for _, item := range results {
		println(item.String())
	}
	if err != nil {
		println("following errors occured:\n", err.Error())
	}

	return err != nil
}

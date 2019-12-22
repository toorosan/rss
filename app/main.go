package main

import "fmt"

const (
	// appName defines application name
	appName = "RSS Reader application"
)

func main() {
	defer fmt.Printf("%q is stopped\n", appName)
	fmt.Printf("%q initiated\n", appName)
}

package rssreader

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
)

// feedGetter holds instance of getter object, defined for testing purposes.
var feedGetter getter = &webGetter{}

// Parse holds main logic for parsing multiple urls.
// For now simply log all the errors, try to return anything.
func Parse(urls []string) []RssItem {
	var (
		// parsedURLs defines map of URLs prepared to be processed.
		parsedURLs = map[string]*url.URL{}
		// resultsTotal defines counter for total amount of results.
		resultsTotal int32 = 0
		// define wait group to ensure that all URLs were processed at the end.
		wg sync.WaitGroup
	)
	for i := range urls {
		rawURL := urls[i]
		feedURL, err := url.Parse(rawURL)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to parse feed url %q: %v", rawURL, err))

			continue
		}
		// this assignment allows us to skip processing in case if URLs are repeated multiple times.
		parsedURLs[feedURL.String()] = feedURL
	}
	// resultGroups defines array of acquired RSS Items, grouped by source URL.
	resultGroups := make([][]RssItem, len(parsedURLs))
	j := int32(-1)
	for i := range parsedURLs {
		atomic.AddInt32(&j, 1)
		k := j + 0
		feedURL := parsedURLs[i]
		feedURLstr := feedURL.String()
		wg.Add(1)
		// run every URL processing in parallel.
		go func() {
			defer wg.Done()
			parser, err := newParser(*feedURL, feedGetter)
			if err != nil {
				fmt.Println(fmt.Errorf("failed to get specific rss parser for url %q: %v", feedURLstr, err))

				return
			}
			resultGroup, err := parser.Parse()
			if err != nil {
				fmt.Println(fmt.Errorf("failed to parse feed from url %q: %v", feedURLstr, err))

				return
			}
			resultGroups[k] = resultGroup
			// increase size of results to make precise array size in the end.
			atomic.AddInt32(&resultsTotal, int32(len(resultGroup)))
		}()
	}
	wg.Wait()

	results := make([]RssItem, 0, resultsTotal)
	for _, r := range resultGroups {
		results = append(results, r...)
	}

	return results
}

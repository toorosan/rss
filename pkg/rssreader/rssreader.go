package rssreader

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
)

// feedGetter holds instance of getter object, defined for testing purposes.
var feedGetter getter = &webGetter{}

// Parse holds main logic for parsing multiple urls
func Parse(urls []string) []RssItem {
	var (
		// errs defines map for potential errors.
		errs = map[string]error{}
		// parsedURLs defines map of URLs prepared to be processed.
		parsedURLs = map[string]*url.URL{}
		// resultsMap defines map for acquired RSS Items.
		resultMap = map[string][]RssItem{}
		// resultsTotal defines counter for total amount of results.
		resultsTotal int32 = 0
		// define wait group to ensure that all URLs were processed at the end.
		wg sync.WaitGroup
	)
	for i := range urls {
		rawURL := urls[i]
		feedURL, err := url.Parse(rawURL)
		if err != nil {
			errs[rawURL] = fmt.Errorf("failed to parse feed url %q: %v", rawURL, err)

			continue
		}
		// this assignment allows us to skip processing in case if URLs are repeated multiple times.
		parsedURLs[feedURL.String()] = feedURL
	}
	for i := range parsedURLs {
		feedURL := parsedURLs[i]
		feedURLstr := feedURL.String()
		wg.Add(1)
		// run every URL processing in parallel.
		go func() {
			defer wg.Done()
			parser, err := newParser(*feedURL, feedGetter)
			if err != nil {
				errs[feedURLstr] = fmt.Errorf("failed to get specific rss parser for url %q: %v", feedURLstr, err)

				return
			}
			resultMap[feedURLstr], err = parser.Parse()
			if err != nil {
				errs[feedURLstr] = fmt.Errorf("failed to parse feed from url %q: %v", feedURLstr, err)

				return
			}
			// increase size of results to make precise array size in the end.
			atomic.AddInt32(&resultsTotal, int32(len(resultMap[feedURLstr])))
		}()
	}
	wg.Wait()

	for u, err := range errs {
		// for now simply log all the errors.
		fmt.Printf("failed to acquire rss items for %q: %v\n", u, err)
	}

	results := make([]RssItem, 0, resultsTotal)
	for _, r := range resultMap {
		results = append(results, r...)
	}

	return results
}

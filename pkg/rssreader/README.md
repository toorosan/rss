RSS Reader package.

Allows to read multiple RSS feeds in parallel. Returns array of RssItem and error object.
Keep in mind, as there are multiple RSS feeds possible, method Parse will not fail in case of errors.
It will return all successfully parsed items and errors for the rest of feeds failed to read.

Supported RSS protocol versions:
- RSS2.0

Usage example:
```
import (
    os

	"github.com/toorosan/rss/pkg/rssreader"
)

func main() {
	urls := []string{"https://www.feedforall.com/sample.xml", "https://www.feedforall.com/sample-feed.xml"}
	results, err := rssreader.Parse(urls)
	for _, item := range results {
		println(item.String())
	}
	if err != nil {
		println("following errors occured:\n", err.Error())
        os.exit(1)
	}
}
```

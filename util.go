package googlenews

import (
	"net/http"

	"github.com/mmcdole/gofeed/rss"
)

func pointer[T int64](i T) *T {
	return &i
}

func handleFeedResponse(resp *http.Response) (*rss.Feed, error) {
	p := rss.Parser{}
	defer resp.Body.Close()
	return p.Parse(resp.Body)
}

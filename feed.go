package googlenews

import (
	"strings"

	"github.com/mmcdole/gofeed/rss"
)

func createFeed(s string) (*rss.Feed, error) {
	p := rss.Parser{}
	feed, err := p.Parse(strings.NewReader(s))
	return feed, err
}

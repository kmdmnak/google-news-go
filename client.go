package googlenews

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mmcdole/gofeed/rss"
)

// addLanguageParams 言語クエリーの追加
func addLanguageParams(val url.Values, lang *langProperty) {
	val.Set("ceid", lang.CeID)
	val.Set("gl", lang.Gl)
	val.Set("hl", lang.Hl)
}

type client struct {
	base    string
	lang    Language
	hclient *http.Client
}

func (c *client) newRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	val := req.URL.Query()
	addLanguageParams(val, getLanguageProperty(c.lang))
	req.URL.RawQuery = val.Encode()
	return req, nil
}

func (c *client) do(req *http.Request) (*rss.Feed, *http.Response, error) {
	resp, err := c.hclient.Do(req)
	if err != nil {
		return nil, resp, err
	}
	feeds, err := handleFeedResponse(resp)
	if err != nil {
		return nil, resp, err
	}
	return feeds, resp, nil
}

// SearchByTopic get headlines of specific topic
func (c *client) SearchByTopic(ctx context.Context, topic NewsTopic) (*rss.Feed, *http.Response, error) {
	url := fmt.Sprintf("%s/headlines/section/topic/%s", c.base, url.QueryEscape(topic.string()))
	req, err := c.newRequest(ctx, url)
	if err != nil {
		return nil, nil, err
	}
	return c.do(req)
}

// SearchByQuery get headlines using query
func (c *client) SearchByQuery(ctx context.Context, params *QueryParameter) (*rss.Feed, *http.Response, error) {
	url := fmt.Sprintf("%s/search", c.base)
	req, err := c.newRequest(ctx, url)
	if err != nil {
		return nil, nil, err
	}
	val := req.URL.Query()
	val.Add("q", params.buildQueryString())
	req.URL.RawQuery = val.Encode()
	return c.do(req)
}

// SearchByGeometry get headlines using place name, country name etc.
func (c *client) SearchByGeometry(ctx context.Context, word string) (*rss.Feed, *http.Response, error) {
	url := fmt.Sprintf("%s/headlines/section/geo/%s", c.base, url.QueryEscape(word))
	req, err := c.newRequest(ctx, url)
	if err != nil {
		return nil, nil, err
	}
	return c.do(req)
}

func NewClient(lang Language) *client {
	c := client{
		base:    "https://news.google.com/news/rss",
		lang:    lang,
		hclient: &http.Client{},
	}
	return &c
}

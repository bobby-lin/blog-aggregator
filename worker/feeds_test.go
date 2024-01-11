package worker

import (
	"testing"
)

func TestFetchDataFromFeed(t *testing.T) {
	tests := []struct {
		url, title string
	}{
		{"https://blog.boot.dev/index.xml", "Boot.dev Blog"},
		{"https://rss.nytimes.com/services/xml/rss/nyt/Technology.xml", "NYT > Technology"},
	}

	for _, tt := range tests {
		rss, err := FetchDataFromFeed(tt.url)

		if err != nil {
			t.Error(err)
		}

		if rss.Channel.Title != tt.title {
			t.Errorf("got %s expect %s", rss.Channel.Title, tt.title)
		}
	}
}

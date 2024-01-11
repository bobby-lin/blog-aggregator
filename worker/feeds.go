package worker

import (
	"context"
	"database/sql"
	"encoding/xml"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"io"
	"log"
	"net/http"
	"time"
)

func (w Worker) GetNextFeedsToFetch(n int32) ([]database.Feed, error) {
	return w.DB.SelectNextFeedsToFetch(context.Background(), n)
}

func (w Worker) MarkFeedFetched(feed database.Feed) (database.Feed, error) {
	t := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	params := database.UpdateLastFetchedAtParams{LastFetchedAt: t, ID: feed.ID, UpdatedAt: time.Now()}
	updatedFeed, err := w.DB.UpdateLastFetchedAt(context.Background(), params)

	if err != nil {
		return database.Feed{}, err
	}

	log.Println("Updated feed " + updatedFeed.ID.String() + " at " + updatedFeed.LastFetchedAt.Time.GoString())
	return updatedFeed, nil
}

type FeedRss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Item        []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func FetchDataFromFeed(url string) (FeedRss, error) {
	resp, _ := http.Get(url)
	rss := FeedRss{}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return FeedRss{}, err
	}

	err = xml.Unmarshal(b, &rss)
	if err != nil {
		return FeedRss{}, err
	}

	return rss, nil
}

package worker

import (
	"context"
	"database/sql"
	"encoding/xml"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"sync"
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

	return updatedFeed, nil
}

type FeedRss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text          string `xml:",chardata"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Description   string `xml:"description"`
		PublishedDate string `xml:"pubDate"`
		Item          []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
				Rel  string `xml:"rel,attr"`
			} `xml:"link"`
			Description   string `xml:"description"`
			PublishedDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func FetchDataFromFeed(url string) (FeedRss, error) {
	resp, err := http.Get(url)
	if err != nil {
		return FeedRss{}, err
	}

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

func (w Worker) Loop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		ctx := context.Background()
		feedBatch, _ := w.DB.SelectNextFeedsToFetch(context.Background(), w.FetchSize)

		wg := sync.WaitGroup{}

		for _, v := range feedBatch {
			wg.Add(1)
			v := v // intermediate variable?

			go func() {
				defer wg.Done()
				feedRss, err := FetchDataFromFeed(v.Url)
				if err != nil {
					log.Println(err)
				}

				for _, i := range feedRss.Channel.Item {
					id, err := uuid.NewUUID()

					parse, err := time.Parse(time.RFC1123Z, i.PublishedDate)
					if err != nil {
						return
					}

					url := i.Link.Text

					if url == "" {
						url = i.Link.Href
					}

					postParam := database.CreatePostParams{
						ID:          id,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
						Title:       i.Title,
						Url:         url,
						Description: i.Description,
						PublishedAt: sql.NullTime{
							Time:  parse,
							Valid: true,
						},
						FeedID: v.ID,
					}

					_, err = w.DB.CreatePost(ctx, postParam)

					if err != nil {
						pgErr, ok := err.(*pq.Error)

						if ok {
							if pgErr.Code != "23505" {
								log.Println(err)
							}
						}
					}
				}

				_, err = w.MarkFeedFetched(v)
				if err != nil {
					log.Println(err)
				}

				log.Print(feedRss.Channel.Title)
			}()
		}

		wg.Wait()
	}

}

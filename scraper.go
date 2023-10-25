package main

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/tintin007k/rss-agg/internal/database"
)

type RSSFeed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	XMLName       xml.Name  `xml:"channel"`
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	Generator     string    `xml:"generator"`
	Language      string    `xml:"language"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func startScraping(db *database.Queries, concurrency int, frequency time.Duration) {
	log.Println("inside start scraping")
	ticker := time.NewTicker(frequency)

	for ; ; <-ticker.C {
		feedsToFetched, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error in fetching the feeds to be fetched")
			continue
		}
		log.Println("got: ", feedsToFetched, " to be fetched")

		wg := &sync.WaitGroup{}

		for _, feed := range feedsToFetched {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.UpdateLastFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("feed with id %s could not be updated: %v", feed.Name, err)
		return
	}
	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("Feed with URL: %s could no be fetched with error: %v", feed.Url, err)
		return
	}
	for _, item := range feedData.Channel.Items {
		log.Println("Found post", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Items))

}

func fetchFeed(feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}

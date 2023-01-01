package googlenews_test

import (
	"context"
	"fmt"
	"log"
	"time"

	googlenews "github.com/kmdmnak/google-news-go"
)

func ExampleSearchByGeometry() {
	client := googlenews.NewClient(googlenews.JPN)
	ctx := context.Background()
	feeds, _, err := client.SearchByGeometry(ctx, "tokyo")
	if err != nil {
		log.Fatal("failed to fetch")
	}
	for _, item := range feeds.Items {
		fmt.Printf("title=%s", item.Title)
	}
}

func ExampleSearchByTopic() {
	client := googlenews.NewClient(googlenews.JPN)
	ctx := context.Background()
	feeds, _, err := client.SearchByTopic(ctx, googlenews.TOPIC_NATION)
	if err != nil {
		log.Fatal("failed to fetch")
	}
	for _, item := range feeds.Items {
		fmt.Printf("title=%s", item.Title)
	}
}

func ExampleSearchByQuery() {
	client := googlenews.NewClient(googlenews.JPN)
	ctx := context.Background()
	date := time.Date(2022, 10, 10, 0, 0, 0, 9, nil)
	feeds, _, err := client.SearchByQuery(ctx, &googlenews.QueryParameter{
		After: &date,
		Words: []string{"soccer", "basketball"},
	})
	if err != nil {
		log.Fatal("failed to fetch")
	}
	for _, item := range feeds.Items {
		fmt.Printf("title=%s", item.Title)
	}
}

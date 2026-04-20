package streaming_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	streaming "github.com/movieofthenight/go-streaming-availability/v4"
)

func ExampleNewAPIClientFromApiKey() {
	apiKey, apiKeyFound := os.LookupEnv("API_KEY")
	if !apiKeyFound {
		log.Fatal("API_KEY not found")
	}
	streaming.NewAPIClientFromApiKey(apiKey, nil)
}

func ExampleShowsAPIService_GetShow() {
	apiKey, apiKeyFound := os.LookupEnv("API_KEY")
	if !apiKeyFound {
		log.Fatal("API_KEY not found")
	}
	client := streaming.NewAPIClientFromApiKey(apiKey, nil)
	show, _, err := client.ShowsAPI.GetShow(context.Background(), "tt0068646").Country("us").Execute()
	if err != nil {
		log.Fatal(err)
	}
	if show == nil {
		log.Fatal("Show not found")
	}
	fmt.Printf("Title: %s\n", show.Title)
	fmt.Printf("Overview: %s\n", show.Overview)
	for _, streamingOption := range show.StreamingOptions["us"] {
		fmt.Printf("Available on %s", streamingOption.Service.Name)
		switch streamingOption.Type {
		case streaming.ADDON:
			fmt.Printf(" via addon %s", streamingOption.Addon.Name)
		case streaming.BUY:
			fmt.Print(" to buy")
		case streaming.RENT:
			fmt.Print(" to rent")
		case streaming.FREE:
			fmt.Print(" for free")
		}
		if streamingOption.HasPrice() {
			fmt.Printf(" for %s", streamingOption.Price.Formatted)
		}
		if streamingOption.HasQuality() {
			fmt.Printf(" in %s quality", strings.ToUpper(*streamingOption.Quality))
		}
		fmt.Printf(" at %s\n", streamingOption.Link)
	}
}

func ExampleShowsAPIService_SearchShowsByFilters() {
	apiKey, apiKeyFound := os.LookupEnv("API_KEY")
	if !apiKeyFound {
		log.Fatal("API_KEY not found")
	}
	client := streaming.NewAPIClientFromApiKey(apiKey, nil)
	searchResult, _, err := client.ShowsAPI.SearchShowsByFilters(context.Background()).
		Genres([]string{"comedy"}).
		OrderBy("popularity_1year").
		Country("us").
		Catalogs([]string{"netflix"}).Execute()
	if err != nil {
		log.Fatal(err)
	}
	if searchResult == nil {
		log.Fatal("Search result not found")
	}
	if len(searchResult.Shows) == 0 {
		log.Fatal("No shows found")
	}
	for _, show := range searchResult.Shows {
		fmt.Printf("Title: %s\n", show.Title)
		fmt.Printf("Overview: %s\n", show.Overview)
		for _, streamingOption := range show.StreamingOptions["us"] {
			if streamingOption.Service.Id != "netflix" {
				continue
			}
			fmt.Printf("Link: %s\n", streamingOption.Link)
		}
	}
}

func ExampleApiSearchShowsByFiltersRequest_ExecuteWithAutoPagination() {
	apiKey, apiKeyFound := os.LookupEnv("API_KEY")
	if !apiKeyFound {
		log.Fatal("API_KEY not found")
	}
	client := streaming.NewAPIClientFromApiKey(apiKey, nil)
	shows, err := client.ShowsAPI.SearchShowsByFilters(context.Background()).
		Genres([]string{"comedy"}).
		OrderBy("popularity_1year").
		Country("us").
		Catalogs([]string{"netflix"}).ExecuteWithAutoPagination(3)
	if err != nil {
		log.Fatal(err)
	}
	for shows.Next() {
		show := shows.Get()
		fmt.Printf("Title: %s\n", show.Title)
		fmt.Printf("Overview: %s\n", show.Overview)
		for _, streamingOption := range show.StreamingOptions["us"] {
			if streamingOption.Service.Id != "netflix" {
				continue
			}
			fmt.Printf("Link: %s\n", streamingOption.Link)
		}
	}
	if shows.Err() != nil {
		log.Fatal(shows.Err())
	}
}

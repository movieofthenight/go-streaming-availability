package streaming_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/movieofthenight/go-streaming-availability/v4"
)

func ExampleNewConfiguration() {
	rapidApiKey, _ := os.LookupEnv("RAPID_API_KEY")
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
}

func ExampleNewAPIClient() {
	rapidApiKey, _ := os.LookupEnv("RAPID_API_KEY")
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	streaming.NewAPIClient(configuration)
}

func ExampleShowsAPIService_GetShow() {
	rapidApiKey, _ := os.LookupEnv("RAPID_API_KEY")
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	client := streaming.NewAPIClient(configuration)
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
	rapidApiKey, _ := os.LookupEnv("RAPID_API_KEY")
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	client := streaming.NewAPIClient(configuration)
	show, _, err := client.ShowsAPI.GetShow(context.Background(), "tt0068646").Country("us").Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Title:", show.Title)
	fmt.Println("Overview:", show.Overview)
	for _, streamingOption := range show.StreamingOptions["us"] {
		fmt.Print("Available on ", streamingOption.Service.Name)
		switch streamingOption.Type {
		case streaming.ADDON:
			fmt.Print(" via addon " + streamingOption.Addon.Name)
		case streaming.BUY:
			fmt.Print(" to buy")
		case streaming.RENT:
			fmt.Print(" to rent")
		case streaming.FREE:
			fmt.Print(" for free")
		}
		if streamingOption.HasPrice() {
			fmt.Print(" for ", streamingOption.Price.Formatted)
		}
		if streamingOption.HasQuality() {
			fmt.Print(" in ", strings.ToUpper(*streamingOption.Quality), " quality")
		}
		fmt.Println(" at " + streamingOption.Link)
	}
}

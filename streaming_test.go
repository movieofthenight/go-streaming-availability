package streaming_test

import (
	"context"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/movieofthenight/go-streaming-availability/v4"
)

func TestGetTheGodfather(t *testing.T) {
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
	t.Logf("Title: %s\n", show.Title)
	t.Logf("Overview: %s\n", show.Overview)
	for _, streamingOption := range show.StreamingOptions["us"] {
		t.Logf("Available on %s", streamingOption.Service.Name)
		switch streamingOption.Type {
		case streaming.ADDON:
			t.Logf(" via addon %s", streamingOption.Addon.Name)
		case streaming.BUY:
			t.Log(" to buy")
		case streaming.RENT:
			t.Log(" to rent")
		case streaming.FREE:
			t.Log(" for free")
		}
		if streamingOption.Price != nil {
			t.Logf(" for %s", streamingOption.Price.Formatted)
		}
		if streamingOption.Quality != nil {
			t.Logf(" in %s quality", strings.ToUpper(*streamingOption.Quality))
		}
		t.Logf(" at %s\n", streamingOption.Link)
	}
}

func TestSearchPopularComedyShowsOnNetflix(t *testing.T) {
	rapidApiKey, _ := os.LookupEnv("RAPID_API_KEY")
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	client := streaming.NewAPIClient(configuration)
	searchResult, _, err := client.ShowsAPI.SearchShowsByFilters(context.Background()).
		Genres([]string{"comedy"}).
		OrderBy("popularity_1year").
		DescendingOrder(true).
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
		t.Logf("Title: %s\n", show.Title)
		t.Logf("Overview: %s\n", show.Overview)
		for _, streamingOption := range show.StreamingOptions["us"] {
			if streamingOption.Service.Id != "netflix" {
				continue
			}
			t.Logf("Link: %s\n", streamingOption.Link)
		}
	}
}

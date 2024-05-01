package streaming_test

import (
	"context"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/movieofthenight/go-streaming-availability/v4"
)

func TestShowsAPIService_GetShow(t *testing.T) {
	rapidApiKey, rapidApiKeyFound := os.LookupEnv("RAPID_API_KEY")
	if !rapidApiKeyFound {
		t.Fatal("RAPID_API_KEY not found")
	}
	client := streaming.NewAPIClientFromRapidAPIKey(rapidApiKey, nil)
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

func TestShowsAPIService_SearchShowsByFilters(t *testing.T) {
	rapidApiKey, rapidApiKeyFound := os.LookupEnv("RAPID_API_KEY")
	if !rapidApiKeyFound {
		t.Fatal("RAPID_API_KEY not found")
	}
	client := streaming.NewAPIClientFromRapidAPIKey(rapidApiKey, nil)
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

func TestApiSearchShowsByFiltersRequest_ExecuteWithAutoPagination(t *testing.T) {
	rapidApiKey, rapidApiKeyFound := os.LookupEnv("RAPID_API_KEY")
	if !rapidApiKeyFound {
		t.Fatal("RAPID_API_KEY not found")
	}
	client := streaming.NewAPIClientFromRapidAPIKey(rapidApiKey, nil)
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
		t.Logf("Title: %s\n", show.Title)
		t.Logf("Overview: %s\n", show.Overview)
		for _, streamingOption := range show.StreamingOptions["us"] {
			if streamingOption.Service.Id != "netflix" {
				continue
			}
			t.Logf("Link: %s\n", streamingOption.Link)
		}
	}
	if shows.Err() != nil {
		log.Fatal(shows.Err())
	}
}

func TestApiSearchShowsByFiltersRequest_ExecuteWithAutoPagination2(t *testing.T) {
	rapidApiKey, rapidApiKeyFound := os.LookupEnv("RAPID_API_KEY")
	if !rapidApiKeyFound {
		t.Fatal("RAPID_API_KEY not found")
	}
	client := streaming.NewAPIClientFromRapidAPIKey(rapidApiKey, nil)
	shows, err := client.ShowsAPI.SearchShowsByFilters(context.Background()).
		Genres([]string{"comedy"}).
		OrderBy("popularity_1year").
		YearMax(2020).
		YearMin(2019).
		Country("us").
		Catalogs([]string{"netflix"}).ExecuteWithAutoPagination(0)
	if err != nil {
		log.Fatal(err)
	}
	for shows.Next() {
		show := shows.Get()
		t.Logf("Title: %s\n", show.Title)
		t.Logf("Overview: %s\n", show.Overview)
		for _, streamingOption := range show.StreamingOptions["us"] {
			if streamingOption.Service.Id != "netflix" {
				continue
			}
			t.Logf("Link: %s\n", streamingOption.Link)
		}
	}
	if shows.Err() != nil {
		log.Fatal(shows.Err())
	}
}

func TestApiGetChangesRequest_ExecuteWithAutoPagination(t *testing.T) {
	rapidApiKey, rapidApiKeyFound := os.LookupEnv("RAPID_API_KEY")
	if !rapidApiKeyFound {
		t.Fatal("RAPID_API_KEY not found")
	}
	client := streaming.NewAPIClientFromRapidAPIKey(rapidApiKey, nil)
	changes, err := client.ChangesAPI.GetChanges(context.Background()).
		Catalogs([]string{"netflix"}).
		ItemType("show").
		ChangeType("new").
		Country("us").
		ExecuteWithAutoPagination(2)
	if err != nil {
		log.Fatal(err)
	}
	for changes.Next() {
		change := changes.Get()
		t.Logf("Title: %s\n", change.Show.Title)
		t.Logf("Overview: %s\n", change.Show.Overview)
	}
	if changes.Err() != nil {
		log.Fatal(changes.Err())
	}
}

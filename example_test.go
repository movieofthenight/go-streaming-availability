package streaming_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/movieofthenight/go-streaming-availability/v3"
)

func ExampleConfiguration() {
	var rapidApiKey, _ = os.LookupEnv("RAPID_API_KEY")
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
}

func ExampleNewAPIClient() {
	var rapidApiKey, _ = os.LookupEnv("RAPID_API_KEY")
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	streaming.NewAPIClient(configuration)
}

// Get The Dark Knight's Streaming Info
func ExampleDefaultAPIService_GetById() {
	var rapidApiKey, _ = os.LookupEnv("RAPID_API_KEY")

	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	client := streaming.NewAPIClient(configuration).DefaultAPI

	country := "us"
	imdbId := "tt0468569" // Imdb id of The Dark Knight

	response, _, err := client.GetById(context.Background()).ImdbId(imdbId).Execute()
	if err != nil {
		log.Fatal(err)
	}
	result := response.Result
	for _, streamingOption := range result.StreamingInfo[country] {
		fmt.Printf("%s is available on %s via %s at link %s", result.GetTitle(), streamingOption.GetService(), streamingOption.GetStreamingType(), streamingOption.GetLink())
		if streamingOption.HasQuality() {
			fmt.Printf(" with %s quality", strings.ToUpper(streamingOption.GetQuality()))
		}
		if streamingOption.HasPrice() {
			fmt.Printf(" for %s", streamingOption.Price.Formatted)
		}
		fmt.Println()
	}
}

// Search for Science Fiction Movies on Netflix and Disney+
func ExampleDefaultAPIService_SearchByFilters_scienceFictionMovies() {
	var rapidApiKey, _ = os.LookupEnv("RAPID_API_KEY")

	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	client := streaming.NewAPIClient(configuration).DefaultAPI

	country := "us"
	services := []string{"netflix", "disney"}
	showType := "movie"
	// 878 is Science-Fiction
	// Use Genres endpoint to get the genre ids
	genre := "878"
	maxPages := 2

	var cursor string
	for i := 0; i < maxPages; i++ {
		fmt.Printf("Page: %d\n", i)
		response, _, err := client.SearchByFilters(
			context.Background()).
			Country(country).
			Services(strings.Join(services, ",")).
			ShowType(showType).
			Genres(genre).
			Cursor(cursor).
			Execute()
		if err != nil {
			log.Fatal(err)
		}
		for _, show := range response.Result {
			for _, streamingOption := range show.StreamingInfo[country] {
				for _, service := range services {
					if streamingOption.GetService() == service {
						fmt.Printf("%s is available on %s: %s\n", show.Title, service, streamingOption.Link)
					}
				}
			}
		}

		// Break out of the loop if there are no more results to load
		if !response.HasMore {
			break
		}

		// Update the cursor for the next request
		cursor = response.NextCursor
	}
}

// Search for Science Fiction Movies on Netflix and Disney+
func ExampleDefaultAPIService_SearchByFilters_zombieActionMovies() {
	var rapidApiKey, _ = os.LookupEnv("RAPID_API_KEY")

	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	client := streaming.NewAPIClient(configuration).DefaultAPI

	country := "us"
	services := []string{"netflix", "disney"}
	showType := "movie"
	// 28 is Science-Fiction
	// Use Genres endpoint to get the genre ids
	genre := "28"
	maxPages := 2
	keyword := "zombie"

	var cursor string
	for i := 0; i < maxPages; i++ {
		fmt.Printf("Page: %d\n", i)
		response, _, err := client.SearchByFilters(
			context.Background()).
			Country(country).
			Services(strings.Join(services, ",")).
			ShowType(showType).
			Genres(genre).
			Keyword(keyword).
			Cursor(cursor).
			Execute()
		if err != nil {
			log.Fatal(err)
		}
		for _, show := range response.Result {
			for _, streamingOption := range show.StreamingInfo[country] {
				for _, service := range services {
					if streamingOption.GetService() == service {
						fmt.Printf("%s is available on %s: %s\n", show.Title, service, streamingOption.Link)
					}
				}
			}
		}

		// Break out of the loop if there are no more results to load
		if !response.HasMore {
			break
		}

		// Update the cursor for the next request
		cursor = response.NextCursor
	}
}

// Search for Spider-Man Movies
func ExampleDefaultAPIService_SearchByTitle() {
	var rapidApiKey, _ = os.LookupEnv("RAPID_API_KEY")

	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	client := streaming.NewAPIClient(configuration).DefaultAPI

	country := "us"
	showType := "movie"
	title := "Spider-Man"

	res, _, err := client.SearchByTitle(context.Background()).
		Title(title).
		ShowType(showType).
		Country(country).
		Execute()

	if err != nil {
		log.Fatal(err)
	}

	// Print the streaming links for each movie
	for _, show := range res.Result {
		// If no streaming options are available, then skip
		if len(show.StreamingInfo[country]) == 0 {
			continue
		}
		fmt.Printf("---%s---\n", show.Title)
		for _, streamingOption := range show.StreamingInfo[country] {
			fmt.Printf("%s: %s\n", streamingOption.Service, streamingOption.Link)
		}
	}
}

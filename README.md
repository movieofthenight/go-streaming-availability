# Streaming Availability API Go Client

[![go-github release (latest SemVer)](https://img.shields.io/github/v/release/movieofthenight/go-streaming-availability?sort=semver&color=darkgreen)](https://github.com/movieofthenight/go-streaming-availability/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/movieofthenight/go-streaming-availability/v4)

## Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 60 countries!

### API Key

To get an instant free subscription to start using the API, you can visit
[our Authorization guide](https://docs.movieofthenight.com/guide/authorization).

You can get started with a free plan without giving any payment information,
and upgrade to paid plans whenever you like.

### Useful Links

- [Official Webpage of the API](https://www.movieofthenight.com/about/api)

- [API Documentation](https://docs.movieofthenight.com/)

- [Contact Form](https://www.movieofthenight.com/contact)

- [Main GitHub Repository of the API](https://github.com/movieofthenight/streaming-availability-api)

### Features

- Query streaming availability info of the movies and series via their TMDb or IMDd ids.
- Search for movies and series via their titles, genres, keywords, release years on
specific streaming services (e.g.: Get all the zombie action movies available
on Netflix and Disney+)
- Order the search results by titles, release year
or popularity over different time periods
(e.g.: get the all-time most popular movies on Netflix US,
get the most popular series in the last 7 days
on Amazon Prime and Disney+ in the United Kingdom)
- Get the list of upcoming & expiring titles
- Get the daily Top 10 lists
- Returned streaming availability info includes:
  - Deep links into the streaming services for
movies, series, seasons and episodes,
  - Available video qualities (eg. SD, HD, UHD),
  - Available subtitles and audios,
  - First detection time of the shows on the streaming services,
  - Expiry date of the shows/seasons/episodes on the streaming services,
  - All the available options to stream a show
(e.g. via subscription, to buy/rent, for free, available via an addons),
  - Price and currency information for buyable/rentable shows
- Channel and addon support (e.g. Apple TV Channels, Hulu Addons, Prime Video Channels)
- Posters, backdrops, cast & director information, genres, rating and many other details of the shows
- Output also includes TMDB and IMDb ids for every show


## Installation

Run

```shell
go get github.com/movieofthenight/go-streaming-availability/v4
```

## Usage

```go
package main

import "github.com/movieofthenight/go-streaming-availability/v4"

const ApiKey = "PUT_YOUR_API_KEY_HERE"

func main() {
	client := streaming.NewAPIClientFromApiKey(ApiKey, nil)
	// Start using the client
}
```

## Examples

### Get The Godfather's Streaming Availability Info

```go
package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/movieofthenight/go-streaming-availability/v4"
)

func main() {
	const ApiKey = "PUT_YOUR_API_KEY_HERE"

	client := streaming.NewAPIClientFromApiKey(ApiKey, nil)
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
```

> Checkout [example_test.go](https://github.com/movieofthenight/go-streaming-availability/blob/main/example_test.go)
file for the rest of the examples.

## Auto Pagination

This client supports auto-pagination for the paginated endpoints.

If you'd like to use auto-pagination,
you can call the `ExecuteWithAutoPagination` function
instead of the `Execute` function.

An example call without auto pagination:

```go
searchResult, _, err := client.ShowsAPI.SearchShowsByFilters(context.Background()).
	Genres([]string{"comedy"}).
	OrderBy("popularity_1year").
	Country("us").
	Catalogs([]string{"netflix"}).Execute()
```

An example call with auto pagination
that fetches at most 3 pages:

```go
shows, err := client.ShowsAPI.SearchShowsByFilters(context.Background()).
	Genres([]string{"comedy"}).
	OrderBy("popularity_1year").
	Country("us").
	Catalogs([]string{"netflix"}).ExecuteWithAutoPagination(3)
```

Then you can iterate over the results in the following way:

```go
for shows.Next() {
	show := shows.Get()
	// Do something with the show
}
// Check if there was an error during pagination
if shows.Err() != nil {
	log.Fatal(shows.Err())
}
```

## Terms & Conditions and Attribution

While the client libraries have MIT licenses,
the Streaming Availability API itself has further
[Terms & Conditions](https://developers.movieofthenight.com/terms-and-conditions).
Make sure to read it before using the API.

Notably, the API requires an attribution to itself, if the data acquired through
is made public. You can read further about the attribution requirement on the
[Terms & Conditions](https://developers.movieofthenight.com/terms-and-conditions)
page.

## Contact Us

If you have any questions or need further assistance, please don't hesitate to reach us via
[our contact form](https://www.movieofthenight.com/contact).

## FAQ

- **How often the data is updated?**
  - The data is updated daily.

- **I run into an issue. How can I get help?**
  - Send us a message via [our contact form](https://www.movieofthenight.com/contact),
and we will get back to you as soon as possible.

- **API returned me some wrong data. What can I do?**
  - Send us a message with details of your findings.
You can reach ous via [our contact form](https://www.movieofthenight.com/contact).
Once we receive the message we will take a look into the problems and fix the data.

- **I have a request to get a new streaming service supported by the API.**
  - Send us a message via [our contact form](https://www.movieofthenight.com/contact),
  and we will get back to you.

- **I need a client library in another language.**
  - Send us a message via [our contact form](https://www.movieofthenight.com/contact),
  and we will get back to you.

## Client Libraries

[Click here to see the full list of official client libraries](https://docs.movieofthenight.com/guide/client-libraries)

## Services & Countries Supported

We support all the major streaming services across the globe, [click here to see the full list!](https://docs.movieofthenight.com/guide/countries-and-services)


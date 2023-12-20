package streaming_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/movieofthenight/go-streaming-availability/v3"
)

func TestStreaming(t *testing.T) {
	rapidApiKey, rapidApiKeyFound := os.LookupEnv("RAPID_API_KEY")
	if !rapidApiKeyFound {
		t.Fatal("RAPID_API_KEY not found")
	}
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	client := streaming.NewAPIClient(configuration).DefaultAPI
	for testName, testFunc := range testFuncMap {
		t.Run(testName, func(t *testing.T) {
			c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			testFunc(t, client, c)
		})
	}
}

type testFunc func(t *testing.T, client *streaming.DefaultAPIService, c context.Context)

var testFuncMap = map[string]testFunc{
	"countries":       testCountries,
	"genres":          testGenres,
	"getById":         testGyById,
	"searchByFilters": testSearchByFilters,
	"searchByTitle":   testSearchByTitle,
	"changes":         testChanges,
	"leaving":         testLeaving,
}

func testCountries(t *testing.T, client *streaming.DefaultAPIService, c context.Context) {
	validate(t, client.Countries(c).Execute)
}

func testGenres(t *testing.T, client *streaming.DefaultAPIService, c context.Context) {
	validate(t, client.Genres(c).Execute)
}

func testGyById(t *testing.T, client *streaming.DefaultAPIService, c context.Context) {
	validate(t, client.GetById(c).ImdbId("tt0120338").Execute)
}

func testSearchByFilters(t *testing.T, client *streaming.DefaultAPIService, c context.Context) {
	validate(t, client.SearchByFilters(c).Country("us").Services("netflix").Execute)
}

func testSearchByTitle(t *testing.T, client *streaming.DefaultAPIService, c context.Context) {
	validate(t, client.SearchByTitle(c).Country("us").Title("batman").Execute)
}

func testChanges(t *testing.T, client *streaming.DefaultAPIService, c context.Context) {
	validate(t, client.Changes(c).Country("us").Services("netflix").ChangeType("updated").TargetType("show").Execute)
}

func testLeaving(t *testing.T, client *streaming.DefaultAPIService, c context.Context) {
	validate(t, client.Leaving(c).Country("us").Services("netflix").TargetType("show").Execute)
}

func validate[T any](t *testing.T, exec func() (*T, *http.Response, error)) {
	apiResponse, httpResponse, err := exec()
	if err != nil {
		t.Fatal(err)
	}
	if httpResponse == nil {
		t.Fatal("httpResponse is nil")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatal("httpResponse.StatusCode != http.StatusOK")
	}
	if apiResponse == nil {
		t.Fatal("apiResponse is nil")
	}
}

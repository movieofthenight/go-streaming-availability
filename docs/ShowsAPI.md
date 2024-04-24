# \ShowsAPI

All URIs are relative to *https://streaming-availability.p.rapidapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShow**](ShowsAPI.md#GetShow) | **Get** /shows/{id} | Get a Show
[**SearchShowsByFilters**](ShowsAPI.md#SearchShowsByFilters) | **Get** /shows/search/filters | Search Shows by filters
[**SearchShowsByTitle**](ShowsAPI.md#SearchShowsByTitle) | **Get** /shows/search/title | Search Shows by title



## GetShow

> Show GetShow(ctx, id).Country(country).SeriesGranularity(seriesGranularity).OutputLanguage(outputLanguage).Execute()

Get a Show



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/movieofthenight/go-streaming-availability"
)

func main() {
	id := "tt0120338" // string | `id` of the show. You can also pass an IMDb Id or a TMDB Id as an `id`. The endpoint will automatically detect the type of the id and fetch the show accordingly.  When passing an IMDb Id, it should be in the format of `tt<numerical_id>`. (e.g. `tt0120338` for Titanic)  When passing a TMDB Id, it should be in the format of `movie/<numerical_id>` for movies and `tv/<numerical_id>` for series. (e.g. `tv/1396` for Breaking Bad and `movie/597` for Titanic)  If you are handcrafting the URL, make sure to encode the `id` parameter. (e.g. final path should look like `/shows/movie%2F597` for Titanic with TMDb id). Here, `%2F` is the encoded version of `/`. To read more about URL encoding, you can check [this link](https://en.wikipedia.org/wiki/Percent-encoding). 
	country := "ca" // string | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the optional target country. If this parameter is not supplied, global streaming availability across all the countries will be returned. If it is supplied, only the streaming availability info from the given country will be returned. If you are only interested in the streaming availability in a single country, then it is recommended to use this parameter to reduce the size of the response and increase the performance of the endpoint. See `/countries` endpoint to get the list of supported countries.  (optional)
	seriesGranularity := "seriesGranularity_example" // string | `series_granularity` determines the level of detail for series. It does not affect movies.  If `series_granularity` is `show`, then the output will not include season and episode info.  If `series_granularity` is `season`, then the output will include season info but not episode info.  If `series_granularity` is `episode`, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set `series_granularity` as `show` to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set `series_granularity` as `episode`. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc.  (optional) (default to "episode")
	outputLanguage := "outputLanguage_example" // string | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShowsAPI.GetShow(context.Background(), id).Country(country).SeriesGranularity(seriesGranularity).OutputLanguage(outputLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowsAPI.GetShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShow`: Show
	fmt.Fprintf(os.Stdout, "Response from `ShowsAPI.GetShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | &#x60;id&#x60; of the show. You can also pass an IMDb Id or a TMDB Id as an &#x60;id&#x60;. The endpoint will automatically detect the type of the id and fetch the show accordingly.  When passing an IMDb Id, it should be in the format of &#x60;tt&lt;numerical_id&gt;&#x60;. (e.g. &#x60;tt0120338&#x60; for Titanic)  When passing a TMDB Id, it should be in the format of &#x60;movie/&lt;numerical_id&gt;&#x60; for movies and &#x60;tv/&lt;numerical_id&gt;&#x60; for series. (e.g. &#x60;tv/1396&#x60; for Breaking Bad and &#x60;movie/597&#x60; for Titanic)  If you are handcrafting the URL, make sure to encode the &#x60;id&#x60; parameter. (e.g. final path should look like &#x60;/shows/movie%2F597&#x60; for Titanic with TMDb id). Here, &#x60;%2F&#x60; is the encoded version of &#x60;/&#x60;. To read more about URL encoding, you can check [this link](https://en.wikipedia.org/wiki/Percent-encoding).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the optional target country. If this parameter is not supplied, global streaming availability across all the countries will be returned. If it is supplied, only the streaming availability info from the given country will be returned. If you are only interested in the streaming availability in a single country, then it is recommended to use this parameter to reduce the size of the response and increase the performance of the endpoint. See &#x60;/countries&#x60; endpoint to get the list of supported countries.  | 
 **seriesGranularity** | **string** | &#x60;series_granularity&#x60; determines the level of detail for series. It does not affect movies.  If &#x60;series_granularity&#x60; is &#x60;show&#x60;, then the output will not include season and episode info.  If &#x60;series_granularity&#x60; is &#x60;season&#x60;, then the output will include season info but not episode info.  If &#x60;series_granularity&#x60; is &#x60;episode&#x60;, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set &#x60;series_granularity&#x60; as &#x60;show&#x60; to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set &#x60;series_granularity&#x60; as &#x60;episode&#x60;. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc.  | [default to &quot;episode&quot;]
 **outputLanguage** | **string** | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  | [default to &quot;en&quot;]

### Return type

[**Show**](Show.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchShowsByFilters

> SearchResult SearchShowsByFilters(ctx).Country(country).Catalogs(catalogs).OutputLanguage(outputLanguage).ShowType(showType).Genres(genres).GenresRelation(genresRelation).ShowOriginalLanguage(showOriginalLanguage).YearMin(yearMin).YearMax(yearMax).RatingMin(ratingMin).RatingMax(ratingMax).Keyword(keyword).SeriesGranularity(seriesGranularity).OrderBy(orderBy).DescendingOrder(descendingOrder).Cursor(cursor).Execute()

Search Shows by filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/movieofthenight/go-streaming-availability"
)

func main() {
	country := "ca" // string | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See `/countries` endpoint to get the list of supported countries. 
	catalogs := []string{"Inner_example"} // []string | A comma separated list of up to 32 catalogs to search in. See `/countries` endpoint to get the supported services in each country and their ids.  When multiple catalogs are passed as a comma separated list, any show that is in at least one of the catalogs will be included in the result.  If no catalogs are passed, the endpoint will search in all the available catalogs in the country.  Syntax of the catalogs supplied in the list can be as the followings:  - `<sevice_id>`: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons e.g. `netflix`, `prime`, `apple`  - `<sevice_id>.<streaming_option_type>`: Only returns the shows that are available in that service with the given streaming option type. Valid streaming option types are `subscription`, `free`, `rent`, `buy` and `addon` e.g. `peacock.free` only returns the shows on Peacock that are free to watch, `prime.subscription` only returns the shows on Prime Video that are available to watch with a Prime subscription. `hulu.addon` only returns the shows on Hulu that are available via an addon, `prime.rent` only returns the shows on Prime Video that are rentable.  - `<sevice_id>.addon.<addon_id>`: Only returns the shows that are available in that service with the given addon. Check `/countries` endpoint to fetch the available addons for a service in each country. Some sample values are: `hulu.addon.hbo`, `prime.addon.hbomaxus`.  (optional)
	outputLanguage := "outputLanguage_example" // string | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  (optional) (default to "en")
	showType := openapiclient.showType("movie") // ShowType | Type of shows to search in. If not supplied, both movies and series will be included in the search results.  (optional)
	genres := []string{"Inner_example"} // []string | A comma seperated list of genre ids to only search within the shows in those genre. See `/genres` endpoint to see the available genres and their ids. Use `genres_relation` parameter to specify between returning shows that have at least one of the given genres or returning shows that have all of the given genres.  (optional)
	genresRelation := "genresRelation_example" // string | Only used when there are multiple genres supplied in `genres` parameter.  When `or`, the endpoint returns any show that has at least one of the given genres. When `and`, it only returns the shows that have all of the given genres.  (optional) (default to "and")
	showOriginalLanguage := "en" // string | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code to only search within the shows whose original language matches with the provided language.  (optional)
	yearMin := int32(1980.0) // int32 | Minimum release/air year of the shows. (optional)
	yearMax := int32(1980.0) // int32 | Maximum release/air year of the shows. (optional)
	ratingMin := int32(75.0) // int32 | Minimum rating of the shows. (optional)
	ratingMax := int32(80.0) // int32 | Maximum rating of the shows. (optional)
	keyword := "zombie" // string | A keyword to only search within the shows have that keyword in their overview or title. (optional)
	seriesGranularity := "seriesGranularity_example" // string | `series_granularity` determines the level of detail for series. It does not affect movies.  If `series_granularity` is `show`, then the output will not include season and episode info.  If `series_granularity` is `season`, then the output will include season info but not episode info.  If `series_granularity` is `episode`, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set `series_granularity` as `show` to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set `series_granularity` as `episode`. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc.  (optional) (default to "show")
	orderBy := "orderBy_example" // string | Determines the ordering of the shows. Make sure to set `descending_order` parameter as `true` when ordering by popularity or rating so that shows with the highest popularity or rating will be returned first.  (optional) (default to "original_title")
	descendingOrder := true // bool | The results are ordered in descending order if set true. (optional) (default to false)
	cursor := "cursor_example" // string | Cursor is used for pagination. After each request, the response includes a `hasMore` boolean field to tell if there are more results that did not fit into the returned list. If it is set as `true`, to get the rest of the result set, send a new request (with the same parameters for other fields), and set the `cursor` parameter as the `nextCursor` value of the response of the previous request. Do not forget to escape the `cursor` value before putting it into a query as it might contain characters such as `?`and `&`.  The first request naturally does not require a `cursor` parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShowsAPI.SearchShowsByFilters(context.Background()).Country(country).Catalogs(catalogs).OutputLanguage(outputLanguage).ShowType(showType).Genres(genres).GenresRelation(genresRelation).ShowOriginalLanguage(showOriginalLanguage).YearMin(yearMin).YearMax(yearMax).RatingMin(ratingMin).RatingMax(ratingMax).Keyword(keyword).SeriesGranularity(seriesGranularity).OrderBy(orderBy).DescendingOrder(descendingOrder).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowsAPI.SearchShowsByFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchShowsByFilters`: SearchResult
	fmt.Fprintf(os.Stdout, "Response from `ShowsAPI.SearchShowsByFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchShowsByFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See &#x60;/countries&#x60; endpoint to get the list of supported countries.  | 
 **catalogs** | **[]string** | A comma separated list of up to 32 catalogs to search in. See &#x60;/countries&#x60; endpoint to get the supported services in each country and their ids.  When multiple catalogs are passed as a comma separated list, any show that is in at least one of the catalogs will be included in the result.  If no catalogs are passed, the endpoint will search in all the available catalogs in the country.  Syntax of the catalogs supplied in the list can be as the followings:  - &#x60;&lt;sevice_id&gt;&#x60;: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons e.g. &#x60;netflix&#x60;, &#x60;prime&#x60;, &#x60;apple&#x60;  - &#x60;&lt;sevice_id&gt;.&lt;streaming_option_type&gt;&#x60;: Only returns the shows that are available in that service with the given streaming option type. Valid streaming option types are &#x60;subscription&#x60;, &#x60;free&#x60;, &#x60;rent&#x60;, &#x60;buy&#x60; and &#x60;addon&#x60; e.g. &#x60;peacock.free&#x60; only returns the shows on Peacock that are free to watch, &#x60;prime.subscription&#x60; only returns the shows on Prime Video that are available to watch with a Prime subscription. &#x60;hulu.addon&#x60; only returns the shows on Hulu that are available via an addon, &#x60;prime.rent&#x60; only returns the shows on Prime Video that are rentable.  - &#x60;&lt;sevice_id&gt;.addon.&lt;addon_id&gt;&#x60;: Only returns the shows that are available in that service with the given addon. Check &#x60;/countries&#x60; endpoint to fetch the available addons for a service in each country. Some sample values are: &#x60;hulu.addon.hbo&#x60;, &#x60;prime.addon.hbomaxus&#x60;.  | 
 **outputLanguage** | **string** | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  | [default to &quot;en&quot;]
 **showType** | [**ShowType**](ShowType.md) | Type of shows to search in. If not supplied, both movies and series will be included in the search results.  | 
 **genres** | **[]string** | A comma seperated list of genre ids to only search within the shows in those genre. See &#x60;/genres&#x60; endpoint to see the available genres and their ids. Use &#x60;genres_relation&#x60; parameter to specify between returning shows that have at least one of the given genres or returning shows that have all of the given genres.  | 
 **genresRelation** | **string** | Only used when there are multiple genres supplied in &#x60;genres&#x60; parameter.  When &#x60;or&#x60;, the endpoint returns any show that has at least one of the given genres. When &#x60;and&#x60;, it only returns the shows that have all of the given genres.  | [default to &quot;and&quot;]
 **showOriginalLanguage** | **string** | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code to only search within the shows whose original language matches with the provided language.  | 
 **yearMin** | **int32** | Minimum release/air year of the shows. | 
 **yearMax** | **int32** | Maximum release/air year of the shows. | 
 **ratingMin** | **int32** | Minimum rating of the shows. | 
 **ratingMax** | **int32** | Maximum rating of the shows. | 
 **keyword** | **string** | A keyword to only search within the shows have that keyword in their overview or title. | 
 **seriesGranularity** | **string** | &#x60;series_granularity&#x60; determines the level of detail for series. It does not affect movies.  If &#x60;series_granularity&#x60; is &#x60;show&#x60;, then the output will not include season and episode info.  If &#x60;series_granularity&#x60; is &#x60;season&#x60;, then the output will include season info but not episode info.  If &#x60;series_granularity&#x60; is &#x60;episode&#x60;, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set &#x60;series_granularity&#x60; as &#x60;show&#x60; to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set &#x60;series_granularity&#x60; as &#x60;episode&#x60;. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc.  | [default to &quot;show&quot;]
 **orderBy** | **string** | Determines the ordering of the shows. Make sure to set &#x60;descending_order&#x60; parameter as &#x60;true&#x60; when ordering by popularity or rating so that shows with the highest popularity or rating will be returned first.  | [default to &quot;original_title&quot;]
 **descendingOrder** | **bool** | The results are ordered in descending order if set true. | [default to false]
 **cursor** | **string** | Cursor is used for pagination. After each request, the response includes a &#x60;hasMore&#x60; boolean field to tell if there are more results that did not fit into the returned list. If it is set as &#x60;true&#x60;, to get the rest of the result set, send a new request (with the same parameters for other fields), and set the &#x60;cursor&#x60; parameter as the &#x60;nextCursor&#x60; value of the response of the previous request. Do not forget to escape the &#x60;cursor&#x60; value before putting it into a query as it might contain characters such as &#x60;?&#x60;and &#x60;&amp;&#x60;.  The first request naturally does not require a &#x60;cursor&#x60; parameter.  | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchShowsByTitle

> []Show SearchShowsByTitle(ctx).Title(title).Country(country).ShowType(showType).SeriesGranularity(seriesGranularity).OutputLanguage(outputLanguage).Execute()

Search Shows by title



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/movieofthenight/go-streaming-availability"
)

func main() {
	title := "Batman" // string | Title phrase to search for.
	country := "ca" // string | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See `/countries` endpoint to get the list of supported countries. 
	showType := openapiclient.showType("movie") // ShowType | Type of shows to search in. If not supplied, both movies and series will be included in the search results.  (optional)
	seriesGranularity := "seriesGranularity_example" // string | `series_granularity` determines the level of detail for series. It does not affect movies.  If `series_granularity` is `show`, then the output will not include season and episode info.  If `series_granularity` is `season`, then the output will include season info but not episode info.  If `series_granularity` is `episode`, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set `series_granularity` as `show` to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set `series_granularity` as `episode`. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc.  (optional) (default to "show")
	outputLanguage := "outputLanguage_example" // string | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShowsAPI.SearchShowsByTitle(context.Background()).Title(title).Country(country).ShowType(showType).SeriesGranularity(seriesGranularity).OutputLanguage(outputLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowsAPI.SearchShowsByTitle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchShowsByTitle`: []Show
	fmt.Fprintf(os.Stdout, "Response from `ShowsAPI.SearchShowsByTitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchShowsByTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | Title phrase to search for. | 
 **country** | **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See &#x60;/countries&#x60; endpoint to get the list of supported countries.  | 
 **showType** | [**ShowType**](ShowType.md) | Type of shows to search in. If not supplied, both movies and series will be included in the search results.  | 
 **seriesGranularity** | **string** | &#x60;series_granularity&#x60; determines the level of detail for series. It does not affect movies.  If &#x60;series_granularity&#x60; is &#x60;show&#x60;, then the output will not include season and episode info.  If &#x60;series_granularity&#x60; is &#x60;season&#x60;, then the output will include season info but not episode info.  If &#x60;series_granularity&#x60; is &#x60;episode&#x60;, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set &#x60;series_granularity&#x60; as &#x60;show&#x60; to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set &#x60;series_granularity&#x60; as &#x60;episode&#x60;. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc.  | [default to &quot;show&quot;]
 **outputLanguage** | **string** | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  | [default to &quot;en&quot;]

### Return type

[**[]Show**](Show.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


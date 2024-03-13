# \DefaultAPI

All URIs are relative to *https://streaming-availability.p.rapidapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Changes**](DefaultAPI.md#Changes) | **Get** /changes | Changes
[**Countries**](DefaultAPI.md#Countries) | **Get** /countries | Countries
[**Genres**](DefaultAPI.md#Genres) | **Get** /genres | Genres
[**GetById**](DefaultAPI.md#GetById) | **Get** /get | Get by Id
[**Leaving**](DefaultAPI.md#Leaving) | **Get** /leaving | Leaving
[**SearchByFilters**](DefaultAPI.md#SearchByFilters) | **Get** /search/filters | Search by Filters
[**SearchByTitle**](DefaultAPI.md#SearchByTitle) | **Get** /search/title | Search by Title
[**Services**](DefaultAPI.md#Services) | **Get** /services | Services



## Changes

> ChangesResponseSchema Changes(ctx).Country(country).Services(services).ChangeType(changeType).TargetType(targetType).Since(since).Cursor(cursor).Desc(desc).OutputLanguage(outputLanguage).Execute()

Changes



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
    country := "ca" // string | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \"/countries\" endpoint to get the list of supported countries. 
    services := "netflix" // string | A comma separated list of up to services to search in. See \"/countries\" endpoint to get the supported services in each country and their ids/names. Maximum amount of services allowed depends on the endpoint. If a service supports both \"free\" and \"subscription\", then results included under \"subscription\" will always include the \"free\" shows as well. When multiple values are passed as a comma separated list, any show that satisfies at least one of the values will be included in the result. Syntax of the values supplied in the list can be as the followings: - \"<sevice_id>\": Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons i.e. \"netflix\", \"prime\", \"apple\" - \"<sevice_id>.<streaming_type>\": Only returns the shows that are available in that service with the given streaming type. Valid streaming type values are \"subscription\", \"free\", \"rent\", \"buy\" and \"addon\" i.e. \"peacock.free\" only returns the shows on Peacock that are free to watch, \"prime.subscription\" only returns the shows on Prime Video that are available to watch with a Prime subscription. \"hulu.addon\" only returns the shows on Hulu that are available via an addon, \"prime.rent\" only returns the shows on Prime Video that are rentable. - \"<sevice_id>.addon.<addon_id>\": Only returns the shows that are available in that service with the given addon. Check \"/countries\" endpoint to fetch the available addons for a service in each country. Some sample values are: \"hulu.addon.hbo\", \"prime.addon.hbomaxus\". 
    changeType := "new" // string | Type of change to query.
    targetType := "show" // string | Type of items to search in.
    since := int32(1672531200) // int32 | [Unix Time Stamp](https://www.unixtimestamp.com/) to only query the changes since then. Must be within the past \"31\" days. If not supplied, the changes withing the past \"15\" days are returned.  (optional)
    cursor := "cursor_example" // string | Cursor is used for pagination. After each request, the response includes a \"hasMore\" boolean field to tell if there are more results that did not fit into the returned list. If it is set as \"true\", to get the rest of the result set, send a new request (with the same parameters for other fields such as \"services\" etc.), and set the \"cursor\" parameter as the \"nextCursor\" value of the previous request response. Do not forget to escape the \"cursor\" value before putting it into a query as it might contain characters such as \"?\"and \"&\". The first request naturally does not require a \"cursor\" parameter.  (optional)
    desc := true // bool | The results are ordered in descending order if set true. (optional) (default to false)
    outputLanguage := "en" // string | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \"title\" will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \"output_language\" is \"es\", then the \"keyword\" will be treated as a Spanish word).  (optional) (default to "en")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.Changes(context.Background()).Country(country).Services(services).ChangeType(changeType).TargetType(targetType).Since(since).Cursor(cursor).Desc(desc).OutputLanguage(outputLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Changes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Changes`: ChangesResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Changes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \&quot;/countries\&quot; endpoint to get the list of supported countries.  | 
 **services** | **string** | A comma separated list of up to services to search in. See \&quot;/countries\&quot; endpoint to get the supported services in each country and their ids/names. Maximum amount of services allowed depends on the endpoint. If a service supports both \&quot;free\&quot; and \&quot;subscription\&quot;, then results included under \&quot;subscription\&quot; will always include the \&quot;free\&quot; shows as well. When multiple values are passed as a comma separated list, any show that satisfies at least one of the values will be included in the result. Syntax of the values supplied in the list can be as the followings: - \&quot;&lt;sevice_id&gt;\&quot;: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons i.e. \&quot;netflix\&quot;, \&quot;prime\&quot;, \&quot;apple\&quot; - \&quot;&lt;sevice_id&gt;.&lt;streaming_type&gt;\&quot;: Only returns the shows that are available in that service with the given streaming type. Valid streaming type values are \&quot;subscription\&quot;, \&quot;free\&quot;, \&quot;rent\&quot;, \&quot;buy\&quot; and \&quot;addon\&quot; i.e. \&quot;peacock.free\&quot; only returns the shows on Peacock that are free to watch, \&quot;prime.subscription\&quot; only returns the shows on Prime Video that are available to watch with a Prime subscription. \&quot;hulu.addon\&quot; only returns the shows on Hulu that are available via an addon, \&quot;prime.rent\&quot; only returns the shows on Prime Video that are rentable. - \&quot;&lt;sevice_id&gt;.addon.&lt;addon_id&gt;\&quot;: Only returns the shows that are available in that service with the given addon. Check \&quot;/countries\&quot; endpoint to fetch the available addons for a service in each country. Some sample values are: \&quot;hulu.addon.hbo\&quot;, \&quot;prime.addon.hbomaxus\&quot;.  | 
 **changeType** | **string** | Type of change to query. | 
 **targetType** | **string** | Type of items to search in. | 
 **since** | **int32** | [Unix Time Stamp](https://www.unixtimestamp.com/) to only query the changes since then. Must be within the past \&quot;31\&quot; days. If not supplied, the changes withing the past \&quot;15\&quot; days are returned.  | 
 **cursor** | **string** | Cursor is used for pagination. After each request, the response includes a \&quot;hasMore\&quot; boolean field to tell if there are more results that did not fit into the returned list. If it is set as \&quot;true\&quot;, to get the rest of the result set, send a new request (with the same parameters for other fields such as \&quot;services\&quot; etc.), and set the \&quot;cursor\&quot; parameter as the \&quot;nextCursor\&quot; value of the previous request response. Do not forget to escape the \&quot;cursor\&quot; value before putting it into a query as it might contain characters such as \&quot;?\&quot;and \&quot;&amp;\&quot;. The first request naturally does not require a \&quot;cursor\&quot; parameter.  | 
 **desc** | **bool** | The results are ordered in descending order if set true. | [default to false]
 **outputLanguage** | **string** | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \&quot;title\&quot; will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \&quot;output_language\&quot; is \&quot;es\&quot;, then the \&quot;keyword\&quot; will be treated as a Spanish word).  | [default to &quot;en&quot;]

### Return type

[**ChangesResponseSchema**](ChangesResponseSchema.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Countries

> CountriesResponseSchema Countries(ctx).Execute()

Countries



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.Countries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Countries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Countries`: CountriesResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Countries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCountriesRequest struct via the builder pattern


### Return type

[**CountriesResponseSchema**](CountriesResponseSchema.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Genres

> GenresResponseSchema Genres(ctx).Execute()

Genres



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.Genres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Genres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Genres`: GenresResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Genres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenresRequest struct via the builder pattern


### Return type

[**GenresResponseSchema**](GenresResponseSchema.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetById

> GetResponseSchema GetById(ctx).ImdbId(imdbId).TmdbId(tmdbId).SeriesGranularity(seriesGranularity).OutputLanguage(outputLanguage).Execute()

Get by Id



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
    imdbId := "tt0120338" // string | IMDb ID of the target show. Either this or \"tmdb_id\" parameter must be supplied. (optional)
    tmdbId := "movie/597" // string | TMDb ID of the target show. Either this or \"imdb_id\" parameter must be supplied (optional)
    seriesGranularity := "show" // string | \"series_granularity\" determines the level of detail for series. It does not affect movies. If \"series_granularity\" is \"show\", then the output will not include season and episode info. If \"series_granularity\" is \"season\", then the output will include season info but not episode info. If \"series_granularity\" is \"episode\", then the output will include season and episode info. If you do not need season and episode info, then it is recommended to set \"series_granularity\" as \"show\" to reduce the size of the response and increase the performance. Depending on the query (i.e. if the result set includes series with high amount of seasons/episodes), this can reduce the response size and the response time drastically. If you need deep links for individual seasons and episodes, then you should set \"series_granularity\" as \"episode\". In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. If \"series_granularity\" is not supplied, then it is set as \"episode\" by default.  (optional) (default to "episode")
    outputLanguage := "en" // string | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \"title\" will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \"output_language\" is \"es\", then the \"keyword\" will be treated as a Spanish word).  (optional) (default to "en")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetById(context.Background()).ImdbId(imdbId).TmdbId(tmdbId).SeriesGranularity(seriesGranularity).OutputLanguage(outputLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetById`: GetResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imdbId** | **string** | IMDb ID of the target show. Either this or \&quot;tmdb_id\&quot; parameter must be supplied. | 
 **tmdbId** | **string** | TMDb ID of the target show. Either this or \&quot;imdb_id\&quot; parameter must be supplied | 
 **seriesGranularity** | **string** | \&quot;series_granularity\&quot; determines the level of detail for series. It does not affect movies. If \&quot;series_granularity\&quot; is \&quot;show\&quot;, then the output will not include season and episode info. If \&quot;series_granularity\&quot; is \&quot;season\&quot;, then the output will include season info but not episode info. If \&quot;series_granularity\&quot; is \&quot;episode\&quot;, then the output will include season and episode info. If you do not need season and episode info, then it is recommended to set \&quot;series_granularity\&quot; as \&quot;show\&quot; to reduce the size of the response and increase the performance. Depending on the query (i.e. if the result set includes series with high amount of seasons/episodes), this can reduce the response size and the response time drastically. If you need deep links for individual seasons and episodes, then you should set \&quot;series_granularity\&quot; as \&quot;episode\&quot;. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. If \&quot;series_granularity\&quot; is not supplied, then it is set as \&quot;episode\&quot; by default.  | [default to &quot;episode&quot;]
 **outputLanguage** | **string** | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \&quot;title\&quot; will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \&quot;output_language\&quot; is \&quot;es\&quot;, then the \&quot;keyword\&quot; will be treated as a Spanish word).  | [default to &quot;en&quot;]

### Return type

[**GetResponseSchema**](GetResponseSchema.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Leaving

> UpcomingChangesResponseSchema Leaving(ctx).Country(country).Services(services).TargetType(targetType).Cursor(cursor).OutputLanguage(outputLanguage).Execute()

Leaving



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
    country := "ca" // string | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \"/countries\" endpoint to get the list of supported countries. 
    services := "netflix" // string | A comma separated list of up to services to search in. See \"/countries\" endpoint to get the supported services in each country and their ids/names. Maximum amount of services allowed depends on the endpoint. If a service supports both \"free\" and \"subscription\", then results included under \"subscription\" will always include the \"free\" shows as well. When multiple values are passed as a comma separated list, any show that satisfies at least one of the values will be included in the result. Syntax of the values supplied in the list can be as the followings: - \"<sevice_id>\": Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons i.e. \"netflix\", \"prime\", \"apple\" - \"<sevice_id>.<streaming_type>\": Only returns the shows that are available in that service with the given streaming type. Valid streaming type values are \"subscription\", \"free\", \"rent\", \"buy\" and \"addon\" i.e. \"peacock.free\" only returns the shows on Peacock that are free to watch, \"prime.subscription\" only returns the shows on Prime Video that are available to watch with a Prime subscription. \"hulu.addon\" only returns the shows on Hulu that are available via an addon, \"prime.rent\" only returns the shows on Prime Video that are rentable. - \"<sevice_id>.addon.<addon_id>\": Only returns the shows that are available in that service with the given addon. Check \"/countries\" endpoint to fetch the available addons for a service in each country. Some sample values are: \"hulu.addon.hbo\", \"prime.addon.hbomaxus\". 
    targetType := "show" // string | Type of items to search in.
    cursor := "cursor_example" // string | Cursor is used for pagination. After each request, the response includes a \"hasMore\" boolean field to tell if there are more results that did not fit into the returned list. If it is set as \"true\", to get the rest of the result set, send a new request (with the same parameters for other fields such as \"services\" etc.), and set the \"cursor\" parameter as the \"nextCursor\" value of the previous request response. Do not forget to escape the \"cursor\" value before putting it into a query as it might contain characters such as \"?\"and \"&\". The first request naturally does not require a \"cursor\" parameter.  (optional)
    outputLanguage := "en" // string | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \"title\" will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \"output_language\" is \"es\", then the \"keyword\" will be treated as a Spanish word).  (optional) (default to "en")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.Leaving(context.Background()).Country(country).Services(services).TargetType(targetType).Cursor(cursor).OutputLanguage(outputLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Leaving``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Leaving`: UpcomingChangesResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Leaving`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLeavingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \&quot;/countries\&quot; endpoint to get the list of supported countries.  | 
 **services** | **string** | A comma separated list of up to services to search in. See \&quot;/countries\&quot; endpoint to get the supported services in each country and their ids/names. Maximum amount of services allowed depends on the endpoint. If a service supports both \&quot;free\&quot; and \&quot;subscription\&quot;, then results included under \&quot;subscription\&quot; will always include the \&quot;free\&quot; shows as well. When multiple values are passed as a comma separated list, any show that satisfies at least one of the values will be included in the result. Syntax of the values supplied in the list can be as the followings: - \&quot;&lt;sevice_id&gt;\&quot;: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons i.e. \&quot;netflix\&quot;, \&quot;prime\&quot;, \&quot;apple\&quot; - \&quot;&lt;sevice_id&gt;.&lt;streaming_type&gt;\&quot;: Only returns the shows that are available in that service with the given streaming type. Valid streaming type values are \&quot;subscription\&quot;, \&quot;free\&quot;, \&quot;rent\&quot;, \&quot;buy\&quot; and \&quot;addon\&quot; i.e. \&quot;peacock.free\&quot; only returns the shows on Peacock that are free to watch, \&quot;prime.subscription\&quot; only returns the shows on Prime Video that are available to watch with a Prime subscription. \&quot;hulu.addon\&quot; only returns the shows on Hulu that are available via an addon, \&quot;prime.rent\&quot; only returns the shows on Prime Video that are rentable. - \&quot;&lt;sevice_id&gt;.addon.&lt;addon_id&gt;\&quot;: Only returns the shows that are available in that service with the given addon. Check \&quot;/countries\&quot; endpoint to fetch the available addons for a service in each country. Some sample values are: \&quot;hulu.addon.hbo\&quot;, \&quot;prime.addon.hbomaxus\&quot;.  | 
 **targetType** | **string** | Type of items to search in. | 
 **cursor** | **string** | Cursor is used for pagination. After each request, the response includes a \&quot;hasMore\&quot; boolean field to tell if there are more results that did not fit into the returned list. If it is set as \&quot;true\&quot;, to get the rest of the result set, send a new request (with the same parameters for other fields such as \&quot;services\&quot; etc.), and set the \&quot;cursor\&quot; parameter as the \&quot;nextCursor\&quot; value of the previous request response. Do not forget to escape the \&quot;cursor\&quot; value before putting it into a query as it might contain characters such as \&quot;?\&quot;and \&quot;&amp;\&quot;. The first request naturally does not require a \&quot;cursor\&quot; parameter.  | 
 **outputLanguage** | **string** | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \&quot;title\&quot; will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \&quot;output_language\&quot; is \&quot;es\&quot;, then the \&quot;keyword\&quot; will be treated as a Spanish word).  | [default to &quot;en&quot;]

### Return type

[**UpcomingChangesResponseSchema**](UpcomingChangesResponseSchema.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchByFilters

> SearchFiltersResponseSchema SearchByFilters(ctx).Country(country).Services(services).OutputLanguage(outputLanguage).ShowType(showType).SeriesGranularity(seriesGranularity).Genres(genres).GenresRelation(genresRelation).ShowOriginalLanguage(showOriginalLanguage).YearMin(yearMin).YearMax(yearMax).Keyword(keyword).OrderBy(orderBy).Desc(desc).Cursor(cursor).Execute()

Search by Filters



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
    country := "ca" // string | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \"/countries\" endpoint to get the list of supported countries. 
    services := "netflix" // string | A comma separated list of up to services to search in. See \"/countries\" endpoint to get the supported services in each country and their ids/names. Maximum amount of services allowed depends on the endpoint. If a service supports both \"free\" and \"subscription\", then results included under \"subscription\" will always include the \"free\" shows as well. When multiple values are passed as a comma separated list, any show that satisfies at least one of the values will be included in the result. Syntax of the values supplied in the list can be as the followings: - \"<sevice_id>\": Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons i.e. \"netflix\", \"prime\", \"apple\" - \"<sevice_id>.<streaming_type>\": Only returns the shows that are available in that service with the given streaming type. Valid streaming type values are \"subscription\", \"free\", \"rent\", \"buy\" and \"addon\" i.e. \"peacock.free\" only returns the shows on Peacock that are free to watch, \"prime.subscription\" only returns the shows on Prime Video that are available to watch with a Prime subscription. \"hulu.addon\" only returns the shows on Hulu that are available via an addon, \"prime.rent\" only returns the shows on Prime Video that are rentable. - \"<sevice_id>.addon.<addon_id>\": Only returns the shows that are available in that service with the given addon. Check \"/countries\" endpoint to fetch the available addons for a service in each country. Some sample values are: \"hulu.addon.hbo\", \"prime.addon.hbomaxus\". 
    outputLanguage := "en" // string | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \"title\" will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \"output_language\" is \"es\", then the \"keyword\" will be treated as a Spanish word).  (optional) (default to "en")
    showType := "movie" // string | Type of shows to search in. (optional) (default to "all")
    seriesGranularity := "show" // string | \"series_granularity\" determines the level of detail for series. It does not affect movies. If \"series_granularity\" is \"show\", then the output will not include season and episode info. If \"series_granularity\" is \"season\", then the output will include season info but not episode info. If \"series_granularity\" is \"episode\", then the output will include season and episode info. If you do not need season and episode info, then it is recommended to set \"series_granularity\" as \"show\" to reduce the size of the response and increase the performance. Depending on the query (i.e. if the result set includes series with high amount of seasons/episodes), this can reduce the response size and the response time drastically. If you need deep links for individual seasons and episodes, then you should set \"series_granularity\" as \"episode\". In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. If \"series_granularity\" is not supplied, then it is set as \"episode\" by default.  (optional) (default to "episode")
    genres := "10749,35" // string | A comma seperated list of genre ids to only search within the shows in those genre. See \"/genres\" endpoint to see the available genres and their ids. Use \"genres_relation\" parameter to specify between returning shows that have at least one of the given genres or returning shows that have all of the given genres.  (optional)
    genresRelation := "and" // string | Only used when there are multiple genres supplied in \"genres\" parameter. When \"or\", the endpoint returns any show that has at least one of the given genres. When \"and\", it only returns the shows that have all of the given genres.  (optional) (default to "and")
    showOriginalLanguage := "en" // string | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code to only search within the shows whose original language matches with the provided language.  (optional)
    yearMin := int32(1980.0) // int32 | Minimum release/air year of the shows. (optional)
    yearMax := int32(1980.0) // int32 | Maximum release/air year of the shows. (optional)
    keyword := "zombie" // string | A keyword to only search within the shows have that keyword in their overview or title. (optional)
    orderBy := "original_title" // string | Determines the ordering of the results. (optional) (default to "original_title")
    desc := true // bool | The results are ordered in descending order if set true. (optional) (default to false)
    cursor := "cursor_example" // string | Cursor is used for pagination. After each request, the response includes a \"hasMore\" boolean field to tell if there are more results that did not fit into the returned list. If it is set as \"true\", to get the rest of the result set, send a new request (with the same parameters for other fields such as \"services\" etc.), and set the \"cursor\" parameter as the \"nextCursor\" value of the previous request response. Do not forget to escape the \"cursor\" value before putting it into a query as it might contain characters such as \"?\"and \"&\". The first request naturally does not require a \"cursor\" parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SearchByFilters(context.Background()).Country(country).Services(services).OutputLanguage(outputLanguage).ShowType(showType).SeriesGranularity(seriesGranularity).Genres(genres).GenresRelation(genresRelation).ShowOriginalLanguage(showOriginalLanguage).YearMin(yearMin).YearMax(yearMax).Keyword(keyword).OrderBy(orderBy).Desc(desc).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SearchByFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchByFilters`: SearchFiltersResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SearchByFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchByFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \&quot;/countries\&quot; endpoint to get the list of supported countries.  | 
 **services** | **string** | A comma separated list of up to services to search in. See \&quot;/countries\&quot; endpoint to get the supported services in each country and their ids/names. Maximum amount of services allowed depends on the endpoint. If a service supports both \&quot;free\&quot; and \&quot;subscription\&quot;, then results included under \&quot;subscription\&quot; will always include the \&quot;free\&quot; shows as well. When multiple values are passed as a comma separated list, any show that satisfies at least one of the values will be included in the result. Syntax of the values supplied in the list can be as the followings: - \&quot;&lt;sevice_id&gt;\&quot;: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons i.e. \&quot;netflix\&quot;, \&quot;prime\&quot;, \&quot;apple\&quot; - \&quot;&lt;sevice_id&gt;.&lt;streaming_type&gt;\&quot;: Only returns the shows that are available in that service with the given streaming type. Valid streaming type values are \&quot;subscription\&quot;, \&quot;free\&quot;, \&quot;rent\&quot;, \&quot;buy\&quot; and \&quot;addon\&quot; i.e. \&quot;peacock.free\&quot; only returns the shows on Peacock that are free to watch, \&quot;prime.subscription\&quot; only returns the shows on Prime Video that are available to watch with a Prime subscription. \&quot;hulu.addon\&quot; only returns the shows on Hulu that are available via an addon, \&quot;prime.rent\&quot; only returns the shows on Prime Video that are rentable. - \&quot;&lt;sevice_id&gt;.addon.&lt;addon_id&gt;\&quot;: Only returns the shows that are available in that service with the given addon. Check \&quot;/countries\&quot; endpoint to fetch the available addons for a service in each country. Some sample values are: \&quot;hulu.addon.hbo\&quot;, \&quot;prime.addon.hbomaxus\&quot;.  | 
 **outputLanguage** | **string** | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \&quot;title\&quot; will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \&quot;output_language\&quot; is \&quot;es\&quot;, then the \&quot;keyword\&quot; will be treated as a Spanish word).  | [default to &quot;en&quot;]
 **showType** | **string** | Type of shows to search in. | [default to &quot;all&quot;]
 **seriesGranularity** | **string** | \&quot;series_granularity\&quot; determines the level of detail for series. It does not affect movies. If \&quot;series_granularity\&quot; is \&quot;show\&quot;, then the output will not include season and episode info. If \&quot;series_granularity\&quot; is \&quot;season\&quot;, then the output will include season info but not episode info. If \&quot;series_granularity\&quot; is \&quot;episode\&quot;, then the output will include season and episode info. If you do not need season and episode info, then it is recommended to set \&quot;series_granularity\&quot; as \&quot;show\&quot; to reduce the size of the response and increase the performance. Depending on the query (i.e. if the result set includes series with high amount of seasons/episodes), this can reduce the response size and the response time drastically. If you need deep links for individual seasons and episodes, then you should set \&quot;series_granularity\&quot; as \&quot;episode\&quot;. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. If \&quot;series_granularity\&quot; is not supplied, then it is set as \&quot;episode\&quot; by default.  | [default to &quot;episode&quot;]
 **genres** | **string** | A comma seperated list of genre ids to only search within the shows in those genre. See \&quot;/genres\&quot; endpoint to see the available genres and their ids. Use \&quot;genres_relation\&quot; parameter to specify between returning shows that have at least one of the given genres or returning shows that have all of the given genres.  | 
 **genresRelation** | **string** | Only used when there are multiple genres supplied in \&quot;genres\&quot; parameter. When \&quot;or\&quot;, the endpoint returns any show that has at least one of the given genres. When \&quot;and\&quot;, it only returns the shows that have all of the given genres.  | [default to &quot;and&quot;]
 **showOriginalLanguage** | **string** | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code to only search within the shows whose original language matches with the provided language.  | 
 **yearMin** | **int32** | Minimum release/air year of the shows. | 
 **yearMax** | **int32** | Maximum release/air year of the shows. | 
 **keyword** | **string** | A keyword to only search within the shows have that keyword in their overview or title. | 
 **orderBy** | **string** | Determines the ordering of the results. | [default to &quot;original_title&quot;]
 **desc** | **bool** | The results are ordered in descending order if set true. | [default to false]
 **cursor** | **string** | Cursor is used for pagination. After each request, the response includes a \&quot;hasMore\&quot; boolean field to tell if there are more results that did not fit into the returned list. If it is set as \&quot;true\&quot;, to get the rest of the result set, send a new request (with the same parameters for other fields such as \&quot;services\&quot; etc.), and set the \&quot;cursor\&quot; parameter as the \&quot;nextCursor\&quot; value of the previous request response. Do not forget to escape the \&quot;cursor\&quot; value before putting it into a query as it might contain characters such as \&quot;?\&quot;and \&quot;&amp;\&quot;. The first request naturally does not require a \&quot;cursor\&quot; parameter.  | 

### Return type

[**SearchFiltersResponseSchema**](SearchFiltersResponseSchema.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchByTitle

> SearchTitleResponseSchema SearchByTitle(ctx).Title(title).Country(country).ShowType(showType).SeriesGranularity(seriesGranularity).OutputLanguage(outputLanguage).Execute()

Search by Title



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
    country := "ca" // string | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \"/countries\" endpoint to get the list of supported countries. 
    showType := "movie" // string | Type of shows to search in. (optional) (default to "all")
    seriesGranularity := "show" // string | \"series_granularity\" determines the level of detail for series. It does not affect movies. If \"series_granularity\" is \"show\", then the output will not include season and episode info. If \"series_granularity\" is \"season\", then the output will include season info but not episode info. If \"series_granularity\" is \"episode\", then the output will include season and episode info. If you do not need season and episode info, then it is recommended to set \"series_granularity\" as \"show\" to reduce the size of the response and increase the performance. Depending on the query (i.e. if the result set includes series with high amount of seasons/episodes), this can reduce the response size and the response time drastically. If you need deep links for individual seasons and episodes, then you should set \"series_granularity\" as \"episode\". In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. If \"series_granularity\" is not supplied, then it is set as \"episode\" by default.  (optional) (default to "episode")
    outputLanguage := "en" // string | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \"title\" will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \"output_language\" is \"es\", then the \"keyword\" will be treated as a Spanish word).  (optional) (default to "en")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SearchByTitle(context.Background()).Title(title).Country(country).ShowType(showType).SeriesGranularity(seriesGranularity).OutputLanguage(outputLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SearchByTitle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchByTitle`: SearchTitleResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SearchByTitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchByTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | Title phrase to search for. | 
 **country** | **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \&quot;/countries\&quot; endpoint to get the list of supported countries.  | 
 **showType** | **string** | Type of shows to search in. | [default to &quot;all&quot;]
 **seriesGranularity** | **string** | \&quot;series_granularity\&quot; determines the level of detail for series. It does not affect movies. If \&quot;series_granularity\&quot; is \&quot;show\&quot;, then the output will not include season and episode info. If \&quot;series_granularity\&quot; is \&quot;season\&quot;, then the output will include season info but not episode info. If \&quot;series_granularity\&quot; is \&quot;episode\&quot;, then the output will include season and episode info. If you do not need season and episode info, then it is recommended to set \&quot;series_granularity\&quot; as \&quot;show\&quot; to reduce the size of the response and increase the performance. Depending on the query (i.e. if the result set includes series with high amount of seasons/episodes), this can reduce the response size and the response time drastically. If you need deep links for individual seasons and episodes, then you should set \&quot;series_granularity\&quot; as \&quot;episode\&quot;. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. If \&quot;series_granularity\&quot; is not supplied, then it is set as \&quot;episode\&quot; by default.  | [default to &quot;episode&quot;]
 **outputLanguage** | **string** | [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code of the output language. Determines in which language the output field \&quot;title\&quot; will be in. It also effects how the keyword input will be handled in search endpoints. (i.e. if \&quot;output_language\&quot; is \&quot;es\&quot;, then the \&quot;keyword\&quot; will be treated as a Spanish word).  | [default to &quot;en&quot;]

### Return type

[**SearchTitleResponseSchema**](SearchTitleResponseSchema.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Services

> ServicesResponseSchema Services(ctx).Execute()

Services



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.Services(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Services``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Services`: ServicesResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Services`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesRequest struct via the builder pattern


### Return type

[**ServicesResponseSchema**](ServicesResponseSchema.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


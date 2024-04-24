# \GenresAPI

All URIs are relative to *https://streaming-availability.p.rapidapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenres**](GenresAPI.md#GetGenres) | **Get** /genres | Get all Genres



## GetGenres

> []Genre GetGenres(ctx).OutputLanguage(outputLanguage).Execute()

Get all Genres



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
	outputLanguage := "outputLanguage_example" // string | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenresAPI.GetGenres(context.Background()).OutputLanguage(outputLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenresAPI.GetGenres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenres`: []Genre
	fmt.Fprintf(os.Stdout, "Response from `GenresAPI.GetGenres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGenresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outputLanguage** | **string** | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  | [default to &quot;en&quot;]

### Return type

[**[]Genre**](Genre.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \CountriesAPI

All URIs are relative to *https://streaming-availability.p.rapidapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](CountriesAPI.md#GetCountries) | **Get** /countries | Get all Countries
[**GetCountry**](CountriesAPI.md#GetCountry) | **Get** /countries/{country-code} | Get a Country



## GetCountries

> map[string]Country GetCountries(ctx).OutputLanguage(outputLanguage).Execute()

Get all Countries



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
	resp, r, err := apiClient.CountriesAPI.GetCountries(context.Background()).OutputLanguage(outputLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: map[string]Country
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outputLanguage** | **string** | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  | [default to &quot;en&quot;]

### Return type

[**map[string]Country**](Country.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountry

> Country GetCountry(ctx, countryCode).OutputLanguage(outputLanguage).Execute()

Get a Country



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
	countryCode := "us" // string | [ISO 3166-1 alpha-2 code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the country. 
	outputLanguage := "outputLanguage_example" // string | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountry(context.Background(), countryCode).OutputLanguage(outputLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountry`: Country
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | [ISO 3166-1 alpha-2 code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the country.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **outputLanguage** | **string** | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  | [default to &quot;en&quot;]

### Return type

[**Country**](Country.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ChangesAPI

All URIs are relative to *https://streaming-availability.p.rapidapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChanges**](ChangesAPI.md#GetChanges) | **Get** /changes | Get Changes



## GetChanges

> ChangesResult GetChanges(ctx).Country(country).ChangeType(changeType).ItemType(itemType).Catalogs(catalogs).ShowType(showType).From(from).To(to).IncludeUnknownDates(includeUnknownDates).Cursor(cursor).DescendingOrder(descendingOrder).OutputLanguage(outputLanguage).Execute()

Get Changes



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
	changeType := openapiclient.changeType("new") // ChangeType | Type of changes to query.
	itemType := openapiclient.itemType("show") // ItemType | Type of items to search in. If `item_type` is `show`, you can use `show_type` parameter to only search for movies or series. 
	catalogs := []string{"Inner_example"} // []string | A comma separated list of up to 32 catalogs to search in. See `/countries` endpoint to get the supported services in each country and their ids.  When multiple catalogs are passed as a comma separated list, any show that is in at least one of the catalogs will be included in the result.  If no catalogs are passed, the endpoint will search in all the available catalogs in the country.  Syntax of the catalogs supplied in the list can be as the followings:  - `<sevice_id>`: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons e.g. `netflix`, `prime`, `apple`  - `<sevice_id>.<streaming_option_type>`: Only returns the shows that are available in that service with the given streaming option type. Valid streaming option types are `subscription`, `free`, `rent`, `buy` and `addon` e.g. `peacock.free` only returns the shows on Peacock that are free to watch, `prime.subscription` only returns the shows on Prime Video that are available to watch with a Prime subscription. `hulu.addon` only returns the shows on Hulu that are available via an addon, `prime.rent` only returns the shows on Prime Video that are rentable.  - `<sevice_id>.addon.<addon_id>`: Only returns the shows that are available in that service with the given addon. Check `/countries` endpoint to fetch the available addons for a service in each country. Some sample values are: `hulu.addon.hbo`, `prime.addon.hbomaxus`.  (optional)
	showType := openapiclient.showType("movie") // ShowType | Type of shows to search in. If not supplied, both movies and series will be included in the search results.  (optional)
	from := int64(1672531200) // int64 | [Unix Time Stamp](https://www.unixtimestamp.com/) to only query the changes happened/happening after this date (inclusive). For `past` changes such as `new`, `removed` or `updated`, the timestamp must be between today and `31` days ago. For `future` changes such as `expiring` or `upcoming`, the timestamp must be between today and `31` days from now in the future.  If not supplied, the default value for `past` changes is `31` days ago, and for `future` changes is today.  (optional)
	to := int64(1672531200) // int64 | [Unix Time Stamp](https://www.unixtimestamp.com/) to only query the changes happened/happening before this date (inclusive). For `past` changes such as `new`, `removed` or `updated`, the timestamp must be between today and `31` days ago. For `future` changes such as `expiring` or `upcoming`, the timestamp must be between today and `31` days from now in the future.  If not supplied, the default value for `past` changes is today, and for `future` changes is `31` days from now.  (optional)
	includeUnknownDates := true // bool | Whether to include the changes with unknown dates. `past` changes such as `new`, `removed` or `updated` will always have a timestamp thus this parameter does not affect them. `future` changes such as `expiring` or `upcoming` may not have a timestamp if the exact date is not known (e.g. some services do not explicitly state the exact date of some of the upcoming/expiring shows). If set as `true`, the changes with unknown dates will be included in the response. If set as `false`, the changes with unknown dates will be excluded from the response.  When ordering, the changes with unknown dates will be treated as if their timestamp is 0. Thus, they will appear first in the ascending order and last in the descending order.  (optional) (default to false)
	cursor := "cursor_example" // string | Cursor is used for pagination. After each request, the response includes a `hasMore` boolean field to tell if there are more results that did not fit into the returned list. If it is set as `true`, to get the rest of the result set, send a new request (with the same parameters for other fields), and set the `cursor` parameter as the `nextCursor` value of the response of the previous request. Do not forget to escape the `cursor` value before putting it into a query as it might contain characters such as `?`and `&`.  The first request naturally does not require a `cursor` parameter.  (optional)
	descendingOrder := true // bool | The results are ordered in descending order if set true. (optional) (default to false)
	outputLanguage := "outputLanguage_example" // string | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangesAPI.GetChanges(context.Background()).Country(country).ChangeType(changeType).ItemType(itemType).Catalogs(catalogs).ShowType(showType).From(from).To(to).IncludeUnknownDates(includeUnknownDates).Cursor(cursor).DescendingOrder(descendingOrder).OutputLanguage(outputLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangesAPI.GetChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChanges`: ChangesResult
	fmt.Fprintf(os.Stdout, "Response from `ChangesAPI.GetChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See &#x60;/countries&#x60; endpoint to get the list of supported countries.  | 
 **changeType** | [**ChangeType**](ChangeType.md) | Type of changes to query. | 
 **itemType** | [**ItemType**](ItemType.md) | Type of items to search in. If &#x60;item_type&#x60; is &#x60;show&#x60;, you can use &#x60;show_type&#x60; parameter to only search for movies or series.  | 
 **catalogs** | **[]string** | A comma separated list of up to 32 catalogs to search in. See &#x60;/countries&#x60; endpoint to get the supported services in each country and their ids.  When multiple catalogs are passed as a comma separated list, any show that is in at least one of the catalogs will be included in the result.  If no catalogs are passed, the endpoint will search in all the available catalogs in the country.  Syntax of the catalogs supplied in the list can be as the followings:  - &#x60;&lt;sevice_id&gt;&#x60;: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons e.g. &#x60;netflix&#x60;, &#x60;prime&#x60;, &#x60;apple&#x60;  - &#x60;&lt;sevice_id&gt;.&lt;streaming_option_type&gt;&#x60;: Only returns the shows that are available in that service with the given streaming option type. Valid streaming option types are &#x60;subscription&#x60;, &#x60;free&#x60;, &#x60;rent&#x60;, &#x60;buy&#x60; and &#x60;addon&#x60; e.g. &#x60;peacock.free&#x60; only returns the shows on Peacock that are free to watch, &#x60;prime.subscription&#x60; only returns the shows on Prime Video that are available to watch with a Prime subscription. &#x60;hulu.addon&#x60; only returns the shows on Hulu that are available via an addon, &#x60;prime.rent&#x60; only returns the shows on Prime Video that are rentable.  - &#x60;&lt;sevice_id&gt;.addon.&lt;addon_id&gt;&#x60;: Only returns the shows that are available in that service with the given addon. Check &#x60;/countries&#x60; endpoint to fetch the available addons for a service in each country. Some sample values are: &#x60;hulu.addon.hbo&#x60;, &#x60;prime.addon.hbomaxus&#x60;.  | 
 **showType** | [**ShowType**](ShowType.md) | Type of shows to search in. If not supplied, both movies and series will be included in the search results.  | 
 **from** | **int64** | [Unix Time Stamp](https://www.unixtimestamp.com/) to only query the changes happened/happening after this date (inclusive). For &#x60;past&#x60; changes such as &#x60;new&#x60;, &#x60;removed&#x60; or &#x60;updated&#x60;, the timestamp must be between today and &#x60;31&#x60; days ago. For &#x60;future&#x60; changes such as &#x60;expiring&#x60; or &#x60;upcoming&#x60;, the timestamp must be between today and &#x60;31&#x60; days from now in the future.  If not supplied, the default value for &#x60;past&#x60; changes is &#x60;31&#x60; days ago, and for &#x60;future&#x60; changes is today.  | 
 **to** | **int64** | [Unix Time Stamp](https://www.unixtimestamp.com/) to only query the changes happened/happening before this date (inclusive). For &#x60;past&#x60; changes such as &#x60;new&#x60;, &#x60;removed&#x60; or &#x60;updated&#x60;, the timestamp must be between today and &#x60;31&#x60; days ago. For &#x60;future&#x60; changes such as &#x60;expiring&#x60; or &#x60;upcoming&#x60;, the timestamp must be between today and &#x60;31&#x60; days from now in the future.  If not supplied, the default value for &#x60;past&#x60; changes is today, and for &#x60;future&#x60; changes is &#x60;31&#x60; days from now.  | 
 **includeUnknownDates** | **bool** | Whether to include the changes with unknown dates. &#x60;past&#x60; changes such as &#x60;new&#x60;, &#x60;removed&#x60; or &#x60;updated&#x60; will always have a timestamp thus this parameter does not affect them. &#x60;future&#x60; changes such as &#x60;expiring&#x60; or &#x60;upcoming&#x60; may not have a timestamp if the exact date is not known (e.g. some services do not explicitly state the exact date of some of the upcoming/expiring shows). If set as &#x60;true&#x60;, the changes with unknown dates will be included in the response. If set as &#x60;false&#x60;, the changes with unknown dates will be excluded from the response.  When ordering, the changes with unknown dates will be treated as if their timestamp is 0. Thus, they will appear first in the ascending order and last in the descending order.  | [default to false]
 **cursor** | **string** | Cursor is used for pagination. After each request, the response includes a &#x60;hasMore&#x60; boolean field to tell if there are more results that did not fit into the returned list. If it is set as &#x60;true&#x60;, to get the rest of the result set, send a new request (with the same parameters for other fields), and set the &#x60;cursor&#x60; parameter as the &#x60;nextCursor&#x60; value of the response of the previous request. Do not forget to escape the &#x60;cursor&#x60; value before putting it into a query as it might contain characters such as &#x60;?&#x60;and &#x60;&amp;&#x60;.  The first request naturally does not require a &#x60;cursor&#x60; parameter.  | 
 **descendingOrder** | **bool** | The results are ordered in descending order if set true. | [default to false]
 **outputLanguage** | **string** | [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in.  | [default to &quot;en&quot;]

### Return type

[**ChangesResult**](ChangesResult.md)

### Authorization

[X-Rapid-API-Key](../README.md#X-Rapid-API-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


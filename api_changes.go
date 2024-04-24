/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 59 countries!

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ChangesAPIService ChangesAPI service
type ChangesAPIService service

type ApiGetChangesRequest struct {
	ctx context.Context
	ApiService *ChangesAPIService
	country *string
	changeType *ChangeType
	itemType *ItemType
	catalogs *[]string
	showType *ShowType
	from *int64
	to *int64
	includeUnknownDates *bool
	cursor *string
	descendingOrder *bool
	outputLanguage *string
}

// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See \&quot;/countries\&quot; endpoint to get the list of supported countries. 
func (r ApiGetChangesRequest) Country(country string) ApiGetChangesRequest {
	r.country = &country
	return r
}

// Type of changes to query.
func (r ApiGetChangesRequest) ChangeType(changeType ChangeType) ApiGetChangesRequest {
	r.changeType = &changeType
	return r
}

// Type of items to search in. If \&quot;item_type\&quot; is \&quot;show\&quot;, you can use \&quot;show_type\&quot; parameter to only search for movies or series. 
func (r ApiGetChangesRequest) ItemType(itemType ItemType) ApiGetChangesRequest {
	r.itemType = &itemType
	return r
}

// A comma separated list of up to 32 catalogs to search in. See \&quot;/countries\&quot; endpoint to get the supported services in each country and their ids.  When multiple catalogs are passed as a comma separated list, any show that is in at least one of the catalogs will be included in the result.  If no catalogs are passed, the endpoint will search in all the available catalogs in the country.  Syntax of the catalogs supplied in the list can be as the followings:  - \&quot;&lt;sevice_id&gt;\&quot;: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons e.g. \&quot;netflix\&quot;, \&quot;prime\&quot;, \&quot;apple\&quot;  - \&quot;&lt;sevice_id&gt;.&lt;streaming_option_type&gt;\&quot;: Only returns the shows that are available in that service with the given streaming option type. Valid streaming option types are \&quot;subscription\&quot;, \&quot;free\&quot;, \&quot;rent\&quot;, \&quot;buy\&quot; and \&quot;addon\&quot; e.g. \&quot;peacock.free\&quot; only returns the shows on Peacock that are free to watch, \&quot;prime.subscription\&quot; only returns the shows on Prime Video that are available to watch with a Prime subscription. \&quot;hulu.addon\&quot; only returns the shows on Hulu that are available via an addon, \&quot;prime.rent\&quot; only returns the shows on Prime Video that are rentable.  - \&quot;&lt;sevice_id&gt;.addon.&lt;addon_id&gt;\&quot;: Only returns the shows that are available in that service with the given addon. Check \&quot;/countries\&quot; endpoint to fetch the available addons for a service in each country. Some sample values are: \&quot;hulu.addon.hbo\&quot;, \&quot;prime.addon.hbomaxus\&quot;. 
func (r ApiGetChangesRequest) Catalogs(catalogs []string) ApiGetChangesRequest {
	r.catalogs = &catalogs
	return r
}

// Type of shows to search in. If not supplied, both movies and series will be included in the search results. 
func (r ApiGetChangesRequest) ShowType(showType ShowType) ApiGetChangesRequest {
	r.showType = &showType
	return r
}

// [Unix Time Stamp](https://www.unixtimestamp.com/) to only query the changes happened/happening after this date (inclusive). For \&quot;past\&quot; changes such as \&quot;new\&quot;, \&quot;removed\&quot; or \&quot;updated\&quot;, the timestamp must be between today and \&quot;31\&quot; days ago. For \&quot;future\&quot; changes such as \&quot;expiring\&quot; or \&quot;upcoming\&quot;, the timestamp must be between today and \&quot;31\&quot; days from now in the future.  If not supplied, the default value for \&quot;past\&quot; changes is \&quot;31\&quot; days ago, and for \&quot;future\&quot; changes is today. 
func (r ApiGetChangesRequest) From(from int64) ApiGetChangesRequest {
	r.from = &from
	return r
}

// [Unix Time Stamp](https://www.unixtimestamp.com/) to only query the changes happened/happening before this date (inclusive). For \&quot;past\&quot; changes such as \&quot;new\&quot;, \&quot;removed\&quot; or \&quot;updated\&quot;, the timestamp must be between today and \&quot;31\&quot; days ago. For \&quot;future\&quot; changes such as \&quot;expiring\&quot; or \&quot;upcoming\&quot;, the timestamp must be between today and \&quot;31\&quot; days from now in the future.  If not supplied, the default value for \&quot;past\&quot; changes is today, and for \&quot;future\&quot; changes is \&quot;31\&quot; days from now. 
func (r ApiGetChangesRequest) To(to int64) ApiGetChangesRequest {
	r.to = &to
	return r
}

// Whether to include the changes with unknown dates. \&quot;past\&quot; changes such as \&quot;new\&quot;, \&quot;removed\&quot; or \&quot;updated\&quot; will always have a timestamp thus this parameter does not affect them. \&quot;future\&quot; changes such as \&quot;expiring\&quot; or \&quot;upcoming\&quot; may not have a timestamp if the exact date is not known (e.g. some services do not explicitly state the exact date of some of the upcoming/expiring shows). If set as \&quot;true\&quot;, the changes with unknown dates will be included in the response. If set as \&quot;false\&quot;, the changes with unknown dates will be excluded from the response.  When ordering, the changes with unknown dates will be treated as if their timestamp is 0. Thus, they will appear first in the ascending order and last in the descending order. 
func (r ApiGetChangesRequest) IncludeUnknownDates(includeUnknownDates bool) ApiGetChangesRequest {
	r.includeUnknownDates = &includeUnknownDates
	return r
}

// Cursor is used for pagination. After each request, the response includes a \&quot;hasMore\&quot; boolean field to tell if there are more results that did not fit into the returned list. If it is set as \&quot;true\&quot;, to get the rest of the result set, send a new request (with the same parameters for other fields), and set the \&quot;cursor\&quot; parameter as the \&quot;nextCursor\&quot; value of the response of the previous request. Do not forget to escape the \&quot;cursor\&quot; value before putting it into a query as it might contain characters such as \&quot;?\&quot;and \&quot;&amp;\&quot;.  The first request naturally does not require a \&quot;cursor\&quot; parameter. 
func (r ApiGetChangesRequest) Cursor(cursor string) ApiGetChangesRequest {
	r.cursor = &cursor
	return r
}

// The results are ordered in descending order if set true.
func (r ApiGetChangesRequest) DescendingOrder(descendingOrder bool) ApiGetChangesRequest {
	r.descendingOrder = &descendingOrder
	return r
}

// [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in. 
func (r ApiGetChangesRequest) OutputLanguage(outputLanguage string) ApiGetChangesRequest {
	r.outputLanguage = &outputLanguage
	return r
}

func (r ApiGetChangesRequest) Execute() (*ChangesResult, *http.Response, error) {
	return r.ApiService.GetChangesExecute(r)
}

/*
GetChanges Get Changes

Query the new, removed, updated, expiring or upcoming movies/series/seasons/episodes in a given list of streaming services.
Results are ordered by the date of the changes.
Changes listed per page is "25".

Changes are listed under "changes" field, and shows affected by these changes are listed under "shows" field.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChangesRequest
*/
func (a *ChangesAPIService) GetChanges(ctx context.Context) ApiGetChangesRequest {
	return ApiGetChangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ChangesResult
func (a *ChangesAPIService) GetChangesExecute(r ApiGetChangesRequest) (*ChangesResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChangesResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangesAPIService.GetChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/changes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.country == nil {
		return localVarReturnValue, nil, reportError("country is required and must be specified")
	}
	if r.changeType == nil {
		return localVarReturnValue, nil, reportError("changeType is required and must be specified")
	}
	if r.itemType == nil {
		return localVarReturnValue, nil, reportError("itemType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "")
	if r.catalogs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "catalogs", r.catalogs, "csv")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "change_type", r.changeType, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "item_type", r.itemType, "")
	if r.showType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_type", r.showType, "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.includeUnknownDates != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_unknown_dates", r.includeUnknownDates, "")
	} else {
		var defaultValue bool = false
		r.includeUnknownDates = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	if r.descendingOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "descending_order", r.descendingOrder, "")
	} else {
		var defaultValue bool = false
		r.descendingOrder = &defaultValue
	}
	if r.outputLanguage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "output_language", r.outputLanguage, "")
	} else {
		var defaultValue string = "en"
		r.outputLanguage = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Rapid-API-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-RapidAPI-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

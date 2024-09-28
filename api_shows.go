/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 60 countries!

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ShowsAPIService ShowsAPI service
type ShowsAPIService service

type ApiGetShowRequest struct {
	ctx context.Context
	ApiService *ShowsAPIService
	id string
	country *string
	seriesGranularity *string
	outputLanguage *string
}

// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the optional target country. If this parameter is not supplied, global streaming availability across all the countries will be returned. If it is supplied, only the streaming availability info from the given country will be returned. If you are only interested in the streaming availability in a single country, then it is recommended to use this parameter to reduce the size of the response and increase the performance of the endpoint. See /countries endpoint to get the list of supported countries. 
func (r ApiGetShowRequest) Country(country string) ApiGetShowRequest {
	r.country = &country
	return r
}

// series_granularity determines the level of detail for series. It does not affect movies.  If series_granularity is show, then the output will not include season and episode info.  If series_granularity is season, then the output will include season info but not episode info.  If series_granularity is episode, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set series_granularity as show to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set series_granularity as episode. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. 
func (r ApiGetShowRequest) SeriesGranularity(seriesGranularity string) ApiGetShowRequest {
	r.seriesGranularity = &seriesGranularity
	return r
}

// [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in. 
func (r ApiGetShowRequest) OutputLanguage(outputLanguage string) ApiGetShowRequest {
	r.outputLanguage = &outputLanguage
	return r
}

func (r ApiGetShowRequest) Execute() (*Show, *http.Response, error) {
	return r.ApiService.GetShowExecute(r)
}

/*
GetShow Get a Show

Get the details of a show via id, imdbId or tmdbId, including the global streaming availability info.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of the show. You can also pass an IMDb Id or a TMDB Id as an id. The endpoint will automatically detect the type of the id and fetch the show accordingly.  When passing an IMDb Id, it should be in the format of tt<numerical_id>. (e.g. tt0120338 for Titanic)  When passing a TMDB Id, it should be in the format of movie/<numerical_id> for movies and tv/<numerical_id> for series. (e.g. tv/1396 for Breaking Bad and movie/597 for Titanic) 
 @return ApiGetShowRequest
*/
func (a *ShowsAPIService) GetShow(ctx context.Context, id string) ApiGetShowRequest {
	return ApiGetShowRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Show
func (a *ShowsAPIService) GetShowExecute(r ApiGetShowRequest) (*Show, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Show
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowsAPIService.GetShow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shows/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "")
	}
	if r.seriesGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "series_granularity", r.seriesGranularity, "")
	} else {
		var defaultValue string = "episode"
		r.seriesGranularity = &defaultValue
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetTopShowsRequest struct {
	ctx context.Context
	ApiService *ShowsAPIService
	country *string
	service *string
	outputLanguage *string
	showType *ShowType
}

// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See /countries endpoint to get the list of supported countries. 
func (r ApiGetTopShowsRequest) Country(country string) ApiGetTopShowsRequest {
	r.country = &country
	return r
}

// Id of the target service. 
func (r ApiGetTopShowsRequest) Service(service string) ApiGetTopShowsRequest {
	r.service = &service
	return r
}

// [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in. 
func (r ApiGetTopShowsRequest) OutputLanguage(outputLanguage string) ApiGetTopShowsRequest {
	r.outputLanguage = &outputLanguage
	return r
}

// Type of shows to search in. If not supplied, both movies and series will be included in the search results. 
func (r ApiGetTopShowsRequest) ShowType(showType ShowType) ApiGetTopShowsRequest {
	r.showType = &showType
	return r
}

func (r ApiGetTopShowsRequest) Execute() ([]Show, *http.Response, error) {
	return r.ApiService.GetTopShowsExecute(r)
}

/*
GetTopShows Get Top Shows

Get the official top shows in a service.
Top shows are determined by the streaming service itself.

Supported streaming services are:
- Netflix: netflix
- Amazon Prime Video: prime
- Apple TV: apple
- Max: hbo

For unsupported services, this endpoint will return an empty list.

Series granularity is always show for this endpoint,
meaning that the output will not include season and episode info.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTopShowsRequest
*/
func (a *ShowsAPIService) GetTopShows(ctx context.Context) ApiGetTopShowsRequest {
	return ApiGetTopShowsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Show
func (a *ShowsAPIService) GetTopShowsExecute(r ApiGetTopShowsRequest) ([]Show, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Show
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowsAPIService.GetTopShows")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shows/top"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.country == nil {
		return localVarReturnValue, nil, reportError("country is required and must be specified")
	}
	if r.service == nil {
		return localVarReturnValue, nil, reportError("service is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "service", r.service, "")
	if r.outputLanguage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "output_language", r.outputLanguage, "")
	} else {
		var defaultValue string = "en"
		r.outputLanguage = &defaultValue
	}
	if r.showType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_type", r.showType, "")
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

type ApiSearchShowsByFiltersRequest struct {
	ctx context.Context
	ApiService *ShowsAPIService
	country *string
	catalogs *[]string
	outputLanguage *string
	showType *ShowType
	genres *[]string
	genresRelation *string
	showOriginalLanguage *string
	yearMin *int32
	yearMax *int32
	ratingMin *int32
	ratingMax *int32
	keyword *string
	seriesGranularity *string
	orderBy *string
	orderDirection *OrderDirection
	cursor *string
}

// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See /countries endpoint to get the list of supported countries. 
func (r ApiSearchShowsByFiltersRequest) Country(country string) ApiSearchShowsByFiltersRequest {
	r.country = &country
	return r
}

// A comma separated list of up to 32 catalogs to search in. See /countries endpoint to get the supported services in each country and their ids.  When multiple catalogs are passed as a comma separated list, any show that is in at least one of the catalogs will be included in the result.  If no catalogs are passed, the endpoint will search in all the available catalogs in the country.  Syntax of the catalogs supplied in the list can be as the followings:  - &lt;sevice_id&gt;: Searches in the entire catalog of that service, including (if applicable) rentable, buyable shows or shows available through addons e.g. netflix, prime, apple  - &lt;sevice_id&gt;.&lt;streaming_option_type&gt;: Only returns the shows that are available in that service with the given streaming option type. Valid streaming option types are subscription, free, rent, buy and addon e.g. peacock.free only returns the shows on Peacock that are free to watch, prime.subscription only returns the shows on Prime Video that are available to watch with a Prime subscription. hulu.addon only returns the shows on Hulu that are available via an addon, prime.rent only returns the shows on Prime Video that are rentable.  - &lt;sevice_id&gt;.addon.&lt;addon_id&gt;: Only returns the shows that are available in that service with the given addon. Check /countries endpoint to fetch the available addons for a service in each country. Some sample values are: hulu.addon.hbo, prime.addon.hbomaxus. 
func (r ApiSearchShowsByFiltersRequest) Catalogs(catalogs []string) ApiSearchShowsByFiltersRequest {
	r.catalogs = &catalogs
	return r
}

// [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in. 
func (r ApiSearchShowsByFiltersRequest) OutputLanguage(outputLanguage string) ApiSearchShowsByFiltersRequest {
	r.outputLanguage = &outputLanguage
	return r
}

// Type of shows to search in. If not supplied, both movies and series will be included in the search results. 
func (r ApiSearchShowsByFiltersRequest) ShowType(showType ShowType) ApiSearchShowsByFiltersRequest {
	r.showType = &showType
	return r
}

// A comma seperated list of genre ids to only search within the shows in those genre. See /genres endpoint to see the available genres and their ids. Use genres_relation parameter to specify between returning shows that have at least one of the given genres or returning shows that have all of the given genres. 
func (r ApiSearchShowsByFiltersRequest) Genres(genres []string) ApiSearchShowsByFiltersRequest {
	r.genres = &genres
	return r
}

// Only used when there are multiple genres supplied in genres parameter.  When or, the endpoint returns any show that has at least one of the given genres. When and, it only returns the shows that have all of the given genres. 
func (r ApiSearchShowsByFiltersRequest) GenresRelation(genresRelation string) ApiSearchShowsByFiltersRequest {
	r.genresRelation = &genresRelation
	return r
}

// [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code to only search within the shows whose original language matches with the provided language. 
func (r ApiSearchShowsByFiltersRequest) ShowOriginalLanguage(showOriginalLanguage string) ApiSearchShowsByFiltersRequest {
	r.showOriginalLanguage = &showOriginalLanguage
	return r
}

// Minimum release/air year of the shows.
func (r ApiSearchShowsByFiltersRequest) YearMin(yearMin int32) ApiSearchShowsByFiltersRequest {
	r.yearMin = &yearMin
	return r
}

// Maximum release/air year of the shows.
func (r ApiSearchShowsByFiltersRequest) YearMax(yearMax int32) ApiSearchShowsByFiltersRequest {
	r.yearMax = &yearMax
	return r
}

// Minimum rating of the shows.
func (r ApiSearchShowsByFiltersRequest) RatingMin(ratingMin int32) ApiSearchShowsByFiltersRequest {
	r.ratingMin = &ratingMin
	return r
}

// Maximum rating of the shows.
func (r ApiSearchShowsByFiltersRequest) RatingMax(ratingMax int32) ApiSearchShowsByFiltersRequest {
	r.ratingMax = &ratingMax
	return r
}

// A keyword to only search within the shows have that keyword in their overview or title.
func (r ApiSearchShowsByFiltersRequest) Keyword(keyword string) ApiSearchShowsByFiltersRequest {
	r.keyword = &keyword
	return r
}

// series_granularity determines the level of detail for series. It does not affect movies.  If series_granularity is show, then the output will not include season and episode info.  If series_granularity is season, then the output will include season info but not episode info.  If series_granularity is episode, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set series_granularity as show to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set series_granularity as episode. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. 
func (r ApiSearchShowsByFiltersRequest) SeriesGranularity(seriesGranularity string) ApiSearchShowsByFiltersRequest {
	r.seriesGranularity = &seriesGranularity
	return r
}

// Determines the ordering of the shows.  You can switch between ascending and descending order by using the order_direction parameter. 
func (r ApiSearchShowsByFiltersRequest) OrderBy(orderBy string) ApiSearchShowsByFiltersRequest {
	r.orderBy = &orderBy
	return r
}

// Determines whether to order the results in ascending or descending order.  Default value when ordering alphabetically or based on dates/times is asc.  Default value when ordering by rating or popularity is desc. 
func (r ApiSearchShowsByFiltersRequest) OrderDirection(orderDirection OrderDirection) ApiSearchShowsByFiltersRequest {
	r.orderDirection = &orderDirection
	return r
}

// Cursor is used for pagination. After each request, the response includes a hasMore boolean field to tell if there are more results that did not fit into the returned list. If it is set as true, to get the rest of the result set, send a new request (with the same parameters for other fields), and set the cursor parameter as the nextCursor value of the response of the previous request. Do not forget to escape the cursor value before putting it into a query as it might contain characters such as ?and &amp;.  The first request naturally does not require a cursor parameter. 
func (r ApiSearchShowsByFiltersRequest) Cursor(cursor string) ApiSearchShowsByFiltersRequest {
	r.cursor = &cursor
	return r
}

func (r ApiSearchShowsByFiltersRequest) Execute() (*SearchResult, *http.Response, error) {
	return r.ApiService.SearchShowsByFiltersExecute(r)
}

/*
SearchShowsByFilters Search Shows by filters

Search through the catalog of the given streaming services in the given country.
Provides filters such as show language, genres, keyword and release year.
Output includes all the information about the shows,
such as title, IMDb ID, TMDb ID, release year,
deep links to streaming services,
available subtitles, audios, available video quality
and many more!
Apart from the info about the given country-service combinations,
output also includes information about streaming availability in the other services for the given country.
Streaming availability info from the other countries are not included in the response.

When show_type is movie or series_granularity is show, items per page is 20.
When show_type is series and series_granularity is episode items per page is 10.
Otherwise, items per page is 15.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchShowsByFiltersRequest
*/
func (a *ShowsAPIService) SearchShowsByFilters(ctx context.Context) ApiSearchShowsByFiltersRequest {
	return ApiSearchShowsByFiltersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchResult
func (a *ShowsAPIService) SearchShowsByFiltersExecute(r ApiSearchShowsByFiltersRequest) (*SearchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowsAPIService.SearchShowsByFilters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shows/search/filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.country == nil {
		return localVarReturnValue, nil, reportError("country is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "")
	if r.catalogs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "catalogs", r.catalogs, "csv")
	}
	if r.outputLanguage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "output_language", r.outputLanguage, "")
	} else {
		var defaultValue string = "en"
		r.outputLanguage = &defaultValue
	}
	if r.showType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_type", r.showType, "")
	}
	if r.genres != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "genres", r.genres, "csv")
	}
	if r.genresRelation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "genres_relation", r.genresRelation, "")
	} else {
		var defaultValue string = "and"
		r.genresRelation = &defaultValue
	}
	if r.showOriginalLanguage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_original_language", r.showOriginalLanguage, "")
	}
	if r.yearMin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "year_min", r.yearMin, "")
	}
	if r.yearMax != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "year_max", r.yearMax, "")
	}
	if r.ratingMin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rating_min", r.ratingMin, "")
	}
	if r.ratingMax != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rating_max", r.ratingMax, "")
	}
	if r.keyword != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", r.keyword, "")
	}
	if r.seriesGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "series_granularity", r.seriesGranularity, "")
	} else {
		var defaultValue string = "show"
		r.seriesGranularity = &defaultValue
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "")
	} else {
		var defaultValue string = "original_title"
		r.orderBy = &defaultValue
	}
	if r.orderDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_direction", r.orderDirection, "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
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

type ApiSearchShowsByTitleRequest struct {
	ctx context.Context
	ApiService *ShowsAPIService
	title *string
	country *string
	showType *ShowType
	seriesGranularity *string
	outputLanguage *string
}

// Title phrase to search for.
func (r ApiSearchShowsByTitleRequest) Title(title string) ApiSearchShowsByTitleRequest {
	r.title = &title
	return r
}

// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code of the target country. See /countries endpoint to get the list of supported countries. 
func (r ApiSearchShowsByTitleRequest) Country(country string) ApiSearchShowsByTitleRequest {
	r.country = &country
	return r
}

// Type of shows to search in. If not supplied, both movies and series will be included in the search results. 
func (r ApiSearchShowsByTitleRequest) ShowType(showType ShowType) ApiSearchShowsByTitleRequest {
	r.showType = &showType
	return r
}

// series_granularity determines the level of detail for series. It does not affect movies.  If series_granularity is show, then the output will not include season and episode info.  If series_granularity is season, then the output will include season info but not episode info.  If series_granularity is episode, then the output will include season and episode info.  If you do not need season and episode info, then it is recommended to set series_granularity as show to reduce the size of the response and increase the performance of the endpoint.  If you need deep links for individual seasons and episodes, then you should set series_granularity as episode. In this case response will include full streaming info for seasons and episodes, similar to the streaming info for movies and series; including deep links into seasons and episodes, individual subtitle/audio and video quality info etc. 
func (r ApiSearchShowsByTitleRequest) SeriesGranularity(seriesGranularity string) ApiSearchShowsByTitleRequest {
	r.seriesGranularity = &seriesGranularity
	return r
}

// [ISO 639-1 code](https://en.wikipedia.org/wiki/ISO_639-1) of the output language. Determines in which language the output  will be in. 
func (r ApiSearchShowsByTitleRequest) OutputLanguage(outputLanguage string) ApiSearchShowsByTitleRequest {
	r.outputLanguage = &outputLanguage
	return r
}

func (r ApiSearchShowsByTitleRequest) Execute() ([]Show, *http.Response, error) {
	return r.ApiService.SearchShowsByTitleExecute(r)
}

/*
SearchShowsByTitle Search Shows by title

Search for movies and series by a title.
Maximum amount of items returned are 20
unless there are more than 20 shows with the exact given title input.
In that case all the items have 100% match with the title will be returned.

Streaming availability info for the target country is included in the response,
but not for the other countries.

Results might include shows that are not streamable in the target country.
Only criteria for the search are the title and the show type.

No pagination is supported.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchShowsByTitleRequest
*/
func (a *ShowsAPIService) SearchShowsByTitle(ctx context.Context) ApiSearchShowsByTitleRequest {
	return ApiSearchShowsByTitleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Show
func (a *ShowsAPIService) SearchShowsByTitleExecute(r ApiSearchShowsByTitleRequest) ([]Show, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Show
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowsAPIService.SearchShowsByTitle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shows/search/title"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.title == nil {
		return localVarReturnValue, nil, reportError("title is required and must be specified")
	}
	if r.country == nil {
		return localVarReturnValue, nil, reportError("country is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "title", r.title, "")
	if r.showType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_type", r.showType, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "")
	if r.seriesGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "series_granularity", r.seriesGranularity, "")
	} else {
		var defaultValue string = "show"
		r.seriesGranularity = &defaultValue
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

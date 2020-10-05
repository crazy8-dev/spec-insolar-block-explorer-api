/*
 * Insolar Explorer API
 *
 * # Insolar Explorer API documentation  Welcome to Insolar documentation for a REST-like API provided by Insolar Explorer.  [Insolar Explorer](https://github.com/insolar/block-explorer) is a service for searching for and viewing the contents of records, lifelines, jet drops, and jets.  Insolar Explorer provides an API that allows to search for said entities as well as request, filter, sort, and paginate their arrays.  ## Entities description  Basic Insolar entities that you can find in this specification are the following:  * Record—minimum unit of storage that contains an associated request, response, and service information. * Lifeline — sequence of records. A record in a lifeline is an object state and an object is a smart contract instance. * Jet drop — unit of storage for jets. * Jet — group of lifelines. * Original-Request -  it is request from outside * Request - it is request inside the platform  Some of the entities have complex identifiers. For example:  * A jet drop ID is a combination of jet ID and pulse number. * A record's index is a combination of pulse number and order. * An order is a record number in a jet drop.  ## Filtering, pagination, sorting  Some requests have filtering, sorting, and pagination capabilities.  ### Filtering  Filtering can be of two types: * By an attribute value: filters out from the returned array all the entities with attribute values other than specified.    For example: if you specify `type=request` in a [request](#operation/jet_drop_records) that returns an array of records of different types, the array doesn't contain the `result` and `state` records.  * By a range of attribute values: filters out from the returned array entities with attribute values that are out of the specified range.    You can specify inclusive, non-inclusive, open-ended, or closed-ended ranges depending on the request.    All range filter parameters have one of the following suffixes:   * `*_gt`—greater than   * `*_lt`—less than   * `*_gte`—greater than or equal to   * `*_lte`—less than or equal to    For example: if you specify `timestamp_gte=1597409219` and `timestamp_lte=1597409241` in a [request](#operation/pulses) that returns an array of pulses, the array doesn't contain pulses older than `51063280` and younger than `51063340`.  ### Sorting  The returned array is sorted by an attribute with a monotonically increasing value (`pulse_number`, `index`), by default in the descending order. Sorting is applied after filtering, if any.  For example: if you specify `sort_by=index_asc` in a [request](#operation/object_lifeline) that returns an array of records, the array is sorted in the ascending order of `index`.  ### Pagination  Pagination can be applied after filtering and sorting, and uses the following 3 parameters:  * `limit`—number of array items to return. * `offset`—number of items to skip. * `from_*`—pagination starting point: the value of a monotonically increasing attribute. The `*` designates the property name to paginate by.  For example: if you specify `?limit=100&offset=5&from_index=47382564:2` in a [request](#operation/jet_drop_records) that returns an array of records, the array:  * Starts with a record with the 5th index relative to the index `47382564:2`. * Can consist of 100 entries maximum.  The number of entries may vary depending on the actual number of records requested. To show if the value of the `limit` parameter is higher or lower than the actual number of entries, each response to a paginated request contains the actual existing number in the `total` property. 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// RecordApiService RecordApi service
type RecordApiService service

// JetDropRecordsOpts Optional parameters for the method 'JetDropRecords'
type JetDropRecordsOpts struct {
    Limit optional.Int32
    Offset optional.Int32
    FromIndex optional.String
    Type_ optional.String
}

/*
JetDropRecords Records
Gets an array of records of a jet drop given a &#x60;jet_drop_id&#x60; as a path parameter.  Optionally, specify filtering and pagination parameters. For more information, refer to the [filtering, pagination, sorting](#section/Insolar-Explorer-API-documentation/Filtering-pagination-sorting) section. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jetDropId Combination of `jet_id` and `pulse_number` separated by a `:`.
 * @param optional nil or *JetDropRecordsOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of entries to show per page.
 * @param "Offset" (optional.Int32) -  Number of entries to skip from the starting point (`from_*`).
 * @param "FromIndex" (optional.String) -  Specific `index` to paginate from.
 * @param "Type_" (optional.String) -  Record type to filter the obtained records by.
@return ObjectLifelineResponse200
*/
func (a *RecordApiService) JetDropRecords(ctx _context.Context, jetDropId string, localVarOptionals *JetDropRecordsOpts) (ObjectLifelineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectLifelineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/jet-drops/{jet_drop_id}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"jet_drop_id"+"}", _neturl.QueryEscape(parameterToString(jetDropId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromIndex.IsSet() {
		localVarQueryParams.Add("from_index", parameterToString(localVarOptionals.FromIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v ObjectLifelineResponse200
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ObjectLifelineOpts Optional parameters for the method 'ObjectLifeline'
type ObjectLifelineOpts struct {
    Limit optional.Int32
    Offset optional.Int32
    FromIndex optional.String
    SortBy optional.String
    PulseNumberGt optional.Int32
    PulseNumberLt optional.Int32
    TimestampGte optional.Int64
    TimestampLte optional.Int64
}

/*
ObjectLifeline Object lifeline
Gets an array of records of a lifeline given an &#x60;object_reference&#x60; as a path parameter.  Optionally, specify filtering, sorting, and pagination parameters. For more information, refer to the [filtering, pagination, sorting](#section/Insolar-Explorer-API-documentation/Filtering-pagination-sorting) section. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param objectReference Object reference.
 * @param optional nil or *ObjectLifelineOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of entries to show per page.
 * @param "Offset" (optional.Int32) -  Number of entries to skip from the starting point (`from_*`).
 * @param "FromIndex" (optional.String) -  Specific `index` to paginate from.
 * @param "SortBy" (optional.String) -  Sorts by the `index` attribute of the returned object. Can take two values that specify the sorting direction: descending (`index_desc`) or ascending (`index_asc`). 
 * @param "PulseNumberGt" (optional.Int32) -  Starting point for a range of pulses—greater than the specified `pulse_number`.
 * @param "PulseNumberLt" (optional.Int32) -  Ending point for a range of pulses—less than the specified `pulse_number`.
 * @param "TimestampGte" (optional.Int64) -  Starting point for a range—greater than or equal to the specified `timestamp` in the Unix format.
 * @param "TimestampLte" (optional.Int64) -  Ending point for a range—less than or equal to the specified `timestamp` in the Unix format.
@return ObjectLifelineResponse200
*/
func (a *RecordApiService) ObjectLifeline(ctx _context.Context, objectReference string, localVarOptionals *ObjectLifelineOpts) (ObjectLifelineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectLifelineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/lifeline/{object_reference}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"object_reference"+"}", _neturl.QueryEscape(parameterToString(objectReference, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromIndex.IsSet() {
		localVarQueryParams.Add("from_index", parameterToString(localVarOptionals.FromIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortBy.IsSet() {
		localVarQueryParams.Add("sort_by", parameterToString(localVarOptionals.SortBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberGt.IsSet() {
		localVarQueryParams.Add("pulse_number_gt", parameterToString(localVarOptionals.PulseNumberGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberLt.IsSet() {
		localVarQueryParams.Add("pulse_number_lt", parameterToString(localVarOptionals.PulseNumberLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimestampGte.IsSet() {
		localVarQueryParams.Add("timestamp_gte", parameterToString(localVarOptionals.TimestampGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimestampLte.IsSet() {
		localVarQueryParams.Add("timestamp_lte", parameterToString(localVarOptionals.TimestampLte.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v ObjectLifelineResponse200
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

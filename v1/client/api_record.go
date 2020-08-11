/*
 * Insolar Explorer API
 *
 * # Insolar Explorer API documentation  Welcome to Insolar documentation for a REST-like API provided by Insolar Explorer.  [Insolar Explorer](https://github.com/insolar/block-explorer) is a service for searching for and viewing the contents of individual transactions, records, lifelines, jet drops and jets.  The API allows to search for, filter and view the contents of said entities.  ## Entities description  * Record—minimum unit of storage that contains an associated request, response, and maintenance details. * Lifeline—sequence of records for object state where an object is a smart contract instance. * Jet drop—unit of storage for jets. * Jet—groups of lifelines.  * Jet drop ID—combination of jet id with pulse number. * Index—combination of pulse number with order (record number in a jet drop).   ## Filtering, pagination, sorting  API provides filtering based on a range of values: greater than and less than, or greater than or equal to and less than or equal to.  API provides a combination of offset and seek pagination.  Pagination can be applied using: * Combination of a starting point (`from_*`), number of entries per page (`limit`) and number of entries to skip from the starting point (`offset`). * Just `limit` to get a limited array of the latest data. * Combination of the filtering parameters `*_gt`/`*_gte` and `*_lt`/`*_lte`, and `limit`.  Some requests can be sorted in the descending (`*_desc`)  or ascending (`*_asc`) order. 
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
Gets records by &#x60;jet_drop_id&#x60; and based on specified filtering and pagination parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jetDropId Combination of `jet_id` with `pulse_number`.
 * @param optional nil or *JetDropRecordsOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of entries to show per page.
 * @param "Offset" (optional.Int32) -  Number of entries to skip from the starting point (`from_*`).
 * @param "FromIndex" (optional.String) -  Specific `index` to paginate from.
 * @param "Type_" (optional.String) -  Record type for filtering records.
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
Gets a lifeline with all its records by &#x60;object_reference&#x60; and based on specified filtering, pagination and sorting parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param objectReference Object reference.
 * @param optional nil or *ObjectLifelineOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of entries to show per page.
 * @param "Offset" (optional.Int32) -  Number of entries to skip from the starting point (`from_*`).
 * @param "FromIndex" (optional.String) -  Specific `index` to paginate from.
 * @param "SortBy" (optional.String) -  Sorts by the `index` attribute of the returned object. Can take two values that specify the sorting direction: descending (`index_desc`) or ascending (`index_asc`). 
 * @param "PulseNumberGt" (optional.Int32) -  Starting point in a range. Greater than this `pulse_number`.
 * @param "PulseNumberLt" (optional.Int32) -  Ending point in a range. Less than this `pulse_number`.
 * @param "TimestampGte" (optional.Int64) -  Starting point in a range. Greater than or equal to this `timestamp` in Unix format.
 * @param "TimestampLte" (optional.Int64) -  Ending point in a range. Greater than or equal to this `timestamp` in Unix format.
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

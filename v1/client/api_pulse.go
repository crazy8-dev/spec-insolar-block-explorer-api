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

// PulseApiService PulseApi service
type PulseApiService service

/*
Pulse Pulse
Gets pulse by &#x60;pulse_number&#x60;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulseNumber Pulse number.
@return PulseResponse200
*/
func (a *PulseApiService) Pulse(ctx _context.Context, pulseNumber int64) (PulseResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PulseResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/pulses/{pulse_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulse_number"+"}", _neturl.QueryEscape(parameterToString(pulseNumber, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if pulseNumber < 0 {
		return localVarReturnValue, nil, reportError("pulseNumber must be greater than 0")
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
			var v PulseResponse200
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

// PulsesOpts Optional parameters for the method 'Pulses'
type PulsesOpts struct {
    Limit optional.Int32
    Offset optional.Int32
    FromPulseNumber optional.Int64
    TimestampGte optional.Int64
    TimestampLte optional.Int64
}

/*
Pulses Get pulses
Gets a range of pulses based on specified filtering and pagination parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PulsesOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of entries to show per page.
 * @param "Offset" (optional.Int32) -  Number of entries to skip from the starting point (`from_*`).
 * @param "FromPulseNumber" (optional.Int64) -  Specific `pulse_number` to paginate from.
 * @param "TimestampGte" (optional.Int64) -  Starting point in a range. Greater than or equal to this `timestamp` in Unix format.
 * @param "TimestampLte" (optional.Int64) -  Ending point in a range. Greater than or equal to this `timestamp` in Unix format.
@return GetPulsesResponse200
*/
func (a *PulseApiService) Pulses(ctx _context.Context, localVarOptionals *PulsesOpts) (GetPulsesResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetPulsesResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/pulses"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromPulseNumber.IsSet() {
		localVarQueryParams.Add("from_pulse_number", parameterToString(localVarOptionals.FromPulseNumber.Value(), ""))
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
			var v GetPulsesResponse200
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

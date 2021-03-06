/*
 * Insolar Explorer API
 *
 * # Insolar Explorer API documentation  Welcome to Insolar documentation for a REST-like API provided by Insolar Explorer.  [Insolar Explorer](https://github.com/insolar/block-explorer) is a service for searching for and viewing the contents of records, lifelines, jet drops, and jets.  Insolar Explorer provides an API that allows to search for said entities as well as request, filter, sort, and paginate their arrays.  ## Entities description  Basic Insolar entities that you can find in this specification are the following:  * Record—minimum unit of storage that contains an associated request, response, and service information.   Records can be of three types:   * Request. Most requests are caused by other requests inside the Plaftorm with the exception of original requests. The latter are invoked by users.   * Result. Each request has a corresponding execution result.   * State. Some execution results change object state. They are caused by so-called mutable requests, while immutable requests are simple read operations. * Lifeline—sequence of records. A record in a lifeline is an object state and an object is a smart contract instance. * Jet drop—unit of storage for jets. * Jet—group of lifelines.  Some of the entities have complex identifiers. For example:  * A jet drop ID is a combination of jet ID and pulse number. * A record's index is a combination of pulse number and order. * An order is a record number in a jet drop.  ## Filtering, pagination, sorting  Some requests have filtering, sorting, and pagination capabilities.  ### Filtering  Filtering can be of two types: * By an attribute value: filters out from the returned array all the entities with attribute values other than specified.    For example: if you specify `type=request` in a [request](#operation/jet_drop_records) that returns an array of records of different types, the array doesn't contain the `result` and `state` records.  * By a range of attribute values: filters out from the returned array entities with attribute values that are out of the specified range.    You can specify inclusive, non-inclusive, open-ended, or closed-ended ranges depending on the request.    All range filter parameters have one of the following suffixes:   * `*_gt`—greater than   * `*_lt`—less than   * `*_gte`—greater than or equal to   * `*_lte`—less than or equal to    For example: if you specify `timestamp_gte=1597409219` and `timestamp_lte=1597409241` in a [request](#operation/pulses) that returns an array of pulses, the array doesn't contain pulses older than `51063280` and younger than `51063340`.  ### Sorting  The returned array is sorted by an attribute with a monotonically increasing value (`pulse_number`, `index`), by default in the descending order. Sorting is applied after filtering, if any.  For example: if you specify `sort_by=index_asc` in a [request](#operation/object_lifeline) that returns an array of records, the array is sorted in the ascending order of `index`.  ### Pagination  Pagination can be applied after filtering and sorting, and uses the following 3 parameters:  * `limit`—number of array items to return. * `offset`—number of items to skip. * `from_*`—pagination starting point: the value of a monotonically increasing attribute. The `*` designates the property name to paginate by.  For example: if you specify `?limit=100&offset=5&from_index=47382564:2` in a [request](#operation/jet_drop_records) that returns an array of records, the array:  * Starts with a record with the 5th index relative to the index `47382564:2`. * Can consist of 100 entries maximum.  The number of entries may vary depending on the actual number of records requested. To show if the value of the `limit` parameter is higher or lower than the actual number of entries, each response to a paginated request contains the actual existing number in the `total` property. 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// JetDrops Response codes.
type JetDrops struct {
	// Actual number of existing entries. May be higher or lower than the specified `limit`.
	Total int64 `json:"total,omitempty"`
	// Array with a number entries as specified by filtering and pagination parameters.
	Result []JetDropByIdResponse200 `json:"result,omitempty"`
	// Error code received from the backend services.
	Code string `json:"code,omitempty"`
	// Short error description.
	Message string `json:"message,omitempty"`
	// Additional information about the error.
	Description string `json:"description,omitempty"`
	// Array containing incorrect parameters/properties.
	ValidationFailures []PulsesResponse200ValidationFailures `json:"validation_failures,omitempty"`
}

/*
 * Insolar Explorer API
 *
 * # Insolar Explorer API documentation  Welcome to Insolar documentation for a REST-like API provided by Insolar Explorer.  [Insolar Explorer](https://github.com/insolar/block-explorer) is a service for searching for and viewing the contents of individual transactions, records, lifelines, jet drops and jets.  The API allows to search for, filter and view the contents of said entities.  ## Entities description  * Record—minimum unit of storage that contains an associated request, response, and maintenance details. * Lifeline—sequence of records for object state where an object is a smart contract instance. * Jet drop—unit of storage for jets. * Jet—group of lifelines.  * Jet drop ID—combination of jet id with pulse number. * Index—combination of pulse number with order (record number in a jet drop).   ## Filtering, pagination, sorting  API provides filtering based on a range of values: greater than and less than, greater than or equal to and less than or equal to.  API provides a combination of offset and seek pagination.  Pagination can be applied using: * Combination of a starting point (`from_*`), number of entries per page (`limit`) and number of entries to skip from the starting point (`offset`). * Just `limit` to get a limited array of the latest data. * Combination of the filtering parameters `*_gt`/`*_gte` and `*_lt`/`*_lte`, and `limit`.  Some requests can be sorted in the descending (`*_desc`)  or ascending (`*_asc`) order. 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// JetDropByIdResponse200 Response codes.
type JetDropByIdResponse200 struct {
	// Combination of `jet_id` with `pulse_number`.
	JetDropId string `json:"jet_drop_id,omitempty"`
	// Next `jet_drop_id`.
	NextJetDropId []JetDropByIdResponse200NextJetDropId `json:"next_jet_drop_id,omitempty"`
	// Previous `jet_drop_id`.
	PrevJetDropId []JetDropByIdResponse200NextJetDropId `json:"prev_jet_drop_id,omitempty"`
	// Jet ID.
	JetId string `json:"jet_id,omitempty"`
	// Pulse number.
	PulseNumber int64 `json:"pulse_number,omitempty"`
	// Number of all records in the pulse.
	RecordAmount int64 `json:"record_amount,omitempty"`
	// Unix timestamp.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Record hash.
	Hash string `json:"hash,omitempty"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Link string `json:"link,omitempty"`
	ValidationFailures []GetPulsesResponse200ValidationFailures `json:"validation_failures,omitempty"`
}

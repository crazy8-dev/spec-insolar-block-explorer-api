/*
 * Insolar Explorer API
 *
 * [Insolar Explorer](https://github.com/insolar/block-explorer)'s REST API documentation.  Insolar Explorer is a service that allows users to search for and view the contents of individual transactions, records, lifelines, jet drops and jets.  * Record—minimum unit of storage that contains an associated request, response, and maintenance details * Lifeline—sequence of records for object state where an object is a smart contract instance * Jet drop—unit of storage for jets * Jet—groups of lifelines  Main API endpoints include: * Search endpoint: `/api/v1/search` * Endpoints for each entity type:   * Records: `/api/v1/jet-drops/{jet_drop_id}/records`   * Lifeline: `/api/v1/lifeline/{object_reference}/records`   * Pulse: `/api/v1/pulses/{pulse_number}`   * Pulses: `/api/v1/pulses`   * Jet drop by jet ID: `/api/v1/jets/{jet_id}/jet-drops`   * Jet drop by its ID: `/api/v1/jet-drops/{jet_drop_id}`   * Jet drop by pulse number: `/api/v1/pulses/{pulse_number}/jet-drops`  A search request first gets to the search endpoint where its the entity type `value` is identified.  It then is redirected to the appropriate endpoint for detailed metadata to show the user on the Insolar Explorer Web UI.  Search results can be paginated. Pagination parameters are optional and include: * Starting point that depends on the searched entity type: `from_index`, `from_pulse_number`, or `from_jet_drop_id` * Number of entries per page: `limit` * Number of entries to skip from the starting point to show on a new page: `offset`  For example, for 10 entries per page starting from the pulse number 431920 for a set of 200 entries: `<example.com>/<api_endpoint>?<params>&from_pulse_number=431920&limit=10&offset=0`.  Search results can additionally be sorted by `index` or `pulse` in the descending or ascending order.  For example, `<example.com>/<api_endpoint>?<params>&sort_by=-index`.  Search results shown to the user on the Insolar Explorer Web UI may miss some data. The `Reload` button gets the missed data via requests to the API with additional filtering parameters: * `jet_drop_id_gt` and `jet_drop_id_lt` to obtain missing data within a range of jet drops (greater than, less then) * `pulse_number_gt` and `pulse_number_lt` to obtain missing data within a range of pulses (greater than, less then) * `timestamp_gte` and `timestamp_lte` to obtain missing data within a timespan (greater than or equal, less then or equal)  For example, `<example.com>/<api_endpoint>?<params>&pulse_number_gt=431902&pulse_number_lt=431904`. 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// SearchRecordMeta Meta data.
type SearchRecordMeta struct {
	// Object reference.
	ObjectReference string `json:"object_reference,omitempty"`
	// Index is combination of `pulse_number` with `order` (record number in a jet drop).
	Index string `json:"index,omitempty"`
}

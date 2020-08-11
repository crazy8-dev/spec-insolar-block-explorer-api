// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// CodeError defines model for code-error.
type CodeError struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Message     *string `json:"message,omitempty"`
}

// CodeValidationError defines model for code-validation-error.
type CodeValidationError struct {
	Code               *string                   `json:"code,omitempty"`
	Description        *string                   `json:"description,omitempty"`
	Link               *string                   `json:"link,omitempty"`
	Message            *string                   `json:"message,omitempty"`
	ValidationFailures *[]CodeValidationFailures `json:"validation_failures,omitempty"`
}

// CodeValidationFailures defines model for code-validation-failures.
type CodeValidationFailures struct {

	// Failure reason.
	FailureReason *string `json:"failure_reason,omitempty"`

	// Property name.
	Property *string `json:"property,omitempty"`
}

// JetDrop defines model for jet-drop.
type JetDrop struct {

	// Record hash.
	Hash *string `json:"hash,omitempty"`

	// Combination of `jet_id` with `pulse_number`.
	JetDropId *string `json:"jet_drop_id,omitempty"`

	// Jet ID.
	JetId *string `json:"jet_id,omitempty"`

	// Next `jet_drop_id`.
	NextJetDropId *[]NextPrevJetDrop `json:"next_jet_drop_id,omitempty"`

	// Previous `jet_drop_id`.
	PrevJetDropId *[]NextPrevJetDrop `json:"prev_jet_drop_id,omitempty"`

	// Pulse number.
	PulseNumber *int64 `json:"pulse_number,omitempty"`

	// Number of all records in the pulse.
	RecordAmount *int64 `json:"record_amount,omitempty"`

	// Unix timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// JetDrops defines model for jet-drops.
type JetDrops struct {
	Result *[]JetDrop `json:"result,omitempty"`

	// Number of entries in the array.
	Total *int64 `json:"total,omitempty"`
}

// NextPrevJetDrop defines model for next-prev-jet-drop.
type NextPrevJetDrop struct {

	// Combination of `jet_id` with `pulse_number`.
	JetDropId *string `json:"jet_drop_id,omitempty"`

	// Jet ID.
	JetId *string `json:"jet_id,omitempty"`

	// Pulse number.
	PulseNumber *int64 `json:"pulse_number,omitempty"`
}

// Pulse defines model for pulse.
type Pulse struct {

	// Pulse completion status.
	IsComplete *bool `json:"is_complete,omitempty"`

	// Number of all jet drops in the pulse.
	JetDropAmount *int64 `json:"jet_drop_amount,omitempty"`

	// Next pulse number.
	NextPulseNumber *int64 `json:"next_pulse_number,omitempty"`

	// Previous pulse number.
	PrevPulseNumber *int64 `json:"prev_pulse_number,omitempty"`

	// Pulse number.
	PulseNumber *int64 `json:"pulse_number,omitempty"`

	// Number of all records in the pulse.
	RecordAmount *int64 `json:"record_amount,omitempty"`

	// Unix timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// Pulses defines model for pulses.
type Pulses struct {
	Result *[]Pulse `json:"result,omitempty"`

	// Number of entries in the array.
	Total *int64 `json:"total,omitempty"`
}

// Record defines model for record.
type Record struct {

	// Record hash.
	Hash *string `json:"hash,omitempty"`

	// Index—combination of `pulse_number` with `order` (record number in a jet drop).
	Index *string `json:"index,omitempty"`

	// Combination of `jet_id` with `pulse_number`.
	JetDropId *string `json:"jet_drop_id,omitempty"`

	// Jet ID.
	JetId *string `json:"jet_id,omitempty"`

	// Object reference.
	ObjectReference *string `json:"object_reference,omitempty"`

	// Order—record number in a `jet drop`.
	Order *int64 `json:"order,omitempty"`

	// Record payload.
	Payload *string `json:"payload,omitempty"`

	// Reference of previous record.
	PrevRecordReference *string `json:"prev_record_reference,omitempty"`

	// Prototype reference.
	PrototypeReference *string `json:"prototype_reference,omitempty"`

	// Pulse number.
	PulseNumber *int64 `json:"pulse_number,omitempty"`

	// Record reference.
	Reference *string `json:"reference,omitempty"`

	// Unix timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`

	// Record type.
	Type *string `json:"type,omitempty"`
}

// Records defines model for records.
type Records struct {
	Result *[]Record `json:"result,omitempty"`

	// Number of entries in the array.
	Total *int64 `json:"total,omitempty"`
}

// SearchJetDrop defines model for search-jet-drop.
type SearchJetDrop struct {

	// Meta data.
	Meta *struct {

		// Combination of `jet_id` with `pulse_number`.
		JetDropId *string `json:"jet_drop_id,omitempty"`
	} `json:"meta,omitempty"`

	// Result type.
	Type *string `json:"type,omitempty"`
}

// SearchLifeline defines model for search-lifeline.
type SearchLifeline struct {

	// Meta data.
	Meta *struct {

		// Object reference.
		ObjectReference *string `json:"object_reference,omitempty"`
	} `json:"meta,omitempty"`

	// Result type.
	Type *string `json:"type,omitempty"`
}

// SearchPulse defines model for search-pulse.
type SearchPulse struct {

	// Meta data.
	Meta *struct {

		// Pulse number.
		PulseNumber *int64 `json:"pulse_number,omitempty"`
	} `json:"meta,omitempty"`

	// Result type.
	Type *string `json:"type,omitempty"`
}

// SearchRecord defines model for search-record.
type SearchRecord struct {

	// Meta data.
	Meta *struct {

		// Index—combination of `pulse_number` with `order` (record number in a jet drop).
		Index *string `json:"index,omitempty"`

		// Object reference.
		ObjectReference *string `json:"object_reference,omitempty"`
	} `json:"meta,omitempty"`

	// Result type.
	Type *string `json:"type,omitempty"`
}

// FromPulseNumberParam defines model for fromPulseNumberParam.
type FromPulseNumberParam int64

// FromIndex defines model for from_index.
type FromIndex string

// FromJetDropId defines model for from_jet_drop_id.
type FromJetDropId string

// JetDropIdPath defines model for jet_drop_id_path.
type JetDropIdPath string

// JetIdPath defines model for jet_id_path.
type JetIdPath string

// Limit defines model for limit.
type Limit int

// ObjectReferencePath defines model for object_reference_path.
type ObjectReferencePath string

// OffsetParam defines model for offsetParam.
type OffsetParam int

// PulseNumberGt defines model for pulse_number_gt.
type PulseNumberGt int

// PulseNumberGte defines model for pulse_number_gte.
type PulseNumberGte int

// PulseNumberLt defines model for pulse_number_lt.
type PulseNumberLt int

// PulseNumberLte defines model for pulse_number_lte.
type PulseNumberLte int

// PulseNumberPath defines model for pulse_number_path.
type PulseNumberPath int64

// RecordTypeParam defines model for recordTypeParam.
type RecordTypeParam string

// List of RecordTypeParam
const (
	RecordTypeParam_request RecordTypeParam = "request"
	RecordTypeParam_result  RecordTypeParam = "result"
	RecordTypeParam_state   RecordTypeParam = "state"
)

// SortByIndex defines model for sort_by_index.
type SortByIndex string

// List of SortByIndex
const (
	SortByIndex_index_asc  SortByIndex = "index_asc"
	SortByIndex_index_desc SortByIndex = "index_desc"
)

// SortByPulse defines model for sort_by_pulse.
type SortByPulse string

// List of SortByPulse
const (
	SortByPulse_pulse_number_asc_jet_id_desc SortByPulse = "pulse_number_asc,jet_id_desc"
	SortByPulse_pulse_number_desc_jet_id_asc SortByPulse = "pulse_number_desc,jet_id_asc"
)

// TimestampGte defines model for timestamp_gte.
type TimestampGte int64

// TimestampLte defines model for timestamp_lte.
type TimestampLte int64

// N400Response defines model for 400Response.
type N400Response CodeValidationError

// N500Response defines model for 500Response.
type N500Response CodeError

// JetDropResponse defines model for jetDropResponse.
type JetDropResponse JetDrop

// JetDropsResponse defines model for jetDropsResponse.
type JetDropsResponse JetDrops

// PulseResponse defines model for pulseResponse.
type PulseResponse Pulse

// PulsesResponse defines model for pulsesResponse.
type PulsesResponse Pulses

// RecordsResponse defines model for recordsResponse.
type RecordsResponse Records

// SearchResponse defines model for searchResponse.
type SearchResponse interface{}

// JetDropRecordsParams defines parameters for JetDropRecords.
type JetDropRecordsParams struct {

	// Defines a number of entries to show per page.
	Limit *Limit `json:"limit,omitempty"`

	// Defines a number of entries to skip from the starting point (`from_*`).
	Offset *OffsetParam `json:"offset,omitempty"`

	// Defines a specific `index` to paginate from.
	FromIndex *FromIndex `json:"from_index,omitempty"`

	// Defines the record type to filter the obtained records by.
	Type *RecordTypeParam `json:"type,omitempty"`
}

// JetDropsByJetIDParams defines parameters for JetDropsByJetID.
type JetDropsByJetIDParams struct {

	// Defines a number of entries to show per page.
	Limit *Limit `json:"limit,omitempty"`

	// Sorts by the `pulse_number` attribute of the returned object.
	// Can take two values that specify the sorting direction: descending (`pulse_number_desc`) or ascending (`pulse_number_asc`).
	SortBy *SortByPulse `json:"sort_by,omitempty"`

	// Defines the starting point for a returned range of pulses—greater than or equal to the specified `pulse_number`.
	PulseNumberGte *PulseNumberGte `json:"pulse_number_gte,omitempty"`

	// Defines the starting point for a returned range of pulses—greater than the specified `pulse_number`.
	PulseNumberGt *PulseNumberGt `json:"pulse_number_gt,omitempty"`

	// Defines the ending point for a returned range of pulses—less than equal to the specified `pulse_number`.
	PulseNumberLte *PulseNumberLte `json:"pulse_number_lte,omitempty"`

	// Defines the ending point for a returned range of pulses—less than the specified `pulse_number`.
	PulseNumberLt *PulseNumberLt `json:"pulse_number_lt,omitempty"`
}

// ObjectLifelineParams defines parameters for ObjectLifeline.
type ObjectLifelineParams struct {

	// Defines a number of entries to show per page.
	Limit *Limit `json:"limit,omitempty"`

	// Defines a number of entries to skip from the starting point (`from_*`).
	Offset *OffsetParam `json:"offset,omitempty"`

	// Defines a specific `index` to paginate from.
	FromIndex *FromIndex `json:"from_index,omitempty"`

	// Sorts by the `index` attribute of the returned object.
	// Can take two values that specify the sorting direction: descending (`index_desc`) or ascending (`index_asc`).
	SortBy *SortByIndex `json:"sort_by,omitempty"`

	// Defines the starting point for a returned range of pulses—greater than the specified `pulse_number`.
	PulseNumberGt *PulseNumberGt `json:"pulse_number_gt,omitempty"`

	// Defines the ending point for a returned range of pulses—less than the specified `pulse_number`.
	PulseNumberLt *PulseNumberLt `json:"pulse_number_lt,omitempty"`

	// Defines the starting point for a returned range—greater than or equal to the specified `timestamp` in the Unix format.
	TimestampGte *TimestampGte `json:"timestamp_gte,omitempty"`

	// Defines the ending point for a returned range—less than or equal to the specified `timestamp` in the Unix format.
	TimestampLte *TimestampLte `json:"timestamp_lte,omitempty"`
}

// PulsesParams defines parameters for Pulses.
type PulsesParams struct {

	// Defines a number of entries to show per page.
	Limit *Limit `json:"limit,omitempty"`

	// Defines a number of entries to skip from the starting point (`from_*`).
	Offset *OffsetParam `json:"offset,omitempty"`

	// Defines a specific `pulse_number` to paginate from.
	FromPulseNumber *FromPulseNumberParam `json:"from_pulse_number,omitempty"`

	// Defines the starting point for a returned range—greater than or equal to the specified `timestamp` in the Unix format.
	TimestampGte *TimestampGte `json:"timestamp_gte,omitempty"`

	// Defines the ending point for a returned range—less than or equal to the specified `timestamp` in the Unix format.
	TimestampLte *TimestampLte `json:"timestamp_lte,omitempty"`
}

// JetDropsByPulseNumberParams defines parameters for JetDropsByPulseNumber.
type JetDropsByPulseNumberParams struct {

	// Defines a number of entries to show per page.
	Limit *Limit `json:"limit,omitempty"`

	// Defines a number of entries to skip from the starting point (`from_*`).
	Offset *OffsetParam `json:"offset,omitempty"`

	// Defins a specific `jet_drop_id` to paginate from.
	FromJetDropId *FromJetDropId `json:"from_jet_drop_id,omitempty"`
}

// SearchParams defines parameters for Search.
type SearchParams struct {

	// Searching value.
	Value string `json:"value"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Jet drop by ID
	// (GET /api/v1/jet-drops/{jet_drop_id})
	JetDropByID(ctx echo.Context, jetDropId JetDropIdPath) error
	// Records
	// (GET /api/v1/jet-drops/{jet_drop_id}/records)
	JetDropRecords(ctx echo.Context, jetDropId JetDropIdPath, params JetDropRecordsParams) error
	// Jet drops by jet ID
	// (GET /api/v1/jets/{jet_id}/jet-drops)
	JetDropsByJetID(ctx echo.Context, jetId JetIdPath, params JetDropsByJetIDParams) error
	// Object lifeline
	// (GET /api/v1/lifeline/{object_reference}/records)
	ObjectLifeline(ctx echo.Context, objectReference ObjectReferencePath, params ObjectLifelineParams) error
	// Get pulses
	// (GET /api/v1/pulses)
	Pulses(ctx echo.Context, params PulsesParams) error
	// Pulse
	// (GET /api/v1/pulses/{pulse_number})
	Pulse(ctx echo.Context, pulseNumber PulseNumberPath) error
	// Jet drops by pulse number
	// (GET /api/v1/pulses/{pulse_number}/jet-drops)
	JetDropsByPulseNumber(ctx echo.Context, pulseNumber PulseNumberPath, params JetDropsByPulseNumberParams) error
	// Search
	// (GET /api/v1/search)
	Search(ctx echo.Context, params SearchParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// JetDropByID converts echo context to params.
func (w *ServerInterfaceWrapper) JetDropByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "jet_drop_id" -------------
	var jetDropId JetDropIdPath

	err = runtime.BindStyledParameter("simple", false, "jet_drop_id", ctx.Param("jet_drop_id"), &jetDropId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter jet_drop_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.JetDropByID(ctx, jetDropId)
	return err
}

// JetDropRecords converts echo context to params.
func (w *ServerInterfaceWrapper) JetDropRecords(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "jet_drop_id" -------------
	var jetDropId JetDropIdPath

	err = runtime.BindStyledParameter("simple", false, "jet_drop_id", ctx.Param("jet_drop_id"), &jetDropId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter jet_drop_id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params JetDropRecordsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "from_index" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_index", ctx.QueryParams(), &params.FromIndex)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_index: %s", err))
	}

	// ------------- Optional query parameter "type" -------------

	err = runtime.BindQueryParameter("form", true, false, "type", ctx.QueryParams(), &params.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter type: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.JetDropRecords(ctx, jetDropId, params)
	return err
}

// JetDropsByJetID converts echo context to params.
func (w *ServerInterfaceWrapper) JetDropsByJetID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "jet_id" -------------
	var jetId JetIdPath

	err = runtime.BindStyledParameter("simple", false, "jet_id", ctx.Param("jet_id"), &jetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter jet_id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params JetDropsByJetIDParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "sort_by" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort_by", ctx.QueryParams(), &params.SortBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort_by: %s", err))
	}

	// ------------- Optional query parameter "pulse_number_gte" -------------

	err = runtime.BindQueryParameter("form", true, false, "pulse_number_gte", ctx.QueryParams(), &params.PulseNumberGte)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pulse_number_gte: %s", err))
	}

	// ------------- Optional query parameter "pulse_number_gt" -------------

	err = runtime.BindQueryParameter("form", true, false, "pulse_number_gt", ctx.QueryParams(), &params.PulseNumberGt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pulse_number_gt: %s", err))
	}

	// ------------- Optional query parameter "pulse_number_lte" -------------

	err = runtime.BindQueryParameter("form", true, false, "pulse_number_lte", ctx.QueryParams(), &params.PulseNumberLte)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pulse_number_lte: %s", err))
	}

	// ------------- Optional query parameter "pulse_number_lt" -------------

	err = runtime.BindQueryParameter("form", true, false, "pulse_number_lt", ctx.QueryParams(), &params.PulseNumberLt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pulse_number_lt: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.JetDropsByJetID(ctx, jetId, params)
	return err
}

// ObjectLifeline converts echo context to params.
func (w *ServerInterfaceWrapper) ObjectLifeline(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "object_reference" -------------
	var objectReference ObjectReferencePath

	err = runtime.BindStyledParameter("simple", false, "object_reference", ctx.Param("object_reference"), &objectReference)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter object_reference: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ObjectLifelineParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "from_index" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_index", ctx.QueryParams(), &params.FromIndex)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_index: %s", err))
	}

	// ------------- Optional query parameter "sort_by" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort_by", ctx.QueryParams(), &params.SortBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort_by: %s", err))
	}

	// ------------- Optional query parameter "pulse_number_gt" -------------

	err = runtime.BindQueryParameter("form", true, false, "pulse_number_gt", ctx.QueryParams(), &params.PulseNumberGt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pulse_number_gt: %s", err))
	}

	// ------------- Optional query parameter "pulse_number_lt" -------------

	err = runtime.BindQueryParameter("form", true, false, "pulse_number_lt", ctx.QueryParams(), &params.PulseNumberLt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pulse_number_lt: %s", err))
	}

	// ------------- Optional query parameter "timestamp_gte" -------------

	err = runtime.BindQueryParameter("form", true, false, "timestamp_gte", ctx.QueryParams(), &params.TimestampGte)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timestamp_gte: %s", err))
	}

	// ------------- Optional query parameter "timestamp_lte" -------------

	err = runtime.BindQueryParameter("form", true, false, "timestamp_lte", ctx.QueryParams(), &params.TimestampLte)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timestamp_lte: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ObjectLifeline(ctx, objectReference, params)
	return err
}

// Pulses converts echo context to params.
func (w *ServerInterfaceWrapper) Pulses(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PulsesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "from_pulse_number" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_pulse_number", ctx.QueryParams(), &params.FromPulseNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_pulse_number: %s", err))
	}

	// ------------- Optional query parameter "timestamp_gte" -------------

	err = runtime.BindQueryParameter("form", true, false, "timestamp_gte", ctx.QueryParams(), &params.TimestampGte)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timestamp_gte: %s", err))
	}

	// ------------- Optional query parameter "timestamp_lte" -------------

	err = runtime.BindQueryParameter("form", true, false, "timestamp_lte", ctx.QueryParams(), &params.TimestampLte)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timestamp_lte: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Pulses(ctx, params)
	return err
}

// Pulse converts echo context to params.
func (w *ServerInterfaceWrapper) Pulse(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pulse_number" -------------
	var pulseNumber PulseNumberPath

	err = runtime.BindStyledParameter("simple", false, "pulse_number", ctx.Param("pulse_number"), &pulseNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pulse_number: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Pulse(ctx, pulseNumber)
	return err
}

// JetDropsByPulseNumber converts echo context to params.
func (w *ServerInterfaceWrapper) JetDropsByPulseNumber(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pulse_number" -------------
	var pulseNumber PulseNumberPath

	err = runtime.BindStyledParameter("simple", false, "pulse_number", ctx.Param("pulse_number"), &pulseNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pulse_number: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params JetDropsByPulseNumberParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "from_jet_drop_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_jet_drop_id", ctx.QueryParams(), &params.FromJetDropId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_jet_drop_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.JetDropsByPulseNumber(ctx, pulseNumber, params)
	return err
}

// Search converts echo context to params.
func (w *ServerInterfaceWrapper) Search(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchParams
	// ------------- Required query parameter "value" -------------

	err = runtime.BindQueryParameter("form", true, true, "value", ctx.QueryParams(), &params.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter value: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Search(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/api/v1/jet-drops/:jet_drop_id", wrapper.JetDropByID)
	router.GET("/api/v1/jet-drops/:jet_drop_id/records", wrapper.JetDropRecords)
	router.GET("/api/v1/jets/:jet_id/jet-drops", wrapper.JetDropsByJetID)
	router.GET("/api/v1/lifeline/:object_reference/records", wrapper.ObjectLifeline)
	router.GET("/api/v1/pulses", wrapper.Pulses)
	router.GET("/api/v1/pulses/:pulse_number", wrapper.Pulse)
	router.GET("/api/v1/pulses/:pulse_number/jet-drops", wrapper.JetDropsByPulseNumber)
	router.GET("/api/v1/search", wrapper.Search)

}

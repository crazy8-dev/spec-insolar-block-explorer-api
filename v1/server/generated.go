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

	// JetDrop ID is concantination of jet_id and pulse_number.
	JetDropId *string `json:"jet_drop_id,omitempty"`

	// Jet ID.
	JetId *string `json:"jet_id,omitempty"`

	// Next jet_drop_id.
	NextJetDropId *[]string `json:"next_jet_drop_id,omitempty"`

	// Previous jet_drop_id.
	PrevJetDropId *[]string `json:"prev_jet_drop_id,omitempty"`

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

	// Total results.
	Total *int64 `json:"total,omitempty"`
}

// Pulse defines model for pulse.
type Pulse struct {

	// Pulse fullness status.
	IsComplete *bool `json:"is_complete,omitempty"`

	// Amount of all Jet Drop in the Pulse.
	JetDropAmount *int64 `json:"jet_drop_amount,omitempty"`

	// Next pulse number.
	NextPulseNumber *int64 `json:"next_pulse_number,omitempty"`

	// Previous pulse number.
	PrevPulseNumber *int64 `json:"prev_pulse_number,omitempty"`

	// Pulse number.
	PulseNumber *int64 `json:"pulse_number,omitempty"`

	// Number of all records in the Pulse.
	RecordAmount *int64 `json:"record_amount,omitempty"`

	// Unix timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// Pulses defines model for pulses.
type Pulses struct {
	Result *[]Pulse `json:"result,omitempty"`

	// Total results.
	Total *int64 `json:"total,omitempty"`
}

// Record defines model for record.
type Record struct {

	// Record hash.
	Hash *string `json:"hash,omitempty"`

	// JetDrop ID is concantination of jet_id and pulse_number.
	JetDropId *string `json:"jet_drop_id,omitempty"`

	// Jet ID.
	JetId *string `json:"jet_id,omitempty"`

	// Object reference.
	ObjectReference *string `json:"object_reference,omitempty"`

	// Order is the record order number in the Jet Drop.
	Order *int64 `json:"order,omitempty"`

	// Record payload.
	Payload *string `json:"payload,omitempty"`

	// Previous record reference.
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

	// Total results.
	Total *int64 `json:"total,omitempty"`
}

// SearchJetDrop defines model for search-jet-drop.
type SearchJetDrop struct {

	// Meta data.
	Meta *struct {

		// JetDrop ID.
		JetDropId *string `json:"jet_drop_id,omitempty"`
	} `json:"meta,omitempty"`

	// Result type.
	Type *string `json:"type,omitempty"`
}

// SearchJetid defines model for search-jetid.
type SearchJetid struct {

	// Meta data.
	Meta *struct {

		// Jet ID.
		JetId *string `json:"jet-id,omitempty"`
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

		// Pulse Number.
		PulseNumber *int64 `json:"pulse_number,omitempty"`
	} `json:"meta,omitempty"`

	// Result type.
	Type *string `json:"type,omitempty"`
}

// SearchRecord defines model for search-record.
type SearchRecord struct {

	// Meta data.
	Meta *struct {

		// Record reference.
		RecordRef *string `json:"record_ref,omitempty"`
	} `json:"meta,omitempty"`

	// Result type.
	Type *string `json:"type,omitempty"`
}

// FromJetDropAmountParam defines model for fromJetDropAmountParam.
type FromJetDropAmountParam int

// FromPulseNumberParam defines model for fromPulseNumberParam.
type FromPulseNumberParam int64

// JetDropIdPathParam defines model for jetDropIdPathParam.
type JetDropIdPathParam string

// JetIdPathParam defines model for jetIdPathParam.
type JetIdPathParam string

// LimitParam defines model for limitParam.
type LimitParam int

// ObjectReferencePathParam defines model for objectReferencePathParam.
type ObjectReferencePathParam string

// OffsetParam defines model for offsetParam.
type OffsetParam int

// PulseNumberPathParam defines model for pulseNumberPathParam.
type PulseNumberPathParam int64

// RecordTypeParam defines model for recordTypeParam.
type RecordTypeParam string

// ToJetDropAmountParam defines model for toJetDropAmountParam.
type ToJetDropAmountParam int

// ToPulseNumberParam defines model for toPulseNumberParam.
type ToPulseNumberParam int64

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

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The pagination starting point. Accepting pulse_number:order.
	FromItem *string `json:"from_item,omitempty"`

	// The record type.
	Type *RecordTypeParam `json:"type,omitempty"`
}

// JetDropsByJetIDParams defines parameters for JetDropsByJetID.
type JetDropsByJetIDParams struct {

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The pagination starting point. Accepting jet_id.
	FromItem *string `json:"from_item,omitempty"`

	// From which Pulse number.
	FromPulseNumber *FromPulseNumberParam `json:"from_pulse_number,omitempty"`

	// To which Pulse number.
	ToPulseNumber *ToPulseNumberParam `json:"to_pulse_number,omitempty"`
}

// ObjectLifelineParams defines parameters for ObjectLifeline.
type ObjectLifelineParams struct {

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The pagination starting point. Accepting pulse_number:order.
	FromItem *string `json:"from_item,omitempty"`

	// The pagination starting point. Accepting record_reference.
	RecordReference *string `json:"record_reference,omitempty"`

	// The record type.
	Type *RecordTypeParam `json:"type,omitempty"`
}

// PulsesParams defines parameters for Pulses.
type PulsesParams struct {

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The pagination starting point. Accepting pulse_number.
	FromItem *int64 `json:"from_item,omitempty"`

	// From which Pulse number.
	FromPulseNumber *FromPulseNumberParam `json:"from_pulse_number,omitempty"`

	// To which Pulse number.
	ToPulseNumber *ToPulseNumberParam `json:"to_pulse_number,omitempty"`

	// To which jet_drop_amount.
	FromJetDropAmount *FromJetDropAmountParam `json:"from_jet_drop_amount,omitempty"`

	// From which jet_drop_amount.
	ToJetDropAmount *ToJetDropAmountParam `json:"to_jet_drop_amount,omitempty"`
}

// JetDropsByPulseNumberParams defines parameters for JetDropsByPulseNumber.
type JetDropsByPulseNumberParams struct {

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The pagination starting point. Accepting jet_drop_id.
	FromItem *string `json:"from_item,omitempty"`
}

// SearchParams defines parameters for Search.
type SearchParams struct {

	// The searching value.
	Value string `json:"value"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Jet drop by ID
	// (GET /api/v1/jet-drops/{jet_drop_id})
	JetDropByID(ctx echo.Context, jetDropId JetDropIdPathParam) error
	// Jet drop records
	// (GET /api/v1/jet-drops/{jet_drop_id}/records)
	JetDropRecords(ctx echo.Context, jetDropId JetDropIdPathParam, params JetDropRecordsParams) error
	// Jet drops by jet ID
	// (GET /api/v1/jets/{jet_id}/jet-drops)
	JetDropsByJetID(ctx echo.Context, jetId JetIdPathParam, params JetDropsByJetIDParams) error
	// Object Lifeline
	// (GET /api/v1/lifeline/{object_reference}/records)
	ObjectLifeline(ctx echo.Context, objectReference ObjectReferencePathParam, params ObjectLifelineParams) error
	// Pulses
	// (GET /api/v1/pulses)
	Pulses(ctx echo.Context, params PulsesParams) error
	// Pulse
	// (GET /api/v1/pulses/{pulse_number})
	Pulse(ctx echo.Context, pulseNumber PulseNumberPathParam) error
	// Jet drops by pulse number
	// (GET /api/v1/pulses/{pulse_number}/jet-drops)
	JetDropsByPulseNumber(ctx echo.Context, pulseNumber PulseNumberPathParam, params JetDropsByPulseNumberParams) error
	// Search Data
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
	var jetDropId JetDropIdPathParam

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
	var jetDropId JetDropIdPathParam

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

	// ------------- Optional query parameter "from_item" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_item", ctx.QueryParams(), &params.FromItem)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_item: %s", err))
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
	var jetId JetIdPathParam

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

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "from_item" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_item", ctx.QueryParams(), &params.FromItem)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_item: %s", err))
	}

	// ------------- Optional query parameter "from_pulse_number" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_pulse_number", ctx.QueryParams(), &params.FromPulseNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_pulse_number: %s", err))
	}

	// ------------- Optional query parameter "to_pulse_number" -------------

	err = runtime.BindQueryParameter("form", true, false, "to_pulse_number", ctx.QueryParams(), &params.ToPulseNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter to_pulse_number: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.JetDropsByJetID(ctx, jetId, params)
	return err
}

// ObjectLifeline converts echo context to params.
func (w *ServerInterfaceWrapper) ObjectLifeline(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "object_reference" -------------
	var objectReference ObjectReferencePathParam

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

	// ------------- Optional query parameter "from_item" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_item", ctx.QueryParams(), &params.FromItem)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_item: %s", err))
	}

	// ------------- Optional query parameter "record_reference" -------------

	err = runtime.BindQueryParameter("form", true, false, "record_reference", ctx.QueryParams(), &params.RecordReference)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter record_reference: %s", err))
	}

	// ------------- Optional query parameter "type" -------------

	err = runtime.BindQueryParameter("form", true, false, "type", ctx.QueryParams(), &params.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter type: %s", err))
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

	// ------------- Optional query parameter "from_item" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_item", ctx.QueryParams(), &params.FromItem)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_item: %s", err))
	}

	// ------------- Optional query parameter "from_pulse_number" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_pulse_number", ctx.QueryParams(), &params.FromPulseNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_pulse_number: %s", err))
	}

	// ------------- Optional query parameter "to_pulse_number" -------------

	err = runtime.BindQueryParameter("form", true, false, "to_pulse_number", ctx.QueryParams(), &params.ToPulseNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter to_pulse_number: %s", err))
	}

	// ------------- Optional query parameter "from_jet_drop_amount" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_jet_drop_amount", ctx.QueryParams(), &params.FromJetDropAmount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_jet_drop_amount: %s", err))
	}

	// ------------- Optional query parameter "to_jet_drop_amount" -------------

	err = runtime.BindQueryParameter("form", true, false, "to_jet_drop_amount", ctx.QueryParams(), &params.ToJetDropAmount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter to_jet_drop_amount: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Pulses(ctx, params)
	return err
}

// Pulse converts echo context to params.
func (w *ServerInterfaceWrapper) Pulse(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pulse_number" -------------
	var pulseNumber PulseNumberPathParam

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
	var pulseNumber PulseNumberPathParam

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

	// ------------- Optional query parameter "from_item" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_item", ctx.QueryParams(), &params.FromItem)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_item: %s", err))
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

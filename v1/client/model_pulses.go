/*
 * Insolar Block Explorer API
 *
 * BE description 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// Pulses Responce codes.
type Pulses struct {
	// Total results.
	Total int64 `json:"total,omitempty"`
	Result []PulsesResponse200Result `json:"result,omitempty"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Link string `json:"link,omitempty"`
	ValidationFailures []PulsesResponse200ValidationFailures `json:"validation_failures,omitempty"`
}

/*
 * Swagger Petstore
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// InfoForASpecificPetResponse200 struct for InfoForASpecificPetResponse200
type InfoForASpecificPetResponse200 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
	Code int32 `json:"code"`
	Message string `json:"message"`
}

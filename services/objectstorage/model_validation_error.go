/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

// ValidationError struct for ValidationError
type ValidationError struct {
	// REQUIRED
	Loc *[]LocationInner `json:"loc"`
	// REQUIRED
	Msg *string `json:"msg"`
	// REQUIRED
	Type *string `json:"type"`
}

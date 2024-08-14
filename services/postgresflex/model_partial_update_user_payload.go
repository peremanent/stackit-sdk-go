/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

// PartialUpdateUserPayload struct for PartialUpdateUserPayload
type PartialUpdateUserPayload struct {
	Database *string   `json:"database,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
}

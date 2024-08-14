/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

// DNS struct for DNS
type DNS struct {
	// Enables the dns extension.
	// REQUIRED
	Enabled *bool `json:"enabled"`
	// Array of domain filters for externalDNS, e.g., *.runs.onstackit.cloud.
	Zones *[]string `json:"zones,omitempty"`
}

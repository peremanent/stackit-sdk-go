/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

type AreaPrefixConfigIPv4 struct {
	// The default prefix length for networks in the network area.
	DefaultPrefixLen *int64 `json:"defaultPrefixLen,omitempty"`
	// The maximal prefix length for networks in the network area.
	MaxPrefixLen *int64 `json:"maxPrefixLen,omitempty"`
	// The minimal prefix length for networks in the network area.
	MinPrefixLen *int64 `json:"minPrefixLen,omitempty"`
}

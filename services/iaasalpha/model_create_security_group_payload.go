/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

// CreateSecurityGroupPayload Object that represents a security group.
type CreateSecurityGroupPayload struct {
	// Description Object. Allows string up to 127 Characters.
	Description *string `json:"description,omitempty"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	Name *string `json:"name"`
	// A list containing security group rule objects.
	Rules *[]SecurityGroupRule `json:"rules,omitempty"`
	// Shows if a security group is stateful or stateless. You can only have one type of security groups per network interface/server.
	Stateful *bool `json:"stateful,omitempty"`
}

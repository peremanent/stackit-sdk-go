/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

import (
	"time"
)

type CreateAccessKeyPayload struct {
	// Expiration date. Null means never expires.
	Expires *time.Time `json:"expires,omitempty"`
}

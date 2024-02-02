/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

type DeleteBucketResponse struct {
	// Name of the bucket
	// REQUIRED
	Bucket *string `json:"bucket"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
}

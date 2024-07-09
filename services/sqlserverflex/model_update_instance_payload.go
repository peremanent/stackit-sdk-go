/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

type UpdateInstancePayload struct {
	// REQUIRED
	Acl *CreateInstancePayloadAcl `json:"acl"`
	// Cronjob for the daily full backup if not provided a job will generated between 00:00 and 04:59
	// REQUIRED
	BackupSchedule *string `json:"backupSchedule"`
	// Id of the selected flavor
	// REQUIRED
	FlavorId *string `json:"flavorId"`
	// REQUIRED
	Labels *map[string]interface{} `json:"labels"`
	// Name of the instance
	// REQUIRED
	Name *string `json:"name"`
	// Version of the MSSQL Server
	// REQUIRED
	Version *string `json:"version"`
}

/*
STACKIT Server Backup Management API

API endpoints for Server Backup Operations on STACKIT Servers.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverbackup

type ListBackupSchedules200Response struct {
	Items *[]BackupSchedule `json:"items,omitempty"`
}

/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

type InstanceParameters struct {
	EnableMonitoring     *bool     `json:"enable_monitoring,omitempty"`
	Graphite             *string   `json:"graphite,omitempty"`
	MaxDiskThreshold     *string   `json:"max_disk_threshold,omitempty"`
	MetricsFrequency     *int64    `json:"metrics_frequency,omitempty"`
	MetricsPrefix        *string   `json:"metrics_prefix,omitempty"`
	MonitoringInstanceId *string   `json:"monitoring_instance_id,omitempty"`
	Plugins              *[]string `json:"plugins,omitempty"`
	SgwAcl               *string   `json:"sgw_acl,omitempty"`
	Syslog               *[]string `json:"syslog,omitempty"`
}

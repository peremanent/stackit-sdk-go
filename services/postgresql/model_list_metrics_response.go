/*
STACKIT PostgreSQL API

The STACKIT PostgreSQL API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresql

type ListMetricsResponse struct {
	CpuIdleTime *int64 `json:"cpuIdleTime,omitempty"`
	// REQUIRED
	CpuLoadPercent *int64 `json:"cpuLoadPercent"`
	CpuSystemTime  *int64 `json:"cpuSystemTime,omitempty"`
	CpuUserTime    *int64 `json:"cpuUserTime,omitempty"`
	// REQUIRED
	DiskEphemeralTotal *int64 `json:"diskEphemeralTotal"`
	// REQUIRED
	DiskEphemeralUsed *int64 `json:"diskEphemeralUsed"`
	// REQUIRED
	DiskPersistentTotal *int64 `json:"diskPersistentTotal"`
	// REQUIRED
	DiskPersistentUsed *int64 `json:"diskPersistentUsed"`
	// REQUIRED
	MemoryTotal *int64 `json:"memoryTotal"`
	// REQUIRED
	MemoryUsed *int64 `json:"memoryUsed"`
	// REQUIRED
	ParachuteDiskEphemeralActivated *int64 `json:"parachuteDiskEphemeralActivated"`
	// REQUIRED
	ParachuteDiskEphemeralTotal *int64 `json:"parachuteDiskEphemeralTotal"`
	// REQUIRED
	ParachuteDiskEphemeralUsed *int64 `json:"parachuteDiskEphemeralUsed"`
	// REQUIRED
	ParachuteDiskEphemeralUsedPercent *int64 `json:"parachuteDiskEphemeralUsedPercent"`
	// REQUIRED
	ParachuteDiskEphemeralUsedThreshold *int64 `json:"parachuteDiskEphemeralUsedThreshold"`
	// REQUIRED
	ParachuteDiskPersistentActivated *int64 `json:"parachuteDiskPersistentActivated"`
	// REQUIRED
	ParachuteDiskPersistentTotal *int64 `json:"parachuteDiskPersistentTotal"`
	// REQUIRED
	ParachuteDiskPersistentUsed *int64 `json:"parachuteDiskPersistentUsed"`
	// REQUIRED
	ParachuteDiskPersistentUsedPercent *int64 `json:"parachuteDiskPersistentUsedPercent"`
	// REQUIRED
	ParachuteDiskPersistentUsedThreshold *int64 `json:"parachuteDiskPersistentUsedThreshold"`
}

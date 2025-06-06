/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lbapplication

import (
	"encoding/json"
)

// checks if the LoadbalancerOptionObservability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadbalancerOptionObservability{}

/*
	types and functions for logs
*/

// isModel
type LoadbalancerOptionObservabilityGetLogsAttributeType = *LoadbalancerOptionLogs
type LoadbalancerOptionObservabilityGetLogsArgType = LoadbalancerOptionLogs
type LoadbalancerOptionObservabilityGetLogsRetType = LoadbalancerOptionLogs

func getLoadbalancerOptionObservabilityGetLogsAttributeTypeOk(arg LoadbalancerOptionObservabilityGetLogsAttributeType) (ret LoadbalancerOptionObservabilityGetLogsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setLoadbalancerOptionObservabilityGetLogsAttributeType(arg *LoadbalancerOptionObservabilityGetLogsAttributeType, val LoadbalancerOptionObservabilityGetLogsRetType) {
	*arg = &val
}

/*
	types and functions for metrics
*/

// isModel
type LoadbalancerOptionObservabilityGetMetricsAttributeType = *LoadbalancerOptionMetrics
type LoadbalancerOptionObservabilityGetMetricsArgType = LoadbalancerOptionMetrics
type LoadbalancerOptionObservabilityGetMetricsRetType = LoadbalancerOptionMetrics

func getLoadbalancerOptionObservabilityGetMetricsAttributeTypeOk(arg LoadbalancerOptionObservabilityGetMetricsAttributeType) (ret LoadbalancerOptionObservabilityGetMetricsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setLoadbalancerOptionObservabilityGetMetricsAttributeType(arg *LoadbalancerOptionObservabilityGetMetricsAttributeType, val LoadbalancerOptionObservabilityGetMetricsRetType) {
	*arg = &val
}

// LoadbalancerOptionObservability We offer Application Load Balancer metrics observability via ARGUS or external solutions. Not changeable after creation.
type LoadbalancerOptionObservability struct {
	Logs    LoadbalancerOptionObservabilityGetLogsAttributeType    `json:"logs,omitempty"`
	Metrics LoadbalancerOptionObservabilityGetMetricsAttributeType `json:"metrics,omitempty"`
}

// NewLoadbalancerOptionObservability instantiates a new LoadbalancerOptionObservability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadbalancerOptionObservability() *LoadbalancerOptionObservability {
	this := LoadbalancerOptionObservability{}
	return &this
}

// NewLoadbalancerOptionObservabilityWithDefaults instantiates a new LoadbalancerOptionObservability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadbalancerOptionObservabilityWithDefaults() *LoadbalancerOptionObservability {
	this := LoadbalancerOptionObservability{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *LoadbalancerOptionObservability) GetLogs() (res LoadbalancerOptionObservabilityGetLogsRetType) {
	res, _ = o.GetLogsOk()
	return
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerOptionObservability) GetLogsOk() (ret LoadbalancerOptionObservabilityGetLogsRetType, ok bool) {
	return getLoadbalancerOptionObservabilityGetLogsAttributeTypeOk(o.Logs)
}

// HasLogs returns a boolean if a field has been set.
func (o *LoadbalancerOptionObservability) HasLogs() bool {
	_, ok := o.GetLogsOk()
	return ok
}

// SetLogs gets a reference to the given LoadbalancerOptionLogs and assigns it to the Logs field.
func (o *LoadbalancerOptionObservability) SetLogs(v LoadbalancerOptionObservabilityGetLogsRetType) {
	setLoadbalancerOptionObservabilityGetLogsAttributeType(&o.Logs, v)
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *LoadbalancerOptionObservability) GetMetrics() (res LoadbalancerOptionObservabilityGetMetricsRetType) {
	res, _ = o.GetMetricsOk()
	return
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerOptionObservability) GetMetricsOk() (ret LoadbalancerOptionObservabilityGetMetricsRetType, ok bool) {
	return getLoadbalancerOptionObservabilityGetMetricsAttributeTypeOk(o.Metrics)
}

// HasMetrics returns a boolean if a field has been set.
func (o *LoadbalancerOptionObservability) HasMetrics() bool {
	_, ok := o.GetMetricsOk()
	return ok
}

// SetMetrics gets a reference to the given LoadbalancerOptionMetrics and assigns it to the Metrics field.
func (o *LoadbalancerOptionObservability) SetMetrics(v LoadbalancerOptionObservabilityGetMetricsRetType) {
	setLoadbalancerOptionObservabilityGetMetricsAttributeType(&o.Metrics, v)
}

func (o LoadbalancerOptionObservability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getLoadbalancerOptionObservabilityGetLogsAttributeTypeOk(o.Logs); ok {
		toSerialize["Logs"] = val
	}
	if val, ok := getLoadbalancerOptionObservabilityGetMetricsAttributeTypeOk(o.Metrics); ok {
		toSerialize["Metrics"] = val
	}
	return toSerialize, nil
}

type NullableLoadbalancerOptionObservability struct {
	value *LoadbalancerOptionObservability
	isSet bool
}

func (v NullableLoadbalancerOptionObservability) Get() *LoadbalancerOptionObservability {
	return v.value
}

func (v *NullableLoadbalancerOptionObservability) Set(val *LoadbalancerOptionObservability) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadbalancerOptionObservability) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadbalancerOptionObservability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadbalancerOptionObservability(val *LoadbalancerOptionObservability) *NullableLoadbalancerOptionObservability {
	return &NullableLoadbalancerOptionObservability{value: val, isSet: true}
}

func (v NullableLoadbalancerOptionObservability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadbalancerOptionObservability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

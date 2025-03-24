/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the QuotaListVcpu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaListVcpu{}

/*
	types and functions for limit
*/

// isLong
type QuotaListVcpuGetLimitAttributeType = *int64
type QuotaListVcpuGetLimitArgType = int64
type QuotaListVcpuGetLimitRetType = int64

func getQuotaListVcpuGetLimitAttributeTypeOk(arg QuotaListVcpuGetLimitAttributeType) (ret QuotaListVcpuGetLimitRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListVcpuGetLimitAttributeType(arg *QuotaListVcpuGetLimitAttributeType, val QuotaListVcpuGetLimitRetType) {
	*arg = &val
}

/*
	types and functions for usage
*/

// isLong
type QuotaListVcpuGetUsageAttributeType = *int64
type QuotaListVcpuGetUsageArgType = int64
type QuotaListVcpuGetUsageRetType = int64

func getQuotaListVcpuGetUsageAttributeTypeOk(arg QuotaListVcpuGetUsageAttributeType) (ret QuotaListVcpuGetUsageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListVcpuGetUsageAttributeType(arg *QuotaListVcpuGetUsageAttributeType, val QuotaListVcpuGetUsageRetType) {
	*arg = &val
}

// QuotaListVcpu Number of server cores.
type QuotaListVcpu struct {
	// REQUIRED
	Limit QuotaListVcpuGetLimitAttributeType `json:"limit"`
	// REQUIRED
	Usage QuotaListVcpuGetUsageAttributeType `json:"usage"`
}

type _QuotaListVcpu QuotaListVcpu

// NewQuotaListVcpu instantiates a new QuotaListVcpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListVcpu(limit QuotaListVcpuGetLimitArgType, usage QuotaListVcpuGetUsageArgType) *QuotaListVcpu {
	this := QuotaListVcpu{}
	setQuotaListVcpuGetLimitAttributeType(&this.Limit, limit)
	setQuotaListVcpuGetUsageAttributeType(&this.Usage, usage)
	return &this
}

// NewQuotaListVcpuWithDefaults instantiates a new QuotaListVcpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListVcpuWithDefaults() *QuotaListVcpu {
	this := QuotaListVcpu{}
	return &this
}

// GetLimit returns the Limit field value
func (o *QuotaListVcpu) GetLimit() (ret QuotaListVcpuGetLimitRetType) {
	ret, _ = o.GetLimitOk()
	return ret
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *QuotaListVcpu) GetLimitOk() (ret QuotaListVcpuGetLimitRetType, ok bool) {
	return getQuotaListVcpuGetLimitAttributeTypeOk(o.Limit)
}

// SetLimit sets field value
func (o *QuotaListVcpu) SetLimit(v QuotaListVcpuGetLimitRetType) {
	setQuotaListVcpuGetLimitAttributeType(&o.Limit, v)
}

// GetUsage returns the Usage field value
func (o *QuotaListVcpu) GetUsage() (ret QuotaListVcpuGetUsageRetType) {
	ret, _ = o.GetUsageOk()
	return ret
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *QuotaListVcpu) GetUsageOk() (ret QuotaListVcpuGetUsageRetType, ok bool) {
	return getQuotaListVcpuGetUsageAttributeTypeOk(o.Usage)
}

// SetUsage sets field value
func (o *QuotaListVcpu) SetUsage(v QuotaListVcpuGetUsageRetType) {
	setQuotaListVcpuGetUsageAttributeType(&o.Usage, v)
}

func (o QuotaListVcpu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getQuotaListVcpuGetLimitAttributeTypeOk(o.Limit); ok {
		toSerialize["Limit"] = val
	}
	if val, ok := getQuotaListVcpuGetUsageAttributeTypeOk(o.Usage); ok {
		toSerialize["Usage"] = val
	}
	return toSerialize, nil
}

type NullableQuotaListVcpu struct {
	value *QuotaListVcpu
	isSet bool
}

func (v NullableQuotaListVcpu) Get() *QuotaListVcpu {
	return v.value
}

func (v *NullableQuotaListVcpu) Set(val *QuotaListVcpu) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaListVcpu) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaListVcpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaListVcpu(val *QuotaListVcpu) *NullableQuotaListVcpu {
	return &NullableQuotaListVcpu{value: val, isSet: true}
}

func (v NullableQuotaListVcpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaListVcpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

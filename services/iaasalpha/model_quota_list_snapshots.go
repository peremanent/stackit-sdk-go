/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the QuotaListSnapshots type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaListSnapshots{}

// QuotaListSnapshots Number of snapshots.
type QuotaListSnapshots struct {
	// REQUIRED
	Limit *int64 `json:"limit"`
	// REQUIRED
	Usage *int64 `json:"usage"`
}

type _QuotaListSnapshots QuotaListSnapshots

// NewQuotaListSnapshots instantiates a new QuotaListSnapshots object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListSnapshots(limit *int64, usage *int64) *QuotaListSnapshots {
	this := QuotaListSnapshots{}
	this.Limit = limit
	this.Usage = usage
	return &this
}

// NewQuotaListSnapshotsWithDefaults instantiates a new QuotaListSnapshots object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListSnapshotsWithDefaults() *QuotaListSnapshots {
	this := QuotaListSnapshots{}
	return &this
}

// GetLimit returns the Limit field value
func (o *QuotaListSnapshots) GetLimit() *int64 {
	if o == nil || IsNil(o.Limit) {
		var ret *int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *QuotaListSnapshots) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit, true
}

// SetLimit sets field value
func (o *QuotaListSnapshots) SetLimit(v *int64) {
	o.Limit = v
}

// GetUsage returns the Usage field value
func (o *QuotaListSnapshots) GetUsage() *int64 {
	if o == nil || IsNil(o.Usage) {
		var ret *int64
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *QuotaListSnapshots) GetUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage, true
}

// SetUsage sets field value
func (o *QuotaListSnapshots) SetUsage(v *int64) {
	o.Usage = v
}

func (o QuotaListSnapshots) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

type NullableQuotaListSnapshots struct {
	value *QuotaListSnapshots
	isSet bool
}

func (v NullableQuotaListSnapshots) Get() *QuotaListSnapshots {
	return v.value
}

func (v *NullableQuotaListSnapshots) Set(val *QuotaListSnapshots) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaListSnapshots) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaListSnapshots) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaListSnapshots(val *QuotaListSnapshots) *NullableQuotaListSnapshots {
	return &NullableQuotaListSnapshots{value: val, isSet: true}
}

func (v NullableQuotaListSnapshots) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaListSnapshots) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

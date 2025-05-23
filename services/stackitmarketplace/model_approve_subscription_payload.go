/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
)

// checks if the ApproveSubscriptionPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApproveSubscriptionPayload{}

/*
	types and functions for instanceTarget
*/

// isNotNullableString
type ApproveSubscriptionPayloadGetInstanceTargetAttributeType = *string

func getApproveSubscriptionPayloadGetInstanceTargetAttributeTypeOk(arg ApproveSubscriptionPayloadGetInstanceTargetAttributeType) (ret ApproveSubscriptionPayloadGetInstanceTargetRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setApproveSubscriptionPayloadGetInstanceTargetAttributeType(arg *ApproveSubscriptionPayloadGetInstanceTargetAttributeType, val ApproveSubscriptionPayloadGetInstanceTargetRetType) {
	*arg = &val
}

type ApproveSubscriptionPayloadGetInstanceTargetArgType = string
type ApproveSubscriptionPayloadGetInstanceTargetRetType = string

// ApproveSubscriptionPayload struct for ApproveSubscriptionPayload
type ApproveSubscriptionPayload struct {
	// Uniform Resource Locator.
	InstanceTarget ApproveSubscriptionPayloadGetInstanceTargetAttributeType `json:"instanceTarget,omitempty"`
}

// NewApproveSubscriptionPayload instantiates a new ApproveSubscriptionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproveSubscriptionPayload() *ApproveSubscriptionPayload {
	this := ApproveSubscriptionPayload{}
	return &this
}

// NewApproveSubscriptionPayloadWithDefaults instantiates a new ApproveSubscriptionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproveSubscriptionPayloadWithDefaults() *ApproveSubscriptionPayload {
	this := ApproveSubscriptionPayload{}
	return &this
}

// GetInstanceTarget returns the InstanceTarget field value if set, zero value otherwise.
func (o *ApproveSubscriptionPayload) GetInstanceTarget() (res ApproveSubscriptionPayloadGetInstanceTargetRetType) {
	res, _ = o.GetInstanceTargetOk()
	return
}

// GetInstanceTargetOk returns a tuple with the InstanceTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproveSubscriptionPayload) GetInstanceTargetOk() (ret ApproveSubscriptionPayloadGetInstanceTargetRetType, ok bool) {
	return getApproveSubscriptionPayloadGetInstanceTargetAttributeTypeOk(o.InstanceTarget)
}

// HasInstanceTarget returns a boolean if a field has been set.
func (o *ApproveSubscriptionPayload) HasInstanceTarget() bool {
	_, ok := o.GetInstanceTargetOk()
	return ok
}

// SetInstanceTarget gets a reference to the given string and assigns it to the InstanceTarget field.
func (o *ApproveSubscriptionPayload) SetInstanceTarget(v ApproveSubscriptionPayloadGetInstanceTargetRetType) {
	setApproveSubscriptionPayloadGetInstanceTargetAttributeType(&o.InstanceTarget, v)
}

func (o ApproveSubscriptionPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getApproveSubscriptionPayloadGetInstanceTargetAttributeTypeOk(o.InstanceTarget); ok {
		toSerialize["InstanceTarget"] = val
	}
	return toSerialize, nil
}

type NullableApproveSubscriptionPayload struct {
	value *ApproveSubscriptionPayload
	isSet bool
}

func (v NullableApproveSubscriptionPayload) Get() *ApproveSubscriptionPayload {
	return v.value
}

func (v *NullableApproveSubscriptionPayload) Set(val *ApproveSubscriptionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApproveSubscriptionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApproveSubscriptionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproveSubscriptionPayload(val *ApproveSubscriptionPayload) *NullableApproveSubscriptionPayload {
	return &NullableApproveSubscriptionPayload{value: val, isSet: true}
}

func (v NullableApproveSubscriptionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproveSubscriptionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

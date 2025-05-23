/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the PutCustomDomainPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutCustomDomainPayload{}

/*
	types and functions for intentId
*/

// isNotNullableString
type PutCustomDomainPayloadGetIntentIdAttributeType = *string

func getPutCustomDomainPayloadGetIntentIdAttributeTypeOk(arg PutCustomDomainPayloadGetIntentIdAttributeType) (ret PutCustomDomainPayloadGetIntentIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPutCustomDomainPayloadGetIntentIdAttributeType(arg *PutCustomDomainPayloadGetIntentIdAttributeType, val PutCustomDomainPayloadGetIntentIdRetType) {
	*arg = &val
}

type PutCustomDomainPayloadGetIntentIdArgType = string
type PutCustomDomainPayloadGetIntentIdRetType = string

// PutCustomDomainPayload struct for PutCustomDomainPayload
type PutCustomDomainPayload struct {
	// While optional, it is greatly encouraged to provide an `intentId`.  This is used to deduplicate requests.   If multiple modifying Requests with the same `intentId` for a given `projectId` are received, all but the first request are dropped.
	IntentId PutCustomDomainPayloadGetIntentIdAttributeType `json:"intentId,omitempty"`
}

// NewPutCustomDomainPayload instantiates a new PutCustomDomainPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutCustomDomainPayload() *PutCustomDomainPayload {
	this := PutCustomDomainPayload{}
	return &this
}

// NewPutCustomDomainPayloadWithDefaults instantiates a new PutCustomDomainPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutCustomDomainPayloadWithDefaults() *PutCustomDomainPayload {
	this := PutCustomDomainPayload{}
	return &this
}

// GetIntentId returns the IntentId field value if set, zero value otherwise.
func (o *PutCustomDomainPayload) GetIntentId() (res PutCustomDomainPayloadGetIntentIdRetType) {
	res, _ = o.GetIntentIdOk()
	return
}

// GetIntentIdOk returns a tuple with the IntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutCustomDomainPayload) GetIntentIdOk() (ret PutCustomDomainPayloadGetIntentIdRetType, ok bool) {
	return getPutCustomDomainPayloadGetIntentIdAttributeTypeOk(o.IntentId)
}

// HasIntentId returns a boolean if a field has been set.
func (o *PutCustomDomainPayload) HasIntentId() bool {
	_, ok := o.GetIntentIdOk()
	return ok
}

// SetIntentId gets a reference to the given string and assigns it to the IntentId field.
func (o *PutCustomDomainPayload) SetIntentId(v PutCustomDomainPayloadGetIntentIdRetType) {
	setPutCustomDomainPayloadGetIntentIdAttributeType(&o.IntentId, v)
}

func (o PutCustomDomainPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPutCustomDomainPayloadGetIntentIdAttributeTypeOk(o.IntentId); ok {
		toSerialize["IntentId"] = val
	}
	return toSerialize, nil
}

type NullablePutCustomDomainPayload struct {
	value *PutCustomDomainPayload
	isSet bool
}

func (v NullablePutCustomDomainPayload) Get() *PutCustomDomainPayload {
	return v.value
}

func (v *NullablePutCustomDomainPayload) Set(val *PutCustomDomainPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePutCustomDomainPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePutCustomDomainPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutCustomDomainPayload(val *PutCustomDomainPayload) *NullablePutCustomDomainPayload {
	return &NullablePutCustomDomainPayload{value: val, isSet: true}
}

func (v NullablePutCustomDomainPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutCustomDomainPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

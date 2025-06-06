/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

import (
	"encoding/json"
)

// checks if the UpdateBackupsConfigPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBackupsConfigPayload{}

/*
	types and functions for encryption_key
*/

// isNotNullableString
type UpdateBackupsConfigPayloadGetEncryptionKeyAttributeType = *string

func getUpdateBackupsConfigPayloadGetEncryptionKeyAttributeTypeOk(arg UpdateBackupsConfigPayloadGetEncryptionKeyAttributeType) (ret UpdateBackupsConfigPayloadGetEncryptionKeyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateBackupsConfigPayloadGetEncryptionKeyAttributeType(arg *UpdateBackupsConfigPayloadGetEncryptionKeyAttributeType, val UpdateBackupsConfigPayloadGetEncryptionKeyRetType) {
	*arg = &val
}

type UpdateBackupsConfigPayloadGetEncryptionKeyArgType = string
type UpdateBackupsConfigPayloadGetEncryptionKeyRetType = string

// UpdateBackupsConfigPayload struct for UpdateBackupsConfigPayload
type UpdateBackupsConfigPayload struct {
	EncryptionKey UpdateBackupsConfigPayloadGetEncryptionKeyAttributeType `json:"encryption_key,omitempty"`
}

// NewUpdateBackupsConfigPayload instantiates a new UpdateBackupsConfigPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBackupsConfigPayload() *UpdateBackupsConfigPayload {
	this := UpdateBackupsConfigPayload{}
	return &this
}

// NewUpdateBackupsConfigPayloadWithDefaults instantiates a new UpdateBackupsConfigPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBackupsConfigPayloadWithDefaults() *UpdateBackupsConfigPayload {
	this := UpdateBackupsConfigPayload{}
	return &this
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *UpdateBackupsConfigPayload) GetEncryptionKey() (res UpdateBackupsConfigPayloadGetEncryptionKeyRetType) {
	res, _ = o.GetEncryptionKeyOk()
	return
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupsConfigPayload) GetEncryptionKeyOk() (ret UpdateBackupsConfigPayloadGetEncryptionKeyRetType, ok bool) {
	return getUpdateBackupsConfigPayloadGetEncryptionKeyAttributeTypeOk(o.EncryptionKey)
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *UpdateBackupsConfigPayload) HasEncryptionKey() bool {
	_, ok := o.GetEncryptionKeyOk()
	return ok
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *UpdateBackupsConfigPayload) SetEncryptionKey(v UpdateBackupsConfigPayloadGetEncryptionKeyRetType) {
	setUpdateBackupsConfigPayloadGetEncryptionKeyAttributeType(&o.EncryptionKey, v)
}

func (o UpdateBackupsConfigPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateBackupsConfigPayloadGetEncryptionKeyAttributeTypeOk(o.EncryptionKey); ok {
		toSerialize["EncryptionKey"] = val
	}
	return toSerialize, nil
}

type NullableUpdateBackupsConfigPayload struct {
	value *UpdateBackupsConfigPayload
	isSet bool
}

func (v NullableUpdateBackupsConfigPayload) Get() *UpdateBackupsConfigPayload {
	return v.value
}

func (v *NullableUpdateBackupsConfigPayload) Set(val *UpdateBackupsConfigPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBackupsConfigPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBackupsConfigPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBackupsConfigPayload(val *UpdateBackupsConfigPayload) *NullableUpdateBackupsConfigPayload {
	return &NullableUpdateBackupsConfigPayload{value: val, isSet: true}
}

func (v NullableUpdateBackupsConfigPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBackupsConfigPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the RawCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RawCredentials{}

/*
	types and functions for credentials
*/

// isModel
type RawCredentialsGetCredentialsAttributeType = *Credentials
type RawCredentialsGetCredentialsArgType = Credentials
type RawCredentialsGetCredentialsRetType = Credentials

func getRawCredentialsGetCredentialsAttributeTypeOk(arg RawCredentialsGetCredentialsAttributeType) (ret RawCredentialsGetCredentialsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRawCredentialsGetCredentialsAttributeType(arg *RawCredentialsGetCredentialsAttributeType, val RawCredentialsGetCredentialsRetType) {
	*arg = &val
}

// RawCredentials struct for RawCredentials
type RawCredentials struct {
	// REQUIRED
	Credentials RawCredentialsGetCredentialsAttributeType `json:"credentials"`
}

type _RawCredentials RawCredentials

// NewRawCredentials instantiates a new RawCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawCredentials(credentials RawCredentialsGetCredentialsArgType) *RawCredentials {
	this := RawCredentials{}
	setRawCredentialsGetCredentialsAttributeType(&this.Credentials, credentials)
	return &this
}

// NewRawCredentialsWithDefaults instantiates a new RawCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawCredentialsWithDefaults() *RawCredentials {
	this := RawCredentials{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *RawCredentials) GetCredentials() (ret RawCredentialsGetCredentialsRetType) {
	ret, _ = o.GetCredentialsOk()
	return ret
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *RawCredentials) GetCredentialsOk() (ret RawCredentialsGetCredentialsRetType, ok bool) {
	return getRawCredentialsGetCredentialsAttributeTypeOk(o.Credentials)
}

// SetCredentials sets field value
func (o *RawCredentials) SetCredentials(v RawCredentialsGetCredentialsRetType) {
	setRawCredentialsGetCredentialsAttributeType(&o.Credentials, v)
}

func (o RawCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getRawCredentialsGetCredentialsAttributeTypeOk(o.Credentials); ok {
		toSerialize["Credentials"] = val
	}
	return toSerialize, nil
}

type NullableRawCredentials struct {
	value *RawCredentials
	isSet bool
}

func (v NullableRawCredentials) Get() *RawCredentials {
	return v.value
}

func (v *NullableRawCredentials) Set(val *RawCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableRawCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableRawCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawCredentials(val *RawCredentials) *NullableRawCredentials {
	return &NullableRawCredentials{value: val, isSet: true}
}

func (v NullableRawCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

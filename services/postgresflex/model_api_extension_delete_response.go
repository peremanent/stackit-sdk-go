/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the ApiExtensionDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiExtensionDeleteResponse{}

/*
	types and functions for isSucceded
*/

// isBoolean
type ApiExtensionDeleteResponsegetIsSuccededAttributeType = *bool
type ApiExtensionDeleteResponsegetIsSuccededArgType = bool
type ApiExtensionDeleteResponsegetIsSuccededRetType = bool

func getApiExtensionDeleteResponsegetIsSuccededAttributeTypeOk(arg ApiExtensionDeleteResponsegetIsSuccededAttributeType) (ret ApiExtensionDeleteResponsegetIsSuccededRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setApiExtensionDeleteResponsegetIsSuccededAttributeType(arg *ApiExtensionDeleteResponsegetIsSuccededAttributeType, val ApiExtensionDeleteResponsegetIsSuccededRetType) {
	*arg = &val
}

// ApiExtensionDeleteResponse struct for ApiExtensionDeleteResponse
type ApiExtensionDeleteResponse struct {
	IsSucceded ApiExtensionDeleteResponsegetIsSuccededAttributeType `json:"isSucceded,omitempty"`
}

// NewApiExtensionDeleteResponse instantiates a new ApiExtensionDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiExtensionDeleteResponse() *ApiExtensionDeleteResponse {
	this := ApiExtensionDeleteResponse{}
	return &this
}

// NewApiExtensionDeleteResponseWithDefaults instantiates a new ApiExtensionDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiExtensionDeleteResponseWithDefaults() *ApiExtensionDeleteResponse {
	this := ApiExtensionDeleteResponse{}
	return &this
}

// GetIsSucceded returns the IsSucceded field value if set, zero value otherwise.
func (o *ApiExtensionDeleteResponse) GetIsSucceded() (res ApiExtensionDeleteResponsegetIsSuccededRetType) {
	res, _ = o.GetIsSuccededOk()
	return
}

// GetIsSuccededOk returns a tuple with the IsSucceded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExtensionDeleteResponse) GetIsSuccededOk() (ret ApiExtensionDeleteResponsegetIsSuccededRetType, ok bool) {
	return getApiExtensionDeleteResponsegetIsSuccededAttributeTypeOk(o.IsSucceded)
}

// HasIsSucceded returns a boolean if a field has been set.
func (o *ApiExtensionDeleteResponse) HasIsSucceded() bool {
	_, ok := o.GetIsSuccededOk()
	return ok
}

// SetIsSucceded gets a reference to the given bool and assigns it to the IsSucceded field.
func (o *ApiExtensionDeleteResponse) SetIsSucceded(v ApiExtensionDeleteResponsegetIsSuccededRetType) {
	setApiExtensionDeleteResponsegetIsSuccededAttributeType(&o.IsSucceded, v)
}

func (o ApiExtensionDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getApiExtensionDeleteResponsegetIsSuccededAttributeTypeOk(o.IsSucceded); ok {
		toSerialize["IsSucceded"] = val
	}
	return toSerialize, nil
}

type NullableApiExtensionDeleteResponse struct {
	value *ApiExtensionDeleteResponse
	isSet bool
}

func (v NullableApiExtensionDeleteResponse) Get() *ApiExtensionDeleteResponse {
	return v.value
}

func (v *NullableApiExtensionDeleteResponse) Set(val *ApiExtensionDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiExtensionDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiExtensionDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiExtensionDeleteResponse(val *ApiExtensionDeleteResponse) *NullableApiExtensionDeleteResponse {
	return &NullableApiExtensionDeleteResponse{value: val, isSet: true}
}

func (v NullableApiExtensionDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiExtensionDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

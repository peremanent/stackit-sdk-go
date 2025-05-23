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

// checks if the ApiConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiConfiguration{}

/*
	types and functions for name
*/

// isNotNullableString
type ApiConfigurationGetNameAttributeType = *string

func getApiConfigurationGetNameAttributeTypeOk(arg ApiConfigurationGetNameAttributeType) (ret ApiConfigurationGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setApiConfigurationGetNameAttributeType(arg *ApiConfigurationGetNameAttributeType, val ApiConfigurationGetNameRetType) {
	*arg = &val
}

type ApiConfigurationGetNameArgType = string
type ApiConfigurationGetNameRetType = string

/*
	types and functions for setting
*/

// isNotNullableString
type ApiConfigurationGetSettingAttributeType = *string

func getApiConfigurationGetSettingAttributeTypeOk(arg ApiConfigurationGetSettingAttributeType) (ret ApiConfigurationGetSettingRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setApiConfigurationGetSettingAttributeType(arg *ApiConfigurationGetSettingAttributeType, val ApiConfigurationGetSettingRetType) {
	*arg = &val
}

type ApiConfigurationGetSettingArgType = string
type ApiConfigurationGetSettingRetType = string

// ApiConfiguration struct for ApiConfiguration
type ApiConfiguration struct {
	Name    ApiConfigurationGetNameAttributeType    `json:"name,omitempty"`
	Setting ApiConfigurationGetSettingAttributeType `json:"setting,omitempty"`
}

// NewApiConfiguration instantiates a new ApiConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiConfiguration() *ApiConfiguration {
	this := ApiConfiguration{}
	return &this
}

// NewApiConfigurationWithDefaults instantiates a new ApiConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiConfigurationWithDefaults() *ApiConfiguration {
	this := ApiConfiguration{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiConfiguration) GetName() (res ApiConfigurationGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfiguration) GetNameOk() (ret ApiConfigurationGetNameRetType, ok bool) {
	return getApiConfigurationGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *ApiConfiguration) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiConfiguration) SetName(v ApiConfigurationGetNameRetType) {
	setApiConfigurationGetNameAttributeType(&o.Name, v)
}

// GetSetting returns the Setting field value if set, zero value otherwise.
func (o *ApiConfiguration) GetSetting() (res ApiConfigurationGetSettingRetType) {
	res, _ = o.GetSettingOk()
	return
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfiguration) GetSettingOk() (ret ApiConfigurationGetSettingRetType, ok bool) {
	return getApiConfigurationGetSettingAttributeTypeOk(o.Setting)
}

// HasSetting returns a boolean if a field has been set.
func (o *ApiConfiguration) HasSetting() bool {
	_, ok := o.GetSettingOk()
	return ok
}

// SetSetting gets a reference to the given string and assigns it to the Setting field.
func (o *ApiConfiguration) SetSetting(v ApiConfigurationGetSettingRetType) {
	setApiConfigurationGetSettingAttributeType(&o.Setting, v)
}

func (o ApiConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getApiConfigurationGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getApiConfigurationGetSettingAttributeTypeOk(o.Setting); ok {
		toSerialize["Setting"] = val
	}
	return toSerialize, nil
}

type NullableApiConfiguration struct {
	value *ApiConfiguration
	isSet bool
}

func (v NullableApiConfiguration) Get() *ApiConfiguration {
	return v.value
}

func (v *NullableApiConfiguration) Set(val *ApiConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableApiConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableApiConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiConfiguration(val *ApiConfiguration) *NullableApiConfiguration {
	return &NullableApiConfiguration{value: val, isSet: true}
}

func (v NullableApiConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
STACKIT Run Commands Service API

API endpoints for the STACKIT Run Commands Service API

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package runcommand

import (
	"encoding/json"
)

// checks if the Properties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Properties{}

/*
	types and functions for ConfirmPassword
*/

// isModel
type PropertiesGetConfirmPasswordAttributeType = *Field
type PropertiesGetConfirmPasswordArgType = Field
type PropertiesGetConfirmPasswordRetType = Field

func getPropertiesGetConfirmPasswordAttributeTypeOk(arg PropertiesGetConfirmPasswordAttributeType) (ret PropertiesGetConfirmPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPropertiesGetConfirmPasswordAttributeType(arg *PropertiesGetConfirmPasswordAttributeType, val PropertiesGetConfirmPasswordRetType) {
	*arg = &val
}

/*
	types and functions for Password
*/

// isModel
type PropertiesGetPasswordAttributeType = *Field
type PropertiesGetPasswordArgType = Field
type PropertiesGetPasswordRetType = Field

func getPropertiesGetPasswordAttributeTypeOk(arg PropertiesGetPasswordAttributeType) (ret PropertiesGetPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPropertiesGetPasswordAttributeType(arg *PropertiesGetPasswordAttributeType, val PropertiesGetPasswordRetType) {
	*arg = &val
}

/*
	types and functions for Script
*/

// isModel
type PropertiesGetScriptAttributeType = *Field
type PropertiesGetScriptArgType = Field
type PropertiesGetScriptRetType = Field

func getPropertiesGetScriptAttributeTypeOk(arg PropertiesGetScriptAttributeType) (ret PropertiesGetScriptRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPropertiesGetScriptAttributeType(arg *PropertiesGetScriptAttributeType, val PropertiesGetScriptRetType) {
	*arg = &val
}

/*
	types and functions for Username
*/

// isModel
type PropertiesGetUsernameAttributeType = *Field
type PropertiesGetUsernameArgType = Field
type PropertiesGetUsernameRetType = Field

func getPropertiesGetUsernameAttributeTypeOk(arg PropertiesGetUsernameAttributeType) (ret PropertiesGetUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPropertiesGetUsernameAttributeType(arg *PropertiesGetUsernameAttributeType, val PropertiesGetUsernameRetType) {
	*arg = &val
}

/*
	types and functions for required
*/

// isArray
type PropertiesGetRequiredAttributeType = *[]string
type PropertiesGetRequiredArgType = []string
type PropertiesGetRequiredRetType = []string

func getPropertiesGetRequiredAttributeTypeOk(arg PropertiesGetRequiredAttributeType) (ret PropertiesGetRequiredRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPropertiesGetRequiredAttributeType(arg *PropertiesGetRequiredAttributeType, val PropertiesGetRequiredRetType) {
	*arg = &val
}

/*
	types and functions for type
*/

// isNotNullableString
type PropertiesGetTypeAttributeType = *string

func getPropertiesGetTypeAttributeTypeOk(arg PropertiesGetTypeAttributeType) (ret PropertiesGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPropertiesGetTypeAttributeType(arg *PropertiesGetTypeAttributeType, val PropertiesGetTypeRetType) {
	*arg = &val
}

type PropertiesGetTypeArgType = string
type PropertiesGetTypeRetType = string

// Properties struct for Properties
type Properties struct {
	ConfirmPassword PropertiesGetConfirmPasswordAttributeType `json:"ConfirmPassword,omitempty"`
	Password        PropertiesGetPasswordAttributeType        `json:"Password,omitempty"`
	Script          PropertiesGetScriptAttributeType          `json:"Script,omitempty"`
	Username        PropertiesGetUsernameAttributeType        `json:"Username,omitempty"`
	Required        PropertiesGetRequiredAttributeType        `json:"required,omitempty"`
	Type            PropertiesGetTypeAttributeType            `json:"type,omitempty"`
}

// NewProperties instantiates a new Properties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProperties() *Properties {
	this := Properties{}
	return &this
}

// NewPropertiesWithDefaults instantiates a new Properties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertiesWithDefaults() *Properties {
	this := Properties{}
	return &this
}

// GetConfirmPassword returns the ConfirmPassword field value if set, zero value otherwise.
func (o *Properties) GetConfirmPassword() (res PropertiesGetConfirmPasswordRetType) {
	res, _ = o.GetConfirmPasswordOk()
	return
}

// GetConfirmPasswordOk returns a tuple with the ConfirmPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetConfirmPasswordOk() (ret PropertiesGetConfirmPasswordRetType, ok bool) {
	return getPropertiesGetConfirmPasswordAttributeTypeOk(o.ConfirmPassword)
}

// HasConfirmPassword returns a boolean if a field has been set.
func (o *Properties) HasConfirmPassword() bool {
	_, ok := o.GetConfirmPasswordOk()
	return ok
}

// SetConfirmPassword gets a reference to the given Field and assigns it to the ConfirmPassword field.
func (o *Properties) SetConfirmPassword(v PropertiesGetConfirmPasswordRetType) {
	setPropertiesGetConfirmPasswordAttributeType(&o.ConfirmPassword, v)
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Properties) GetPassword() (res PropertiesGetPasswordRetType) {
	res, _ = o.GetPasswordOk()
	return
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetPasswordOk() (ret PropertiesGetPasswordRetType, ok bool) {
	return getPropertiesGetPasswordAttributeTypeOk(o.Password)
}

// HasPassword returns a boolean if a field has been set.
func (o *Properties) HasPassword() bool {
	_, ok := o.GetPasswordOk()
	return ok
}

// SetPassword gets a reference to the given Field and assigns it to the Password field.
func (o *Properties) SetPassword(v PropertiesGetPasswordRetType) {
	setPropertiesGetPasswordAttributeType(&o.Password, v)
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *Properties) GetScript() (res PropertiesGetScriptRetType) {
	res, _ = o.GetScriptOk()
	return
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetScriptOk() (ret PropertiesGetScriptRetType, ok bool) {
	return getPropertiesGetScriptAttributeTypeOk(o.Script)
}

// HasScript returns a boolean if a field has been set.
func (o *Properties) HasScript() bool {
	_, ok := o.GetScriptOk()
	return ok
}

// SetScript gets a reference to the given Field and assigns it to the Script field.
func (o *Properties) SetScript(v PropertiesGetScriptRetType) {
	setPropertiesGetScriptAttributeType(&o.Script, v)
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Properties) GetUsername() (res PropertiesGetUsernameRetType) {
	res, _ = o.GetUsernameOk()
	return
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetUsernameOk() (ret PropertiesGetUsernameRetType, ok bool) {
	return getPropertiesGetUsernameAttributeTypeOk(o.Username)
}

// HasUsername returns a boolean if a field has been set.
func (o *Properties) HasUsername() bool {
	_, ok := o.GetUsernameOk()
	return ok
}

// SetUsername gets a reference to the given Field and assigns it to the Username field.
func (o *Properties) SetUsername(v PropertiesGetUsernameRetType) {
	setPropertiesGetUsernameAttributeType(&o.Username, v)
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *Properties) GetRequired() (res PropertiesGetRequiredRetType) {
	res, _ = o.GetRequiredOk()
	return
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetRequiredOk() (ret PropertiesGetRequiredRetType, ok bool) {
	return getPropertiesGetRequiredAttributeTypeOk(o.Required)
}

// HasRequired returns a boolean if a field has been set.
func (o *Properties) HasRequired() bool {
	_, ok := o.GetRequiredOk()
	return ok
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *Properties) SetRequired(v PropertiesGetRequiredRetType) {
	setPropertiesGetRequiredAttributeType(&o.Required, v)
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Properties) GetType() (res PropertiesGetTypeRetType) {
	res, _ = o.GetTypeOk()
	return
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetTypeOk() (ret PropertiesGetTypeRetType, ok bool) {
	return getPropertiesGetTypeAttributeTypeOk(o.Type)
}

// HasType returns a boolean if a field has been set.
func (o *Properties) HasType() bool {
	_, ok := o.GetTypeOk()
	return ok
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Properties) SetType(v PropertiesGetTypeRetType) {
	setPropertiesGetTypeAttributeType(&o.Type, v)
}

func (o Properties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPropertiesGetConfirmPasswordAttributeTypeOk(o.ConfirmPassword); ok {
		toSerialize["ConfirmPassword"] = val
	}
	if val, ok := getPropertiesGetPasswordAttributeTypeOk(o.Password); ok {
		toSerialize["Password"] = val
	}
	if val, ok := getPropertiesGetScriptAttributeTypeOk(o.Script); ok {
		toSerialize["Script"] = val
	}
	if val, ok := getPropertiesGetUsernameAttributeTypeOk(o.Username); ok {
		toSerialize["Username"] = val
	}
	if val, ok := getPropertiesGetRequiredAttributeTypeOk(o.Required); ok {
		toSerialize["Required"] = val
	}
	if val, ok := getPropertiesGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	return toSerialize, nil
}

type NullableProperties struct {
	value *Properties
	isSet bool
}

func (v NullableProperties) Get() *Properties {
	return v.value
}

func (v *NullableProperties) Set(val *Properties) {
	v.value = val
	v.isSet = true
}

func (v NullableProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProperties(val *Properties) *NullableProperties {
	return &NullableProperties{value: val, isSet: true}
}

func (v NullableProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

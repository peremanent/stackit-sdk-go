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

// checks if the UpdateCredentialsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCredentialsPayload{}

/*
	types and functions for displayName
*/

// isNotNullableString
type UpdateCredentialsPayloadGetDisplayNameAttributeType = *string

func getUpdateCredentialsPayloadGetDisplayNameAttributeTypeOk(arg UpdateCredentialsPayloadGetDisplayNameAttributeType) (ret UpdateCredentialsPayloadGetDisplayNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateCredentialsPayloadGetDisplayNameAttributeType(arg *UpdateCredentialsPayloadGetDisplayNameAttributeType, val UpdateCredentialsPayloadGetDisplayNameRetType) {
	*arg = &val
}

type UpdateCredentialsPayloadGetDisplayNameArgType = string
type UpdateCredentialsPayloadGetDisplayNameRetType = string

/*
	types and functions for password
*/

// isNotNullableString
type UpdateCredentialsPayloadGetPasswordAttributeType = *string

func getUpdateCredentialsPayloadGetPasswordAttributeTypeOk(arg UpdateCredentialsPayloadGetPasswordAttributeType) (ret UpdateCredentialsPayloadGetPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateCredentialsPayloadGetPasswordAttributeType(arg *UpdateCredentialsPayloadGetPasswordAttributeType, val UpdateCredentialsPayloadGetPasswordRetType) {
	*arg = &val
}

type UpdateCredentialsPayloadGetPasswordArgType = string
type UpdateCredentialsPayloadGetPasswordRetType = string

/*
	types and functions for username
*/

// isNotNullableString
type UpdateCredentialsPayloadGetUsernameAttributeType = *string

func getUpdateCredentialsPayloadGetUsernameAttributeTypeOk(arg UpdateCredentialsPayloadGetUsernameAttributeType) (ret UpdateCredentialsPayloadGetUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateCredentialsPayloadGetUsernameAttributeType(arg *UpdateCredentialsPayloadGetUsernameAttributeType, val UpdateCredentialsPayloadGetUsernameRetType) {
	*arg = &val
}

type UpdateCredentialsPayloadGetUsernameArgType = string
type UpdateCredentialsPayloadGetUsernameRetType = string

// UpdateCredentialsPayload struct for UpdateCredentialsPayload
type UpdateCredentialsPayload struct {
	// Credential name
	DisplayName UpdateCredentialsPayloadGetDisplayNameAttributeType `json:"displayName,omitempty"`
	// A valid password used for an existing ARGUS instance, which is used during basic auth.
	Password UpdateCredentialsPayloadGetPasswordAttributeType `json:"password,omitempty"`
	// A valid username used for an existing ARGUS instance, which is used during basic auth.
	Username UpdateCredentialsPayloadGetUsernameAttributeType `json:"username,omitempty"`
}

// NewUpdateCredentialsPayload instantiates a new UpdateCredentialsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCredentialsPayload() *UpdateCredentialsPayload {
	this := UpdateCredentialsPayload{}
	return &this
}

// NewUpdateCredentialsPayloadWithDefaults instantiates a new UpdateCredentialsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCredentialsPayloadWithDefaults() *UpdateCredentialsPayload {
	this := UpdateCredentialsPayload{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateCredentialsPayload) GetDisplayName() (res UpdateCredentialsPayloadGetDisplayNameRetType) {
	res, _ = o.GetDisplayNameOk()
	return
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialsPayload) GetDisplayNameOk() (ret UpdateCredentialsPayloadGetDisplayNameRetType, ok bool) {
	return getUpdateCredentialsPayloadGetDisplayNameAttributeTypeOk(o.DisplayName)
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateCredentialsPayload) HasDisplayName() bool {
	_, ok := o.GetDisplayNameOk()
	return ok
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateCredentialsPayload) SetDisplayName(v UpdateCredentialsPayloadGetDisplayNameRetType) {
	setUpdateCredentialsPayloadGetDisplayNameAttributeType(&o.DisplayName, v)
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateCredentialsPayload) GetPassword() (res UpdateCredentialsPayloadGetPasswordRetType) {
	res, _ = o.GetPasswordOk()
	return
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialsPayload) GetPasswordOk() (ret UpdateCredentialsPayloadGetPasswordRetType, ok bool) {
	return getUpdateCredentialsPayloadGetPasswordAttributeTypeOk(o.Password)
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateCredentialsPayload) HasPassword() bool {
	_, ok := o.GetPasswordOk()
	return ok
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateCredentialsPayload) SetPassword(v UpdateCredentialsPayloadGetPasswordRetType) {
	setUpdateCredentialsPayloadGetPasswordAttributeType(&o.Password, v)
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateCredentialsPayload) GetUsername() (res UpdateCredentialsPayloadGetUsernameRetType) {
	res, _ = o.GetUsernameOk()
	return
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialsPayload) GetUsernameOk() (ret UpdateCredentialsPayloadGetUsernameRetType, ok bool) {
	return getUpdateCredentialsPayloadGetUsernameAttributeTypeOk(o.Username)
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateCredentialsPayload) HasUsername() bool {
	_, ok := o.GetUsernameOk()
	return ok
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateCredentialsPayload) SetUsername(v UpdateCredentialsPayloadGetUsernameRetType) {
	setUpdateCredentialsPayloadGetUsernameAttributeType(&o.Username, v)
}

func (o UpdateCredentialsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateCredentialsPayloadGetDisplayNameAttributeTypeOk(o.DisplayName); ok {
		toSerialize["DisplayName"] = val
	}
	if val, ok := getUpdateCredentialsPayloadGetPasswordAttributeTypeOk(o.Password); ok {
		toSerialize["Password"] = val
	}
	if val, ok := getUpdateCredentialsPayloadGetUsernameAttributeTypeOk(o.Username); ok {
		toSerialize["Username"] = val
	}
	return toSerialize, nil
}

type NullableUpdateCredentialsPayload struct {
	value *UpdateCredentialsPayload
	isSet bool
}

func (v NullableUpdateCredentialsPayload) Get() *UpdateCredentialsPayload {
	return v.value
}

func (v *NullableUpdateCredentialsPayload) Set(val *UpdateCredentialsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCredentialsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCredentialsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCredentialsPayload(val *UpdateCredentialsPayload) *NullableUpdateCredentialsPayload {
	return &NullableUpdateCredentialsPayload{value: val, isSet: true}
}

func (v NullableUpdateCredentialsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCredentialsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

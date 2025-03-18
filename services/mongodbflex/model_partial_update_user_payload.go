/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"encoding/json"
)

// checks if the PartialUpdateUserPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialUpdateUserPayload{}

/*
	types and functions for database
*/

// isNotNullableString
type PartialUpdateUserPayloadGetDatabaseAttributeType = *string

func getPartialUpdateUserPayloadGetDatabaseAttributeTypeOk(arg PartialUpdateUserPayloadGetDatabaseAttributeType) (ret PartialUpdateUserPayloadGetDatabaseRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPartialUpdateUserPayloadGetDatabaseAttributeType(arg *PartialUpdateUserPayloadGetDatabaseAttributeType, val PartialUpdateUserPayloadGetDatabaseRetType) {
	*arg = &val
}

type PartialUpdateUserPayloadGetDatabaseArgType = string
type PartialUpdateUserPayloadGetDatabaseRetType = string

/*
	types and functions for roles
*/

// isArray
type PartialUpdateUserPayloadGetRolesAttributeType = *[]string
type PartialUpdateUserPayloadGetRolesArgType = []string
type PartialUpdateUserPayloadGetRolesRetType = []string

func getPartialUpdateUserPayloadGetRolesAttributeTypeOk(arg PartialUpdateUserPayloadGetRolesAttributeType) (ret PartialUpdateUserPayloadGetRolesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPartialUpdateUserPayloadGetRolesAttributeType(arg *PartialUpdateUserPayloadGetRolesAttributeType, val PartialUpdateUserPayloadGetRolesRetType) {
	*arg = &val
}

// PartialUpdateUserPayload struct for PartialUpdateUserPayload
type PartialUpdateUserPayload struct {
	Database PartialUpdateUserPayloadGetDatabaseAttributeType `json:"database,omitempty"`
	Roles    PartialUpdateUserPayloadGetRolesAttributeType    `json:"roles,omitempty"`
}

// NewPartialUpdateUserPayload instantiates a new PartialUpdateUserPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialUpdateUserPayload() *PartialUpdateUserPayload {
	this := PartialUpdateUserPayload{}
	return &this
}

// NewPartialUpdateUserPayloadWithDefaults instantiates a new PartialUpdateUserPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialUpdateUserPayloadWithDefaults() *PartialUpdateUserPayload {
	this := PartialUpdateUserPayload{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *PartialUpdateUserPayload) GetDatabase() (res PartialUpdateUserPayloadGetDatabaseRetType) {
	res, _ = o.GetDatabaseOk()
	return
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateUserPayload) GetDatabaseOk() (ret PartialUpdateUserPayloadGetDatabaseRetType, ok bool) {
	return getPartialUpdateUserPayloadGetDatabaseAttributeTypeOk(o.Database)
}

// HasDatabase returns a boolean if a field has been set.
func (o *PartialUpdateUserPayload) HasDatabase() bool {
	_, ok := o.GetDatabaseOk()
	return ok
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *PartialUpdateUserPayload) SetDatabase(v PartialUpdateUserPayloadGetDatabaseRetType) {
	setPartialUpdateUserPayloadGetDatabaseAttributeType(&o.Database, v)
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *PartialUpdateUserPayload) GetRoles() (res PartialUpdateUserPayloadGetRolesRetType) {
	res, _ = o.GetRolesOk()
	return
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateUserPayload) GetRolesOk() (ret PartialUpdateUserPayloadGetRolesRetType, ok bool) {
	return getPartialUpdateUserPayloadGetRolesAttributeTypeOk(o.Roles)
}

// HasRoles returns a boolean if a field has been set.
func (o *PartialUpdateUserPayload) HasRoles() bool {
	_, ok := o.GetRolesOk()
	return ok
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *PartialUpdateUserPayload) SetRoles(v PartialUpdateUserPayloadGetRolesRetType) {
	setPartialUpdateUserPayloadGetRolesAttributeType(&o.Roles, v)
}

func (o PartialUpdateUserPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPartialUpdateUserPayloadGetDatabaseAttributeTypeOk(o.Database); ok {
		toSerialize["Database"] = val
	}
	if val, ok := getPartialUpdateUserPayloadGetRolesAttributeTypeOk(o.Roles); ok {
		toSerialize["Roles"] = val
	}
	return toSerialize, nil
}

type NullablePartialUpdateUserPayload struct {
	value *PartialUpdateUserPayload
	isSet bool
}

func (v NullablePartialUpdateUserPayload) Get() *PartialUpdateUserPayload {
	return v.value
}

func (v *NullablePartialUpdateUserPayload) Set(val *PartialUpdateUserPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialUpdateUserPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialUpdateUserPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialUpdateUserPayload(val *PartialUpdateUserPayload) *NullablePartialUpdateUserPayload {
	return &NullablePartialUpdateUserPayload{value: val, isSet: true}
}

func (v NullablePartialUpdateUserPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialUpdateUserPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

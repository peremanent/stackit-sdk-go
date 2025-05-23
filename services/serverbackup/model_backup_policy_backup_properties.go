/*
STACKIT Server Backup Management API

API endpoints for Server Backup Operations on STACKIT Servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverbackup

import (
	"encoding/json"
)

// checks if the BackupPolicyBackupProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupPolicyBackupProperties{}

/*
	types and functions for name
*/

// isNotNullableString
type BackupPolicyBackupPropertiesGetNameAttributeType = *string

func getBackupPolicyBackupPropertiesGetNameAttributeTypeOk(arg BackupPolicyBackupPropertiesGetNameAttributeType) (ret BackupPolicyBackupPropertiesGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupPolicyBackupPropertiesGetNameAttributeType(arg *BackupPolicyBackupPropertiesGetNameAttributeType, val BackupPolicyBackupPropertiesGetNameRetType) {
	*arg = &val
}

type BackupPolicyBackupPropertiesGetNameArgType = string
type BackupPolicyBackupPropertiesGetNameRetType = string

/*
	types and functions for retentionPeriod
*/

// isInteger
type BackupPolicyBackupPropertiesGetRetentionPeriodAttributeType = *int64
type BackupPolicyBackupPropertiesGetRetentionPeriodArgType = int64
type BackupPolicyBackupPropertiesGetRetentionPeriodRetType = int64

func getBackupPolicyBackupPropertiesGetRetentionPeriodAttributeTypeOk(arg BackupPolicyBackupPropertiesGetRetentionPeriodAttributeType) (ret BackupPolicyBackupPropertiesGetRetentionPeriodRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupPolicyBackupPropertiesGetRetentionPeriodAttributeType(arg *BackupPolicyBackupPropertiesGetRetentionPeriodAttributeType, val BackupPolicyBackupPropertiesGetRetentionPeriodRetType) {
	*arg = &val
}

// BackupPolicyBackupProperties struct for BackupPolicyBackupProperties
type BackupPolicyBackupProperties struct {
	Name BackupPolicyBackupPropertiesGetNameAttributeType `json:"name,omitempty"`
	// Can be cast to int32 without loss of precision.
	RetentionPeriod BackupPolicyBackupPropertiesGetRetentionPeriodAttributeType `json:"retentionPeriod,omitempty"`
}

// NewBackupPolicyBackupProperties instantiates a new BackupPolicyBackupProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupPolicyBackupProperties() *BackupPolicyBackupProperties {
	this := BackupPolicyBackupProperties{}
	return &this
}

// NewBackupPolicyBackupPropertiesWithDefaults instantiates a new BackupPolicyBackupProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPolicyBackupPropertiesWithDefaults() *BackupPolicyBackupProperties {
	this := BackupPolicyBackupProperties{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackupPolicyBackupProperties) GetName() (res BackupPolicyBackupPropertiesGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicyBackupProperties) GetNameOk() (ret BackupPolicyBackupPropertiesGetNameRetType, ok bool) {
	return getBackupPolicyBackupPropertiesGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *BackupPolicyBackupProperties) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackupPolicyBackupProperties) SetName(v BackupPolicyBackupPropertiesGetNameRetType) {
	setBackupPolicyBackupPropertiesGetNameAttributeType(&o.Name, v)
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *BackupPolicyBackupProperties) GetRetentionPeriod() (res BackupPolicyBackupPropertiesGetRetentionPeriodRetType) {
	res, _ = o.GetRetentionPeriodOk()
	return
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicyBackupProperties) GetRetentionPeriodOk() (ret BackupPolicyBackupPropertiesGetRetentionPeriodRetType, ok bool) {
	return getBackupPolicyBackupPropertiesGetRetentionPeriodAttributeTypeOk(o.RetentionPeriod)
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *BackupPolicyBackupProperties) HasRetentionPeriod() bool {
	_, ok := o.GetRetentionPeriodOk()
	return ok
}

// SetRetentionPeriod gets a reference to the given int64 and assigns it to the RetentionPeriod field.
func (o *BackupPolicyBackupProperties) SetRetentionPeriod(v BackupPolicyBackupPropertiesGetRetentionPeriodRetType) {
	setBackupPolicyBackupPropertiesGetRetentionPeriodAttributeType(&o.RetentionPeriod, v)
}

func (o BackupPolicyBackupProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getBackupPolicyBackupPropertiesGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getBackupPolicyBackupPropertiesGetRetentionPeriodAttributeTypeOk(o.RetentionPeriod); ok {
		toSerialize["RetentionPeriod"] = val
	}
	return toSerialize, nil
}

type NullableBackupPolicyBackupProperties struct {
	value *BackupPolicyBackupProperties
	isSet bool
}

func (v NullableBackupPolicyBackupProperties) Get() *BackupPolicyBackupProperties {
	return v.value
}

func (v *NullableBackupPolicyBackupProperties) Set(val *BackupPolicyBackupProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupPolicyBackupProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupPolicyBackupProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupPolicyBackupProperties(val *BackupPolicyBackupProperties) *NullableBackupPolicyBackupProperties {
	return &NullableBackupPolicyBackupProperties{value: val, isSet: true}
}

func (v NullableBackupPolicyBackupProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupPolicyBackupProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

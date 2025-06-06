/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the ListBackupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBackupsResponse{}

/*
	types and functions for databases
*/

// isArray
type ListBackupsResponseGetDatabasesAttributeType = *[]BackupListBackupsResponseGrouped
type ListBackupsResponseGetDatabasesArgType = []BackupListBackupsResponseGrouped
type ListBackupsResponseGetDatabasesRetType = []BackupListBackupsResponseGrouped

func getListBackupsResponseGetDatabasesAttributeTypeOk(arg ListBackupsResponseGetDatabasesAttributeType) (ret ListBackupsResponseGetDatabasesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListBackupsResponseGetDatabasesAttributeType(arg *ListBackupsResponseGetDatabasesAttributeType, val ListBackupsResponseGetDatabasesRetType) {
	*arg = &val
}

// ListBackupsResponse struct for ListBackupsResponse
type ListBackupsResponse struct {
	Databases ListBackupsResponseGetDatabasesAttributeType `json:"databases,omitempty"`
}

// NewListBackupsResponse instantiates a new ListBackupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBackupsResponse() *ListBackupsResponse {
	this := ListBackupsResponse{}
	return &this
}

// NewListBackupsResponseWithDefaults instantiates a new ListBackupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBackupsResponseWithDefaults() *ListBackupsResponse {
	this := ListBackupsResponse{}
	return &this
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *ListBackupsResponse) GetDatabases() (res ListBackupsResponseGetDatabasesRetType) {
	res, _ = o.GetDatabasesOk()
	return
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupsResponse) GetDatabasesOk() (ret ListBackupsResponseGetDatabasesRetType, ok bool) {
	return getListBackupsResponseGetDatabasesAttributeTypeOk(o.Databases)
}

// HasDatabases returns a boolean if a field has been set.
func (o *ListBackupsResponse) HasDatabases() bool {
	_, ok := o.GetDatabasesOk()
	return ok
}

// SetDatabases gets a reference to the given []BackupListBackupsResponseGrouped and assigns it to the Databases field.
func (o *ListBackupsResponse) SetDatabases(v ListBackupsResponseGetDatabasesRetType) {
	setListBackupsResponseGetDatabasesAttributeType(&o.Databases, v)
}

func (o ListBackupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListBackupsResponseGetDatabasesAttributeTypeOk(o.Databases); ok {
		toSerialize["Databases"] = val
	}
	return toSerialize, nil
}

type NullableListBackupsResponse struct {
	value *ListBackupsResponse
	isSet bool
}

func (v NullableListBackupsResponse) Get() *ListBackupsResponse {
	return v.value
}

func (v *NullableListBackupsResponse) Set(val *ListBackupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListBackupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListBackupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBackupsResponse(val *ListBackupsResponse) *NullableListBackupsResponse {
	return &NullableListBackupsResponse{value: val, isSet: true}
}

func (v NullableListBackupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBackupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

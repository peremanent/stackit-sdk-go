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

// checks if the CreateInstancePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInstancePayload{}

// CreateInstancePayload struct for CreateInstancePayload
type CreateInstancePayload struct {
	Acl *CreateInstancePayloadAcl `json:"acl,omitempty"`
	// Cronjob for the daily full backup if not provided a job will generated between 00:00 and 04:59
	BackupSchedule *string `json:"backupSchedule,omitempty"`
	// Id of the selected flavor
	// REQUIRED
	FlavorId *string                 `json:"flavorId"`
	Labels   *map[string]interface{} `json:"labels,omitempty"`
	// Name of the instance
	// REQUIRED
	Name    *string                       `json:"name"`
	Options *CreateInstancePayloadOptions `json:"options,omitempty"`
	Storage *CreateInstancePayloadStorage `json:"storage,omitempty"`
	// Version of the MSSQL Server
	Version *string `json:"version,omitempty"`
}

type _CreateInstancePayload CreateInstancePayload

// NewCreateInstancePayload instantiates a new CreateInstancePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstancePayload(flavorId *string, name *string) *CreateInstancePayload {
	this := CreateInstancePayload{}
	this.FlavorId = flavorId
	this.Name = name
	var version string = "2022"
	this.Version = &version
	return &this
}

// NewCreateInstancePayloadWithDefaults instantiates a new CreateInstancePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstancePayloadWithDefaults() *CreateInstancePayload {
	this := CreateInstancePayload{}
	var version string = "2022"
	this.Version = &version
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *CreateInstancePayload) GetAcl() *CreateInstancePayloadAcl {
	if o == nil || IsNil(o.Acl) {
		var ret *CreateInstancePayloadAcl
		return ret
	}
	return o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetAclOk() (*CreateInstancePayloadAcl, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *CreateInstancePayload) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given CreateInstancePayloadAcl and assigns it to the Acl field.
func (o *CreateInstancePayload) SetAcl(v *CreateInstancePayloadAcl) {
	o.Acl = v
}

// GetBackupSchedule returns the BackupSchedule field value if set, zero value otherwise.
func (o *CreateInstancePayload) GetBackupSchedule() *string {
	if o == nil || IsNil(o.BackupSchedule) {
		var ret *string
		return ret
	}
	return o.BackupSchedule
}

// GetBackupScheduleOk returns a tuple with the BackupSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetBackupScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.BackupSchedule) {
		return nil, false
	}
	return o.BackupSchedule, true
}

// HasBackupSchedule returns a boolean if a field has been set.
func (o *CreateInstancePayload) HasBackupSchedule() bool {
	if o != nil && !IsNil(o.BackupSchedule) {
		return true
	}

	return false
}

// SetBackupSchedule gets a reference to the given string and assigns it to the BackupSchedule field.
func (o *CreateInstancePayload) SetBackupSchedule(v *string) {
	o.BackupSchedule = v
}

// GetFlavorId returns the FlavorId field value
func (o *CreateInstancePayload) GetFlavorId() *string {
	if o == nil || IsNil(o.FlavorId) {
		var ret *string
		return ret
	}

	return o.FlavorId
}

// GetFlavorIdOk returns a tuple with the FlavorId field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetFlavorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlavorId, true
}

// SetFlavorId sets field value
func (o *CreateInstancePayload) SetFlavorId(v *string) {
	o.FlavorId = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateInstancePayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateInstancePayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *CreateInstancePayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *CreateInstancePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CreateInstancePayload) SetName(v *string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreateInstancePayload) GetOptions() *CreateInstancePayloadOptions {
	if o == nil || IsNil(o.Options) {
		var ret *CreateInstancePayloadOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetOptionsOk() (*CreateInstancePayloadOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateInstancePayload) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given CreateInstancePayloadOptions and assigns it to the Options field.
func (o *CreateInstancePayload) SetOptions(v *CreateInstancePayloadOptions) {
	o.Options = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *CreateInstancePayload) GetStorage() *CreateInstancePayloadStorage {
	if o == nil || IsNil(o.Storage) {
		var ret *CreateInstancePayloadStorage
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetStorageOk() (*CreateInstancePayloadStorage, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *CreateInstancePayload) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given CreateInstancePayloadStorage and assigns it to the Storage field.
func (o *CreateInstancePayload) SetStorage(v *CreateInstancePayloadStorage) {
	o.Storage = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateInstancePayload) GetVersion() *string {
	if o == nil || IsNil(o.Version) {
		var ret *string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateInstancePayload) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CreateInstancePayload) SetVersion(v *string) {
	o.Version = v
}

func (o CreateInstancePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	if !IsNil(o.BackupSchedule) {
		toSerialize["backupSchedule"] = o.BackupSchedule
	}
	toSerialize["flavorId"] = o.FlavorId
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableCreateInstancePayload struct {
	value *CreateInstancePayload
	isSet bool
}

func (v NullableCreateInstancePayload) Get() *CreateInstancePayload {
	return v.value
}

func (v *NullableCreateInstancePayload) Set(val *CreateInstancePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstancePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstancePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstancePayload(val *CreateInstancePayload) *NullableCreateInstancePayload {
	return &NullableCreateInstancePayload{value: val, isSet: true}
}

func (v NullableCreateInstancePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstancePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

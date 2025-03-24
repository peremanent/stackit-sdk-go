/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
	"time"
)

// checks if the Backup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Backup{}

/*
	types and functions for availabilityZone
*/

// isNotNullableString
type BackupGetAvailabilityZoneAttributeType = *string

func getBackupGetAvailabilityZoneAttributeTypeOk(arg BackupGetAvailabilityZoneAttributeType) (ret BackupGetAvailabilityZoneRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetAvailabilityZoneAttributeType(arg *BackupGetAvailabilityZoneAttributeType, val BackupGetAvailabilityZoneRetType) {
	*arg = &val
}

type BackupGetAvailabilityZoneArgType = string
type BackupGetAvailabilityZoneRetType = string

/*
	types and functions for createdAt
*/

// isDateTime
type BackupGetCreatedAtAttributeType = *time.Time
type BackupGetCreatedAtArgType = time.Time
type BackupGetCreatedAtRetType = time.Time

func getBackupGetCreatedAtAttributeTypeOk(arg BackupGetCreatedAtAttributeType) (ret BackupGetCreatedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetCreatedAtAttributeType(arg *BackupGetCreatedAtAttributeType, val BackupGetCreatedAtRetType) {
	*arg = &val
}

/*
	types and functions for id
*/

// isNotNullableString
type BackupGetIdAttributeType = *string

func getBackupGetIdAttributeTypeOk(arg BackupGetIdAttributeType) (ret BackupGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetIdAttributeType(arg *BackupGetIdAttributeType, val BackupGetIdRetType) {
	*arg = &val
}

type BackupGetIdArgType = string
type BackupGetIdRetType = string

/*
	types and functions for labels
*/

// isFreeform
type BackupGetLabelsAttributeType = *map[string]interface{}
type BackupGetLabelsArgType = map[string]interface{}
type BackupGetLabelsRetType = map[string]interface{}

func getBackupGetLabelsAttributeTypeOk(arg BackupGetLabelsAttributeType) (ret BackupGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetLabelsAttributeType(arg *BackupGetLabelsAttributeType, val BackupGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type BackupGetNameAttributeType = *string

func getBackupGetNameAttributeTypeOk(arg BackupGetNameAttributeType) (ret BackupGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetNameAttributeType(arg *BackupGetNameAttributeType, val BackupGetNameRetType) {
	*arg = &val
}

type BackupGetNameArgType = string
type BackupGetNameRetType = string

/*
	types and functions for size
*/

// isLong
type BackupGetSizeAttributeType = *int64
type BackupGetSizeArgType = int64
type BackupGetSizeRetType = int64

func getBackupGetSizeAttributeTypeOk(arg BackupGetSizeAttributeType) (ret BackupGetSizeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetSizeAttributeType(arg *BackupGetSizeAttributeType, val BackupGetSizeRetType) {
	*arg = &val
}

/*
	types and functions for snapshotId
*/

// isNotNullableString
type BackupGetSnapshotIdAttributeType = *string

func getBackupGetSnapshotIdAttributeTypeOk(arg BackupGetSnapshotIdAttributeType) (ret BackupGetSnapshotIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetSnapshotIdAttributeType(arg *BackupGetSnapshotIdAttributeType, val BackupGetSnapshotIdRetType) {
	*arg = &val
}

type BackupGetSnapshotIdArgType = string
type BackupGetSnapshotIdRetType = string

/*
	types and functions for status
*/

// isNotNullableString
type BackupGetStatusAttributeType = *string

func getBackupGetStatusAttributeTypeOk(arg BackupGetStatusAttributeType) (ret BackupGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetStatusAttributeType(arg *BackupGetStatusAttributeType, val BackupGetStatusRetType) {
	*arg = &val
}

type BackupGetStatusArgType = string
type BackupGetStatusRetType = string

/*
	types and functions for updatedAt
*/

// isDateTime
type BackupGetUpdatedAtAttributeType = *time.Time
type BackupGetUpdatedAtArgType = time.Time
type BackupGetUpdatedAtRetType = time.Time

func getBackupGetUpdatedAtAttributeTypeOk(arg BackupGetUpdatedAtAttributeType) (ret BackupGetUpdatedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetUpdatedAtAttributeType(arg *BackupGetUpdatedAtAttributeType, val BackupGetUpdatedAtRetType) {
	*arg = &val
}

/*
	types and functions for volumeId
*/

// isNotNullableString
type BackupGetVolumeIdAttributeType = *string

func getBackupGetVolumeIdAttributeTypeOk(arg BackupGetVolumeIdAttributeType) (ret BackupGetVolumeIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetVolumeIdAttributeType(arg *BackupGetVolumeIdAttributeType, val BackupGetVolumeIdRetType) {
	*arg = &val
}

type BackupGetVolumeIdArgType = string
type BackupGetVolumeIdRetType = string

// Backup Object that represents a backup.
type Backup struct {
	// Object that represents an availability zone.
	AvailabilityZone BackupGetAvailabilityZoneAttributeType `json:"availabilityZone,omitempty"`
	// Date-time when resource was created.
	CreatedAt BackupGetCreatedAtAttributeType `json:"createdAt,omitempty"`
	// Universally Unique Identifier (UUID).
	Id BackupGetIdAttributeType `json:"id,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels BackupGetLabelsAttributeType `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name BackupGetNameAttributeType `json:"name,omitempty"`
	// Size in Gigabyte.
	Size BackupGetSizeAttributeType `json:"size,omitempty"`
	// Universally Unique Identifier (UUID).
	SnapshotId BackupGetSnapshotIdAttributeType `json:"snapshotId,omitempty"`
	// The status of a backup object. Possible values: `AVAILABLE`, `CREATING`, `DELETED`, `DELETING`, `ERROR`, `RESTORING`.
	Status BackupGetStatusAttributeType `json:"status,omitempty"`
	// Date-time when resource was last updated.
	UpdatedAt BackupGetUpdatedAtAttributeType `json:"updatedAt,omitempty"`
	// Universally Unique Identifier (UUID).
	VolumeId BackupGetVolumeIdAttributeType `json:"volumeId,omitempty"`
}

// NewBackup instantiates a new Backup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackup() *Backup {
	this := Backup{}
	return &this
}

// NewBackupWithDefaults instantiates a new Backup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupWithDefaults() *Backup {
	this := Backup{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *Backup) GetAvailabilityZone() (res BackupGetAvailabilityZoneRetType) {
	res, _ = o.GetAvailabilityZoneOk()
	return
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetAvailabilityZoneOk() (ret BackupGetAvailabilityZoneRetType, ok bool) {
	return getBackupGetAvailabilityZoneAttributeTypeOk(o.AvailabilityZone)
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *Backup) HasAvailabilityZone() bool {
	_, ok := o.GetAvailabilityZoneOk()
	return ok
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *Backup) SetAvailabilityZone(v BackupGetAvailabilityZoneRetType) {
	setBackupGetAvailabilityZoneAttributeType(&o.AvailabilityZone, v)
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Backup) GetCreatedAt() (res BackupGetCreatedAtRetType) {
	res, _ = o.GetCreatedAtOk()
	return
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetCreatedAtOk() (ret BackupGetCreatedAtRetType, ok bool) {
	return getBackupGetCreatedAtAttributeTypeOk(o.CreatedAt)
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Backup) HasCreatedAt() bool {
	_, ok := o.GetCreatedAtOk()
	return ok
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Backup) SetCreatedAt(v BackupGetCreatedAtRetType) {
	setBackupGetCreatedAtAttributeType(&o.CreatedAt, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Backup) GetId() (res BackupGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetIdOk() (ret BackupGetIdRetType, ok bool) {
	return getBackupGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *Backup) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Backup) SetId(v BackupGetIdRetType) {
	setBackupGetIdAttributeType(&o.Id, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Backup) GetLabels() (res BackupGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetLabelsOk() (ret BackupGetLabelsRetType, ok bool) {
	return getBackupGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *Backup) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *Backup) SetLabels(v BackupGetLabelsRetType) {
	setBackupGetLabelsAttributeType(&o.Labels, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Backup) GetName() (res BackupGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetNameOk() (ret BackupGetNameRetType, ok bool) {
	return getBackupGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *Backup) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Backup) SetName(v BackupGetNameRetType) {
	setBackupGetNameAttributeType(&o.Name, v)
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Backup) GetSize() (res BackupGetSizeRetType) {
	res, _ = o.GetSizeOk()
	return
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetSizeOk() (ret BackupGetSizeRetType, ok bool) {
	return getBackupGetSizeAttributeTypeOk(o.Size)
}

// HasSize returns a boolean if a field has been set.
func (o *Backup) HasSize() bool {
	_, ok := o.GetSizeOk()
	return ok
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Backup) SetSize(v BackupGetSizeRetType) {
	setBackupGetSizeAttributeType(&o.Size, v)
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *Backup) GetSnapshotId() (res BackupGetSnapshotIdRetType) {
	res, _ = o.GetSnapshotIdOk()
	return
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetSnapshotIdOk() (ret BackupGetSnapshotIdRetType, ok bool) {
	return getBackupGetSnapshotIdAttributeTypeOk(o.SnapshotId)
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *Backup) HasSnapshotId() bool {
	_, ok := o.GetSnapshotIdOk()
	return ok
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *Backup) SetSnapshotId(v BackupGetSnapshotIdRetType) {
	setBackupGetSnapshotIdAttributeType(&o.SnapshotId, v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Backup) GetStatus() (res BackupGetStatusRetType) {
	res, _ = o.GetStatusOk()
	return
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetStatusOk() (ret BackupGetStatusRetType, ok bool) {
	return getBackupGetStatusAttributeTypeOk(o.Status)
}

// HasStatus returns a boolean if a field has been set.
func (o *Backup) HasStatus() bool {
	_, ok := o.GetStatusOk()
	return ok
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Backup) SetStatus(v BackupGetStatusRetType) {
	setBackupGetStatusAttributeType(&o.Status, v)
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Backup) GetUpdatedAt() (res BackupGetUpdatedAtRetType) {
	res, _ = o.GetUpdatedAtOk()
	return
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetUpdatedAtOk() (ret BackupGetUpdatedAtRetType, ok bool) {
	return getBackupGetUpdatedAtAttributeTypeOk(o.UpdatedAt)
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Backup) HasUpdatedAt() bool {
	_, ok := o.GetUpdatedAtOk()
	return ok
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Backup) SetUpdatedAt(v BackupGetUpdatedAtRetType) {
	setBackupGetUpdatedAtAttributeType(&o.UpdatedAt, v)
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *Backup) GetVolumeId() (res BackupGetVolumeIdRetType) {
	res, _ = o.GetVolumeIdOk()
	return
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetVolumeIdOk() (ret BackupGetVolumeIdRetType, ok bool) {
	return getBackupGetVolumeIdAttributeTypeOk(o.VolumeId)
}

// HasVolumeId returns a boolean if a field has been set.
func (o *Backup) HasVolumeId() bool {
	_, ok := o.GetVolumeIdOk()
	return ok
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *Backup) SetVolumeId(v BackupGetVolumeIdRetType) {
	setBackupGetVolumeIdAttributeType(&o.VolumeId, v)
}

func (o Backup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getBackupGetAvailabilityZoneAttributeTypeOk(o.AvailabilityZone); ok {
		toSerialize["AvailabilityZone"] = val
	}
	if val, ok := getBackupGetCreatedAtAttributeTypeOk(o.CreatedAt); ok {
		toSerialize["CreatedAt"] = val
	}
	if val, ok := getBackupGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getBackupGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getBackupGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getBackupGetSizeAttributeTypeOk(o.Size); ok {
		toSerialize["Size"] = val
	}
	if val, ok := getBackupGetSnapshotIdAttributeTypeOk(o.SnapshotId); ok {
		toSerialize["SnapshotId"] = val
	}
	if val, ok := getBackupGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	if val, ok := getBackupGetUpdatedAtAttributeTypeOk(o.UpdatedAt); ok {
		toSerialize["UpdatedAt"] = val
	}
	if val, ok := getBackupGetVolumeIdAttributeTypeOk(o.VolumeId); ok {
		toSerialize["VolumeId"] = val
	}
	return toSerialize, nil
}

type NullableBackup struct {
	value *Backup
	isSet bool
}

func (v NullableBackup) Get() *Backup {
	return v.value
}

func (v *NullableBackup) Set(val *Backup) {
	v.value = val
	v.isSet = true
}

func (v NullableBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackup(val *Backup) *NullableBackup {
	return &NullableBackup{value: val, isSet: true}
}

func (v NullableBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the BootVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootVolume{}

/*
	types and functions for deleteOnTermination
*/

// isBoolean
type BootVolumegetDeleteOnTerminationAttributeType = *bool
type BootVolumegetDeleteOnTerminationArgType = bool
type BootVolumegetDeleteOnTerminationRetType = bool

func getBootVolumegetDeleteOnTerminationAttributeTypeOk(arg BootVolumegetDeleteOnTerminationAttributeType) (ret BootVolumegetDeleteOnTerminationRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBootVolumegetDeleteOnTerminationAttributeType(arg *BootVolumegetDeleteOnTerminationAttributeType, val BootVolumegetDeleteOnTerminationRetType) {
	*arg = &val
}

/*
	types and functions for id
*/

// isNotNullableString
type BootVolumeGetIdAttributeType = *string

func getBootVolumeGetIdAttributeTypeOk(arg BootVolumeGetIdAttributeType) (ret BootVolumeGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBootVolumeGetIdAttributeType(arg *BootVolumeGetIdAttributeType, val BootVolumeGetIdRetType) {
	*arg = &val
}

type BootVolumeGetIdArgType = string
type BootVolumeGetIdRetType = string

/*
	types and functions for performanceClass
*/

// isNotNullableString
type BootVolumeGetPerformanceClassAttributeType = *string

func getBootVolumeGetPerformanceClassAttributeTypeOk(arg BootVolumeGetPerformanceClassAttributeType) (ret BootVolumeGetPerformanceClassRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBootVolumeGetPerformanceClassAttributeType(arg *BootVolumeGetPerformanceClassAttributeType, val BootVolumeGetPerformanceClassRetType) {
	*arg = &val
}

type BootVolumeGetPerformanceClassArgType = string
type BootVolumeGetPerformanceClassRetType = string

/*
	types and functions for size
*/

// isLong
type BootVolumeGetSizeAttributeType = *int64
type BootVolumeGetSizeArgType = int64
type BootVolumeGetSizeRetType = int64

func getBootVolumeGetSizeAttributeTypeOk(arg BootVolumeGetSizeAttributeType) (ret BootVolumeGetSizeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBootVolumeGetSizeAttributeType(arg *BootVolumeGetSizeAttributeType, val BootVolumeGetSizeRetType) {
	*arg = &val
}

/*
	types and functions for source
*/

// isModel
type BootVolumeGetSourceAttributeType = *BootVolumeSource
type BootVolumeGetSourceArgType = BootVolumeSource
type BootVolumeGetSourceRetType = BootVolumeSource

func getBootVolumeGetSourceAttributeTypeOk(arg BootVolumeGetSourceAttributeType) (ret BootVolumeGetSourceRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBootVolumeGetSourceAttributeType(arg *BootVolumeGetSourceAttributeType, val BootVolumeGetSourceRetType) {
	*arg = &val
}

// BootVolume The boot device for the server.
type BootVolume struct {
	// Delete the volume during the termination of the server. Defaults to false.
	DeleteOnTermination BootVolumegetDeleteOnTerminationAttributeType `json:"deleteOnTermination,omitempty"`
	// Universally Unique Identifier (UUID).
	Id BootVolumeGetIdAttributeType `json:"id,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	PerformanceClass BootVolumeGetPerformanceClassAttributeType `json:"performanceClass,omitempty"`
	// Size in Gigabyte.
	Size   BootVolumeGetSizeAttributeType   `json:"size,omitempty"`
	Source BootVolumeGetSourceAttributeType `json:"source,omitempty"`
}

// NewBootVolume instantiates a new BootVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootVolume() *BootVolume {
	this := BootVolume{}
	return &this
}

// NewBootVolumeWithDefaults instantiates a new BootVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootVolumeWithDefaults() *BootVolume {
	this := BootVolume{}
	return &this
}

// GetDeleteOnTermination returns the DeleteOnTermination field value if set, zero value otherwise.
func (o *BootVolume) GetDeleteOnTermination() (res BootVolumegetDeleteOnTerminationRetType) {
	res, _ = o.GetDeleteOnTerminationOk()
	return
}

// GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetDeleteOnTerminationOk() (ret BootVolumegetDeleteOnTerminationRetType, ok bool) {
	return getBootVolumegetDeleteOnTerminationAttributeTypeOk(o.DeleteOnTermination)
}

// HasDeleteOnTermination returns a boolean if a field has been set.
func (o *BootVolume) HasDeleteOnTermination() bool {
	_, ok := o.GetDeleteOnTerminationOk()
	return ok
}

// SetDeleteOnTermination gets a reference to the given bool and assigns it to the DeleteOnTermination field.
func (o *BootVolume) SetDeleteOnTermination(v BootVolumegetDeleteOnTerminationRetType) {
	setBootVolumegetDeleteOnTerminationAttributeType(&o.DeleteOnTermination, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BootVolume) GetId() (res BootVolumeGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetIdOk() (ret BootVolumeGetIdRetType, ok bool) {
	return getBootVolumeGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *BootVolume) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BootVolume) SetId(v BootVolumeGetIdRetType) {
	setBootVolumeGetIdAttributeType(&o.Id, v)
}

// GetPerformanceClass returns the PerformanceClass field value if set, zero value otherwise.
func (o *BootVolume) GetPerformanceClass() (res BootVolumeGetPerformanceClassRetType) {
	res, _ = o.GetPerformanceClassOk()
	return
}

// GetPerformanceClassOk returns a tuple with the PerformanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetPerformanceClassOk() (ret BootVolumeGetPerformanceClassRetType, ok bool) {
	return getBootVolumeGetPerformanceClassAttributeTypeOk(o.PerformanceClass)
}

// HasPerformanceClass returns a boolean if a field has been set.
func (o *BootVolume) HasPerformanceClass() bool {
	_, ok := o.GetPerformanceClassOk()
	return ok
}

// SetPerformanceClass gets a reference to the given string and assigns it to the PerformanceClass field.
func (o *BootVolume) SetPerformanceClass(v BootVolumeGetPerformanceClassRetType) {
	setBootVolumeGetPerformanceClassAttributeType(&o.PerformanceClass, v)
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BootVolume) GetSize() (res BootVolumeGetSizeRetType) {
	res, _ = o.GetSizeOk()
	return
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetSizeOk() (ret BootVolumeGetSizeRetType, ok bool) {
	return getBootVolumeGetSizeAttributeTypeOk(o.Size)
}

// HasSize returns a boolean if a field has been set.
func (o *BootVolume) HasSize() bool {
	_, ok := o.GetSizeOk()
	return ok
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *BootVolume) SetSize(v BootVolumeGetSizeRetType) {
	setBootVolumeGetSizeAttributeType(&o.Size, v)
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BootVolume) GetSource() (res BootVolumeGetSourceRetType) {
	res, _ = o.GetSourceOk()
	return
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetSourceOk() (ret BootVolumeGetSourceRetType, ok bool) {
	return getBootVolumeGetSourceAttributeTypeOk(o.Source)
}

// HasSource returns a boolean if a field has been set.
func (o *BootVolume) HasSource() bool {
	_, ok := o.GetSourceOk()
	return ok
}

// SetSource gets a reference to the given BootVolumeSource and assigns it to the Source field.
func (o *BootVolume) SetSource(v BootVolumeGetSourceRetType) {
	setBootVolumeGetSourceAttributeType(&o.Source, v)
}

func (o BootVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getBootVolumegetDeleteOnTerminationAttributeTypeOk(o.DeleteOnTermination); ok {
		toSerialize["DeleteOnTermination"] = val
	}
	if val, ok := getBootVolumeGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getBootVolumeGetPerformanceClassAttributeTypeOk(o.PerformanceClass); ok {
		toSerialize["PerformanceClass"] = val
	}
	if val, ok := getBootVolumeGetSizeAttributeTypeOk(o.Size); ok {
		toSerialize["Size"] = val
	}
	if val, ok := getBootVolumeGetSourceAttributeTypeOk(o.Source); ok {
		toSerialize["Source"] = val
	}
	return toSerialize, nil
}

type NullableBootVolume struct {
	value *BootVolume
	isSet bool
}

func (v NullableBootVolume) Get() *BootVolume {
	return v.value
}

func (v *NullableBootVolume) Set(val *BootVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableBootVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableBootVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootVolume(val *BootVolume) *NullableBootVolume {
	return &NullableBootVolume{value: val, isSet: true}
}

func (v NullableBootVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

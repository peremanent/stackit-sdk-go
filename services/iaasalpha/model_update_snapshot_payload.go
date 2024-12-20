/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the UpdateSnapshotPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSnapshotPayload{}

// UpdateSnapshotPayload Object that represents an update request body of a snapshot.
type UpdateSnapshotPayload struct {
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
}

// NewUpdateSnapshotPayload instantiates a new UpdateSnapshotPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSnapshotPayload() *UpdateSnapshotPayload {
	this := UpdateSnapshotPayload{}
	return &this
}

// NewUpdateSnapshotPayloadWithDefaults instantiates a new UpdateSnapshotPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSnapshotPayloadWithDefaults() *UpdateSnapshotPayload {
	this := UpdateSnapshotPayload{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateSnapshotPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateSnapshotPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateSnapshotPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSnapshotPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSnapshotPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateSnapshotPayload) SetName(v *string) {
	o.Name = v
}

func (o UpdateSnapshotPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateSnapshotPayload struct {
	value *UpdateSnapshotPayload
	isSet bool
}

func (v NullableUpdateSnapshotPayload) Get() *UpdateSnapshotPayload {
	return v.value
}

func (v *NullableUpdateSnapshotPayload) Set(val *UpdateSnapshotPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSnapshotPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSnapshotPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSnapshotPayload(val *UpdateSnapshotPayload) *NullableUpdateSnapshotPayload {
	return &NullableUpdateSnapshotPayload{value: val, isSet: true}
}

func (v NullableUpdateSnapshotPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSnapshotPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

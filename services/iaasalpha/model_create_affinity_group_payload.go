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

// checks if the CreateAffinityGroupPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAffinityGroupPayload{}

// CreateAffinityGroupPayload Definition of an affinity group.
type CreateAffinityGroupPayload struct {
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// The servers that are part of the affinity group.
	Members *[]string `json:"members,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	Name *string `json:"name"`
	// The affinity group policy.
	// REQUIRED
	Policy *string `json:"policy"`
}

type _CreateAffinityGroupPayload CreateAffinityGroupPayload

// NewCreateAffinityGroupPayload instantiates a new CreateAffinityGroupPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAffinityGroupPayload(name *string, policy *string) *CreateAffinityGroupPayload {
	this := CreateAffinityGroupPayload{}
	this.Name = name
	this.Policy = policy
	return &this
}

// NewCreateAffinityGroupPayloadWithDefaults instantiates a new CreateAffinityGroupPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAffinityGroupPayloadWithDefaults() *CreateAffinityGroupPayload {
	this := CreateAffinityGroupPayload{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateAffinityGroupPayload) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAffinityGroupPayload) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateAffinityGroupPayload) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateAffinityGroupPayload) SetId(v *string) {
	o.Id = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *CreateAffinityGroupPayload) GetMembers() *[]string {
	if o == nil || IsNil(o.Members) {
		var ret *[]string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAffinityGroupPayload) GetMembersOk() (*[]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *CreateAffinityGroupPayload) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *CreateAffinityGroupPayload) SetMembers(v *[]string) {
	o.Members = v
}

// GetName returns the Name field value
func (o *CreateAffinityGroupPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAffinityGroupPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CreateAffinityGroupPayload) SetName(v *string) {
	o.Name = v
}

// GetPolicy returns the Policy field value
func (o *CreateAffinityGroupPayload) GetPolicy() *string {
	if o == nil || IsNil(o.Policy) {
		var ret *string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *CreateAffinityGroupPayload) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy, true
}

// SetPolicy sets field value
func (o *CreateAffinityGroupPayload) SetPolicy(v *string) {
	o.Policy = v
}

func (o CreateAffinityGroupPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	toSerialize["name"] = o.Name
	toSerialize["policy"] = o.Policy
	return toSerialize, nil
}

type NullableCreateAffinityGroupPayload struct {
	value *CreateAffinityGroupPayload
	isSet bool
}

func (v NullableCreateAffinityGroupPayload) Get() *CreateAffinityGroupPayload {
	return v.value
}

func (v *NullableCreateAffinityGroupPayload) Set(val *CreateAffinityGroupPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAffinityGroupPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAffinityGroupPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAffinityGroupPayload(val *CreateAffinityGroupPayload) *NullableCreateAffinityGroupPayload {
	return &NullableCreateAffinityGroupPayload{value: val, isSet: true}
}

func (v NullableCreateAffinityGroupPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAffinityGroupPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

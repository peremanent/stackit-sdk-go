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

// checks if the CreateNetworkAddressFamily type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkAddressFamily{}

/*
	types and functions for ipv4
*/

// isModel
type CreateNetworkAddressFamilyGetIpv4AttributeType = *CreateNetworkIPv4Body
type CreateNetworkAddressFamilyGetIpv4ArgType = CreateNetworkIPv4Body
type CreateNetworkAddressFamilyGetIpv4RetType = CreateNetworkIPv4Body

func getCreateNetworkAddressFamilyGetIpv4AttributeTypeOk(arg CreateNetworkAddressFamilyGetIpv4AttributeType) (ret CreateNetworkAddressFamilyGetIpv4RetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateNetworkAddressFamilyGetIpv4AttributeType(arg *CreateNetworkAddressFamilyGetIpv4AttributeType, val CreateNetworkAddressFamilyGetIpv4RetType) {
	*arg = &val
}

/*
	types and functions for ipv6
*/

// isModel
type CreateNetworkAddressFamilyGetIpv6AttributeType = *CreateNetworkIPv6Body
type CreateNetworkAddressFamilyGetIpv6ArgType = CreateNetworkIPv6Body
type CreateNetworkAddressFamilyGetIpv6RetType = CreateNetworkIPv6Body

func getCreateNetworkAddressFamilyGetIpv6AttributeTypeOk(arg CreateNetworkAddressFamilyGetIpv6AttributeType) (ret CreateNetworkAddressFamilyGetIpv6RetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateNetworkAddressFamilyGetIpv6AttributeType(arg *CreateNetworkAddressFamilyGetIpv6AttributeType, val CreateNetworkAddressFamilyGetIpv6RetType) {
	*arg = &val
}

// CreateNetworkAddressFamily The addressFamily object for a network create request.
type CreateNetworkAddressFamily struct {
	Ipv4 CreateNetworkAddressFamilyGetIpv4AttributeType `json:"ipv4,omitempty"`
	Ipv6 CreateNetworkAddressFamilyGetIpv6AttributeType `json:"ipv6,omitempty"`
}

// NewCreateNetworkAddressFamily instantiates a new CreateNetworkAddressFamily object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkAddressFamily() *CreateNetworkAddressFamily {
	this := CreateNetworkAddressFamily{}
	return &this
}

// NewCreateNetworkAddressFamilyWithDefaults instantiates a new CreateNetworkAddressFamily object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkAddressFamilyWithDefaults() *CreateNetworkAddressFamily {
	this := CreateNetworkAddressFamily{}
	return &this
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *CreateNetworkAddressFamily) GetIpv4() (res CreateNetworkAddressFamilyGetIpv4RetType) {
	res, _ = o.GetIpv4Ok()
	return
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkAddressFamily) GetIpv4Ok() (ret CreateNetworkAddressFamilyGetIpv4RetType, ok bool) {
	return getCreateNetworkAddressFamilyGetIpv4AttributeTypeOk(o.Ipv4)
}

// HasIpv4 returns a boolean if a field has been set.
func (o *CreateNetworkAddressFamily) HasIpv4() bool {
	_, ok := o.GetIpv4Ok()
	return ok
}

// SetIpv4 gets a reference to the given CreateNetworkIPv4Body and assigns it to the Ipv4 field.
func (o *CreateNetworkAddressFamily) SetIpv4(v CreateNetworkAddressFamilyGetIpv4RetType) {
	setCreateNetworkAddressFamilyGetIpv4AttributeType(&o.Ipv4, v)
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *CreateNetworkAddressFamily) GetIpv6() (res CreateNetworkAddressFamilyGetIpv6RetType) {
	res, _ = o.GetIpv6Ok()
	return
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkAddressFamily) GetIpv6Ok() (ret CreateNetworkAddressFamilyGetIpv6RetType, ok bool) {
	return getCreateNetworkAddressFamilyGetIpv6AttributeTypeOk(o.Ipv6)
}

// HasIpv6 returns a boolean if a field has been set.
func (o *CreateNetworkAddressFamily) HasIpv6() bool {
	_, ok := o.GetIpv6Ok()
	return ok
}

// SetIpv6 gets a reference to the given CreateNetworkIPv6Body and assigns it to the Ipv6 field.
func (o *CreateNetworkAddressFamily) SetIpv6(v CreateNetworkAddressFamilyGetIpv6RetType) {
	setCreateNetworkAddressFamilyGetIpv6AttributeType(&o.Ipv6, v)
}

func (o CreateNetworkAddressFamily) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateNetworkAddressFamilyGetIpv4AttributeTypeOk(o.Ipv4); ok {
		toSerialize["Ipv4"] = val
	}
	if val, ok := getCreateNetworkAddressFamilyGetIpv6AttributeTypeOk(o.Ipv6); ok {
		toSerialize["Ipv6"] = val
	}
	return toSerialize, nil
}

type NullableCreateNetworkAddressFamily struct {
	value *CreateNetworkAddressFamily
	isSet bool
}

func (v NullableCreateNetworkAddressFamily) Get() *CreateNetworkAddressFamily {
	return v.value
}

func (v *NullableCreateNetworkAddressFamily) Set(val *CreateNetworkAddressFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkAddressFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkAddressFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkAddressFamily(val *CreateNetworkAddressFamily) *NullableCreateNetworkAddressFamily {
	return &NullableCreateNetworkAddressFamily{value: val, isSet: true}
}

func (v NullableCreateNetworkAddressFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkAddressFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

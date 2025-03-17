/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the CustomDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDomain{}

/*
	types and functions for errors
*/

// isArray
type CustomDomainGetErrorsAttributeType = *[]StatusError
type CustomDomainGetErrorsArgType = []StatusError
type CustomDomainGetErrorsRetType = []StatusError

func getCustomDomainGetErrorsAttributeTypeOk(arg CustomDomainGetErrorsAttributeType) (ret CustomDomainGetErrorsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCustomDomainGetErrorsAttributeType(arg *CustomDomainGetErrorsAttributeType, val CustomDomainGetErrorsRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type CustomDomainGetNameAttributeType = *string

func getCustomDomainGetNameAttributeTypeOk(arg CustomDomainGetNameAttributeType) (ret CustomDomainGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCustomDomainGetNameAttributeType(arg *CustomDomainGetNameAttributeType, val CustomDomainGetNameRetType) {
	*arg = &val
}

type CustomDomainGetNameArgType = string
type CustomDomainGetNameRetType = string

/*
	types and functions for status
*/

// isEnumRef
type CustomDomainGetStatusAttributeType = *DomainStatus
type CustomDomainGetStatusArgType = DomainStatus
type CustomDomainGetStatusRetType = DomainStatus

func getCustomDomainGetStatusAttributeTypeOk(arg CustomDomainGetStatusAttributeType) (ret CustomDomainGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCustomDomainGetStatusAttributeType(arg *CustomDomainGetStatusAttributeType, val CustomDomainGetStatusRetType) {
	*arg = &val
}

// CustomDomain Definition of a custom domain
type CustomDomain struct {
	// This object is present if the custom domain has errors.
	Errors CustomDomainGetErrorsAttributeType `json:"errors,omitempty"`
	// The domain. Can be used as input for the GetCustomDomain endpoint
	// REQUIRED
	Name CustomDomainGetNameAttributeType `json:"name"`
	// REQUIRED
	Status CustomDomainGetStatusAttributeType `json:"status"`
}

type _CustomDomain CustomDomain

// NewCustomDomain instantiates a new CustomDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomain(name CustomDomainGetNameArgType, status CustomDomainGetStatusArgType) *CustomDomain {
	this := CustomDomain{}
	setCustomDomainGetNameAttributeType(&this.Name, name)
	setCustomDomainGetStatusAttributeType(&this.Status, status)
	return &this
}

// NewCustomDomainWithDefaults instantiates a new CustomDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainWithDefaults() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CustomDomain) GetErrors() (res CustomDomainGetErrorsRetType) {
	res, _ = o.GetErrorsOk()
	return
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetErrorsOk() (ret CustomDomainGetErrorsRetType, ok bool) {
	return getCustomDomainGetErrorsAttributeTypeOk(o.Errors)
}

// HasErrors returns a boolean if a field has been set.
func (o *CustomDomain) HasErrors() bool {
	_, ok := o.GetErrorsOk()
	return ok
}

// SetErrors gets a reference to the given []StatusError and assigns it to the Errors field.
func (o *CustomDomain) SetErrors(v CustomDomainGetErrorsRetType) {
	setCustomDomainGetErrorsAttributeType(&o.Errors, v)
}

// GetName returns the Name field value
func (o *CustomDomain) GetName() (ret CustomDomainGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetNameOk() (ret CustomDomainGetNameRetType, ok bool) {
	return getCustomDomainGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *CustomDomain) SetName(v CustomDomainGetNameRetType) {
	setCustomDomainGetNameAttributeType(&o.Name, v)
}

// GetStatus returns the Status field value
func (o *CustomDomain) GetStatus() (ret CustomDomainGetStatusRetType) {
	ret, _ = o.GetStatusOk()
	return ret
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetStatusOk() (ret CustomDomainGetStatusRetType, ok bool) {
	return getCustomDomainGetStatusAttributeTypeOk(o.Status)
}

// SetStatus sets field value
func (o *CustomDomain) SetStatus(v CustomDomainGetStatusRetType) {
	setCustomDomainGetStatusAttributeType(&o.Status, v)
}

func (o CustomDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCustomDomainGetErrorsAttributeTypeOk(o.Errors); ok {
		toSerialize["Errors"] = val
	}
	if val, ok := getCustomDomainGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getCustomDomainGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	return toSerialize, nil
}

type NullableCustomDomain struct {
	value *CustomDomain
	isSet bool
}

func (v NullableCustomDomain) Get() *CustomDomain {
	return v.value
}

func (v *NullableCustomDomain) Set(val *CustomDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomain(val *CustomDomain) *NullableCustomDomain {
	return &NullableCustomDomain{value: val, isSet: true}
}

func (v NullableCustomDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

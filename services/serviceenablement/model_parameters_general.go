/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

import (
	"encoding/json"
)

// checks if the ParametersGeneral type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParametersGeneral{}

/*
	types and functions for organizationId
*/

// isNotNullableString
type ParametersGeneralGetOrganizationIdAttributeType = *string

func getParametersGeneralGetOrganizationIdAttributeTypeOk(arg ParametersGeneralGetOrganizationIdAttributeType) (ret ParametersGeneralGetOrganizationIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setParametersGeneralGetOrganizationIdAttributeType(arg *ParametersGeneralGetOrganizationIdAttributeType, val ParametersGeneralGetOrganizationIdRetType) {
	*arg = &val
}

type ParametersGeneralGetOrganizationIdArgType = string
type ParametersGeneralGetOrganizationIdRetType = string

/*
	types and functions for projectName
*/

// isNotNullableString
type ParametersGeneralGetProjectNameAttributeType = *string

func getParametersGeneralGetProjectNameAttributeTypeOk(arg ParametersGeneralGetProjectNameAttributeType) (ret ParametersGeneralGetProjectNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setParametersGeneralGetProjectNameAttributeType(arg *ParametersGeneralGetProjectNameAttributeType, val ParametersGeneralGetProjectNameRetType) {
	*arg = &val
}

type ParametersGeneralGetProjectNameArgType = string
type ParametersGeneralGetProjectNameRetType = string

/*
	types and functions for projectScope
*/

// isEnumRef
type ParametersGeneralGetProjectScopeAttributeType = *string
type ParametersGeneralGetProjectScopeArgType = string
type ParametersGeneralGetProjectScopeRetType = string

func getParametersGeneralGetProjectScopeAttributeTypeOk(arg ParametersGeneralGetProjectScopeAttributeType) (ret ParametersGeneralGetProjectScopeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setParametersGeneralGetProjectScopeAttributeType(arg *ParametersGeneralGetProjectScopeAttributeType, val ParametersGeneralGetProjectScopeRetType) {
	*arg = &val
}

// ParametersGeneral struct for ParametersGeneral
type ParametersGeneral struct {
	OrganizationId ParametersGeneralGetOrganizationIdAttributeType `json:"organizationId,omitempty"`
	ProjectName    ParametersGeneralGetProjectNameAttributeType    `json:"projectName,omitempty"`
	ProjectScope   ParametersGeneralGetProjectScopeAttributeType   `json:"projectScope,omitempty"`
}

// NewParametersGeneral instantiates a new ParametersGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParametersGeneral() *ParametersGeneral {
	this := ParametersGeneral{}
	return &this
}

// NewParametersGeneralWithDefaults instantiates a new ParametersGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersGeneralWithDefaults() *ParametersGeneral {
	this := ParametersGeneral{}
	var projectScope string = "PUBLIC"
	this.ProjectScope = &projectScope
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ParametersGeneral) GetOrganizationId() (res ParametersGeneralGetOrganizationIdRetType) {
	res, _ = o.GetOrganizationIdOk()
	return
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParametersGeneral) GetOrganizationIdOk() (ret ParametersGeneralGetOrganizationIdRetType, ok bool) {
	return getParametersGeneralGetOrganizationIdAttributeTypeOk(o.OrganizationId)
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ParametersGeneral) HasOrganizationId() bool {
	_, ok := o.GetOrganizationIdOk()
	return ok
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ParametersGeneral) SetOrganizationId(v ParametersGeneralGetOrganizationIdRetType) {
	setParametersGeneralGetOrganizationIdAttributeType(&o.OrganizationId, v)
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *ParametersGeneral) GetProjectName() (res ParametersGeneralGetProjectNameRetType) {
	res, _ = o.GetProjectNameOk()
	return
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParametersGeneral) GetProjectNameOk() (ret ParametersGeneralGetProjectNameRetType, ok bool) {
	return getParametersGeneralGetProjectNameAttributeTypeOk(o.ProjectName)
}

// HasProjectName returns a boolean if a field has been set.
func (o *ParametersGeneral) HasProjectName() bool {
	_, ok := o.GetProjectNameOk()
	return ok
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *ParametersGeneral) SetProjectName(v ParametersGeneralGetProjectNameRetType) {
	setParametersGeneralGetProjectNameAttributeType(&o.ProjectName, v)
}

// GetProjectScope returns the ProjectScope field value if set, zero value otherwise.
func (o *ParametersGeneral) GetProjectScope() (res ParametersGeneralGetProjectScopeRetType) {
	res, _ = o.GetProjectScopeOk()
	return
}

// GetProjectScopeOk returns a tuple with the ProjectScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParametersGeneral) GetProjectScopeOk() (ret ParametersGeneralGetProjectScopeRetType, ok bool) {
	return getParametersGeneralGetProjectScopeAttributeTypeOk(o.ProjectScope)
}

// HasProjectScope returns a boolean if a field has been set.
func (o *ParametersGeneral) HasProjectScope() bool {
	_, ok := o.GetProjectScopeOk()
	return ok
}

// SetProjectScope gets a reference to the given string and assigns it to the ProjectScope field.
func (o *ParametersGeneral) SetProjectScope(v ParametersGeneralGetProjectScopeRetType) {
	setParametersGeneralGetProjectScopeAttributeType(&o.ProjectScope, v)
}

func (o ParametersGeneral) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getParametersGeneralGetOrganizationIdAttributeTypeOk(o.OrganizationId); ok {
		toSerialize["OrganizationId"] = val
	}
	if val, ok := getParametersGeneralGetProjectNameAttributeTypeOk(o.ProjectName); ok {
		toSerialize["ProjectName"] = val
	}
	if val, ok := getParametersGeneralGetProjectScopeAttributeTypeOk(o.ProjectScope); ok {
		toSerialize["ProjectScope"] = val
	}
	return toSerialize, nil
}

type NullableParametersGeneral struct {
	value *ParametersGeneral
	isSet bool
}

func (v NullableParametersGeneral) Get() *ParametersGeneral {
	return v.value
}

func (v *NullableParametersGeneral) Set(val *ParametersGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableParametersGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableParametersGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParametersGeneral(val *ParametersGeneral) *NullableParametersGeneral {
	return &NullableParametersGeneral{value: val, isSet: true}
}

func (v NullableParametersGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParametersGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

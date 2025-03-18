/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"encoding/json"
)

// checks if the HandlersInfraGetFlavorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlersInfraGetFlavorsResponse{}

/*
	types and functions for flavors
*/

// isArray
type HandlersInfraGetFlavorsResponseGetFlavorsAttributeType = *[]HandlersInfraFlavor
type HandlersInfraGetFlavorsResponseGetFlavorsArgType = []HandlersInfraFlavor
type HandlersInfraGetFlavorsResponseGetFlavorsRetType = []HandlersInfraFlavor

func getHandlersInfraGetFlavorsResponseGetFlavorsAttributeTypeOk(arg HandlersInfraGetFlavorsResponseGetFlavorsAttributeType) (ret HandlersInfraGetFlavorsResponseGetFlavorsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHandlersInfraGetFlavorsResponseGetFlavorsAttributeType(arg *HandlersInfraGetFlavorsResponseGetFlavorsAttributeType, val HandlersInfraGetFlavorsResponseGetFlavorsRetType) {
	*arg = &val
}

// HandlersInfraGetFlavorsResponse struct for HandlersInfraGetFlavorsResponse
type HandlersInfraGetFlavorsResponse struct {
	Flavors HandlersInfraGetFlavorsResponseGetFlavorsAttributeType `json:"flavors,omitempty"`
}

// NewHandlersInfraGetFlavorsResponse instantiates a new HandlersInfraGetFlavorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlersInfraGetFlavorsResponse() *HandlersInfraGetFlavorsResponse {
	this := HandlersInfraGetFlavorsResponse{}
	return &this
}

// NewHandlersInfraGetFlavorsResponseWithDefaults instantiates a new HandlersInfraGetFlavorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlersInfraGetFlavorsResponseWithDefaults() *HandlersInfraGetFlavorsResponse {
	this := HandlersInfraGetFlavorsResponse{}
	return &this
}

// GetFlavors returns the Flavors field value if set, zero value otherwise.
func (o *HandlersInfraGetFlavorsResponse) GetFlavors() (res HandlersInfraGetFlavorsResponseGetFlavorsRetType) {
	res, _ = o.GetFlavorsOk()
	return
}

// GetFlavorsOk returns a tuple with the Flavors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersInfraGetFlavorsResponse) GetFlavorsOk() (ret HandlersInfraGetFlavorsResponseGetFlavorsRetType, ok bool) {
	return getHandlersInfraGetFlavorsResponseGetFlavorsAttributeTypeOk(o.Flavors)
}

// HasFlavors returns a boolean if a field has been set.
func (o *HandlersInfraGetFlavorsResponse) HasFlavors() bool {
	_, ok := o.GetFlavorsOk()
	return ok
}

// SetFlavors gets a reference to the given []HandlersInfraFlavor and assigns it to the Flavors field.
func (o *HandlersInfraGetFlavorsResponse) SetFlavors(v HandlersInfraGetFlavorsResponseGetFlavorsRetType) {
	setHandlersInfraGetFlavorsResponseGetFlavorsAttributeType(&o.Flavors, v)
}

func (o HandlersInfraGetFlavorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getHandlersInfraGetFlavorsResponseGetFlavorsAttributeTypeOk(o.Flavors); ok {
		toSerialize["Flavors"] = val
	}
	return toSerialize, nil
}

type NullableHandlersInfraGetFlavorsResponse struct {
	value *HandlersInfraGetFlavorsResponse
	isSet bool
}

func (v NullableHandlersInfraGetFlavorsResponse) Get() *HandlersInfraGetFlavorsResponse {
	return v.value
}

func (v *NullableHandlersInfraGetFlavorsResponse) Set(val *HandlersInfraGetFlavorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlersInfraGetFlavorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlersInfraGetFlavorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlersInfraGetFlavorsResponse(val *HandlersInfraGetFlavorsResponse) *NullableHandlersInfraGetFlavorsResponse {
	return &NullableHandlersInfraGetFlavorsResponse{value: val, isSet: true}
}

func (v NullableHandlersInfraGetFlavorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlersInfraGetFlavorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

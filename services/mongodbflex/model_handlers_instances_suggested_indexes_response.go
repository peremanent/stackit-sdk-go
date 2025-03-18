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

// checks if the HandlersInstancesSuggestedIndexesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlersInstancesSuggestedIndexesResponse{}

/*
	types and functions for shapes
*/

// isArray
type HandlersInstancesSuggestedIndexesResponseGetShapesAttributeType = *[]Shape
type HandlersInstancesSuggestedIndexesResponseGetShapesArgType = []Shape
type HandlersInstancesSuggestedIndexesResponseGetShapesRetType = []Shape

func getHandlersInstancesSuggestedIndexesResponseGetShapesAttributeTypeOk(arg HandlersInstancesSuggestedIndexesResponseGetShapesAttributeType) (ret HandlersInstancesSuggestedIndexesResponseGetShapesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHandlersInstancesSuggestedIndexesResponseGetShapesAttributeType(arg *HandlersInstancesSuggestedIndexesResponseGetShapesAttributeType, val HandlersInstancesSuggestedIndexesResponseGetShapesRetType) {
	*arg = &val
}

/*
	types and functions for suggestedIndexes
*/

// isArray
type HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeType = *[]SuggestedIndex
type HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesArgType = []SuggestedIndex
type HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesRetType = []SuggestedIndex

func getHandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeTypeOk(arg HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeType) (ret HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeType(arg *HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeType, val HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesRetType) {
	*arg = &val
}

// HandlersInstancesSuggestedIndexesResponse struct for HandlersInstancesSuggestedIndexesResponse
type HandlersInstancesSuggestedIndexesResponse struct {
	// Documents with information about the query shapes that are served by the suggested indexes.
	Shapes HandlersInstancesSuggestedIndexesResponseGetShapesAttributeType `json:"shapes,omitempty"`
	// Documents with information about the indexes suggested by the Performance Advisor.
	SuggestedIndexes HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeType `json:"suggestedIndexes,omitempty"`
}

// NewHandlersInstancesSuggestedIndexesResponse instantiates a new HandlersInstancesSuggestedIndexesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlersInstancesSuggestedIndexesResponse() *HandlersInstancesSuggestedIndexesResponse {
	this := HandlersInstancesSuggestedIndexesResponse{}
	return &this
}

// NewHandlersInstancesSuggestedIndexesResponseWithDefaults instantiates a new HandlersInstancesSuggestedIndexesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlersInstancesSuggestedIndexesResponseWithDefaults() *HandlersInstancesSuggestedIndexesResponse {
	this := HandlersInstancesSuggestedIndexesResponse{}
	return &this
}

// GetShapes returns the Shapes field value if set, zero value otherwise.
func (o *HandlersInstancesSuggestedIndexesResponse) GetShapes() (res HandlersInstancesSuggestedIndexesResponseGetShapesRetType) {
	res, _ = o.GetShapesOk()
	return
}

// GetShapesOk returns a tuple with the Shapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersInstancesSuggestedIndexesResponse) GetShapesOk() (ret HandlersInstancesSuggestedIndexesResponseGetShapesRetType, ok bool) {
	return getHandlersInstancesSuggestedIndexesResponseGetShapesAttributeTypeOk(o.Shapes)
}

// HasShapes returns a boolean if a field has been set.
func (o *HandlersInstancesSuggestedIndexesResponse) HasShapes() bool {
	_, ok := o.GetShapesOk()
	return ok
}

// SetShapes gets a reference to the given []Shape and assigns it to the Shapes field.
func (o *HandlersInstancesSuggestedIndexesResponse) SetShapes(v HandlersInstancesSuggestedIndexesResponseGetShapesRetType) {
	setHandlersInstancesSuggestedIndexesResponseGetShapesAttributeType(&o.Shapes, v)
}

// GetSuggestedIndexes returns the SuggestedIndexes field value if set, zero value otherwise.
func (o *HandlersInstancesSuggestedIndexesResponse) GetSuggestedIndexes() (res HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesRetType) {
	res, _ = o.GetSuggestedIndexesOk()
	return
}

// GetSuggestedIndexesOk returns a tuple with the SuggestedIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlersInstancesSuggestedIndexesResponse) GetSuggestedIndexesOk() (ret HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesRetType, ok bool) {
	return getHandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeTypeOk(o.SuggestedIndexes)
}

// HasSuggestedIndexes returns a boolean if a field has been set.
func (o *HandlersInstancesSuggestedIndexesResponse) HasSuggestedIndexes() bool {
	_, ok := o.GetSuggestedIndexesOk()
	return ok
}

// SetSuggestedIndexes gets a reference to the given []SuggestedIndex and assigns it to the SuggestedIndexes field.
func (o *HandlersInstancesSuggestedIndexesResponse) SetSuggestedIndexes(v HandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesRetType) {
	setHandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeType(&o.SuggestedIndexes, v)
}

func (o HandlersInstancesSuggestedIndexesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getHandlersInstancesSuggestedIndexesResponseGetShapesAttributeTypeOk(o.Shapes); ok {
		toSerialize["Shapes"] = val
	}
	if val, ok := getHandlersInstancesSuggestedIndexesResponseGetSuggestedIndexesAttributeTypeOk(o.SuggestedIndexes); ok {
		toSerialize["SuggestedIndexes"] = val
	}
	return toSerialize, nil
}

type NullableHandlersInstancesSuggestedIndexesResponse struct {
	value *HandlersInstancesSuggestedIndexesResponse
	isSet bool
}

func (v NullableHandlersInstancesSuggestedIndexesResponse) Get() *HandlersInstancesSuggestedIndexesResponse {
	return v.value
}

func (v *NullableHandlersInstancesSuggestedIndexesResponse) Set(val *HandlersInstancesSuggestedIndexesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlersInstancesSuggestedIndexesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlersInstancesSuggestedIndexesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlersInstancesSuggestedIndexesResponse(val *HandlersInstancesSuggestedIndexesResponse) *NullableHandlersInstancesSuggestedIndexesResponse {
	return &NullableHandlersInstancesSuggestedIndexesResponse{value: val, isSet: true}
}

func (v NullableHandlersInstancesSuggestedIndexesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlersInstancesSuggestedIndexesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

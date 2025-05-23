/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ZoneModelsImportRecordModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneModelsImportRecordModel{}

/*
	types and functions for comment
*/

// isNotNullableString
type ZoneModelsImportRecordModelGetCommentAttributeType = *string

func getZoneModelsImportRecordModelGetCommentAttributeTypeOk(arg ZoneModelsImportRecordModelGetCommentAttributeType) (ret ZoneModelsImportRecordModelGetCommentRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setZoneModelsImportRecordModelGetCommentAttributeType(arg *ZoneModelsImportRecordModelGetCommentAttributeType, val ZoneModelsImportRecordModelGetCommentRetType) {
	*arg = &val
}

type ZoneModelsImportRecordModelGetCommentArgType = string
type ZoneModelsImportRecordModelGetCommentRetType = string

/*
	types and functions for content
*/

// isArray
type ZoneModelsImportRecordModelGetContentAttributeType = *[]string
type ZoneModelsImportRecordModelGetContentArgType = []string
type ZoneModelsImportRecordModelGetContentRetType = []string

func getZoneModelsImportRecordModelGetContentAttributeTypeOk(arg ZoneModelsImportRecordModelGetContentAttributeType) (ret ZoneModelsImportRecordModelGetContentRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setZoneModelsImportRecordModelGetContentAttributeType(arg *ZoneModelsImportRecordModelGetContentAttributeType, val ZoneModelsImportRecordModelGetContentRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type ZoneModelsImportRecordModelGetNameAttributeType = *string

func getZoneModelsImportRecordModelGetNameAttributeTypeOk(arg ZoneModelsImportRecordModelGetNameAttributeType) (ret ZoneModelsImportRecordModelGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setZoneModelsImportRecordModelGetNameAttributeType(arg *ZoneModelsImportRecordModelGetNameAttributeType, val ZoneModelsImportRecordModelGetNameRetType) {
	*arg = &val
}

type ZoneModelsImportRecordModelGetNameArgType = string
type ZoneModelsImportRecordModelGetNameRetType = string

/*
	types and functions for ttl
*/

// isInteger
type ZoneModelsImportRecordModelGetTtlAttributeType = *int64
type ZoneModelsImportRecordModelGetTtlArgType = int64
type ZoneModelsImportRecordModelGetTtlRetType = int64

func getZoneModelsImportRecordModelGetTtlAttributeTypeOk(arg ZoneModelsImportRecordModelGetTtlAttributeType) (ret ZoneModelsImportRecordModelGetTtlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setZoneModelsImportRecordModelGetTtlAttributeType(arg *ZoneModelsImportRecordModelGetTtlAttributeType, val ZoneModelsImportRecordModelGetTtlRetType) {
	*arg = &val
}

/*
	types and functions for type
*/

// isNotNullableString
type ZoneModelsImportRecordModelGetTypeAttributeType = *string

func getZoneModelsImportRecordModelGetTypeAttributeTypeOk(arg ZoneModelsImportRecordModelGetTypeAttributeType) (ret ZoneModelsImportRecordModelGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setZoneModelsImportRecordModelGetTypeAttributeType(arg *ZoneModelsImportRecordModelGetTypeAttributeType, val ZoneModelsImportRecordModelGetTypeRetType) {
	*arg = &val
}

type ZoneModelsImportRecordModelGetTypeArgType = string
type ZoneModelsImportRecordModelGetTypeRetType = string

// ZoneModelsImportRecordModel struct for ZoneModelsImportRecordModel
type ZoneModelsImportRecordModel struct {
	Comment ZoneModelsImportRecordModelGetCommentAttributeType `json:"comment,omitempty"`
	Content ZoneModelsImportRecordModelGetContentAttributeType `json:"content,omitempty"`
	Name    ZoneModelsImportRecordModelGetNameAttributeType    `json:"name,omitempty"`
	// Can be cast to int32 without loss of precision.
	Ttl  ZoneModelsImportRecordModelGetTtlAttributeType  `json:"ttl,omitempty"`
	Type ZoneModelsImportRecordModelGetTypeAttributeType `json:"type,omitempty"`
}

// NewZoneModelsImportRecordModel instantiates a new ZoneModelsImportRecordModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneModelsImportRecordModel() *ZoneModelsImportRecordModel {
	this := ZoneModelsImportRecordModel{}
	return &this
}

// NewZoneModelsImportRecordModelWithDefaults instantiates a new ZoneModelsImportRecordModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneModelsImportRecordModelWithDefaults() *ZoneModelsImportRecordModel {
	this := ZoneModelsImportRecordModel{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ZoneModelsImportRecordModel) GetComment() (res ZoneModelsImportRecordModelGetCommentRetType) {
	res, _ = o.GetCommentOk()
	return
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModelsImportRecordModel) GetCommentOk() (ret ZoneModelsImportRecordModelGetCommentRetType, ok bool) {
	return getZoneModelsImportRecordModelGetCommentAttributeTypeOk(o.Comment)
}

// HasComment returns a boolean if a field has been set.
func (o *ZoneModelsImportRecordModel) HasComment() bool {
	_, ok := o.GetCommentOk()
	return ok
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ZoneModelsImportRecordModel) SetComment(v ZoneModelsImportRecordModelGetCommentRetType) {
	setZoneModelsImportRecordModelGetCommentAttributeType(&o.Comment, v)
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ZoneModelsImportRecordModel) GetContent() (res ZoneModelsImportRecordModelGetContentRetType) {
	res, _ = o.GetContentOk()
	return
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModelsImportRecordModel) GetContentOk() (ret ZoneModelsImportRecordModelGetContentRetType, ok bool) {
	return getZoneModelsImportRecordModelGetContentAttributeTypeOk(o.Content)
}

// HasContent returns a boolean if a field has been set.
func (o *ZoneModelsImportRecordModel) HasContent() bool {
	_, ok := o.GetContentOk()
	return ok
}

// SetContent gets a reference to the given []string and assigns it to the Content field.
func (o *ZoneModelsImportRecordModel) SetContent(v ZoneModelsImportRecordModelGetContentRetType) {
	setZoneModelsImportRecordModelGetContentAttributeType(&o.Content, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneModelsImportRecordModel) GetName() (res ZoneModelsImportRecordModelGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModelsImportRecordModel) GetNameOk() (ret ZoneModelsImportRecordModelGetNameRetType, ok bool) {
	return getZoneModelsImportRecordModelGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *ZoneModelsImportRecordModel) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneModelsImportRecordModel) SetName(v ZoneModelsImportRecordModelGetNameRetType) {
	setZoneModelsImportRecordModelGetNameAttributeType(&o.Name, v)
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ZoneModelsImportRecordModel) GetTtl() (res ZoneModelsImportRecordModelGetTtlRetType) {
	res, _ = o.GetTtlOk()
	return
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModelsImportRecordModel) GetTtlOk() (ret ZoneModelsImportRecordModelGetTtlRetType, ok bool) {
	return getZoneModelsImportRecordModelGetTtlAttributeTypeOk(o.Ttl)
}

// HasTtl returns a boolean if a field has been set.
func (o *ZoneModelsImportRecordModel) HasTtl() bool {
	_, ok := o.GetTtlOk()
	return ok
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *ZoneModelsImportRecordModel) SetTtl(v ZoneModelsImportRecordModelGetTtlRetType) {
	setZoneModelsImportRecordModelGetTtlAttributeType(&o.Ttl, v)
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ZoneModelsImportRecordModel) GetType() (res ZoneModelsImportRecordModelGetTypeRetType) {
	res, _ = o.GetTypeOk()
	return
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModelsImportRecordModel) GetTypeOk() (ret ZoneModelsImportRecordModelGetTypeRetType, ok bool) {
	return getZoneModelsImportRecordModelGetTypeAttributeTypeOk(o.Type)
}

// HasType returns a boolean if a field has been set.
func (o *ZoneModelsImportRecordModel) HasType() bool {
	_, ok := o.GetTypeOk()
	return ok
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ZoneModelsImportRecordModel) SetType(v ZoneModelsImportRecordModelGetTypeRetType) {
	setZoneModelsImportRecordModelGetTypeAttributeType(&o.Type, v)
}

func (o ZoneModelsImportRecordModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getZoneModelsImportRecordModelGetCommentAttributeTypeOk(o.Comment); ok {
		toSerialize["Comment"] = val
	}
	if val, ok := getZoneModelsImportRecordModelGetContentAttributeTypeOk(o.Content); ok {
		toSerialize["Content"] = val
	}
	if val, ok := getZoneModelsImportRecordModelGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getZoneModelsImportRecordModelGetTtlAttributeTypeOk(o.Ttl); ok {
		toSerialize["Ttl"] = val
	}
	if val, ok := getZoneModelsImportRecordModelGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	return toSerialize, nil
}

type NullableZoneModelsImportRecordModel struct {
	value *ZoneModelsImportRecordModel
	isSet bool
}

func (v NullableZoneModelsImportRecordModel) Get() *ZoneModelsImportRecordModel {
	return v.value
}

func (v *NullableZoneModelsImportRecordModel) Set(val *ZoneModelsImportRecordModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneModelsImportRecordModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneModelsImportRecordModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneModelsImportRecordModel(val *ZoneModelsImportRecordModel) *NullableZoneModelsImportRecordModel {
	return &NullableZoneModelsImportRecordModel{value: val, isSet: true}
}

func (v NullableZoneModelsImportRecordModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneModelsImportRecordModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

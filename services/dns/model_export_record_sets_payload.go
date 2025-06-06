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

// checks if the ExportRecordSetsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportRecordSetsPayload{}

/*
	types and functions for exportAsFQDN
*/

// isBoolean
type ExportRecordSetsPayloadgetExportAsFQDNAttributeType = *bool
type ExportRecordSetsPayloadgetExportAsFQDNArgType = bool
type ExportRecordSetsPayloadgetExportAsFQDNRetType = bool

func getExportRecordSetsPayloadgetExportAsFQDNAttributeTypeOk(arg ExportRecordSetsPayloadgetExportAsFQDNAttributeType) (ret ExportRecordSetsPayloadgetExportAsFQDNRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setExportRecordSetsPayloadgetExportAsFQDNAttributeType(arg *ExportRecordSetsPayloadgetExportAsFQDNAttributeType, val ExportRecordSetsPayloadgetExportAsFQDNRetType) {
	*arg = &val
}

/*
	types and functions for format
*/

// isEnumRef
type ExportRecordSetsPayloadGetFormatAttributeType = *string
type ExportRecordSetsPayloadGetFormatArgType = string
type ExportRecordSetsPayloadGetFormatRetType = string

func getExportRecordSetsPayloadGetFormatAttributeTypeOk(arg ExportRecordSetsPayloadGetFormatAttributeType) (ret ExportRecordSetsPayloadGetFormatRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setExportRecordSetsPayloadGetFormatAttributeType(arg *ExportRecordSetsPayloadGetFormatAttributeType, val ExportRecordSetsPayloadGetFormatRetType) {
	*arg = &val
}

// ExportRecordSetsPayload struct for ExportRecordSetsPayload
type ExportRecordSetsPayload struct {
	ExportAsFQDN ExportRecordSetsPayloadgetExportAsFQDNAttributeType `json:"exportAsFQDN,omitempty"`
	Format       ExportRecordSetsPayloadGetFormatAttributeType       `json:"format,omitempty"`
}

// NewExportRecordSetsPayload instantiates a new ExportRecordSetsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportRecordSetsPayload() *ExportRecordSetsPayload {
	this := ExportRecordSetsPayload{}
	return &this
}

// NewExportRecordSetsPayloadWithDefaults instantiates a new ExportRecordSetsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportRecordSetsPayloadWithDefaults() *ExportRecordSetsPayload {
	this := ExportRecordSetsPayload{}
	var exportAsFQDN bool = true
	this.ExportAsFQDN = &exportAsFQDN
	var format string = "csv"
	this.Format = &format
	return &this
}

// GetExportAsFQDN returns the ExportAsFQDN field value if set, zero value otherwise.
func (o *ExportRecordSetsPayload) GetExportAsFQDN() (res ExportRecordSetsPayloadgetExportAsFQDNRetType) {
	res, _ = o.GetExportAsFQDNOk()
	return
}

// GetExportAsFQDNOk returns a tuple with the ExportAsFQDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportRecordSetsPayload) GetExportAsFQDNOk() (ret ExportRecordSetsPayloadgetExportAsFQDNRetType, ok bool) {
	return getExportRecordSetsPayloadgetExportAsFQDNAttributeTypeOk(o.ExportAsFQDN)
}

// HasExportAsFQDN returns a boolean if a field has been set.
func (o *ExportRecordSetsPayload) HasExportAsFQDN() bool {
	_, ok := o.GetExportAsFQDNOk()
	return ok
}

// SetExportAsFQDN gets a reference to the given bool and assigns it to the ExportAsFQDN field.
func (o *ExportRecordSetsPayload) SetExportAsFQDN(v ExportRecordSetsPayloadgetExportAsFQDNRetType) {
	setExportRecordSetsPayloadgetExportAsFQDNAttributeType(&o.ExportAsFQDN, v)
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ExportRecordSetsPayload) GetFormat() (res ExportRecordSetsPayloadGetFormatRetType) {
	res, _ = o.GetFormatOk()
	return
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportRecordSetsPayload) GetFormatOk() (ret ExportRecordSetsPayloadGetFormatRetType, ok bool) {
	return getExportRecordSetsPayloadGetFormatAttributeTypeOk(o.Format)
}

// HasFormat returns a boolean if a field has been set.
func (o *ExportRecordSetsPayload) HasFormat() bool {
	_, ok := o.GetFormatOk()
	return ok
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ExportRecordSetsPayload) SetFormat(v ExportRecordSetsPayloadGetFormatRetType) {
	setExportRecordSetsPayloadGetFormatAttributeType(&o.Format, v)
}

func (o ExportRecordSetsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getExportRecordSetsPayloadgetExportAsFQDNAttributeTypeOk(o.ExportAsFQDN); ok {
		toSerialize["ExportAsFQDN"] = val
	}
	if val, ok := getExportRecordSetsPayloadGetFormatAttributeTypeOk(o.Format); ok {
		toSerialize["Format"] = val
	}
	return toSerialize, nil
}

type NullableExportRecordSetsPayload struct {
	value *ExportRecordSetsPayload
	isSet bool
}

func (v NullableExportRecordSetsPayload) Get() *ExportRecordSetsPayload {
	return v.value
}

func (v *NullableExportRecordSetsPayload) Set(val *ExportRecordSetsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableExportRecordSetsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableExportRecordSetsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportRecordSetsPayload(val *ExportRecordSetsPayload) *NullableExportRecordSetsPayload {
	return &NullableExportRecordSetsPayload{value: val, isSet: true}
}

func (v NullableExportRecordSetsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportRecordSetsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

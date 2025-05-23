/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alb

import (
	"encoding/json"
)

// checks if the UpdateCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCredentialsResponse{}

/*
	types and functions for credential
*/

// isModel
type UpdateCredentialsResponseGetCredentialAttributeType = *CredentialsResponse
type UpdateCredentialsResponseGetCredentialArgType = CredentialsResponse
type UpdateCredentialsResponseGetCredentialRetType = CredentialsResponse

func getUpdateCredentialsResponseGetCredentialAttributeTypeOk(arg UpdateCredentialsResponseGetCredentialAttributeType) (ret UpdateCredentialsResponseGetCredentialRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateCredentialsResponseGetCredentialAttributeType(arg *UpdateCredentialsResponseGetCredentialAttributeType, val UpdateCredentialsResponseGetCredentialRetType) {
	*arg = &val
}

// UpdateCredentialsResponse struct for UpdateCredentialsResponse
type UpdateCredentialsResponse struct {
	Credential UpdateCredentialsResponseGetCredentialAttributeType `json:"credential,omitempty"`
}

// NewUpdateCredentialsResponse instantiates a new UpdateCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCredentialsResponse() *UpdateCredentialsResponse {
	this := UpdateCredentialsResponse{}
	return &this
}

// NewUpdateCredentialsResponseWithDefaults instantiates a new UpdateCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCredentialsResponseWithDefaults() *UpdateCredentialsResponse {
	this := UpdateCredentialsResponse{}
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *UpdateCredentialsResponse) GetCredential() (res UpdateCredentialsResponseGetCredentialRetType) {
	res, _ = o.GetCredentialOk()
	return
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialsResponse) GetCredentialOk() (ret UpdateCredentialsResponseGetCredentialRetType, ok bool) {
	return getUpdateCredentialsResponseGetCredentialAttributeTypeOk(o.Credential)
}

// HasCredential returns a boolean if a field has been set.
func (o *UpdateCredentialsResponse) HasCredential() bool {
	_, ok := o.GetCredentialOk()
	return ok
}

// SetCredential gets a reference to the given CredentialsResponse and assigns it to the Credential field.
func (o *UpdateCredentialsResponse) SetCredential(v UpdateCredentialsResponseGetCredentialRetType) {
	setUpdateCredentialsResponseGetCredentialAttributeType(&o.Credential, v)
}

func (o UpdateCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateCredentialsResponseGetCredentialAttributeTypeOk(o.Credential); ok {
		toSerialize["Credential"] = val
	}
	return toSerialize, nil
}

type NullableUpdateCredentialsResponse struct {
	value *UpdateCredentialsResponse
	isSet bool
}

func (v NullableUpdateCredentialsResponse) Get() *UpdateCredentialsResponse {
	return v.value
}

func (v *NullableUpdateCredentialsResponse) Set(val *UpdateCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCredentialsResponse(val *UpdateCredentialsResponse) *NullableUpdateCredentialsResponse {
	return &NullableUpdateCredentialsResponse{value: val, isSet: true}
}

func (v NullableUpdateCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

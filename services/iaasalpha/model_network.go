/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"time"
)

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network Object that represents a network.
type Network struct {
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The gateway of a network. If not specified the first IP of the network will be assigned as the gateway. If 'null' is sent, then the network doesn't have a gateway.
	Gateway *NullableString `json:"gateway,omitempty"`
	// The gateway of a network. If not specified the first IP of the network will be assigned as the gateway. If 'null' is sent, then the network doesn't have a gateway.
	Gatewayv6 *NullableString `json:"gatewayv6,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// A list containing DNS Servers/Nameservers for IPv4.
	Nameservers *[]string `json:"nameservers,omitempty"`
	// A list containing DNS Servers/Nameservers for IPv6.
	NameserversV6 *[]string `json:"nameserversV6,omitempty"`
	// Universally Unique Identifier (UUID).
	// REQUIRED
	NetworkId  *string   `json:"networkId"`
	Prefixes   *[]string `json:"prefixes,omitempty"`
	PrefixesV6 *[]string `json:"prefixesV6,omitempty"`
	// Object that represents an IP address.
	PublicIp *string `json:"publicIp,omitempty"`
	// Shows if the network is routed and therefore accessible from other networks.
	Routed *bool `json:"routed,omitempty"`
	// The state of a resource object. Possible values: `CREATING`, `CREATED`, `DELETING`, `DELETED`, `FAILED`, `UPDATED`, `UPDATING`.
	// REQUIRED
	State *string `json:"state"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _Network Network

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork(name *string, networkId *string, state *string) *Network {
	this := Network{}
	this.Name = name
	this.NetworkId = networkId
	this.State = state
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Network) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Network) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Network) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetGateway() *string {
	if o == nil || IsNil(o.Gateway) || IsNil(o.Gateway.Get()) {
		var ret *string
		return ret
	}
	return o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *Network) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *Network) SetGateway(v *string) {
	if IsNil(o.Gateway) {
		o.Gateway = new(NullableString)
	}
	o.Gateway.Set(v)
}

// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *Network) SetGatewayNil() {
	if IsNil(o.Gateway) {
		o.Gateway = new(NullableString)
	}
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *Network) UnsetGateway() {
	if IsNil(o.Gateway) {
		o.Gateway = new(NullableString)
	}
	o.Gateway.Unset()
}

// GetGatewayv6 returns the Gatewayv6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetGatewayv6() *string {
	if o == nil || IsNil(o.Gatewayv6) || IsNil(o.Gatewayv6.Get()) {
		var ret *string
		return ret
	}
	return o.Gatewayv6.Get()
}

// GetGatewayv6Ok returns a tuple with the Gatewayv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetGatewayv6Ok() (*string, bool) {
	if o == nil || IsNil(o.Gatewayv6) {
		return nil, false
	}
	return o.Gatewayv6.Get(), o.Gatewayv6.IsSet()
}

// HasGatewayv6 returns a boolean if a field has been set.
func (o *Network) HasGatewayv6() bool {
	if o != nil && !IsNil(o.Gatewayv6) && o.Gatewayv6.IsSet() {
		return true
	}

	return false
}

// SetGatewayv6 gets a reference to the given string and assigns it to the Gatewayv6 field.
func (o *Network) SetGatewayv6(v *string) {
	if IsNil(o.Gatewayv6) {
		o.Gatewayv6 = new(NullableString)
	}
	o.Gatewayv6.Set(v)
}

// SetGatewayv6Nil sets the value for Gatewayv6 to be an explicit nil
func (o *Network) SetGatewayv6Nil() {
	if IsNil(o.Gatewayv6) {
		o.Gatewayv6 = new(NullableString)
	}
	o.Gatewayv6.Set(nil)
}

// UnsetGatewayv6 ensures that no value is present for Gatewayv6, not even an explicit nil
func (o *Network) UnsetGatewayv6() {
	if IsNil(o.Gatewayv6) {
		o.Gatewayv6 = new(NullableString)
	}
	o.Gatewayv6.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Network) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Network) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *Network) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *Network) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Network) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Network) SetName(v *string) {
	o.Name = v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *Network) GetNameservers() *[]string {
	if o == nil || IsNil(o.Nameservers) {
		var ret *[]string
		return ret
	}
	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetNameserversOk() (*[]string, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *Network) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *Network) SetNameservers(v *[]string) {
	o.Nameservers = v
}

// GetNameserversV6 returns the NameserversV6 field value if set, zero value otherwise.
func (o *Network) GetNameserversV6() *[]string {
	if o == nil || IsNil(o.NameserversV6) {
		var ret *[]string
		return ret
	}
	return o.NameserversV6
}

// GetNameserversV6Ok returns a tuple with the NameserversV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetNameserversV6Ok() (*[]string, bool) {
	if o == nil || IsNil(o.NameserversV6) {
		return nil, false
	}
	return o.NameserversV6, true
}

// HasNameserversV6 returns a boolean if a field has been set.
func (o *Network) HasNameserversV6() bool {
	if o != nil && !IsNil(o.NameserversV6) {
		return true
	}

	return false
}

// SetNameserversV6 gets a reference to the given []string and assigns it to the NameserversV6 field.
func (o *Network) SetNameserversV6(v *[]string) {
	o.NameserversV6 = v
}

// GetNetworkId returns the NetworkId field value
func (o *Network) GetNetworkId() *string {
	if o == nil || IsNil(o.NetworkId) {
		var ret *string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *Network) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkId, true
}

// SetNetworkId sets field value
func (o *Network) SetNetworkId(v *string) {
	o.NetworkId = v
}

// GetPrefixes returns the Prefixes field value if set, zero value otherwise.
func (o *Network) GetPrefixes() *[]string {
	if o == nil || IsNil(o.Prefixes) {
		var ret *[]string
		return ret
	}
	return o.Prefixes
}

// GetPrefixesOk returns a tuple with the Prefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetPrefixesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Prefixes) {
		return nil, false
	}
	return o.Prefixes, true
}

// HasPrefixes returns a boolean if a field has been set.
func (o *Network) HasPrefixes() bool {
	if o != nil && !IsNil(o.Prefixes) {
		return true
	}

	return false
}

// SetPrefixes gets a reference to the given []string and assigns it to the Prefixes field.
func (o *Network) SetPrefixes(v *[]string) {
	o.Prefixes = v
}

// GetPrefixesV6 returns the PrefixesV6 field value if set, zero value otherwise.
func (o *Network) GetPrefixesV6() *[]string {
	if o == nil || IsNil(o.PrefixesV6) {
		var ret *[]string
		return ret
	}
	return o.PrefixesV6
}

// GetPrefixesV6Ok returns a tuple with the PrefixesV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetPrefixesV6Ok() (*[]string, bool) {
	if o == nil || IsNil(o.PrefixesV6) {
		return nil, false
	}
	return o.PrefixesV6, true
}

// HasPrefixesV6 returns a boolean if a field has been set.
func (o *Network) HasPrefixesV6() bool {
	if o != nil && !IsNil(o.PrefixesV6) {
		return true
	}

	return false
}

// SetPrefixesV6 gets a reference to the given []string and assigns it to the PrefixesV6 field.
func (o *Network) SetPrefixesV6(v *[]string) {
	o.PrefixesV6 = v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *Network) GetPublicIp() *string {
	if o == nil || IsNil(o.PublicIp) {
		var ret *string
		return ret
	}
	return o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *Network) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *Network) SetPublicIp(v *string) {
	o.PublicIp = v
}

// GetRouted returns the Routed field value if set, zero value otherwise.
func (o *Network) GetRouted() *bool {
	if o == nil || IsNil(o.Routed) {
		var ret *bool
		return ret
	}
	return o.Routed
}

// GetRoutedOk returns a tuple with the Routed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetRoutedOk() (*bool, bool) {
	if o == nil || IsNil(o.Routed) {
		return nil, false
	}
	return o.Routed, true
}

// HasRouted returns a boolean if a field has been set.
func (o *Network) HasRouted() bool {
	if o != nil && !IsNil(o.Routed) {
		return true
	}

	return false
}

// SetRouted gets a reference to the given bool and assigns it to the Routed field.
func (o *Network) SetRouted(v *bool) {
	o.Routed = v
}

// GetState returns the State field value
func (o *Network) GetState() *string {
	if o == nil || IsNil(o.State) {
		var ret *string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Network) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State, true
}

// SetState sets field value
func (o *Network) SetState(v *string) {
	o.State = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Network) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Network) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Network) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if o.Gatewayv6.IsSet() {
		toSerialize["gatewayv6"] = o.Gatewayv6.Get()
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	if !IsNil(o.NameserversV6) {
		toSerialize["nameserversV6"] = o.NameserversV6
	}
	toSerialize["networkId"] = o.NetworkId
	if !IsNil(o.Prefixes) {
		toSerialize["prefixes"] = o.Prefixes
	}
	if !IsNil(o.PrefixesV6) {
		toSerialize["prefixesV6"] = o.PrefixesV6
	}
	if !IsNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if !IsNil(o.Routed) {
		toSerialize["routed"] = o.Routed
	}
	toSerialize["state"] = o.State
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

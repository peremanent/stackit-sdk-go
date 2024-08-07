/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type Network struct {
	// Openstack network ID
	NetworkId *string `json:"networkId,omitempty"`
	// The role defines how the load balancer is using the network. Currently only ROLE_LISTENERS_AND_TARGETS is supported.
	Role *string `json:"role,omitempty"`
}

/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type LoadBalancer struct {
	// Reports all errors a load balancer has.
	Errors *[]LoadBalancerError `json:"errors,omitempty"`
	// External load balancer IP address where this load balancer is exposed. Not changeable after creation.
	ExternalAddress *string `json:"externalAddress,omitempty"`
	// There is a maximum listener count of 20.  Port and protocol limitations:  - UDP listeners cannot have the same port. - TCP-derived listeners cannot have the same port. A TCP-derived listener is any listener that listens on a TCP port. As of now those are: TCP, TCP_PROXY, and PROTOCOL_TLS_PASSTHROUGH. The only exception is, if all listeners for the same port are PROTOCOL_TLS_PASSTHROUGH. - PROTOCOL_TLS_PASSTHROUGH listeners cannot have the same port and at least one common domain name. - PROTOCOL_TLS_PASSTHROUGH listeners can have the same domain name and different ports though (e.g. ports 443 and 8443 for domain example.com). - PROTOCOL_TLS_PASSTHROUGH listeners without a domain name serve as a default listener and you can have only one default listener.
	Listeners *[]Listener `json:"listeners,omitempty"`
	// Load balancer name. Not changeable after creation.
	Name *string `json:"name,omitempty"`
	// List of networks that listeners and targets reside in. Currently limited to one. Not changeable after creation.
	Networks *[]Network           `json:"networks,omitempty"`
	Options  *LoadBalancerOptions `json:"options,omitempty"`
	// Service Plan configures the size of the Load Balancer. Currently supported plans are p10, p50, p250 and p750. This list can change in the future where plan ids will be removed and new plans by added. That is the reason this is not an enum.
	PlanId *string `json:"planId,omitempty"`
	// Transient private load balancer IP address that can change any time.
	PrivateAddress *string `json:"privateAddress,omitempty"`
	Status         *string `json:"status,omitempty"`
	// List of all target pools which will be used in the load balancer. Limited to 20.
	TargetPools *[]TargetPool `json:"targetPools,omitempty"`
	// Load balancer resource version. Must be empty or unset for creating load balancers, non-empty for updating load balancers. Semantics: While retrieving load balancers, this is the current version of this load balancer resource that changes during updates of the load balancers. On updates this field specified the load balancer version you calculated your update for instead of the future version to enable concurrency safe updates. Update calls will then report the new version in their result as you would see with a load balancer retrieval call later. There exist no total order of the version, so you can only compare it for equality, but not for less/greater than another version. Since the creation of load balancer is always intended to create the first version of it, there should be no existing version. That's why this field must by empty of not present in that case.
	Version *string `json:"version,omitempty"`
}

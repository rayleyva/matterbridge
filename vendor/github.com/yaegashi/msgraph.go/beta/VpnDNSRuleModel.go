// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VpnDNSRule undocumented
type VpnDNSRule struct {
	// Object is the base model of VpnDNSRule
	Object
	// Name Name.
	Name *string `json:"name,omitempty"`
	// Servers Servers.
	Servers []string `json:"servers,omitempty"`
	// ProxyServerURI Proxy Server Uri.
	ProxyServerURI *string `json:"proxyServerUri,omitempty"`
	// AutoTrigger Automatically connect to the VPN when the device connects to this domain: Default False.
	AutoTrigger *bool `json:"autoTrigger,omitempty"`
	// Persistent Keep this rule active even when the VPN is not connected: Default False
	Persistent *bool `json:"persistent,omitempty"`
}

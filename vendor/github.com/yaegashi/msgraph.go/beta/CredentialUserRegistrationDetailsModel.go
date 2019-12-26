// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CredentialUserRegistrationDetails undocumented
type CredentialUserRegistrationDetails struct {
	// Entity is the base model of CredentialUserRegistrationDetails
	Entity
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserDisplayName undocumented
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// AuthMethods undocumented
	AuthMethods []RegistrationAuthMethod `json:"authMethods,omitempty"`
	// IsRegistered undocumented
	IsRegistered *bool `json:"isRegistered,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// IsCapable undocumented
	IsCapable *bool `json:"isCapable,omitempty"`
	// IsMFARegistered undocumented
	IsMFARegistered *bool `json:"isMfaRegistered,omitempty"`
}

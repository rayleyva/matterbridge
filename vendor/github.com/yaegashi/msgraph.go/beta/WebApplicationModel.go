// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WebApplication undocumented
type WebApplication struct {
	// Object is the base model of WebApplication
	Object
	// HomePageURL undocumented
	HomePageURL *string `json:"homePageUrl,omitempty"`
	// RedirectUris undocumented
	RedirectUris []string `json:"redirectUris,omitempty"`
	// Oauth2AllowImplicitFlow undocumented
	Oauth2AllowImplicitFlow *bool `json:"oauth2AllowImplicitFlow,omitempty"`
	// LogoutURL undocumented
	LogoutURL *string `json:"logoutUrl,omitempty"`
	// ImplicitGrantSettings undocumented
	ImplicitGrantSettings *ImplicitGrantSettings `json:"implicitGrantSettings,omitempty"`
}

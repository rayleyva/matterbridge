// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// NonEapAuthenticationMethodForEapTtlsType undocumented
type NonEapAuthenticationMethodForEapTtlsType int

const (
	// NonEapAuthenticationMethodForEapTtlsTypeVUnencryptedPassword undocumented
	NonEapAuthenticationMethodForEapTtlsTypeVUnencryptedPassword NonEapAuthenticationMethodForEapTtlsType = 0
	// NonEapAuthenticationMethodForEapTtlsTypeVChallengeHandshakeAuthenticationProtocol undocumented
	NonEapAuthenticationMethodForEapTtlsTypeVChallengeHandshakeAuthenticationProtocol NonEapAuthenticationMethodForEapTtlsType = 1
	// NonEapAuthenticationMethodForEapTtlsTypeVMicrosoftChap undocumented
	NonEapAuthenticationMethodForEapTtlsTypeVMicrosoftChap NonEapAuthenticationMethodForEapTtlsType = 2
	// NonEapAuthenticationMethodForEapTtlsTypeVMicrosoftChapVersionTwo undocumented
	NonEapAuthenticationMethodForEapTtlsTypeVMicrosoftChapVersionTwo NonEapAuthenticationMethodForEapTtlsType = 3
)

// NonEapAuthenticationMethodForEapTtlsTypePUnencryptedPassword returns a pointer to NonEapAuthenticationMethodForEapTtlsTypeVUnencryptedPassword
func NonEapAuthenticationMethodForEapTtlsTypePUnencryptedPassword() *NonEapAuthenticationMethodForEapTtlsType {
	v := NonEapAuthenticationMethodForEapTtlsTypeVUnencryptedPassword
	return &v
}

// NonEapAuthenticationMethodForEapTtlsTypePChallengeHandshakeAuthenticationProtocol returns a pointer to NonEapAuthenticationMethodForEapTtlsTypeVChallengeHandshakeAuthenticationProtocol
func NonEapAuthenticationMethodForEapTtlsTypePChallengeHandshakeAuthenticationProtocol() *NonEapAuthenticationMethodForEapTtlsType {
	v := NonEapAuthenticationMethodForEapTtlsTypeVChallengeHandshakeAuthenticationProtocol
	return &v
}

// NonEapAuthenticationMethodForEapTtlsTypePMicrosoftChap returns a pointer to NonEapAuthenticationMethodForEapTtlsTypeVMicrosoftChap
func NonEapAuthenticationMethodForEapTtlsTypePMicrosoftChap() *NonEapAuthenticationMethodForEapTtlsType {
	v := NonEapAuthenticationMethodForEapTtlsTypeVMicrosoftChap
	return &v
}

// NonEapAuthenticationMethodForEapTtlsTypePMicrosoftChapVersionTwo returns a pointer to NonEapAuthenticationMethodForEapTtlsTypeVMicrosoftChapVersionTwo
func NonEapAuthenticationMethodForEapTtlsTypePMicrosoftChapVersionTwo() *NonEapAuthenticationMethodForEapTtlsType {
	v := NonEapAuthenticationMethodForEapTtlsTypeVMicrosoftChapVersionTwo
	return &v
}

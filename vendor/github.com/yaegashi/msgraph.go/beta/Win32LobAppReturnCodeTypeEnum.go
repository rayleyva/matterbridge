// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Win32LobAppReturnCodeType undocumented
type Win32LobAppReturnCodeType int

const (
	// Win32LobAppReturnCodeTypeVFailed undocumented
	Win32LobAppReturnCodeTypeVFailed Win32LobAppReturnCodeType = 0
	// Win32LobAppReturnCodeTypeVSuccess undocumented
	Win32LobAppReturnCodeTypeVSuccess Win32LobAppReturnCodeType = 1
	// Win32LobAppReturnCodeTypeVSoftReboot undocumented
	Win32LobAppReturnCodeTypeVSoftReboot Win32LobAppReturnCodeType = 2
	// Win32LobAppReturnCodeTypeVHardReboot undocumented
	Win32LobAppReturnCodeTypeVHardReboot Win32LobAppReturnCodeType = 3
	// Win32LobAppReturnCodeTypeVRetry undocumented
	Win32LobAppReturnCodeTypeVRetry Win32LobAppReturnCodeType = 4
)

// Win32LobAppReturnCodeTypePFailed returns a pointer to Win32LobAppReturnCodeTypeVFailed
func Win32LobAppReturnCodeTypePFailed() *Win32LobAppReturnCodeType {
	v := Win32LobAppReturnCodeTypeVFailed
	return &v
}

// Win32LobAppReturnCodeTypePSuccess returns a pointer to Win32LobAppReturnCodeTypeVSuccess
func Win32LobAppReturnCodeTypePSuccess() *Win32LobAppReturnCodeType {
	v := Win32LobAppReturnCodeTypeVSuccess
	return &v
}

// Win32LobAppReturnCodeTypePSoftReboot returns a pointer to Win32LobAppReturnCodeTypeVSoftReboot
func Win32LobAppReturnCodeTypePSoftReboot() *Win32LobAppReturnCodeType {
	v := Win32LobAppReturnCodeTypeVSoftReboot
	return &v
}

// Win32LobAppReturnCodeTypePHardReboot returns a pointer to Win32LobAppReturnCodeTypeVHardReboot
func Win32LobAppReturnCodeTypePHardReboot() *Win32LobAppReturnCodeType {
	v := Win32LobAppReturnCodeTypeVHardReboot
	return &v
}

// Win32LobAppReturnCodeTypePRetry returns a pointer to Win32LobAppReturnCodeTypeVRetry
func Win32LobAppReturnCodeTypePRetry() *Win32LobAppReturnCodeType {
	v := Win32LobAppReturnCodeTypeVRetry
	return &v
}

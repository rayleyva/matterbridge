// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsStartMenuAppListVisibilityType undocumented
type WindowsStartMenuAppListVisibilityType int

const (
	// WindowsStartMenuAppListVisibilityTypeVUserDefined undocumented
	WindowsStartMenuAppListVisibilityTypeVUserDefined WindowsStartMenuAppListVisibilityType = 0
	// WindowsStartMenuAppListVisibilityTypeVCollapse undocumented
	WindowsStartMenuAppListVisibilityTypeVCollapse WindowsStartMenuAppListVisibilityType = 1
	// WindowsStartMenuAppListVisibilityTypeVRemove undocumented
	WindowsStartMenuAppListVisibilityTypeVRemove WindowsStartMenuAppListVisibilityType = 2
	// WindowsStartMenuAppListVisibilityTypeVDisableSettingsApp undocumented
	WindowsStartMenuAppListVisibilityTypeVDisableSettingsApp WindowsStartMenuAppListVisibilityType = 4
)

// WindowsStartMenuAppListVisibilityTypePUserDefined returns a pointer to WindowsStartMenuAppListVisibilityTypeVUserDefined
func WindowsStartMenuAppListVisibilityTypePUserDefined() *WindowsStartMenuAppListVisibilityType {
	v := WindowsStartMenuAppListVisibilityTypeVUserDefined
	return &v
}

// WindowsStartMenuAppListVisibilityTypePCollapse returns a pointer to WindowsStartMenuAppListVisibilityTypeVCollapse
func WindowsStartMenuAppListVisibilityTypePCollapse() *WindowsStartMenuAppListVisibilityType {
	v := WindowsStartMenuAppListVisibilityTypeVCollapse
	return &v
}

// WindowsStartMenuAppListVisibilityTypePRemove returns a pointer to WindowsStartMenuAppListVisibilityTypeVRemove
func WindowsStartMenuAppListVisibilityTypePRemove() *WindowsStartMenuAppListVisibilityType {
	v := WindowsStartMenuAppListVisibilityTypeVRemove
	return &v
}

// WindowsStartMenuAppListVisibilityTypePDisableSettingsApp returns a pointer to WindowsStartMenuAppListVisibilityTypeVDisableSettingsApp
func WindowsStartMenuAppListVisibilityTypePDisableSettingsApp() *WindowsStartMenuAppListVisibilityType {
	v := WindowsStartMenuAppListVisibilityTypeVDisableSettingsApp
	return &v
}

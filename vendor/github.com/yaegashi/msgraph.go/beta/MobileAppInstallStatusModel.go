// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// MobileAppInstallStatus Contains properties for the installation state of a mobile app for a device.
type MobileAppInstallStatus struct {
	// Entity is the base model of MobileAppInstallStatus
	Entity
	// DeviceName Device name
	DeviceName *string `json:"deviceName,omitempty"`
	// DeviceID Device ID
	DeviceID *string `json:"deviceId,omitempty"`
	// LastSyncDateTime Last sync date time
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// MobileAppInstallStatusValue The install state of the app.
	MobileAppInstallStatusValue *ResultantAppState `json:"mobileAppInstallStatusValue,omitempty"`
	// InstallState The install state of the app.
	InstallState *ResultantAppState `json:"installState,omitempty"`
	// InstallStateDetail The install state detail of the app.
	InstallStateDetail *ResultantAppStateDetail `json:"installStateDetail,omitempty"`
	// ErrorCode The error code for install or uninstall failures.
	ErrorCode *int `json:"errorCode,omitempty"`
	// OsVersion OS Version
	OsVersion *string `json:"osVersion,omitempty"`
	// OsDescription OS Description
	OsDescription *string `json:"osDescription,omitempty"`
	// UserName Device User Name
	UserName *string `json:"userName,omitempty"`
	// UserPrincipalName User Principal Name
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// DisplayVersion Human readable version of the application
	DisplayVersion *string `json:"displayVersion,omitempty"`
	// App undocumented
	App *MobileApp `json:"app,omitempty"`
}

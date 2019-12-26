// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AppleUserInitiatedEnrollmentProfile The enrollmentProfile resource represents a collection of configurations which must be provided pre-enrollment to enable enrolling certain devices whose identities have been pre-staged. Pre-staged device identities are assigned to this type of profile to apply the profile's configurations at enrollment of the corresponding device.
type AppleUserInitiatedEnrollmentProfile struct {
	// Entity is the base model of AppleUserInitiatedEnrollmentProfile
	Entity
	// DefaultEnrollmentType The default profile enrollment type.
	DefaultEnrollmentType *AppleUserInitiatedEnrollmentType `json:"defaultEnrollmentType,omitempty"`
	// AvailableEnrollmentTypeOptions List of available enrollment type options
	AvailableEnrollmentTypeOptions []AppleOwnerTypeEnrollmentType `json:"availableEnrollmentTypeOptions,omitempty"`
	// DisplayName Name of the profile
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the profile
	Description *string `json:"description,omitempty"`
	// Priority Priority, 0 is highest
	Priority *int `json:"priority,omitempty"`
	// Platform The platform of the Device.
	Platform *DevicePlatformType `json:"platform,omitempty"`
	// CreatedDateTime Profile creation time
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime Profile last modified time
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Assignments undocumented
	Assignments []AppleEnrollmentProfileAssignment `json:"assignments,omitempty"`
}

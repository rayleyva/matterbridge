// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SchedulingGroup undocumented
type SchedulingGroup struct {
	// ChangeTrackedEntity is the base model of SchedulingGroup
	ChangeTrackedEntity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
	// UserIDs undocumented
	UserIDs []string `json:"userIds,omitempty"`
}

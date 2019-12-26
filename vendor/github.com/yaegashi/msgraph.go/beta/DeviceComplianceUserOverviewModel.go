// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DeviceComplianceUserOverview undocumented
type DeviceComplianceUserOverview struct {
	// Entity is the base model of DeviceComplianceUserOverview
	Entity
	// PendingCount Number of pending Users
	PendingCount *int `json:"pendingCount,omitempty"`
	// NotApplicableCount Number of not applicable users
	NotApplicableCount *int `json:"notApplicableCount,omitempty"`
	// SuccessCount Number of succeeded Users
	SuccessCount *int `json:"successCount,omitempty"`
	// ErrorCount Number of error Users
	ErrorCount *int `json:"errorCount,omitempty"`
	// FailedCount Number of failed Users
	FailedCount *int `json:"failedCount,omitempty"`
	// ConflictCount Number of users in conflict
	ConflictCount *int `json:"conflictCount,omitempty"`
	// LastUpdateDateTime Last update time
	LastUpdateDateTime *time.Time `json:"lastUpdateDateTime,omitempty"`
	// ConfigurationVersion Version of the policy for that overview
	ConfigurationVersion *int `json:"configurationVersion,omitempty"`
}

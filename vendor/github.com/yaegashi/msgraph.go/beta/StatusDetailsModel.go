// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// StatusDetails undocumented
type StatusDetails struct {
	// StatusBase is the base model of StatusDetails
	StatusBase
	// ErrorCode undocumented
	ErrorCode *string `json:"errorCode,omitempty"`
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// AdditionalDetails undocumented
	AdditionalDetails *string `json:"additionalDetails,omitempty"`
	// ErrorCategory undocumented
	ErrorCategory *string `json:"errorCategory,omitempty"`
	// RecommendedAction undocumented
	RecommendedAction *string `json:"recommendedAction,omitempty"`
}

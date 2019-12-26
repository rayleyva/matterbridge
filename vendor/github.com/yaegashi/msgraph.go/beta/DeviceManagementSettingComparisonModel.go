// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementSettingComparison undocumented
type DeviceManagementSettingComparison struct {
	// Object is the base model of DeviceManagementSettingComparison
	Object
	// ID The setting ID
	ID *string `json:"id,omitempty"`
	// DisplayName The setting's display name
	DisplayName *string `json:"displayName,omitempty"`
	// DefinitionID The ID of the setting definition for this instance
	DefinitionID *string `json:"definitionId,omitempty"`
	// CurrentValueJSON JSON representation of current intent (or) template setting's value
	CurrentValueJSON *string `json:"currentValueJson,omitempty"`
	// NewValueJSON JSON representation of new template setting's value
	NewValueJSON *string `json:"newValueJson,omitempty"`
	// ComparisonResult Setting comparison result
	ComparisonResult *DeviceManagementComparisonResult `json:"comparisonResult,omitempty"`
}

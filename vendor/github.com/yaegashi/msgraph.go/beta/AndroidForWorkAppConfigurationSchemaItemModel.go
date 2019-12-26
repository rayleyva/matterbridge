// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkAppConfigurationSchemaItem undocumented
type AndroidForWorkAppConfigurationSchemaItem struct {
	// Object is the base model of AndroidForWorkAppConfigurationSchemaItem
	Object
	// SchemaItemKey Unique key the application uses to identify the item
	SchemaItemKey *string `json:"schemaItemKey,omitempty"`
	// DisplayName Human readable name
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of what the item controls within the application
	Description *string `json:"description,omitempty"`
	// DefaultBoolValue Default value for boolean type items, if specified by the app developer
	DefaultBoolValue *bool `json:"defaultBoolValue,omitempty"`
	// DefaultIntValue Default value for integer type items, if specified by the app developer
	DefaultIntValue *int `json:"defaultIntValue,omitempty"`
	// DefaultStringValue Default value for string type items, if specified by the app developer
	DefaultStringValue *string `json:"defaultStringValue,omitempty"`
	// DefaultStringArrayValue Default value for string array type items, if specified by the app developer
	DefaultStringArrayValue []string `json:"defaultStringArrayValue,omitempty"`
	// DataType The type of value this item describes
	DataType *AndroidForWorkAppConfigurationSchemaItemDataType `json:"dataType,omitempty"`
	// Selections List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
	Selections []KeyValuePair `json:"selections,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "encoding/json"

// WorkbookChartPoint undocumented
type WorkbookChartPoint struct {
	// Entity is the base model of WorkbookChartPoint
	Entity
	// Value undocumented
	Value json.RawMessage `json:"value,omitempty"`
	// Format undocumented
	Format *WorkbookChartPointFormat `json:"format,omitempty"`
}

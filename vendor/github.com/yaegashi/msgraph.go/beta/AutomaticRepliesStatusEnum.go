// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AutomaticRepliesStatus undocumented
type AutomaticRepliesStatus int

const (
	// AutomaticRepliesStatusVDisabled undocumented
	AutomaticRepliesStatusVDisabled AutomaticRepliesStatus = 0
	// AutomaticRepliesStatusVAlwaysEnabled undocumented
	AutomaticRepliesStatusVAlwaysEnabled AutomaticRepliesStatus = 1
	// AutomaticRepliesStatusVScheduled undocumented
	AutomaticRepliesStatusVScheduled AutomaticRepliesStatus = 2
)

// AutomaticRepliesStatusPDisabled returns a pointer to AutomaticRepliesStatusVDisabled
func AutomaticRepliesStatusPDisabled() *AutomaticRepliesStatus {
	v := AutomaticRepliesStatusVDisabled
	return &v
}

// AutomaticRepliesStatusPAlwaysEnabled returns a pointer to AutomaticRepliesStatusVAlwaysEnabled
func AutomaticRepliesStatusPAlwaysEnabled() *AutomaticRepliesStatus {
	v := AutomaticRepliesStatusVAlwaysEnabled
	return &v
}

// AutomaticRepliesStatusPScheduled returns a pointer to AutomaticRepliesStatusVScheduled
func AutomaticRepliesStatusPScheduled() *AutomaticRepliesStatus {
	v := AutomaticRepliesStatusVScheduled
	return &v
}

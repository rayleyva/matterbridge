// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SecurityBaselineDeviceStateRequestBuilder is request builder for SecurityBaselineDeviceState
type SecurityBaselineDeviceStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityBaselineDeviceStateRequest
func (b *SecurityBaselineDeviceStateRequestBuilder) Request() *SecurityBaselineDeviceStateRequest {
	return &SecurityBaselineDeviceStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityBaselineDeviceStateRequest is request for SecurityBaselineDeviceState
type SecurityBaselineDeviceStateRequest struct{ BaseRequest }

// Get performs GET request for SecurityBaselineDeviceState
func (r *SecurityBaselineDeviceStateRequest) Get(ctx context.Context) (resObj *SecurityBaselineDeviceState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityBaselineDeviceState
func (r *SecurityBaselineDeviceStateRequest) Update(ctx context.Context, reqObj *SecurityBaselineDeviceState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityBaselineDeviceState
func (r *SecurityBaselineDeviceStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

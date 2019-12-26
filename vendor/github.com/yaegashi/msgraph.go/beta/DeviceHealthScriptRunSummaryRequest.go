// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceHealthScriptRunSummaryRequestBuilder is request builder for DeviceHealthScriptRunSummary
type DeviceHealthScriptRunSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceHealthScriptRunSummaryRequest
func (b *DeviceHealthScriptRunSummaryRequestBuilder) Request() *DeviceHealthScriptRunSummaryRequest {
	return &DeviceHealthScriptRunSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceHealthScriptRunSummaryRequest is request for DeviceHealthScriptRunSummary
type DeviceHealthScriptRunSummaryRequest struct{ BaseRequest }

// Get performs GET request for DeviceHealthScriptRunSummary
func (r *DeviceHealthScriptRunSummaryRequest) Get(ctx context.Context) (resObj *DeviceHealthScriptRunSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceHealthScriptRunSummary
func (r *DeviceHealthScriptRunSummaryRequest) Update(ctx context.Context, reqObj *DeviceHealthScriptRunSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceHealthScriptRunSummary
func (r *DeviceHealthScriptRunSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

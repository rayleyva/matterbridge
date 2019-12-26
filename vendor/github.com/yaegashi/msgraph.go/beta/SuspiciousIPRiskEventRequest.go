// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SuspiciousIPRiskEventRequestBuilder is request builder for SuspiciousIPRiskEvent
type SuspiciousIPRiskEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns SuspiciousIPRiskEventRequest
func (b *SuspiciousIPRiskEventRequestBuilder) Request() *SuspiciousIPRiskEventRequest {
	return &SuspiciousIPRiskEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SuspiciousIPRiskEventRequest is request for SuspiciousIPRiskEvent
type SuspiciousIPRiskEventRequest struct{ BaseRequest }

// Get performs GET request for SuspiciousIPRiskEvent
func (r *SuspiciousIPRiskEventRequest) Get(ctx context.Context) (resObj *SuspiciousIPRiskEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SuspiciousIPRiskEvent
func (r *SuspiciousIPRiskEventRequest) Update(ctx context.Context, reqObj *SuspiciousIPRiskEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SuspiciousIPRiskEvent
func (r *SuspiciousIPRiskEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

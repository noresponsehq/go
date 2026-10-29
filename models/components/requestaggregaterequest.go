package components

// RequestAggregateRequest - Count every retained matching attempt in [from,to). Capture is best effort; missing attempts cannot be inferred. UTC buckets may be partial at the edges. The window is at most 7 days.
type RequestAggregateRequest struct {
	From     string                         `json:"from"`
	To       string                         `json:"to"`
	Bucket   *RequestAggregateRequestBucket `json:"bucket,omitzero"`
	Status   *int64                         `json:"status,omitzero"`
	Route    *string                        `json:"route,omitzero"`
	APIKeyID *string                        `json:"api_key_id,omitzero"`
}

func (r *RequestAggregateRequest) GetFrom() string {
	if r == nil {
		return ""
	}
	return r.From
}

func (r *RequestAggregateRequest) GetTo() string {
	if r == nil {
		return ""
	}
	return r.To
}

func (r *RequestAggregateRequest) GetBucket() *RequestAggregateRequestBucket {
	if r == nil {
		return nil
	}
	return r.Bucket
}

func (r *RequestAggregateRequest) GetStatus() *int64 {
	if r == nil {
		return nil
	}
	return r.Status
}

func (r *RequestAggregateRequest) GetRoute() *string {
	if r == nil {
		return nil
	}
	return r.Route
}

func (r *RequestAggregateRequest) GetAPIKeyID() *string {
	if r == nil {
		return nil
	}
	return r.APIKeyID
}

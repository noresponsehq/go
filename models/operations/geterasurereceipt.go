package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type GetErasureReceiptRequest struct {
	ErasureID string `pathParam:"style=simple,explode=false,name=erasure_id"`
}

func (g *GetErasureReceiptRequest) GetErasureID() string {
	if g == nil {
		return ""
	}
	return g.ErasureID
}

type GetErasureReceiptResponse struct {
	Headers map[string][]string
	Result  components.ErasureReceipt
}

func (g *GetErasureReceiptResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetErasureReceiptResponse) GetResult() components.ErasureReceipt {
	if g == nil {
		return components.ErasureReceipt{}
	}
	return g.Result
}

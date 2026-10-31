package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// ListBulkRequest - At most 100 additions and removals combined. One transaction applies all changes or none. Idempotency-Key prevents duplicate changes for 24 hours; a different body under the same key returns 409. Matching reads see the new list immediately, while bounded durable jobs propagate account risk notifications.
type ListBulkRequest struct {
	Add    []ListEntryRequest `json:"add,omitzero"`
	Remove []string           `json:"remove,omitzero"`
}

func (l ListBulkRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListBulkRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListBulkRequest) GetAdd() []ListEntryRequest {
	if l == nil {
		return nil
	}
	return l.Add
}

func (l *ListBulkRequest) GetRemove() []string {
	if l == nil {
		return nil
	}
	return l.Remove
}

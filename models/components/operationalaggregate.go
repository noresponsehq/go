package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// OperationalAggregate - One repeatable-read snapshot of retained request diagnostics or webhook attempts. Capture can have gaps and these counts are not a billing ledger. Unattributed keys include older requests and member requests. Queue state is current retained work, not a delivery success denominator.
type OperationalAggregate struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string                        `const:"1" json:"schema_version"`
	Source        Source                        `json:"source"`
	From          string                        `json:"from"`
	To            string                        `json:"to"`
	AsOf          string                        `json:"as_of"`
	RetainedFrom  string                        `json:"retained_from"`
	RetentionDays int64                         `json:"retention_days"`
	Bucket        RequestAggregateRequestBucket `json:"bucket"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	capture string `const:"best_effort" json:"capture"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	population string                     `const:"retained_attempts" json:"population"`
	Totals     OperationalAggregateTotals `json:"totals"`
	// Nonempty UTC buckets in ascending order. Counts include all retained matching attempts, without the log page limit.
	Buckets []OperationalAggregateBucket `json:"buckets"`
	// Current retained delivery queue, independent of the requested attempt window and status filter. Endpoint filter applies. Does not prove customer enforcement execution.
	Queue *OperationalAggregateQueue `json:"queue,omitzero"`
}

func (o OperationalAggregate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OperationalAggregate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (o *OperationalAggregate) GetSchemaVersion() string {
	return "1"
}

func (o *OperationalAggregate) GetSource() Source {
	if o == nil {
		return Source("")
	}
	return o.Source
}

func (o *OperationalAggregate) GetFrom() string {
	if o == nil {
		return ""
	}
	return o.From
}

func (o *OperationalAggregate) GetTo() string {
	if o == nil {
		return ""
	}
	return o.To
}

func (o *OperationalAggregate) GetAsOf() string {
	if o == nil {
		return ""
	}
	return o.AsOf
}

func (o *OperationalAggregate) GetRetainedFrom() string {
	if o == nil {
		return ""
	}
	return o.RetainedFrom
}

func (o *OperationalAggregate) GetRetentionDays() int64 {
	if o == nil {
		return 0
	}
	return o.RetentionDays
}

func (o *OperationalAggregate) GetBucket() RequestAggregateRequestBucket {
	if o == nil {
		return RequestAggregateRequestBucket("")
	}
	return o.Bucket
}

func (o *OperationalAggregate) GetCapture() string {
	return "best_effort"
}

func (o *OperationalAggregate) GetPopulation() string {
	return "retained_attempts"
}

func (o *OperationalAggregate) GetTotals() OperationalAggregateTotals {
	if o == nil {
		return OperationalAggregateTotals{}
	}
	return o.Totals
}

func (o *OperationalAggregate) GetBuckets() []OperationalAggregateBucket {
	if o == nil {
		return []OperationalAggregateBucket{}
	}
	return o.Buckets
}

func (o *OperationalAggregate) GetQueue() *OperationalAggregateQueue {
	if o == nil {
		return nil
	}
	return o.Queue
}

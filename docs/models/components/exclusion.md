# Exclusion

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.ExclusionExpiredOrErasedRecords

// Open enum: custom values can be created with a direct type cast
custom := components.Exclusion("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ExclusionExpiredOrErasedRecords`                 | expired_or_erased_records                         |
| `ExclusionSecretsAndCredentialDigests`            | secrets_and_credential_digests                    |
| `ExclusionPrivateSourceAndWorkPayloads`           | private_source_and_work_payloads                  |
| `ExclusionCrossCustomerData`                      | cross_customer_data                               |
| `ExclusionLegacyOutcomeNotes`                     | legacy_outcome_notes                              |
| `ExclusionUncapturedHistoricalDecisions`          | uncaptured_historical_decisions                   |
| `ExclusionInvestigatorReviewsRequireReviewExport` | investigator_reviews_require_review_export        |
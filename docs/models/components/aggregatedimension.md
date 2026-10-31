# AggregateDimension

decision, recommendation and band: the first answer's. flow: the assessment's. list_match: the action of the list entry that decided the account when it was first answered. reason: a reason the first answer fired, with its impact band. lane: an evidence lane the first answer listed in incomplete_sources. country and asn: the submitted IP's, from the IP data.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AggregateDimensionDecision

// Open enum: custom values can be created with a direct type cast
custom := components.AggregateDimension("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `AggregateDimensionDecision`       | decision                           |
| `AggregateDimensionFlow`           | flow                               |
| `AggregateDimensionRecommendation` | recommendation                     |
| `AggregateDimensionBand`           | band                               |
| `AggregateDimensionReason`         | reason                             |
| `AggregateDimensionLane`           | lane                               |
| `AggregateDimensionCountry`        | country                            |
| `AggregateDimensionAsn`            | asn                                |
| `AggregateDimensionListMatch`      | list_match                         |
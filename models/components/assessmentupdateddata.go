package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/noresponsehq/go/v10/internal/utils"
)

type AssessmentUpdatedDataType string

const (
	AssessmentUpdatedDataTypeComplete AssessmentUpdatedDataType = "complete"
	AssessmentUpdatedDataTypePartial  AssessmentUpdatedDataType = "partial"
	AssessmentUpdatedDataTypePending  AssessmentUpdatedDataType = "pending"
	AssessmentUpdatedDataTypeExpired  AssessmentUpdatedDataType = "expired"
)

// AssessmentUpdatedData - The assessment, or the minimal notice that it expired.
type AssessmentUpdatedData struct {
	Assessment        *Assessment        `queryParam:"inline" union:"member"`
	AssessmentExpired *AssessmentExpired `queryParam:"inline" union:"member"`

	Type AssessmentUpdatedDataType
}

func CreateAssessmentUpdatedDataComplete(complete Assessment) AssessmentUpdatedData {
	typ := AssessmentUpdatedDataTypeComplete

	typStr := Status(typ)
	complete.Status = typStr

	return AssessmentUpdatedData{
		Assessment: &complete,
		Type:       typ,
	}
}

func CreateAssessmentUpdatedDataPartial(partial Assessment) AssessmentUpdatedData {
	typ := AssessmentUpdatedDataTypePartial

	typStr := Status(typ)
	partial.Status = typStr

	return AssessmentUpdatedData{
		Assessment: &partial,
		Type:       typ,
	}
}

func CreateAssessmentUpdatedDataPending(pending Assessment) AssessmentUpdatedData {
	typ := AssessmentUpdatedDataTypePending

	typStr := Status(typ)
	pending.Status = typStr

	return AssessmentUpdatedData{
		Assessment: &pending,
		Type:       typ,
	}
}

func CreateAssessmentUpdatedDataExpired(expired AssessmentExpired) AssessmentUpdatedData {
	typ := AssessmentUpdatedDataTypeExpired

	return AssessmentUpdatedData{
		AssessmentExpired: &expired,
		Type:              typ,
	}
}

func (u *AssessmentUpdatedData) UnmarshalJSON(data []byte) (err error) {
	previous := *u
	*u = AssessmentUpdatedData{}
	defer func() {
		if err != nil {
			*u = previous
		}
	}()

	type discriminator struct {
		Status string `json:"status"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Status {
	case "complete":
		assessment := new(Assessment)
		if err := utils.UnmarshalJSON(data, &assessment, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Status == complete) type Assessment within AssessmentUpdatedData: %w", string(data), err)
		}

		u.Assessment = assessment
		u.Type = AssessmentUpdatedDataTypeComplete
		return nil
	case "partial":
		assessment := new(Assessment)
		if err := utils.UnmarshalJSON(data, &assessment, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Status == partial) type Assessment within AssessmentUpdatedData: %w", string(data), err)
		}

		u.Assessment = assessment
		u.Type = AssessmentUpdatedDataTypePartial
		return nil
	case "pending":
		assessment := new(Assessment)
		if err := utils.UnmarshalJSON(data, &assessment, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Status == pending) type Assessment within AssessmentUpdatedData: %w", string(data), err)
		}

		u.Assessment = assessment
		u.Type = AssessmentUpdatedDataTypePending
		return nil
	case "expired":
		assessmentExpired := new(AssessmentExpired)
		if err := utils.UnmarshalJSON(data, &assessmentExpired, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Status == expired) type AssessmentExpired within AssessmentUpdatedData: %w", string(data), err)
		}

		u.AssessmentExpired = assessmentExpired
		u.Type = AssessmentUpdatedDataTypeExpired
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for AssessmentUpdatedData", string(data))
}

func (u AssessmentUpdatedData) MarshalJSON() ([]byte, error) {
	if u.Assessment != nil {
		return utils.MarshalJSON(u.Assessment, "", true)
	}

	if u.AssessmentExpired != nil {
		return utils.MarshalJSON(u.AssessmentExpired, "", true)
	}

	return nil, errors.New("could not marshal union type AssessmentUpdatedData: all fields are null")
}

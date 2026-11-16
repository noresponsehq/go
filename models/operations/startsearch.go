package operations

import (
	"github.com/noresponsehq/go/v15/models/components"
)

type StartSearchRequest struct {
	GroupID  string `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID string `pathParam:"style=simple,explode=false,name=person_id"`
	// Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person.
	IdempotencyKey *string                  `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.SearchRequest `request:"mediaType=application/json"`
}

func (s *StartSearchRequest) GetGroupID() string {
	if s == nil {
		return ""
	}
	return s.GroupID
}

func (s *StartSearchRequest) GetPersonID() string {
	if s == nil {
		return ""
	}
	return s.PersonID
}

func (s *StartSearchRequest) GetIdempotencyKey() *string {
	if s == nil {
		return nil
	}
	return s.IdempotencyKey
}

func (s *StartSearchRequest) GetBody() components.SearchRequest {
	if s == nil {
		return components.SearchRequest{}
	}
	return s.Body
}

type StartSearchResponse struct {
	Headers map[string][]string
	Result  components.Search
}

func (s *StartSearchResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *StartSearchResponse) GetResult() components.Search {
	if s == nil {
		return components.Search{}
	}
	return s.Result
}

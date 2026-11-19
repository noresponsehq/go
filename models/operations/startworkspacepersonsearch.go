package operations

import (
	"github.com/noresponsehq/go/v18/models/components"
)

type StartWorkspacePersonSearchRequest struct {
	PersonID string `pathParam:"style=simple,explode=false,name=person_id"`
	// Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person.
	IdempotencyKey *string                  `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.SearchRequest `request:"mediaType=application/json"`
}

func (s *StartWorkspacePersonSearchRequest) GetPersonID() string {
	if s == nil {
		return ""
	}
	return s.PersonID
}

func (s *StartWorkspacePersonSearchRequest) GetIdempotencyKey() *string {
	if s == nil {
		return nil
	}
	return s.IdempotencyKey
}

func (s *StartWorkspacePersonSearchRequest) GetBody() components.SearchRequest {
	if s == nil {
		return components.SearchRequest{}
	}
	return s.Body
}

type StartWorkspacePersonSearchResponse struct {
	Headers map[string][]string
	Result  components.Search
}

func (s *StartWorkspacePersonSearchResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *StartWorkspacePersonSearchResponse) GetResult() components.Search {
	if s == nil {
		return components.Search{}
	}
	return s.Result
}

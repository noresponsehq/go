package operations

import (
	"fmt"
	"github.com/noresponsehq/go/v20/internal/utils"
)

const (
	SupportedOptionGroupsCreateIdempotencyKey = "GroupsCreate:idempotencyKey"
)

// WithGroupsCreateIdempotencyKey sets the Idempotency-Key Header argument for GroupsCreate.
//
// Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person.
//
// If omitted, the argument remains unset.
// If supplied more than once, the last value wins.
func WithGroupsCreateIdempotencyKey(idempotencyKey string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionGroupsCreateIdempotencyKey) {
			return fmt.Errorf("%w: WithGroupsCreateIdempotencyKey", ErrUnsupportedOption)
		}
		opts.setMethodArgument(SupportedOptionGroupsCreateIdempotencyKey, &idempotencyKey)
		return nil
	}
}

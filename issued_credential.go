package noresponse

// Hash an issued API key. Hand-written beside the generated client.

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// HashIssuedCredential returns the SHA-256 of an API key your signup gate
// issued to the new account, as 64 lower-case hex characters. The key never
// leaves your backend: send only this hash, as issued_credential.sha256 on the
// assessment request. The key is hashed as its UTF-8 bytes exactly as given,
// with no trimming or normalisation. An empty key is an error.
func HashIssuedCredential(key string) (string, error) {
	if key == "" {
		return "", errors.New("noresponse: issued key must not be empty")
	}
	sum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(sum[:]), nil
}

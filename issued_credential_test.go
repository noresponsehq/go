package noresponse

import "testing"

func TestHashIssuedCredentialVectors(t *testing.T) {
	var file struct {
		Cases []struct {
			Key    string `json:"key"`
			Sha256 string `json:"sha256"`
		} `json:"cases"`
	}
	readAuthFixture(t, "issued-credential-vectors.json", &file)
	if len(file.Cases) == 0 {
		t.Fatal("no vectors")
	}
	for _, c := range file.Cases {
		got, err := HashIssuedCredential(c.Key)
		if err != nil || got != c.Sha256 {
			t.Errorf("%q: got %q, %v; want %q", c.Key, got, err, c.Sha256)
		}
	}
}

func TestHashIssuedCredentialRefusesEmptyKey(t *testing.T) {
	if got, err := HashIssuedCredential(""); err == nil || got != "" {
		t.Fatalf("empty key: got %q, %v", got, err)
	}
}

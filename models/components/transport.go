package components

// Transport - Where the context was derived: backend (also when the request named no transport), browser, or provider_hook for a hosted provider hook.
type Transport string

const (
	TransportBackend      Transport = "backend"
	TransportBrowser      Transport = "browser"
	TransportProviderHook Transport = "provider_hook"
)

func (e Transport) ToPointer() *Transport {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Transport) IsExact() bool {
	if e != nil {
		switch *e {
		case "backend", "browser", "provider_hook":
			return true
		}
	}
	return false
}

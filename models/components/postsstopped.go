package components

// PostsStopped - Why fewer posts were read than the screening's depth asks for while the account holds more. platform_limit: the platform shows the screening no more of them. page_limit: the pages a screening reads ran out first. rate_limited or unavailable: a later page was held back or failed; the posts before it are kept and the person is partial. Absent when the account was read to its end or to the depth.
type PostsStopped string

const (
	PostsStoppedPlatformLimit PostsStopped = "platform_limit"
	PostsStoppedPageLimit     PostsStopped = "page_limit"
	PostsStoppedRateLimited   PostsStopped = "rate_limited"
	PostsStoppedUnavailable   PostsStopped = "unavailable"
)

func (e PostsStopped) ToPointer() *PostsStopped {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PostsStopped) IsExact() bool {
	if e != nil {
		switch *e {
		case "platform_limit", "page_limit", "rate_limited", "unavailable":
			return true
		}
	}
	return false
}

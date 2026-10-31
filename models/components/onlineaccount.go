package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type SignIn string

const (
	SignInPassword  SignIn = "password"
	SignInGoogle    SignIn = "google"
	SignInFacebook  SignIn = "facebook"
	SignInApple     SignIn = "apple"
	SignInMicrosoft SignIn = "microsoft"
	SignInSso       SignIn = "sso"
	SignInOther     SignIn = "other"
)

func (e SignIn) ToPointer() *SignIn {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SignIn) IsExact() bool {
	if e != nil {
		switch *e {
		case "password", "google", "facebook", "apple", "microsoft", "sso", "other":
			return true
		}
	}
	return false
}

// OnlineAccount - One online-account search of the signup's email or username, detailed view only. name and domain name the platform searched, not a provider. outcome is the search's signal outcome: found, not_found, or a failure (blocked, unavailable, inconclusive) with its reason; a failure is not reported as not_found. The members after reason are what the platform returned beside an account it found, each absent when it did not say.
type OnlineAccount struct {
	// The platform's name, such as Adobe or GitHub.
	Name string `json:"name"`
	// The platform's domain, lower case: the host its logo is fetched for.
	Domain     string     `json:"domain"`
	Identifier Identifier `json:"identifier"`
	// Per-signal outcome. A source that timed out, was blocked or errored reports blocked, unavailable or inconclusive with a reason and is listed in incomplete_sources; not_found is a definite absence.
	Outcome SignalOutcome `json:"outcome"`
	// Why a signal or source did not produce a definite found/not_found. timed_out means the source missed the request deadline; a continuation may still settle it. deferred means the (source, identifier) pair is async-only and was not attempted on the synchronous path; a continuation settles it.
	Reason *FailureReason `json:"reason,omitzero"`
	// How the account signs in, as the platform reports it, sorted: a password, a Google, Facebook, Apple or Microsoft account, the organisation's single sign-on, or another method.
	SignIn []SignIn `json:"sign_in,omitzero"`
	// Whether the platform says the account's email is verified.
	EmailVerified *bool `json:"email_verified,omitzero"`
	// Whether the platform says the account is blocked or suspended.
	Blocked *bool `json:"blocked,omitzero"`
	// personal, or business: an account a company or school manages.
	AccountType *AccountType `json:"account_type,omitzero"`
	// Whether the account has its own profile picture rather than the platform's default. The picture is not returned.
	HasAvatar *bool `json:"has_avatar,omitzero"`
	// The first letters of the account holder's first name, as the platform shows them (Life360 shows two and hides the rest). It is evidence only: it does not change the score or the consistency results.
	FirstNamePrefix *string `json:"first_name_prefix,omitzero"`
}

func (o OnlineAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OnlineAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (o *OnlineAccount) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OnlineAccount) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *OnlineAccount) GetIdentifier() Identifier {
	if o == nil {
		return Identifier("")
	}
	return o.Identifier
}

func (o *OnlineAccount) GetOutcome() SignalOutcome {
	if o == nil {
		return SignalOutcome("")
	}
	return o.Outcome
}

func (o *OnlineAccount) GetReason() *FailureReason {
	if o == nil {
		return nil
	}
	return o.Reason
}

func (o *OnlineAccount) GetSignIn() []SignIn {
	if o == nil {
		return nil
	}
	return o.SignIn
}

func (o *OnlineAccount) GetEmailVerified() *bool {
	if o == nil {
		return nil
	}
	return o.EmailVerified
}

func (o *OnlineAccount) GetBlocked() *bool {
	if o == nil {
		return nil
	}
	return o.Blocked
}

func (o *OnlineAccount) GetAccountType() *AccountType {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *OnlineAccount) GetHasAvatar() *bool {
	if o == nil {
		return nil
	}
	return o.HasAvatar
}

func (o *OnlineAccount) GetFirstNamePrefix() *string {
	if o == nil {
		return nil
	}
	return o.FirstNamePrefix
}

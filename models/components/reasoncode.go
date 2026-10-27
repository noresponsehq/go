package components

// ReasonCode - A stable reason code. New codes may appear; ignore any you do not know.
type ReasonCode string

const (
	ReasonCodeNoDigitalFootprint         ReasonCode = "no_digital_footprint"
	ReasonCodeVoipNumber                 ReasonCode = "voip_number"
	ReasonCodeDisposableEmail            ReasonCode = "disposable_email"
	ReasonCodeNewEmail                   ReasonCode = "new_email"
	ReasonCodeSpamReports                ReasonCode = "spam_reports"
	ReasonCodeFraudFlag                  ReasonCode = "fraud_flag"
	ReasonCodeIdentityMismatch           ReasonCode = "identity_mismatch"
	ReasonCodeCountryMismatch            ReasonCode = "country_mismatch"
	ReasonCodeProxyIP                    ReasonCode = "proxy_ip"
	ReasonCodeSeenBefore                 ReasonCode = "seen_before"
	ReasonCodeSharedIP                   ReasonCode = "shared_ip"
	ReasonCodeManyRecentAccounts         ReasonCode = "many_recent_accounts"
	ReasonCodeEstablishedAccounts        ReasonCode = "established_accounts"
	ReasonCodeAgedIdentity               ReasonCode = "aged_identity"
	ReasonCodeBreachNameMatch            ReasonCode = "breach_name_match"
	ReasonCodeBreachCooccurrence         ReasonCode = "breach_cooccurrence"
	ReasonCodeBreachNameMismatch         ReasonCode = "breach_name_mismatch"
	ReasonCodeBreachPhoneMismatch        ReasonCode = "breach_phone_mismatch"
	ReasonCodeRecentPasswordExposure     ReasonCode = "recent_password_exposure"
	ReasonCodeCredentialListExposure     ReasonCode = "credential_list_exposure"
	ReasonCodeAbuseEcosystemMember       ReasonCode = "abuse_ecosystem_member"
	ReasonCodeNoBreachHistory            ReasonCode = "no_breach_history"
	ReasonCodeAutomationFastSubmit       ReasonCode = "automation_fast_submit"
	ReasonCodeClientClockSkew            ReasonCode = "client_clock_skew"
	ReasonCodeBotOrAutomation            ReasonCode = "bot_or_automation"
	ReasonCodeAuthEmailVerified          ReasonCode = "auth_email_verified"
	ReasonCodeAuthPhoneVerified          ReasonCode = "auth_phone_verified"
	ReasonCodeAuthMfaRecent              ReasonCode = "auth_mfa_recent"
	ReasonCodeAuthNewSession             ReasonCode = "auth_new_session"
	ReasonCodeAuthNewAccount             ReasonCode = "auth_new_account"
	ReasonCodeAuthImpersonated           ReasonCode = "auth_impersonated"
	ReasonCodeContextTimezoneMismatch    ReasonCode = "context_timezone_mismatch"
	ReasonCodeContextLocaleMismatch      ReasonCode = "context_locale_mismatch"
	ReasonCodeContextLineCountryMismatch ReasonCode = "context_line_country_mismatch"
	ReasonCodePriorAbuseLink             ReasonCode = "prior_abuse_link"
	ReasonCodeTestScenario               ReasonCode = "test_scenario"
	ReasonCodeSharedAPIKey               ReasonCode = "shared_api_key"
	ReasonCodeListMatch                  ReasonCode = "list_match"
	ReasonCodeYoungEmailDomain           ReasonCode = "young_email_domain"
	ReasonCodeEmailDomainNoMx            ReasonCode = "email_domain_no_mx"
	ReasonCodeEmailDomainNonexistent     ReasonCode = "email_domain_nonexistent"
	ReasonCodeForumSanctioned            ReasonCode = "forum_sanctioned"
	ReasonCodePriorPaymentAbuse          ReasonCode = "prior_payment_abuse"
	ReasonCodeIPHistoryMatch             ReasonCode = "ip_history_match"
	ReasonCodeIdentityAttributeConflict  ReasonCode = "identity_attribute_conflict"
	ReasonCodeImpossibleIdentityAge      ReasonCode = "impossible_identity_age"
	ReasonCodeBreachSharedIP             ReasonCode = "breach_shared_ip"
	ReasonCodeNationalityMismatch        ReasonCode = "nationality_mismatch"
	ReasonCodeMinorIdentity              ReasonCode = "minor_identity"
)

func (e ReasonCode) ToPointer() *ReasonCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ReasonCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "no_digital_footprint", "voip_number", "disposable_email", "new_email", "spam_reports", "fraud_flag", "identity_mismatch", "country_mismatch", "proxy_ip", "seen_before", "shared_ip", "many_recent_accounts", "established_accounts", "aged_identity", "breach_name_match", "breach_cooccurrence", "breach_name_mismatch", "breach_phone_mismatch", "recent_password_exposure", "credential_list_exposure", "abuse_ecosystem_member", "no_breach_history", "automation_fast_submit", "client_clock_skew", "bot_or_automation", "auth_email_verified", "auth_phone_verified", "auth_mfa_recent", "auth_new_session", "auth_new_account", "auth_impersonated", "context_timezone_mismatch", "context_locale_mismatch", "context_line_country_mismatch", "prior_abuse_link", "test_scenario", "shared_api_key", "list_match", "young_email_domain", "email_domain_no_mx", "email_domain_nonexistent", "forum_sanctioned", "prior_payment_abuse", "ip_history_match", "identity_attribute_conflict", "impossible_identity_age", "breach_shared_ip", "nationality_mismatch", "minor_identity":
			return true
		}
	}
	return false
}

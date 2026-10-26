# ReasonCode

A stable reason code. New codes may appear; ignore any you do not know.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.ReasonCodeNoDigitalFootprint

// Open enum: custom values can be created with a direct type cast
custom := components.ReasonCode("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `ReasonCodeNoDigitalFootprint`         | no_digital_footprint                   |
| `ReasonCodeVoipNumber`                 | voip_number                            |
| `ReasonCodeDisposableEmail`            | disposable_email                       |
| `ReasonCodeNewEmail`                   | new_email                              |
| `ReasonCodeSpamReports`                | spam_reports                           |
| `ReasonCodeFraudFlag`                  | fraud_flag                             |
| `ReasonCodeIdentityMismatch`           | identity_mismatch                      |
| `ReasonCodeCountryMismatch`            | country_mismatch                       |
| `ReasonCodeProxyIP`                    | proxy_ip                               |
| `ReasonCodeSeenBefore`                 | seen_before                            |
| `ReasonCodeSharedIP`                   | shared_ip                              |
| `ReasonCodeManyRecentAccounts`         | many_recent_accounts                   |
| `ReasonCodeEstablishedAccounts`        | established_accounts                   |
| `ReasonCodeAgedIdentity`               | aged_identity                          |
| `ReasonCodeBreachNameMatch`            | breach_name_match                      |
| `ReasonCodeBreachCooccurrence`         | breach_cooccurrence                    |
| `ReasonCodeBreachNameMismatch`         | breach_name_mismatch                   |
| `ReasonCodeBreachPhoneMismatch`        | breach_phone_mismatch                  |
| `ReasonCodeRecentPasswordExposure`     | recent_password_exposure               |
| `ReasonCodeCredentialListExposure`     | credential_list_exposure               |
| `ReasonCodeAbuseEcosystemMember`       | abuse_ecosystem_member                 |
| `ReasonCodeNoBreachHistory`            | no_breach_history                      |
| `ReasonCodeAutomationFastSubmit`       | automation_fast_submit                 |
| `ReasonCodeClientClockSkew`            | client_clock_skew                      |
| `ReasonCodeBotOrAutomation`            | bot_or_automation                      |
| `ReasonCodeAuthEmailVerified`          | auth_email_verified                    |
| `ReasonCodeAuthPhoneVerified`          | auth_phone_verified                    |
| `ReasonCodeAuthMfaRecent`              | auth_mfa_recent                        |
| `ReasonCodeAuthNewSession`             | auth_new_session                       |
| `ReasonCodeAuthNewAccount`             | auth_new_account                       |
| `ReasonCodeAuthImpersonated`           | auth_impersonated                      |
| `ReasonCodeContextTimezoneMismatch`    | context_timezone_mismatch              |
| `ReasonCodeContextLocaleMismatch`      | context_locale_mismatch                |
| `ReasonCodeContextLineCountryMismatch` | context_line_country_mismatch          |
| `ReasonCodePriorAbuseLink`             | prior_abuse_link                       |
| `ReasonCodeTestScenario`               | test_scenario                          |
| `ReasonCodeSharedAPIKey`               | shared_api_key                         |
| `ReasonCodeYoungEmailDomain`           | young_email_domain                     |
| `ReasonCodeEmailDomainNoMx`            | email_domain_no_mx                     |
| `ReasonCodeEmailDomainNonexistent`     | email_domain_nonexistent               |
| `ReasonCodeForumSanctioned`            | forum_sanctioned                       |
| `ReasonCodePriorPaymentAbuse`          | prior_payment_abuse                    |
| `ReasonCodeIPHistoryMatch`             | ip_history_match                       |
| `ReasonCodeIdentityAttributeConflict`  | identity_attribute_conflict            |
| `ReasonCodeImpossibleIdentityAge`      | impossible_identity_age                |
| `ReasonCodeBreachSharedIP`             | breach_shared_ip                       |
| `ReasonCodeNationalityMismatch`        | nationality_mismatch                   |
| `ReasonCodeMinorIdentity`              | minor_identity                         |
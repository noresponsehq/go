package components

// RecordRole - Which common field a value is, where many sources share it: a join date and a birth date are both dates, with the roles joined and born. A platform account's field and a breach row's column with the same meaning carry the same kind and role. A linked_account field is named by its network. Absent for a field only its source returns. Each role's description names the kinds its values have.
type RecordRole string

const (
	// RecordRoleAccountID The platform's own id for the account (`account_id`).
	RecordRoleAccountID RecordRole = "account_id"
	// RecordRoleAccountStatus The state of the account, such as active or suspended (`label`).
	RecordRoleAccountStatus RecordRole = "account_status"
	// RecordRoleAccountType The sort of account, such as personal or business (`label`).
	RecordRoleAccountType RecordRole = "account_type"
	// RecordRoleAlternatePhone A mobile, alternate or fax number beside the main one (`phone`).
	RecordRoleAlternatePhone RecordRole = "alternate_phone"
	// RecordRoleBalance Money held on the account (`amount`).
	RecordRoleBalance RecordRole = "balance"
	// RecordRoleBanned Whether the site banned the account (`flag`).
	RecordRoleBanned RecordRole = "banned"
	// RecordRoleBio What the person wrote about themselves (`long_text`).
	RecordRoleBio RecordRole = "bio"
	// RecordRoleBorn The date of birth (`date` or `date_time`).
	RecordRoleBorn RecordRole = "born"
	// RecordRoleBot Whether the site marks the account as a bot (`flag`).
	RecordRoleBot RecordRole = "bot"
	// RecordRoleCallerCategory The category callers give a phone number, such as telemarketer (`label`).
	RecordRoleCallerCategory RecordRole = "caller_category"
	// RecordRoleChargebacks How many chargebacks a shop recorded against the account, or whether it has one (`count` or `flag`).
	RecordRoleChargebacks RecordRole = "chargebacks"
	// RecordRoleContact Another account on the same platform the account deals with, such as an opponent it plays most or a follower (`username`). Fields beside it under the same name say how, how often and when last. It names another person, never the account's owner.
	RecordRoleContact RecordRole = "contact"
	// RecordRoleCoverPicture The banner picture of the profile (`image`).
	RecordRoleCoverPicture RecordRole = "cover_picture"
	// RecordRoleDeclines How many of the account's payments a shop declined (`count`).
	RecordRoleDeclines RecordRole = "declines"
	// RecordRoleDeleted Whether the account is deleted or inactive (`flag`).
	RecordRoleDeleted RecordRole = "deleted"
	// RecordRoleDeletedOn When the account was deleted (`date` or `date_time`).
	RecordRoleDeletedOn RecordRole = "deleted_on"
	// RecordRoleDepartment A department the source names (`text`).
	RecordRoleDepartment RecordRole = "department"
	// RecordRoleDisplayName The name the account shows (`person_name`).
	RecordRoleDisplayName RecordRole = "display_name"
	// RecordRoleEmailVerified Whether the site verified the email (`flag`).
	RecordRoleEmailVerified RecordRole = "email_verified"
	// RecordRoleExpires When the account or its subscription expires (`date` or `date_time`).
	RecordRoleExpires RecordRole = "expires"
	// RecordRoleFirstName A first name (`person_name`).
	RecordRoleFirstName RecordRole = "first_name"
	// RecordRoleFollowers How many accounts follow the account (`count`).
	RecordRoleFollowers RecordRole = "followers"
	// RecordRoleFollowing How many accounts the account follows (`count`).
	RecordRoleFollowing RecordRole = "following"
	// RecordRoleFriends How many friends the account has (`count`).
	RecordRoleFriends RecordRole = "friends"
	// RecordRoleFullAddress A whole postal address on one line (`address`).
	RecordRoleFullAddress RecordRole = "full_address"
	// RecordRoleFullName A whole name (`person_name`).
	RecordRoleFullName RecordRole = "full_name"
	// RecordRoleHandle The handle of the account (`username`).
	RecordRoleHandle RecordRole = "handle"
	// RecordRoleHeadline A headline or user title shown with the name (`text`).
	RecordRoleHeadline RecordRole = "headline"
	// RecordRoleInterests Interests the person lists (`text`).
	RecordRoleInterests RecordRole = "interests"
	// RecordRoleJobTitle A job title the source states (`job_title`).
	RecordRoleJobTitle RecordRole = "job_title"
	// RecordRoleJoined When the account was made (`date` or `date_time`).
	RecordRoleJoined RecordRole = "joined"
	// RecordRoleLastActive When the account was last active (`date` or `date_time`).
	RecordRoleLastActive RecordRole = "last_active"
	// RecordRoleLastIP The IP the account last used (`ip`).
	RecordRoleLastIP RecordRole = "last_ip"
	// RecordRoleLastName A last name (`person_name`).
	RecordRoleLastName RecordRole = "last_name"
	// RecordRoleLastPost When the account last posted (`date` or `date_time`).
	RecordRoleLastPost RecordRole = "last_post"
	// RecordRoleLastUpdated When the account or its profile last changed (`date` or `date_time`).
	RecordRoleLastUpdated RecordRole = "last_updated"
	// RecordRoleLevel A level or rank on the site (`label`).
	RecordRoleLevel RecordRole = "level"
	// RecordRoleLikesGiven How many likes the account gave (`count`).
	RecordRoleLikesGiven RecordRole = "likes_given"
	// RecordRoleLikesReceived How many likes the account received (`count`).
	RecordRoleLikesReceived RecordRole = "likes_received"
	// RecordRoleLineType The line type of a phone number, such as mobile or VoIP (`label`).
	RecordRoleLineType RecordRole = "line_type"
	// RecordRoleLinkedAccount An account on another network. The field's name names the network (`username`, `account_id` or `url`).
	RecordRoleLinkedAccount RecordRole = "linked_account"
	// RecordRoleLocation A place as the person wrote it (`text`).
	RecordRoleLocation RecordRole = "location"
	// RecordRoleMarkedSpam Whether callers mark a phone number as spam (`flag`).
	RecordRoleMarkedSpam RecordRole = "marked_spam"
	// RecordRoleMiddleName A middle name (`person_name`).
	RecordRoleMiddleName RecordRole = "middle_name"
	// RecordRoleNationality A nationality (`country`).
	RecordRoleNationality RecordRole = "nationality"
	// RecordRoleOrganisation An organisation the source names, such as a company (`organisation`).
	RecordRoleOrganisation RecordRole = "organisation"
	// RecordRoleOtherIP An IP whose use the source does not say (`ip`).
	RecordRoleOtherIP RecordRole = "other_ip"
	// RecordRolePaid Whether the account is paid or premium (`flag`).
	RecordRolePaid RecordRole = "paid"
	// RecordRolePasswordChanged When the password last changed (`date` or `date_time`).
	RecordRolePasswordChanged RecordRole = "password_changed"
	// RecordRolePaymentMethod The payment method or card brand the account paid with (`label`).
	RecordRolePaymentMethod RecordRole = "payment_method"
	// RecordRolePicture The profile picture (`image`).
	RecordRolePicture RecordRole = "picture"
	// RecordRolePoints Points on the site (`score`).
	RecordRolePoints RecordRole = "points"
	// RecordRolePosts How many posts the account made (`count`).
	RecordRolePosts RecordRole = "posts"
	// RecordRolePrivate Whether the account is private (`flag`).
	RecordRolePrivate RecordRole = "private"
	// RecordRoleProfileLink The link to the profile (`url`).
	RecordRoleProfileLink RecordRole = "profile_link"
	// RecordRoleProfileViews How many times people viewed the profile (`count`).
	RecordRoleProfileViews RecordRole = "profile_views"
	// RecordRoleRating A rating of the account (`score`).
	RecordRoleRating RecordRole = "rating"
	// RecordRoleReferrals How many people the account referred (`count`).
	RecordRoleReferrals RecordRole = "referrals"
	// RecordRoleReferrer Who or what referred the account (`text`).
	RecordRoleReferrer RecordRole = "referrer"
	// RecordRoleReportCount How many reports callers made on a phone number (`count`).
	RecordRoleReportCount RecordRole = "report_count"
	// RecordRoleReputation Reputation on the site (`score`).
	RecordRoleReputation RecordRole = "reputation"
	// RecordRoleReputationLabel A reputation label for a phone number (`label`).
	RecordRoleReputationLabel RecordRole = "reputation_label"
	// RecordRoleResidence The country the person lives in (`country`).
	RecordRoleResidence RecordRole = "residence"
	// RecordRoleSalutation A salutation or suffix, such as Dr or Jr (`text`).
	RecordRoleSalutation RecordRole = "salutation"
	// RecordRoleSchool A school the source names (`organisation`).
	RecordRoleSchool RecordRole = "school"
	// RecordRoleSignature A signature the person posts under (`long_text`).
	RecordRoleSignature RecordRole = "signature"
	// RecordRoleSignupIP The IP the account was made from (`ip`).
	RecordRoleSignupIP RecordRole = "signup_ip"
	// RecordRoleSiteRole A role on the site, such as admin or moderator (`label`).
	RecordRoleSiteRole RecordRole = "site_role"
	// RecordRoleSpamScore How likely a phone number is spam, as the source scores it (`score`).
	RecordRoleSpamScore RecordRole = "spam_score"
	// RecordRoleStreet A street line of an address (`address`).
	RecordRoleStreet RecordRole = "street"
	// RecordRoleTrades How many trades a forum account completed (`count`).
	RecordRoleTrades RecordRole = "trades"
	// RecordRoleVerified Whether the site verified the account (`flag`).
	RecordRoleVerified RecordRole = "verified"
	// RecordRoleWarnings Warnings and infractions on the account (`count`).
	RecordRoleWarnings RecordRole = "warnings"
	// RecordRoleWebsite A website the person lists (`url`).
	RecordRoleWebsite RecordRole = "website"
)

func (e RecordRole) ToPointer() *RecordRole {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RecordRole) IsExact() bool {
	if e != nil {
		switch *e {
		case "account_id", "account_status", "account_type", "alternate_phone", "balance", "banned", "bio", "born", "bot", "caller_category", "chargebacks", "contact", "cover_picture", "declines", "deleted", "deleted_on", "department", "display_name", "email_verified", "expires", "first_name", "followers", "following", "friends", "full_address", "full_name", "handle", "headline", "interests", "job_title", "joined", "last_active", "last_ip", "last_name", "last_post", "last_updated", "level", "likes_given", "likes_received", "line_type", "linked_account", "location", "marked_spam", "middle_name", "nationality", "organisation", "other_ip", "paid", "password_changed", "payment_method", "picture", "points", "posts", "private", "profile_link", "profile_views", "rating", "referrals", "referrer", "report_count", "reputation", "reputation_label", "residence", "salutation", "school", "signature", "signup_ip", "site_role", "spam_score", "street", "trades", "verified", "warnings", "website":
			return true
		}
	}
	return false
}

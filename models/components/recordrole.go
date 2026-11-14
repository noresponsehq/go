package components

// RecordRole - Which common field a value is, where many sources share it: a join date and a birth date are both dates, with the roles joined and born. A platform account's field and a breach row's column with the same meaning carry the same kind and role. A linked_account field is named by its network. Absent for a field only its source returns.
type RecordRole string

const (
	RecordRoleAccountID       RecordRole = "account_id"
	RecordRoleAccountStatus   RecordRole = "account_status"
	RecordRoleAccountType     RecordRole = "account_type"
	RecordRoleAlternatePhone  RecordRole = "alternate_phone"
	RecordRoleBalance         RecordRole = "balance"
	RecordRoleBanned          RecordRole = "banned"
	RecordRoleBio             RecordRole = "bio"
	RecordRoleBorn            RecordRole = "born"
	RecordRoleBot             RecordRole = "bot"
	RecordRoleCallerCategory  RecordRole = "caller_category"
	RecordRoleCoverPicture    RecordRole = "cover_picture"
	RecordRoleDeleted         RecordRole = "deleted"
	RecordRoleDeletedOn       RecordRole = "deleted_on"
	RecordRoleDepartment      RecordRole = "department"
	RecordRoleDisplayName     RecordRole = "display_name"
	RecordRoleEmailVerified   RecordRole = "email_verified"
	RecordRoleExpires         RecordRole = "expires"
	RecordRoleFirstName       RecordRole = "first_name"
	RecordRoleFollowers       RecordRole = "followers"
	RecordRoleFollowing       RecordRole = "following"
	RecordRoleFriends         RecordRole = "friends"
	RecordRoleFullAddress     RecordRole = "full_address"
	RecordRoleFullName        RecordRole = "full_name"
	RecordRoleHandle          RecordRole = "handle"
	RecordRoleHeadline        RecordRole = "headline"
	RecordRoleInterests       RecordRole = "interests"
	RecordRoleJobTitle        RecordRole = "job_title"
	RecordRoleJoined          RecordRole = "joined"
	RecordRoleLastActive      RecordRole = "last_active"
	RecordRoleLastIP          RecordRole = "last_ip"
	RecordRoleLastName        RecordRole = "last_name"
	RecordRoleLastPost        RecordRole = "last_post"
	RecordRoleLastUpdated     RecordRole = "last_updated"
	RecordRoleLevel           RecordRole = "level"
	RecordRoleLikesGiven      RecordRole = "likes_given"
	RecordRoleLikesReceived   RecordRole = "likes_received"
	RecordRoleLineType        RecordRole = "line_type"
	RecordRoleLinkedAccount   RecordRole = "linked_account"
	RecordRoleLocation        RecordRole = "location"
	RecordRoleMarkedSpam      RecordRole = "marked_spam"
	RecordRoleMiddleName      RecordRole = "middle_name"
	RecordRoleNationality     RecordRole = "nationality"
	RecordRoleOrganisation    RecordRole = "organisation"
	RecordRoleOtherIP         RecordRole = "other_ip"
	RecordRolePaid            RecordRole = "paid"
	RecordRolePasswordChanged RecordRole = "password_changed"
	RecordRolePicture         RecordRole = "picture"
	RecordRolePoints          RecordRole = "points"
	RecordRolePosts           RecordRole = "posts"
	RecordRolePrivate         RecordRole = "private"
	RecordRoleProfileLink     RecordRole = "profile_link"
	RecordRoleProfileViews    RecordRole = "profile_views"
	RecordRoleRating          RecordRole = "rating"
	RecordRoleReferrals       RecordRole = "referrals"
	RecordRoleReferrer        RecordRole = "referrer"
	RecordRoleReportCount     RecordRole = "report_count"
	RecordRoleReputation      RecordRole = "reputation"
	RecordRoleReputationLabel RecordRole = "reputation_label"
	RecordRoleResidence       RecordRole = "residence"
	RecordRoleSalutation      RecordRole = "salutation"
	RecordRoleSchool          RecordRole = "school"
	RecordRoleSignature       RecordRole = "signature"
	RecordRoleSignupIP        RecordRole = "signup_ip"
	RecordRoleSiteRole        RecordRole = "site_role"
	RecordRoleSpamScore       RecordRole = "spam_score"
	RecordRoleStreet          RecordRole = "street"
	RecordRoleVerified        RecordRole = "verified"
	RecordRoleWarnings        RecordRole = "warnings"
	RecordRoleWebsite         RecordRole = "website"
)

func (e RecordRole) ToPointer() *RecordRole {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RecordRole) IsExact() bool {
	if e != nil {
		switch *e {
		case "account_id", "account_status", "account_type", "alternate_phone", "balance", "banned", "bio", "born", "bot", "caller_category", "cover_picture", "deleted", "deleted_on", "department", "display_name", "email_verified", "expires", "first_name", "followers", "following", "friends", "full_address", "full_name", "handle", "headline", "interests", "job_title", "joined", "last_active", "last_ip", "last_name", "last_post", "last_updated", "level", "likes_given", "likes_received", "line_type", "linked_account", "location", "marked_spam", "middle_name", "nationality", "organisation", "other_ip", "paid", "password_changed", "picture", "points", "posts", "private", "profile_link", "profile_views", "rating", "referrals", "referrer", "report_count", "reputation", "reputation_label", "residence", "salutation", "school", "signature", "signup_ip", "site_role", "spam_score", "street", "verified", "warnings", "website":
			return true
		}
	}
	return false
}

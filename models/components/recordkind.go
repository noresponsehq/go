package components

// RecordKind - What sort of value a field holds, which decides how it is cleaned, compared and searched again. It comes from the value itself where its shape is certain (an email, an IP, a URL, an international phone number), else from the answer stored for the source's column, else text. A value of any kind is returned.
type RecordKind string

const (
	RecordKindEmail          RecordKind = "email"
	RecordKindPhone          RecordKind = "phone"
	RecordKindUsername       RecordKind = "username"
	RecordKindAccountID      RecordKind = "account_id"
	RecordKindURL            RecordKind = "url"
	RecordKindDomain         RecordKind = "domain"
	RecordKindImage          RecordKind = "image"
	RecordKindPersonName     RecordKind = "person_name"
	RecordKindGender         RecordKind = "gender"
	RecordKindAge            RecordKind = "age"
	RecordKindLanguage       RecordKind = "language"
	RecordKindOrganisation   RecordKind = "organisation"
	RecordKindJobTitle       RecordKind = "job_title"
	RecordKindCountry        RecordKind = "country"
	RecordKindRegion         RecordKind = "region"
	RecordKindCity           RecordKind = "city"
	RecordKindAddress        RecordKind = "address"
	RecordKindPostcode       RecordKind = "postcode"
	RecordKindCoordinates    RecordKind = "coordinates"
	RecordKindTimezone       RecordKind = "timezone"
	RecordKindDate           RecordKind = "date"
	RecordKindDateTime       RecordKind = "date_time"
	RecordKindDuration       RecordKind = "duration"
	RecordKindIP             RecordKind = "ip"
	RecordKindMacAddress     RecordKind = "mac_address"
	RecordKindDeviceID       RecordKind = "device_id"
	RecordKindUserAgent      RecordKind = "user_agent"
	RecordKindPassword       RecordKind = "password"
	RecordKindHash           RecordKind = "hash"
	RecordKindSalt           RecordKind = "salt"
	RecordKindToken          RecordKind = "token"
	RecordKindSecurityAnswer RecordKind = "security_answer"
	RecordKindCardNumber     RecordKind = "card_number"
	RecordKindBankAccount    RecordKind = "bank_account"
	RecordKindCryptoAddress  RecordKind = "crypto_address"
	RecordKindAmount         RecordKind = "amount"
	RecordKindIDNumber       RecordKind = "id_number"
	RecordKindVehicle        RecordKind = "vehicle"
	RecordKindLabel          RecordKind = "label"
	RecordKindLongText       RecordKind = "long_text"
	RecordKindCount          RecordKind = "count"
	RecordKindScore          RecordKind = "score"
	RecordKindFlag           RecordKind = "flag"
	RecordKindStructured     RecordKind = "structured"
	RecordKindText           RecordKind = "text"
)

func (e RecordKind) ToPointer() *RecordKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RecordKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone", "username", "account_id", "url", "domain", "image", "person_name", "gender", "age", "language", "organisation", "job_title", "country", "region", "city", "address", "postcode", "coordinates", "timezone", "date", "date_time", "duration", "ip", "mac_address", "device_id", "user_agent", "password", "hash", "salt", "token", "security_answer", "card_number", "bank_account", "crypto_address", "amount", "id_number", "vehicle", "label", "long_text", "count", "score", "flag", "structured", "text":
			return true
		}
	}
	return false
}

package components

// RecordKind - What sort of value a field holds, which decides how it is cleaned, compared and searched again. It comes from the value itself where its shape is certain (an email, an IP, a URL, an international phone number), else from the answer stored for the source's column, else text. A value of any kind is returned. Each kind's description says what it holds and the form value takes.
type RecordKind string

const (
	// RecordKindEmail An email address. value: in lower case, without a mailto: prefix.
	RecordKindEmail RecordKind = "email"
	// RecordKindPhone A phone number. value: in E.164 form, such as +447700900123. A national number is cleaned only when the country of the source's site is known.
	RecordKindPhone RecordKind = "phone"
	// RecordKindUsername A handle a person chose. value: as written. Compare it in lower case.
	RecordKindUsername RecordKind = "username"
	// RecordKindAccountID A platform's own id for an account. value: as written.
	RecordKindAccountID RecordKind = "account_id"
	// RecordKindURL A web address. value: the scheme and host in lower case, the path, query and fragment as written.
	RecordKindURL RecordKind = "url"
	// RecordKindDomain A host name with no path. value: in lower case, also when the source wrote a URL.
	RecordKindDomain RecordKind = "domain"
	// RecordKindImage A picture, as a link or a file. value: the link, cleaned as a `url`.
	RecordKindImage RecordKind = "image"
	// RecordKindPersonName A whole name or a part of one. value: with each run of spaces made one space.
	RecordKindPersonName RecordKind = "person_name"
	// RecordKindGender A stated gender. value: `female`, `male` or `other` when the value names one of them, else as written.
	RecordKindGender RecordKind = "gender"
	// RecordKindAge An age in years. value: a whole number from 0 to 150.
	RecordKindAge RecordKind = "age"
	// RecordKindLanguage A language. value: its ISO 639 code, two letters when the language has them.
	RecordKindLanguage RecordKind = "language"
	// RecordKindOrganisation A company, school or group by name. value: with each run of spaces made one space.
	RecordKindOrganisation RecordKind = "organisation"
	// RecordKindJobTitle A position or occupation. value: as written.
	RecordKindJobTitle RecordKind = "job_title"
	// RecordKindCountry A country. value: its ISO 3166-1 alpha-2 code, such as GB.
	RecordKindCountry RecordKind = "country"
	// RecordKindRegion A state, province or county. value: as written.
	RecordKindRegion RecordKind = "region"
	// RecordKindCity A city or town. value: as written.
	RecordKindCity RecordKind = "city"
	// RecordKindAddress A street address or one line of it. value: as written.
	RecordKindAddress RecordKind = "address"
	// RecordKindPostcode A postal code. value: in upper case with single spaces.
	RecordKindPostcode RecordKind = "postcode"
	// RecordKindCoordinates A latitude and longitude. value: `latitude,longitude` in decimal degrees.
	RecordKindCoordinates RecordKind = "coordinates"
	// RecordKindTimezone A timezone or an offset. value: an IANA zone name, UTC, or an offset such as +05:30.
	RecordKindTimezone RecordKind = "timezone"
	// RecordKindDate A calendar date. value: in ISO form at the precision the source gives: 2014, 2014-03 or 2014-03-07.
	RecordKindDate RecordKind = "date"
	// RecordKindDateTime A moment. value: in ISO form in UTC, such as 2014-03-07T09:30:00Z, or a date when the source gives no time.
	RecordKindDateTime RecordKind = "date_time"
	// RecordKindDuration A length of time. value: a number of seconds.
	RecordKindDuration RecordKind = "duration"
	// RecordKindIP An IPv4 or IPv6 address. value: its standard text form. An IPv4 address stored as a whole number is written as IPv4.
	RecordKindIP RecordKind = "ip"
	// RecordKindMacAddress A network card address. value: in lower case with colons.
	RecordKindMacAddress RecordKind = "mac_address"
	// RecordKindDeviceID An id of a phone, computer or browser. value: as written.
	RecordKindDeviceID RecordKind = "device_id"
	// RecordKindUserAgent A browser or app signature. value: as written.
	RecordKindUserAgent RecordKind = "user_agent"
	// RecordKindPassword A password in plain text. value: exactly as written, spaces included.
	RecordKindPassword RecordKind = "password"
	// RecordKindHash A password hash. value: as written. detail names the likely algorithm, such as bcrypt.
	RecordKindHash RecordKind = "hash"
	// RecordKindSalt The salt for a hash. value: exactly as written.
	RecordKindSalt RecordKind = "salt"
	// RecordKindToken A session, reset or API token. value: as written.
	RecordKindToken RecordKind = "token"
	// RecordKindSecurityAnswer A security question or its answer. value: exactly as written.
	RecordKindSecurityAnswer RecordKind = "security_answer"
	// RecordKindCardNumber A payment card number. value: its 12 to 19 digits only.
	RecordKindCardNumber RecordKind = "card_number"
	// RecordKindBankAccount An IBAN or account number. value: its letters and digits only, in upper case.
	RecordKindBankAccount RecordKind = "bank_account"
	// RecordKindCryptoAddress A wallet address. value: as written. detail names the likely coin, such as BTC.
	RecordKindCryptoAddress RecordKind = "crypto_address"
	// RecordKindAmount Money, with its currency when given. value: a plain decimal. detail names the currency, such as USD.
	RecordKindAmount RecordKind = "amount"
	// RecordKindIDNumber A passport, tax, national or licence number. value: as written.
	RecordKindIDNumber RecordKind = "id_number"
	// RecordKindVehicle A number plate or a VIN. value: in upper case with no spaces.
	RecordKindVehicle RecordKind = "vehicle"
	// RecordKindLabel One of a small set, such as a status or level. value: as written.
	RecordKindLabel RecordKind = "label"
	// RecordKindLongText Free writing, such as a bio or signature. value: as written.
	RecordKindLongText RecordKind = "long_text"
	// RecordKindCount A whole number of things. value: a whole number with no thousands separators.
	RecordKindCount RecordKind = "count"
	// RecordKindScore A rating, rank or points. value: a plain decimal.
	RecordKindScore RecordKind = "score"
	// RecordKindFlag A yes or no. value: `true` or `false`.
	RecordKindFlag RecordKind = "flag"
	// RecordKindStructured A nested value that could not be split. value: as written.
	RecordKindStructured RecordKind = "structured"
	// RecordKindText Anything else. value: as written.
	RecordKindText RecordKind = "text"
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

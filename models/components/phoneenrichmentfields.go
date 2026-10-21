package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/noresponsehq/go/v2/internal/utils"
)

type AlternateName struct {
}

func (a AlternateName) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AlternateName) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

type AlternateNameUnionType string

const (
	AlternateNameUnionTypeStr           AlternateNameUnionType = "str"
	AlternateNameUnionTypeNumber        AlternateNameUnionType = "number"
	AlternateNameUnionTypeBoolean       AlternateNameUnionType = "boolean"
	AlternateNameUnionTypeAlternateName AlternateNameUnionType = "alternate_name"
	AlternateNameUnionTypeArrayOfAny    AlternateNameUnionType = "arrayOfAny"
	AlternateNameUnionTypeUnknown       AlternateNameUnionType = "Unknown"
)

type AlternateNameUnion struct {
	Str           *string         `queryParam:"inline" union:"member"`
	Number        *float64        `queryParam:"inline" union:"member"`
	Boolean       *bool           `queryParam:"inline" union:"member"`
	AlternateName *AlternateName  `queryParam:"inline" union:"member"`
	ArrayOfAny    []any           `queryParam:"inline" union:"member"`
	UnknownRaw    json.RawMessage `json:"-" union:"unknown"`

	Type AlternateNameUnionType
}

func CreateAlternateNameUnionStr(str string) AlternateNameUnion {
	typ := AlternateNameUnionTypeStr

	return AlternateNameUnion{
		Str:  &str,
		Type: typ,
	}
}

func CreateAlternateNameUnionNumber(number float64) AlternateNameUnion {
	typ := AlternateNameUnionTypeNumber

	return AlternateNameUnion{
		Number: &number,
		Type:   typ,
	}
}

func CreateAlternateNameUnionBoolean(boolean bool) AlternateNameUnion {
	typ := AlternateNameUnionTypeBoolean

	return AlternateNameUnion{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateAlternateNameUnionAlternateName(alternateName AlternateName) AlternateNameUnion {
	typ := AlternateNameUnionTypeAlternateName

	return AlternateNameUnion{
		AlternateName: &alternateName,
		Type:          typ,
	}
}

func CreateAlternateNameUnionArrayOfAny(arrayOfAny []any) AlternateNameUnion {
	typ := AlternateNameUnionTypeArrayOfAny

	return AlternateNameUnion{
		ArrayOfAny: arrayOfAny,
		Type:       typ,
	}
}

func CreateAlternateNameUnionUnknown(raw json.RawMessage) AlternateNameUnion {
	return AlternateNameUnion{
		UnknownRaw: raw,
		Type:       AlternateNameUnionTypeUnknown,
	}
}

func (u AlternateNameUnion) GetUnknownRaw() json.RawMessage {
	return u.UnknownRaw
}

func (u AlternateNameUnion) IsUnknown() bool {
	return u.Type == AlternateNameUnionTypeUnknown
}

func (u *AlternateNameUnion) UnmarshalJSON(data []byte) error {
	*u = AlternateNameUnion{}

	var candidates []utils.UnionCandidate

	// Collect all valid candidates
	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  AlternateNameUnionTypeStr,
			Value: &str,
		})
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  AlternateNameUnionTypeNumber,
			Value: &number,
		})
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  AlternateNameUnionTypeBoolean,
			Value: &boolean,
		})
	}

	var alternateName AlternateName = AlternateName{}
	if err := utils.UnmarshalJSON(data, &alternateName, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  AlternateNameUnionTypeAlternateName,
			Value: &alternateName,
		})
	}

	var arrayOfAny []any = []any{}
	if err := utils.UnmarshalJSON(data, &arrayOfAny, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  AlternateNameUnionTypeArrayOfAny,
			Value: arrayOfAny,
		})
	}

	if len(candidates) == 0 {
		if !json.Valid(data) {
			return fmt.Errorf("could not unmarshal `%s` into any supported union types for AlternateNameUnion", string(data))
		}
		u.UnknownRaw = json.RawMessage(data)
		u.Type = AlternateNameUnionTypeUnknown
		return nil
	}

	// Pick the best candidate using multi-stage filtering
	best := utils.PickBestUnionCandidate(candidates, data)
	if best == nil {
		u.UnknownRaw = json.RawMessage(data)
		u.Type = AlternateNameUnionTypeUnknown
		return nil
	}

	// Set the union type and value based on the best candidate
	u.Type = best.Type.(AlternateNameUnionType)
	switch best.Type {
	case AlternateNameUnionTypeStr:
		u.Str = best.Value.(*string)
		return nil
	case AlternateNameUnionTypeNumber:
		u.Number = best.Value.(*float64)
		return nil
	case AlternateNameUnionTypeBoolean:
		u.Boolean = best.Value.(*bool)
		return nil
	case AlternateNameUnionTypeAlternateName:
		u.AlternateName = best.Value.(*AlternateName)
		return nil
	case AlternateNameUnionTypeArrayOfAny:
		u.ArrayOfAny = best.Value.([]any)
		return nil
	}

	if !json.Valid(data) {
		return fmt.Errorf("could not unmarshal `%s` into any supported union types for AlternateNameUnion", string(data))
	}
	u.UnknownRaw = json.RawMessage(data)
	u.Type = AlternateNameUnionTypeUnknown
	return nil
}

func (u AlternateNameUnion) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.AlternateName != nil {
		return utils.MarshalJSON(u.AlternateName, "", true)
	}

	if u.ArrayOfAny != nil {
		return utils.MarshalJSON(u.ArrayOfAny, "", true)
	}

	if u.UnknownRaw != nil {
		return json.RawMessage(u.UnknownRaw), nil
	}
	return nil, errors.New("could not marshal union type AlternateNameUnion: all fields are null")
}

type NameCandidate struct {
}

func (n NameCandidate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NameCandidate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, nil); err != nil {
		return err
	}
	return nil
}

type NameCandidateUnionType string

const (
	NameCandidateUnionTypeStr           NameCandidateUnionType = "str"
	NameCandidateUnionTypeNumber        NameCandidateUnionType = "number"
	NameCandidateUnionTypeBoolean       NameCandidateUnionType = "boolean"
	NameCandidateUnionTypeNameCandidate NameCandidateUnionType = "name_candidate"
	NameCandidateUnionTypeArrayOfAny    NameCandidateUnionType = "arrayOfAny"
	NameCandidateUnionTypeUnknown       NameCandidateUnionType = "Unknown"
)

type NameCandidateUnion struct {
	Str           *string         `queryParam:"inline" union:"member"`
	Number        *float64        `queryParam:"inline" union:"member"`
	Boolean       *bool           `queryParam:"inline" union:"member"`
	NameCandidate *NameCandidate  `queryParam:"inline" union:"member"`
	ArrayOfAny    []any           `queryParam:"inline" union:"member"`
	UnknownRaw    json.RawMessage `json:"-" union:"unknown"`

	Type NameCandidateUnionType
}

func CreateNameCandidateUnionStr(str string) NameCandidateUnion {
	typ := NameCandidateUnionTypeStr

	return NameCandidateUnion{
		Str:  &str,
		Type: typ,
	}
}

func CreateNameCandidateUnionNumber(number float64) NameCandidateUnion {
	typ := NameCandidateUnionTypeNumber

	return NameCandidateUnion{
		Number: &number,
		Type:   typ,
	}
}

func CreateNameCandidateUnionBoolean(boolean bool) NameCandidateUnion {
	typ := NameCandidateUnionTypeBoolean

	return NameCandidateUnion{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateNameCandidateUnionNameCandidate(nameCandidate NameCandidate) NameCandidateUnion {
	typ := NameCandidateUnionTypeNameCandidate

	return NameCandidateUnion{
		NameCandidate: &nameCandidate,
		Type:          typ,
	}
}

func CreateNameCandidateUnionArrayOfAny(arrayOfAny []any) NameCandidateUnion {
	typ := NameCandidateUnionTypeArrayOfAny

	return NameCandidateUnion{
		ArrayOfAny: arrayOfAny,
		Type:       typ,
	}
}

func CreateNameCandidateUnionUnknown(raw json.RawMessage) NameCandidateUnion {
	return NameCandidateUnion{
		UnknownRaw: raw,
		Type:       NameCandidateUnionTypeUnknown,
	}
}

func (u NameCandidateUnion) GetUnknownRaw() json.RawMessage {
	return u.UnknownRaw
}

func (u NameCandidateUnion) IsUnknown() bool {
	return u.Type == NameCandidateUnionTypeUnknown
}

func (u *NameCandidateUnion) UnmarshalJSON(data []byte) error {
	*u = NameCandidateUnion{}

	var candidates []utils.UnionCandidate

	// Collect all valid candidates
	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  NameCandidateUnionTypeStr,
			Value: &str,
		})
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  NameCandidateUnionTypeNumber,
			Value: &number,
		})
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  NameCandidateUnionTypeBoolean,
			Value: &boolean,
		})
	}

	var nameCandidate NameCandidate = NameCandidate{}
	if err := utils.UnmarshalJSON(data, &nameCandidate, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  NameCandidateUnionTypeNameCandidate,
			Value: &nameCandidate,
		})
	}

	var arrayOfAny []any = []any{}
	if err := utils.UnmarshalJSON(data, &arrayOfAny, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  NameCandidateUnionTypeArrayOfAny,
			Value: arrayOfAny,
		})
	}

	if len(candidates) == 0 {
		if !json.Valid(data) {
			return fmt.Errorf("could not unmarshal `%s` into any supported union types for NameCandidateUnion", string(data))
		}
		u.UnknownRaw = json.RawMessage(data)
		u.Type = NameCandidateUnionTypeUnknown
		return nil
	}

	// Pick the best candidate using multi-stage filtering
	best := utils.PickBestUnionCandidate(candidates, data)
	if best == nil {
		u.UnknownRaw = json.RawMessage(data)
		u.Type = NameCandidateUnionTypeUnknown
		return nil
	}

	// Set the union type and value based on the best candidate
	u.Type = best.Type.(NameCandidateUnionType)
	switch best.Type {
	case NameCandidateUnionTypeStr:
		u.Str = best.Value.(*string)
		return nil
	case NameCandidateUnionTypeNumber:
		u.Number = best.Value.(*float64)
		return nil
	case NameCandidateUnionTypeBoolean:
		u.Boolean = best.Value.(*bool)
		return nil
	case NameCandidateUnionTypeNameCandidate:
		u.NameCandidate = best.Value.(*NameCandidate)
		return nil
	case NameCandidateUnionTypeArrayOfAny:
		u.ArrayOfAny = best.Value.([]any)
		return nil
	}

	if !json.Valid(data) {
		return fmt.Errorf("could not unmarshal `%s` into any supported union types for NameCandidateUnion", string(data))
	}
	u.UnknownRaw = json.RawMessage(data)
	u.Type = NameCandidateUnionTypeUnknown
	return nil
}

func (u NameCandidateUnion) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.NameCandidate != nil {
		return utils.MarshalJSON(u.NameCandidate, "", true)
	}

	if u.ArrayOfAny != nil {
		return utils.MarshalJSON(u.ArrayOfAny, "", true)
	}

	if u.UnknownRaw != nil {
		return json.RawMessage(u.UnknownRaw), nil
	}
	return nil, errors.New("could not marshal union type NameCandidateUnion: all fields are null")
}

type Timezone struct {
}

func (t Timezone) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Timezone) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, nil); err != nil {
		return err
	}
	return nil
}

type TimezoneUnionType string

const (
	TimezoneUnionTypeStr        TimezoneUnionType = "str"
	TimezoneUnionTypeNumber     TimezoneUnionType = "number"
	TimezoneUnionTypeBoolean    TimezoneUnionType = "boolean"
	TimezoneUnionTypeTimezone   TimezoneUnionType = "timezone"
	TimezoneUnionTypeArrayOfAny TimezoneUnionType = "arrayOfAny"
	TimezoneUnionTypeUnknown    TimezoneUnionType = "Unknown"
)

type TimezoneUnion struct {
	Str        *string         `queryParam:"inline" union:"member"`
	Number     *float64        `queryParam:"inline" union:"member"`
	Boolean    *bool           `queryParam:"inline" union:"member"`
	Timezone   *Timezone       `queryParam:"inline" union:"member"`
	ArrayOfAny []any           `queryParam:"inline" union:"member"`
	UnknownRaw json.RawMessage `json:"-" union:"unknown"`

	Type TimezoneUnionType
}

func CreateTimezoneUnionStr(str string) TimezoneUnion {
	typ := TimezoneUnionTypeStr

	return TimezoneUnion{
		Str:  &str,
		Type: typ,
	}
}

func CreateTimezoneUnionNumber(number float64) TimezoneUnion {
	typ := TimezoneUnionTypeNumber

	return TimezoneUnion{
		Number: &number,
		Type:   typ,
	}
}

func CreateTimezoneUnionBoolean(boolean bool) TimezoneUnion {
	typ := TimezoneUnionTypeBoolean

	return TimezoneUnion{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateTimezoneUnionTimezone(timezone Timezone) TimezoneUnion {
	typ := TimezoneUnionTypeTimezone

	return TimezoneUnion{
		Timezone: &timezone,
		Type:     typ,
	}
}

func CreateTimezoneUnionArrayOfAny(arrayOfAny []any) TimezoneUnion {
	typ := TimezoneUnionTypeArrayOfAny

	return TimezoneUnion{
		ArrayOfAny: arrayOfAny,
		Type:       typ,
	}
}

func CreateTimezoneUnionUnknown(raw json.RawMessage) TimezoneUnion {
	return TimezoneUnion{
		UnknownRaw: raw,
		Type:       TimezoneUnionTypeUnknown,
	}
}

func (u TimezoneUnion) GetUnknownRaw() json.RawMessage {
	return u.UnknownRaw
}

func (u TimezoneUnion) IsUnknown() bool {
	return u.Type == TimezoneUnionTypeUnknown
}

func (u *TimezoneUnion) UnmarshalJSON(data []byte) error {
	*u = TimezoneUnion{}

	var candidates []utils.UnionCandidate

	// Collect all valid candidates
	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  TimezoneUnionTypeStr,
			Value: &str,
		})
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  TimezoneUnionTypeNumber,
			Value: &number,
		})
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  TimezoneUnionTypeBoolean,
			Value: &boolean,
		})
	}

	var timezone Timezone = Timezone{}
	if err := utils.UnmarshalJSON(data, &timezone, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  TimezoneUnionTypeTimezone,
			Value: &timezone,
		})
	}

	var arrayOfAny []any = []any{}
	if err := utils.UnmarshalJSON(data, &arrayOfAny, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  TimezoneUnionTypeArrayOfAny,
			Value: arrayOfAny,
		})
	}

	if len(candidates) == 0 {
		if !json.Valid(data) {
			return fmt.Errorf("could not unmarshal `%s` into any supported union types for TimezoneUnion", string(data))
		}
		u.UnknownRaw = json.RawMessage(data)
		u.Type = TimezoneUnionTypeUnknown
		return nil
	}

	// Pick the best candidate using multi-stage filtering
	best := utils.PickBestUnionCandidate(candidates, data)
	if best == nil {
		u.UnknownRaw = json.RawMessage(data)
		u.Type = TimezoneUnionTypeUnknown
		return nil
	}

	// Set the union type and value based on the best candidate
	u.Type = best.Type.(TimezoneUnionType)
	switch best.Type {
	case TimezoneUnionTypeStr:
		u.Str = best.Value.(*string)
		return nil
	case TimezoneUnionTypeNumber:
		u.Number = best.Value.(*float64)
		return nil
	case TimezoneUnionTypeBoolean:
		u.Boolean = best.Value.(*bool)
		return nil
	case TimezoneUnionTypeTimezone:
		u.Timezone = best.Value.(*Timezone)
		return nil
	case TimezoneUnionTypeArrayOfAny:
		u.ArrayOfAny = best.Value.([]any)
		return nil
	}

	if !json.Valid(data) {
		return fmt.Errorf("could not unmarshal `%s` into any supported union types for TimezoneUnion", string(data))
	}
	u.UnknownRaw = json.RawMessage(data)
	u.Type = TimezoneUnionTypeUnknown
	return nil
}

func (u TimezoneUnion) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Timezone != nil {
		return utils.MarshalJSON(u.Timezone, "", true)
	}

	if u.ArrayOfAny != nil {
		return utils.MarshalJSON(u.ArrayOfAny, "", true)
	}

	if u.UnknownRaw != nil {
		return json.RawMessage(u.UnknownRaw), nil
	}
	return nil, errors.New("could not marshal union type TimezoneUnion: all fields are null")
}

type Email struct {
}

func (e Email) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *Email) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, nil); err != nil {
		return err
	}
	return nil
}

type EmailUnionType string

const (
	EmailUnionTypeStr        EmailUnionType = "str"
	EmailUnionTypeNumber     EmailUnionType = "number"
	EmailUnionTypeBoolean    EmailUnionType = "boolean"
	EmailUnionTypeEmail      EmailUnionType = "email"
	EmailUnionTypeArrayOfAny EmailUnionType = "arrayOfAny"
	EmailUnionTypeUnknown    EmailUnionType = "Unknown"
)

type EmailUnion struct {
	Str        *string         `queryParam:"inline" union:"member"`
	Number     *float64        `queryParam:"inline" union:"member"`
	Boolean    *bool           `queryParam:"inline" union:"member"`
	Email      *Email          `queryParam:"inline" union:"member"`
	ArrayOfAny []any           `queryParam:"inline" union:"member"`
	UnknownRaw json.RawMessage `json:"-" union:"unknown"`

	Type EmailUnionType
}

func CreateEmailUnionStr(str string) EmailUnion {
	typ := EmailUnionTypeStr

	return EmailUnion{
		Str:  &str,
		Type: typ,
	}
}

func CreateEmailUnionNumber(number float64) EmailUnion {
	typ := EmailUnionTypeNumber

	return EmailUnion{
		Number: &number,
		Type:   typ,
	}
}

func CreateEmailUnionBoolean(boolean bool) EmailUnion {
	typ := EmailUnionTypeBoolean

	return EmailUnion{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateEmailUnionEmail(email Email) EmailUnion {
	typ := EmailUnionTypeEmail

	return EmailUnion{
		Email: &email,
		Type:  typ,
	}
}

func CreateEmailUnionArrayOfAny(arrayOfAny []any) EmailUnion {
	typ := EmailUnionTypeArrayOfAny

	return EmailUnion{
		ArrayOfAny: arrayOfAny,
		Type:       typ,
	}
}

func CreateEmailUnionUnknown(raw json.RawMessage) EmailUnion {
	return EmailUnion{
		UnknownRaw: raw,
		Type:       EmailUnionTypeUnknown,
	}
}

func (u EmailUnion) GetUnknownRaw() json.RawMessage {
	return u.UnknownRaw
}

func (u EmailUnion) IsUnknown() bool {
	return u.Type == EmailUnionTypeUnknown
}

func (u *EmailUnion) UnmarshalJSON(data []byte) error {
	*u = EmailUnion{}

	var candidates []utils.UnionCandidate

	// Collect all valid candidates
	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  EmailUnionTypeStr,
			Value: &str,
		})
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  EmailUnionTypeNumber,
			Value: &number,
		})
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  EmailUnionTypeBoolean,
			Value: &boolean,
		})
	}

	var email Email = Email{}
	if err := utils.UnmarshalJSON(data, &email, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  EmailUnionTypeEmail,
			Value: &email,
		})
	}

	var arrayOfAny []any = []any{}
	if err := utils.UnmarshalJSON(data, &arrayOfAny, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  EmailUnionTypeArrayOfAny,
			Value: arrayOfAny,
		})
	}

	if len(candidates) == 0 {
		if !json.Valid(data) {
			return fmt.Errorf("could not unmarshal `%s` into any supported union types for EmailUnion", string(data))
		}
		u.UnknownRaw = json.RawMessage(data)
		u.Type = EmailUnionTypeUnknown
		return nil
	}

	// Pick the best candidate using multi-stage filtering
	best := utils.PickBestUnionCandidate(candidates, data)
	if best == nil {
		u.UnknownRaw = json.RawMessage(data)
		u.Type = EmailUnionTypeUnknown
		return nil
	}

	// Set the union type and value based on the best candidate
	u.Type = best.Type.(EmailUnionType)
	switch best.Type {
	case EmailUnionTypeStr:
		u.Str = best.Value.(*string)
		return nil
	case EmailUnionTypeNumber:
		u.Number = best.Value.(*float64)
		return nil
	case EmailUnionTypeBoolean:
		u.Boolean = best.Value.(*bool)
		return nil
	case EmailUnionTypeEmail:
		u.Email = best.Value.(*Email)
		return nil
	case EmailUnionTypeArrayOfAny:
		u.ArrayOfAny = best.Value.([]any)
		return nil
	}

	if !json.Valid(data) {
		return fmt.Errorf("could not unmarshal `%s` into any supported union types for EmailUnion", string(data))
	}
	u.UnknownRaw = json.RawMessage(data)
	u.Type = EmailUnionTypeUnknown
	return nil
}

func (u EmailUnion) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Email != nil {
		return utils.MarshalJSON(u.Email, "", true)
	}

	if u.ArrayOfAny != nil {
		return utils.MarshalJSON(u.ArrayOfAny, "", true)
	}

	if u.UnknownRaw != nil {
		return json.RawMessage(u.UnknownRaw), nil
	}
	return nil, errors.New("could not marshal union type EmailUnion: all fields are null")
}

type Website struct {
}

func (w Website) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *Website) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

type WebsiteUnionType string

const (
	WebsiteUnionTypeStr        WebsiteUnionType = "str"
	WebsiteUnionTypeNumber     WebsiteUnionType = "number"
	WebsiteUnionTypeBoolean    WebsiteUnionType = "boolean"
	WebsiteUnionTypeWebsite    WebsiteUnionType = "website"
	WebsiteUnionTypeArrayOfAny WebsiteUnionType = "arrayOfAny"
	WebsiteUnionTypeUnknown    WebsiteUnionType = "Unknown"
)

type WebsiteUnion struct {
	Str        *string         `queryParam:"inline" union:"member"`
	Number     *float64        `queryParam:"inline" union:"member"`
	Boolean    *bool           `queryParam:"inline" union:"member"`
	Website    *Website        `queryParam:"inline" union:"member"`
	ArrayOfAny []any           `queryParam:"inline" union:"member"`
	UnknownRaw json.RawMessage `json:"-" union:"unknown"`

	Type WebsiteUnionType
}

func CreateWebsiteUnionStr(str string) WebsiteUnion {
	typ := WebsiteUnionTypeStr

	return WebsiteUnion{
		Str:  &str,
		Type: typ,
	}
}

func CreateWebsiteUnionNumber(number float64) WebsiteUnion {
	typ := WebsiteUnionTypeNumber

	return WebsiteUnion{
		Number: &number,
		Type:   typ,
	}
}

func CreateWebsiteUnionBoolean(boolean bool) WebsiteUnion {
	typ := WebsiteUnionTypeBoolean

	return WebsiteUnion{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateWebsiteUnionWebsite(website Website) WebsiteUnion {
	typ := WebsiteUnionTypeWebsite

	return WebsiteUnion{
		Website: &website,
		Type:    typ,
	}
}

func CreateWebsiteUnionArrayOfAny(arrayOfAny []any) WebsiteUnion {
	typ := WebsiteUnionTypeArrayOfAny

	return WebsiteUnion{
		ArrayOfAny: arrayOfAny,
		Type:       typ,
	}
}

func CreateWebsiteUnionUnknown(raw json.RawMessage) WebsiteUnion {
	return WebsiteUnion{
		UnknownRaw: raw,
		Type:       WebsiteUnionTypeUnknown,
	}
}

func (u WebsiteUnion) GetUnknownRaw() json.RawMessage {
	return u.UnknownRaw
}

func (u WebsiteUnion) IsUnknown() bool {
	return u.Type == WebsiteUnionTypeUnknown
}

func (u *WebsiteUnion) UnmarshalJSON(data []byte) error {
	*u = WebsiteUnion{}

	var candidates []utils.UnionCandidate

	// Collect all valid candidates
	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  WebsiteUnionTypeStr,
			Value: &str,
		})
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  WebsiteUnionTypeNumber,
			Value: &number,
		})
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  WebsiteUnionTypeBoolean,
			Value: &boolean,
		})
	}

	var website Website = Website{}
	if err := utils.UnmarshalJSON(data, &website, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  WebsiteUnionTypeWebsite,
			Value: &website,
		})
	}

	var arrayOfAny []any = []any{}
	if err := utils.UnmarshalJSON(data, &arrayOfAny, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  WebsiteUnionTypeArrayOfAny,
			Value: arrayOfAny,
		})
	}

	if len(candidates) == 0 {
		if !json.Valid(data) {
			return fmt.Errorf("could not unmarshal `%s` into any supported union types for WebsiteUnion", string(data))
		}
		u.UnknownRaw = json.RawMessage(data)
		u.Type = WebsiteUnionTypeUnknown
		return nil
	}

	// Pick the best candidate using multi-stage filtering
	best := utils.PickBestUnionCandidate(candidates, data)
	if best == nil {
		u.UnknownRaw = json.RawMessage(data)
		u.Type = WebsiteUnionTypeUnknown
		return nil
	}

	// Set the union type and value based on the best candidate
	u.Type = best.Type.(WebsiteUnionType)
	switch best.Type {
	case WebsiteUnionTypeStr:
		u.Str = best.Value.(*string)
		return nil
	case WebsiteUnionTypeNumber:
		u.Number = best.Value.(*float64)
		return nil
	case WebsiteUnionTypeBoolean:
		u.Boolean = best.Value.(*bool)
		return nil
	case WebsiteUnionTypeWebsite:
		u.Website = best.Value.(*Website)
		return nil
	case WebsiteUnionTypeArrayOfAny:
		u.ArrayOfAny = best.Value.([]any)
		return nil
	}

	if !json.Valid(data) {
		return fmt.Errorf("could not unmarshal `%s` into any supported union types for WebsiteUnion", string(data))
	}
	u.UnknownRaw = json.RawMessage(data)
	u.Type = WebsiteUnionTypeUnknown
	return nil
}

func (u WebsiteUnion) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Website != nil {
		return utils.MarshalJSON(u.Website, "", true)
	}

	if u.ArrayOfAny != nil {
		return utils.MarshalJSON(u.ArrayOfAny, "", true)
	}

	if u.UnknownRaw != nil {
		return json.RawMessage(u.UnknownRaw), nil
	}
	return nil, errors.New("could not marshal union type WebsiteUnion: all fields are null")
}

type Location struct {
}

func (l Location) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Location) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

type LocationUnionType string

const (
	LocationUnionTypeStr        LocationUnionType = "str"
	LocationUnionTypeNumber     LocationUnionType = "number"
	LocationUnionTypeBoolean    LocationUnionType = "boolean"
	LocationUnionTypeLocation   LocationUnionType = "location"
	LocationUnionTypeArrayOfAny LocationUnionType = "arrayOfAny"
	LocationUnionTypeUnknown    LocationUnionType = "Unknown"
)

type LocationUnion struct {
	Str        *string         `queryParam:"inline" union:"member"`
	Number     *float64        `queryParam:"inline" union:"member"`
	Boolean    *bool           `queryParam:"inline" union:"member"`
	Location   *Location       `queryParam:"inline" union:"member"`
	ArrayOfAny []any           `queryParam:"inline" union:"member"`
	UnknownRaw json.RawMessage `json:"-" union:"unknown"`

	Type LocationUnionType
}

func CreateLocationUnionStr(str string) LocationUnion {
	typ := LocationUnionTypeStr

	return LocationUnion{
		Str:  &str,
		Type: typ,
	}
}

func CreateLocationUnionNumber(number float64) LocationUnion {
	typ := LocationUnionTypeNumber

	return LocationUnion{
		Number: &number,
		Type:   typ,
	}
}

func CreateLocationUnionBoolean(boolean bool) LocationUnion {
	typ := LocationUnionTypeBoolean

	return LocationUnion{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateLocationUnionLocation(location Location) LocationUnion {
	typ := LocationUnionTypeLocation

	return LocationUnion{
		Location: &location,
		Type:     typ,
	}
}

func CreateLocationUnionArrayOfAny(arrayOfAny []any) LocationUnion {
	typ := LocationUnionTypeArrayOfAny

	return LocationUnion{
		ArrayOfAny: arrayOfAny,
		Type:       typ,
	}
}

func CreateLocationUnionUnknown(raw json.RawMessage) LocationUnion {
	return LocationUnion{
		UnknownRaw: raw,
		Type:       LocationUnionTypeUnknown,
	}
}

func (u LocationUnion) GetUnknownRaw() json.RawMessage {
	return u.UnknownRaw
}

func (u LocationUnion) IsUnknown() bool {
	return u.Type == LocationUnionTypeUnknown
}

func (u *LocationUnion) UnmarshalJSON(data []byte) error {
	*u = LocationUnion{}

	var candidates []utils.UnionCandidate

	// Collect all valid candidates
	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  LocationUnionTypeStr,
			Value: &str,
		})
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  LocationUnionTypeNumber,
			Value: &number,
		})
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  LocationUnionTypeBoolean,
			Value: &boolean,
		})
	}

	var location Location = Location{}
	if err := utils.UnmarshalJSON(data, &location, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  LocationUnionTypeLocation,
			Value: &location,
		})
	}

	var arrayOfAny []any = []any{}
	if err := utils.UnmarshalJSON(data, &arrayOfAny, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  LocationUnionTypeArrayOfAny,
			Value: arrayOfAny,
		})
	}

	if len(candidates) == 0 {
		if !json.Valid(data) {
			return fmt.Errorf("could not unmarshal `%s` into any supported union types for LocationUnion", string(data))
		}
		u.UnknownRaw = json.RawMessage(data)
		u.Type = LocationUnionTypeUnknown
		return nil
	}

	// Pick the best candidate using multi-stage filtering
	best := utils.PickBestUnionCandidate(candidates, data)
	if best == nil {
		u.UnknownRaw = json.RawMessage(data)
		u.Type = LocationUnionTypeUnknown
		return nil
	}

	// Set the union type and value based on the best candidate
	u.Type = best.Type.(LocationUnionType)
	switch best.Type {
	case LocationUnionTypeStr:
		u.Str = best.Value.(*string)
		return nil
	case LocationUnionTypeNumber:
		u.Number = best.Value.(*float64)
		return nil
	case LocationUnionTypeBoolean:
		u.Boolean = best.Value.(*bool)
		return nil
	case LocationUnionTypeLocation:
		u.Location = best.Value.(*Location)
		return nil
	case LocationUnionTypeArrayOfAny:
		u.ArrayOfAny = best.Value.([]any)
		return nil
	}

	if !json.Valid(data) {
		return fmt.Errorf("could not unmarshal `%s` into any supported union types for LocationUnion", string(data))
	}
	u.UnknownRaw = json.RawMessage(data)
	u.Type = LocationUnionTypeUnknown
	return nil
}

func (u LocationUnion) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Location != nil {
		return utils.MarshalJSON(u.Location, "", true)
	}

	if u.ArrayOfAny != nil {
		return utils.MarshalJSON(u.ArrayOfAny, "", true)
	}

	if u.UnknownRaw != nil {
		return json.RawMessage(u.UnknownRaw), nil
	}
	return nil, errors.New("could not marshal union type LocationUnion: all fields are null")
}

type SocialProfile struct {
}

func (s SocialProfile) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SocialProfile) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

type SocialProfileUnionType string

const (
	SocialProfileUnionTypeStr           SocialProfileUnionType = "str"
	SocialProfileUnionTypeNumber        SocialProfileUnionType = "number"
	SocialProfileUnionTypeBoolean       SocialProfileUnionType = "boolean"
	SocialProfileUnionTypeSocialProfile SocialProfileUnionType = "social_profile"
	SocialProfileUnionTypeArrayOfAny    SocialProfileUnionType = "arrayOfAny"
	SocialProfileUnionTypeUnknown       SocialProfileUnionType = "Unknown"
)

type SocialProfileUnion struct {
	Str           *string         `queryParam:"inline" union:"member"`
	Number        *float64        `queryParam:"inline" union:"member"`
	Boolean       *bool           `queryParam:"inline" union:"member"`
	SocialProfile *SocialProfile  `queryParam:"inline" union:"member"`
	ArrayOfAny    []any           `queryParam:"inline" union:"member"`
	UnknownRaw    json.RawMessage `json:"-" union:"unknown"`

	Type SocialProfileUnionType
}

func CreateSocialProfileUnionStr(str string) SocialProfileUnion {
	typ := SocialProfileUnionTypeStr

	return SocialProfileUnion{
		Str:  &str,
		Type: typ,
	}
}

func CreateSocialProfileUnionNumber(number float64) SocialProfileUnion {
	typ := SocialProfileUnionTypeNumber

	return SocialProfileUnion{
		Number: &number,
		Type:   typ,
	}
}

func CreateSocialProfileUnionBoolean(boolean bool) SocialProfileUnion {
	typ := SocialProfileUnionTypeBoolean

	return SocialProfileUnion{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateSocialProfileUnionSocialProfile(socialProfile SocialProfile) SocialProfileUnion {
	typ := SocialProfileUnionTypeSocialProfile

	return SocialProfileUnion{
		SocialProfile: &socialProfile,
		Type:          typ,
	}
}

func CreateSocialProfileUnionArrayOfAny(arrayOfAny []any) SocialProfileUnion {
	typ := SocialProfileUnionTypeArrayOfAny

	return SocialProfileUnion{
		ArrayOfAny: arrayOfAny,
		Type:       typ,
	}
}

func CreateSocialProfileUnionUnknown(raw json.RawMessage) SocialProfileUnion {
	return SocialProfileUnion{
		UnknownRaw: raw,
		Type:       SocialProfileUnionTypeUnknown,
	}
}

func (u SocialProfileUnion) GetUnknownRaw() json.RawMessage {
	return u.UnknownRaw
}

func (u SocialProfileUnion) IsUnknown() bool {
	return u.Type == SocialProfileUnionTypeUnknown
}

func (u *SocialProfileUnion) UnmarshalJSON(data []byte) error {
	*u = SocialProfileUnion{}

	var candidates []utils.UnionCandidate

	// Collect all valid candidates
	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  SocialProfileUnionTypeStr,
			Value: &str,
		})
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  SocialProfileUnionTypeNumber,
			Value: &number,
		})
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  SocialProfileUnionTypeBoolean,
			Value: &boolean,
		})
	}

	var socialProfile SocialProfile = SocialProfile{}
	if err := utils.UnmarshalJSON(data, &socialProfile, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  SocialProfileUnionTypeSocialProfile,
			Value: &socialProfile,
		})
	}

	var arrayOfAny []any = []any{}
	if err := utils.UnmarshalJSON(data, &arrayOfAny, "", true, nil); err == nil {
		candidates = append(candidates, utils.UnionCandidate{
			Type:  SocialProfileUnionTypeArrayOfAny,
			Value: arrayOfAny,
		})
	}

	if len(candidates) == 0 {
		if !json.Valid(data) {
			return fmt.Errorf("could not unmarshal `%s` into any supported union types for SocialProfileUnion", string(data))
		}
		u.UnknownRaw = json.RawMessage(data)
		u.Type = SocialProfileUnionTypeUnknown
		return nil
	}

	// Pick the best candidate using multi-stage filtering
	best := utils.PickBestUnionCandidate(candidates, data)
	if best == nil {
		u.UnknownRaw = json.RawMessage(data)
		u.Type = SocialProfileUnionTypeUnknown
		return nil
	}

	// Set the union type and value based on the best candidate
	u.Type = best.Type.(SocialProfileUnionType)
	switch best.Type {
	case SocialProfileUnionTypeStr:
		u.Str = best.Value.(*string)
		return nil
	case SocialProfileUnionTypeNumber:
		u.Number = best.Value.(*float64)
		return nil
	case SocialProfileUnionTypeBoolean:
		u.Boolean = best.Value.(*bool)
		return nil
	case SocialProfileUnionTypeSocialProfile:
		u.SocialProfile = best.Value.(*SocialProfile)
		return nil
	case SocialProfileUnionTypeArrayOfAny:
		u.ArrayOfAny = best.Value.([]any)
		return nil
	}

	if !json.Valid(data) {
		return fmt.Errorf("could not unmarshal `%s` into any supported union types for SocialProfileUnion", string(data))
	}
	u.UnknownRaw = json.RawMessage(data)
	u.Type = SocialProfileUnionTypeUnknown
	return nil
}

func (u SocialProfileUnion) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.SocialProfile != nil {
		return utils.MarshalJSON(u.SocialProfile, "", true)
	}

	if u.ArrayOfAny != nil {
		return utils.MarshalJSON(u.ArrayOfAny, "", true)
	}

	if u.UnknownRaw != nil {
		return json.RawMessage(u.UnknownRaw), nil
	}
	return nil, errors.New("could not marshal union type SocialProfileUnion: all fields are null")
}

type PhoneEnrichmentFields struct {
	Phone           *string              `json:"phone,omitzero"`
	Name            *string              `json:"name,omitzero"`
	Carrier         *string              `json:"carrier,omitzero"`
	LineType        *string              `json:"line_type,omitzero"`
	EntityType      *string              `json:"entity_type,omitzero"`
	CountryCode     *string              `json:"country_code,omitzero"`
	Country         *string              `json:"country,omitzero"`
	State           *string              `json:"state,omitzero"`
	City            *string              `json:"city,omitzero"`
	Location        *string              `json:"location,omitzero"`
	NationalFormat  *string              `json:"national_format,omitzero"`
	JobTitle        *string              `json:"job_title,omitzero"`
	CompanyName     *string              `json:"company_name,omitzero"`
	Email           *string              `json:"email,omitzero"`
	Website         *string              `json:"website,omitzero"`
	FirstCallAt     *string              `json:"first_call_at,omitzero"`
	LastCallAt      *string              `json:"last_call_at,omitzero"`
	MatchedName     *string              `json:"matched_name,omitzero"`
	AlternateNames  []AlternateNameUnion `json:"alternate_names,omitzero"`
	NameCandidates  []NameCandidateUnion `json:"name_candidates,omitzero"`
	Timezones       []TimezoneUnion      `json:"timezones,omitzero"`
	Emails          []EmailUnion         `json:"emails,omitzero"`
	Websites        []WebsiteUnion       `json:"websites,omitzero"`
	Locations       []LocationUnion      `json:"locations,omitzero"`
	SocialProfiles  []SocialProfileUnion `json:"social_profiles,omitzero"`
	ReportCount     *int64               `json:"report_count,omitzero"`
	TagCount        *int64               `json:"tag_count,omitzero"`
	TotalCalls      *int64               `json:"total_calls,omitzero"`
	Calls60Days     *int64               `json:"calls_60_days,omitzero"`
	Reports60Days   *int64               `json:"reports_60_days,omitzero"`
	Searches60Days  *int64               `json:"searches_60_days,omitzero"`
	CommentCount    *int64               `json:"comment_count,omitzero"`
	NamesCount      *int64               `json:"names_count,omitzero"`
	SpamReportCount *int64               `json:"spam_report_count,omitzero"`
}

func (p PhoneEnrichmentFields) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PhoneEnrichmentFields) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PhoneEnrichmentFields) GetPhone() *string {
	if p == nil {
		return nil
	}
	return p.Phone
}

func (p *PhoneEnrichmentFields) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PhoneEnrichmentFields) GetCarrier() *string {
	if p == nil {
		return nil
	}
	return p.Carrier
}

func (p *PhoneEnrichmentFields) GetLineType() *string {
	if p == nil {
		return nil
	}
	return p.LineType
}

func (p *PhoneEnrichmentFields) GetEntityType() *string {
	if p == nil {
		return nil
	}
	return p.EntityType
}

func (p *PhoneEnrichmentFields) GetCountryCode() *string {
	if p == nil {
		return nil
	}
	return p.CountryCode
}

func (p *PhoneEnrichmentFields) GetCountry() *string {
	if p == nil {
		return nil
	}
	return p.Country
}

func (p *PhoneEnrichmentFields) GetState() *string {
	if p == nil {
		return nil
	}
	return p.State
}

func (p *PhoneEnrichmentFields) GetCity() *string {
	if p == nil {
		return nil
	}
	return p.City
}

func (p *PhoneEnrichmentFields) GetLocation() *string {
	if p == nil {
		return nil
	}
	return p.Location
}

func (p *PhoneEnrichmentFields) GetNationalFormat() *string {
	if p == nil {
		return nil
	}
	return p.NationalFormat
}

func (p *PhoneEnrichmentFields) GetJobTitle() *string {
	if p == nil {
		return nil
	}
	return p.JobTitle
}

func (p *PhoneEnrichmentFields) GetCompanyName() *string {
	if p == nil {
		return nil
	}
	return p.CompanyName
}

func (p *PhoneEnrichmentFields) GetEmail() *string {
	if p == nil {
		return nil
	}
	return p.Email
}

func (p *PhoneEnrichmentFields) GetWebsite() *string {
	if p == nil {
		return nil
	}
	return p.Website
}

func (p *PhoneEnrichmentFields) GetFirstCallAt() *string {
	if p == nil {
		return nil
	}
	return p.FirstCallAt
}

func (p *PhoneEnrichmentFields) GetLastCallAt() *string {
	if p == nil {
		return nil
	}
	return p.LastCallAt
}

func (p *PhoneEnrichmentFields) GetMatchedName() *string {
	if p == nil {
		return nil
	}
	return p.MatchedName
}

func (p *PhoneEnrichmentFields) GetAlternateNames() []AlternateNameUnion {
	if p == nil {
		return nil
	}
	return p.AlternateNames
}

func (p *PhoneEnrichmentFields) GetNameCandidates() []NameCandidateUnion {
	if p == nil {
		return nil
	}
	return p.NameCandidates
}

func (p *PhoneEnrichmentFields) GetTimezones() []TimezoneUnion {
	if p == nil {
		return nil
	}
	return p.Timezones
}

func (p *PhoneEnrichmentFields) GetEmails() []EmailUnion {
	if p == nil {
		return nil
	}
	return p.Emails
}

func (p *PhoneEnrichmentFields) GetWebsites() []WebsiteUnion {
	if p == nil {
		return nil
	}
	return p.Websites
}

func (p *PhoneEnrichmentFields) GetLocations() []LocationUnion {
	if p == nil {
		return nil
	}
	return p.Locations
}

func (p *PhoneEnrichmentFields) GetSocialProfiles() []SocialProfileUnion {
	if p == nil {
		return nil
	}
	return p.SocialProfiles
}

func (p *PhoneEnrichmentFields) GetReportCount() *int64 {
	if p == nil {
		return nil
	}
	return p.ReportCount
}

func (p *PhoneEnrichmentFields) GetTagCount() *int64 {
	if p == nil {
		return nil
	}
	return p.TagCount
}

func (p *PhoneEnrichmentFields) GetTotalCalls() *int64 {
	if p == nil {
		return nil
	}
	return p.TotalCalls
}

func (p *PhoneEnrichmentFields) GetCalls60Days() *int64 {
	if p == nil {
		return nil
	}
	return p.Calls60Days
}

func (p *PhoneEnrichmentFields) GetReports60Days() *int64 {
	if p == nil {
		return nil
	}
	return p.Reports60Days
}

func (p *PhoneEnrichmentFields) GetSearches60Days() *int64 {
	if p == nil {
		return nil
	}
	return p.Searches60Days
}

func (p *PhoneEnrichmentFields) GetCommentCount() *int64 {
	if p == nil {
		return nil
	}
	return p.CommentCount
}

func (p *PhoneEnrichmentFields) GetNamesCount() *int64 {
	if p == nil {
		return nil
	}
	return p.NamesCount
}

func (p *PhoneEnrichmentFields) GetSpamReportCount() *int64 {
	if p == nil {
		return nil
	}
	return p.SpamReportCount
}

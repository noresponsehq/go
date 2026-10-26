package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
)

type FootprintCoherence struct {
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	BirthDate *AttributeAgreement `json:"birth_date,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	Gender *AttributeAgreement `json:"gender,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	Nationality *AttributeAgreement `json:"nationality,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	Country *AttributeAgreement `json:"country,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	City *AttributeAgreement `json:"city,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	Zip *AttributeAgreement `json:"zip,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	Name *AttributeAgreement `json:"name,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	Company *AttributeAgreement `json:"company,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	Website *AttributeAgreement `json:"website,omitzero"`
	// How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
	Social *AttributeAgreement `json:"social,omitzero"`
}

func (f FootprintCoherence) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FootprintCoherence) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (f *FootprintCoherence) GetBirthDate() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.BirthDate
}

func (f *FootprintCoherence) GetGender() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.Gender
}

func (f *FootprintCoherence) GetNationality() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.Nationality
}

func (f *FootprintCoherence) GetCountry() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.Country
}

func (f *FootprintCoherence) GetCity() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.City
}

func (f *FootprintCoherence) GetZip() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.Zip
}

func (f *FootprintCoherence) GetName() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.Name
}

func (f *FootprintCoherence) GetCompany() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.Company
}

func (f *FootprintCoherence) GetWebsite() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.Website
}

func (f *FootprintCoherence) GetSocial() *AttributeAgreement {
	if f == nil {
		return nil
	}
	return f.Social
}

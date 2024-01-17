// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Country string

const (
	CountryDe Country = "DE"
	CountryAt Country = "AT"
	CountryCh Country = "CH"
)

func (e Country) ToPointer() *Country {
	return &e
}

func (e *Country) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DE":
		fallthrough
	case "AT":
		fallthrough
	case "CH":
		*e = Country(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Country: %v", v)
	}
}

type BaseAddress struct {
	ID             *string  `json:"_id,omitempty"`
	Tags           []string `json:"_tags,omitempty"`
	AdditionalInfo *string  `json:"additional_info,omitempty"`
	City           *string  `json:"city"`
	Country        *Country `json:"country"`
	PostalCode     *string  `json:"postal_code"`
	Street         *string  `json:"street"`
	StreetNumber   *string  `json:"street_number"`
}

func (o *BaseAddress) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BaseAddress) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *BaseAddress) GetAdditionalInfo() *string {
	if o == nil {
		return nil
	}
	return o.AdditionalInfo
}

func (o *BaseAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *BaseAddress) GetCountry() *Country {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *BaseAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *BaseAddress) GetStreet() *string {
	if o == nil {
		return nil
	}
	return o.Street
}

func (o *BaseAddress) GetStreetNumber() *string {
	if o == nil {
		return nil
	}
	return o.StreetNumber
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CustomerOrganizationEntityTypeEnum - Organization entity type. Only applicable for B2B customers.
type CustomerOrganizationEntityTypeEnum string

const (
	CustomerOrganizationEntityTypeEnumLimitedCompany                   CustomerOrganizationEntityTypeEnum = "LIMITED_COMPANY"
	CustomerOrganizationEntityTypeEnumPublicLimitedCompany             CustomerOrganizationEntityTypeEnum = "PUBLIC_LIMITED_COMPANY"
	CustomerOrganizationEntityTypeEnumEntrepreneurialCompany           CustomerOrganizationEntityTypeEnum = "ENTREPRENEURIAL_COMPANY"
	CustomerOrganizationEntityTypeEnumLimitedPartnershipLimitedCompany CustomerOrganizationEntityTypeEnum = "LIMITED_PARTNERSHIP_LIMITED_COMPANY"
	CustomerOrganizationEntityTypeEnumLimitedPartnership               CustomerOrganizationEntityTypeEnum = "LIMITED_PARTNERSHIP"
	CustomerOrganizationEntityTypeEnumGeneralPartnership               CustomerOrganizationEntityTypeEnum = "GENERAL_PARTNERSHIP"
	CustomerOrganizationEntityTypeEnumRegisteredSoleTrader             CustomerOrganizationEntityTypeEnum = "REGISTERED_SOLE_TRADER"
	CustomerOrganizationEntityTypeEnumSoleTrader                       CustomerOrganizationEntityTypeEnum = "SOLE_TRADER"
	CustomerOrganizationEntityTypeEnumCivilLawPartnership              CustomerOrganizationEntityTypeEnum = "CIVIL_LAW_PARTNERSHIP"
	CustomerOrganizationEntityTypeEnumPublicInstitution                CustomerOrganizationEntityTypeEnum = "PUBLIC_INSTITUTION"
	CustomerOrganizationEntityTypeEnumOther                            CustomerOrganizationEntityTypeEnum = "OTHER"
)

func (e CustomerOrganizationEntityTypeEnum) ToPointer() *CustomerOrganizationEntityTypeEnum {
	return &e
}

func (e *CustomerOrganizationEntityTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LIMITED_COMPANY":
		fallthrough
	case "PUBLIC_LIMITED_COMPANY":
		fallthrough
	case "ENTREPRENEURIAL_COMPANY":
		fallthrough
	case "LIMITED_PARTNERSHIP_LIMITED_COMPANY":
		fallthrough
	case "LIMITED_PARTNERSHIP":
		fallthrough
	case "GENERAL_PARTNERSHIP":
		fallthrough
	case "REGISTERED_SOLE_TRADER":
		fallthrough
	case "SOLE_TRADER":
		fallthrough
	case "CIVIL_LAW_PARTNERSHIP":
		fallthrough
	case "PUBLIC_INSTITUTION":
		fallthrough
	case "OTHER":
		*e = CustomerOrganizationEntityTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CustomerOrganizationEntityTypeEnum: %v", v)
	}
}

type Customer struct {
	// Customer’s date of birth. The format is ‘yyyy-mm-dd’
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	// Customer’s gender - ‘male’ or ‘female’
	Gender *string `json:"gender,omitempty"`
	// Last four digits of the customer's social security number. This value is available for US customers.
	LastFourSsn *string `json:"last_four_ssn,omitempty"`
	// The customer's national identification number. This value is available for EU customers utilizing national identification numbers.
	NationalIdentificationNumber *string `json:"national_identification_number,omitempty"`
	// Organization entity type. Only applicable for B2B customers.
	OrganizationEntityType *CustomerOrganizationEntityTypeEnum `json:"organization_entity_type,omitempty"`
	// Organization registration id. Only applicable for B2B customers.
	OrganizationRegistrationID *string `json:"organization_registration_id,omitempty"`
	// Customer’s Title. Allowed values per country:
	// UK - "Mr", "Ms"
	// DE - "Herr", "Frau"
	// AT: "Herr, "Frau"
	// CH: de-CH: "Herr, "Frau" it-CH: "Sig.", "Sig.ra" fr-CH: "M", "Mme"
	// BE: "Dhr.", "Mevr."
	// NL: "Dhr.", "Mevr."
	Title *string `json:"title,omitempty"`
	// Type of customer in the session. If nothing is added, a B2C session will be the default. If it is a b2b-session, you should enter organization to trigger a B2B session.
	Type *string `json:"type,omitempty"`
	// VAT ID. Only applicable for B2B customers.
	VatID *string `json:"vat_id,omitempty"`
}

package models

type Medicament struct {
	AMM                             string `json:"amm"`
	ProductDesignation              string `json:"product_designation"`
	InternationalNonProprietaryName string `json:"international_nonProprietary_name"`
	LaboratoryAuthorizationHolder   string `json:"laboratory_authorization_holder"`
	LaboratoryManufacturer          string `json:"laboratory_manufacturer"`
	AuthorizationStartDate          string `json:"authorization_date"`
	AuthorizationEndDate            string `json:"authorization_end_date"`
}

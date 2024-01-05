package swagger

import "medicaments-cm/pkg/models"

// swagger:route GET /medicaments medicaments idOfmedicamentsWithoutID
// Medicaments returns the list of medicaments based on search keywords
// responses:
//   200: medicaments-Response
//   500: internalError

// This text will appear as description of your response body.
// swagger:response medicaments-Response
type MedicamentsResponseWrapper struct {
	// in:body
	Body []models.Medicament
}

// swagger:response internalError
type MedicamentsInternalErrorWrapper struct {
	// in:body
	Body models.ApiError
}

// swagger:parameters idOfmedicamentsWithoutID
type RatingsParam struct {
	//in:query
	//example: doliprane
	Search string `json:"search"`
}

package swagger

// swagger:route GET /health health idOfHealth
// Health endpoint to check if the API is up and running
// responses:
//   200: healthResponse

// This message will always be the same
// swagger:response healthResponse
type HealthResponseWrapper struct {
	// in:body
	Body struct {
		//Example: Medicament.cm, backend system good!
		Message string `json:"message"`
	}
}

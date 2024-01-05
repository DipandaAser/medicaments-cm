// Package classification Werr API.
//
// Documentation of Medicament.cm API.
//
//	Schemes: https
//	BasePath: /
//	Version: 1.0.0
//	Host: api.medicament.cm
//
//	Consumes:
//	- application/json; charset=utf-8
//
//	Produces:
//	- application/json
//
// swagger:meta
package swagger

import "medicaments-cm/pkg/models"

// swagger:response internalError
type InternalErrorWrapper struct {
	// in:body
	Body models.ApiError
}

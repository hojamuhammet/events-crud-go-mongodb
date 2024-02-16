package status

import "net/http"

const (
	BadRequest          = http.StatusBadRequest
	NotFound            = http.StatusNotFound
	OK                  = http.StatusOK
	InternalServerError = http.StatusInternalServerError
	Forbidden           = http.StatusForbidden
	Conflict            = http.StatusConflict
)

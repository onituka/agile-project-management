package apperrors

import "net/http"

var (
	InvalidParameter    = &appError{code: InvalidParameterCode, httpStatus: http.StatusBadRequest}
	NotFound            = &appError{code: NotFoundCode, httpStatus: http.StatusNotFound}
	Conflict            = &appError{code: ConflictCode, httpStatus: http.StatusConflict}
	InternalServerError = &appError{code: InternalServerErrorCode, httpStatus: http.StatusInternalServerError}
)

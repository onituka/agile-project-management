package apperrors

type code string

const (
	InvalidParameterCode    code = "InvalidParameter"
	ConflictCode            code = "Conflict"
	InternalServerErrorCode code = "InternalServerError"
	NotFoundCode            code = "projectNotFound"
)

func (c code) value() string {
	return string(c)
}

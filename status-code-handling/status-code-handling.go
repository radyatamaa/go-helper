package status_code_handling

import (
	models2 "go-helper/models"
	"net/http"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case models2.ErrInternalServerError:
		return http.StatusInternalServerError
	case models2.ErrForbidden:
		return http.StatusForbidden
	case models2.ErrNotFound:
		return http.StatusNotFound
	case models2.PersonalNumberNotFound:
		return http.StatusUnauthorized
	case models2.LoginFailedMessageLockedUser:
		return http.StatusBadRequest
	case models2.ErrUnAuthorize:
		return http.StatusUnauthorized
	case models2.ErrConflict:
		return http.StatusConflict
	case models2.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}


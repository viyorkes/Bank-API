package errs

import "net/http"

type AppError struct{
	Code int  `json:"omitempy"`
	Message string `json:"message"`
}


func  NewNotFoundErrors(message string) *AppError{

	return &AppError{
		Message: message,
		Code: http.StatusFound,

	}
}

func (e AppError) AsMessage() *AppError {

	return &AppError{

		Message: e.Message,
	}
}


func  NewInternalServerErrors(message string) *AppError{

	return &AppError{
		Message: message,
		Code: http.StatusInternalServerError,

	}
}
package errs

import "net/http"

type AppError struct{
	Code int
	Message string
}


func  NewNotFoundErrors(message string) *AppError{

	return &AppError{
		Message: message,
		Code: http.StatusFound,

	}
}

func  NewInternalServerErrors(message string) *AppError{

	return &AppError{
		Message: message,
		Code: http.StatusInternalServerError,

	}
}
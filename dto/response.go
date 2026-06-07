package dto

import "net/http"

type ResponseDto struct {
	Status      int    `json:"status"`
	ErrorSchema string `json:"error_schema"`
	Payload     any    `json:"payload"`
	Validation  any    `json:"validation"`
}

func SuccessResponse(data any) ResponseDto {
	return ResponseDto{
		Status:      http.StatusOK,
		ErrorSchema: "Success",
		Payload:     data,
		Validation:  nil,
	}
}

func ErrorResponse() ResponseDto {
	return ResponseDto{
		Status:      http.StatusInternalServerError,
		ErrorSchema: "Internal Server Error",
		Payload:     nil,
		Validation:  nil,
	}
}

func FailedResponse(msg string, status int) ResponseDto {
	return ResponseDto{
		Status:      status,
		ErrorSchema: msg,
		Payload:     nil,
		Validation:  nil,
	}
}

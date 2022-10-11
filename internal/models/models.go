package models

import "rest-calculator/internal/models/operator"

type Operation struct {
	NumberOne *float64          `json:"number_one" binding:"required"`
	NumberTwo *float64          `json:"number_two" binding:"required"`
	Operator  operator.Operator `json:"operator" binding:"required"`
}

type ServerErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ServerSuccessResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

func NewServerSuccessResponse(result string) *ServerSuccessResponse {
	return &ServerSuccessResponse{
		Success: true,
		Result:  result,
	}
}

func NewServerErrorResponse(error string) *ServerErrorResponse {
	return &ServerErrorResponse{
		Success: false,
		Error:   error,
	}
}

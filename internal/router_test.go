package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"rest-calculator/internal/models"
	"rest-calculator/internal/models/error_messages"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := Inject()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"status\":\"alive\"}", w.Body.String())
}

func TestExecuteOperationRoute_Add(t *testing.T) {
	router := Inject()

	object, buffer := map[string]interface{}{
		"number_one": 12.2,
		"number_two": 11.1,
		"operator":   "+",
	}, new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(object)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/execute", buffer)
	router.ServeHTTP(w, req)

	var response models.ServerSuccessResponse
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, models.NewServerSuccessResponse(fmt.Sprintf("%.5f", 12.2+11.1)), response)
}

func TestExecuteOperationRoute_Minus(t *testing.T) {
	router := Inject()

	object, buffer := map[string]interface{}{
		"number_one": 12.2,
		"number_two": 11.1,
		"operator":   "-",
	}, new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(object)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/execute", buffer)
	router.ServeHTTP(w, req)

	var response models.ServerSuccessResponse
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, models.NewServerSuccessResponse(fmt.Sprintf("%.5f", 12.2-11.1)), response)
}

func TestExecuteOperationRoute_Multiply(t *testing.T) {
	router := Inject()

	object, buffer := map[string]interface{}{
		"number_one": 12.2,
		"number_two": 11.1,
		"operator":   "*",
	}, new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(object)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/execute", buffer)
	router.ServeHTTP(w, req)

	var response models.ServerSuccessResponse
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, models.NewServerSuccessResponse(fmt.Sprintf("%.5f", 12.2*11.1)), response)
}

func TestExecuteOperationRoute_Divide(t *testing.T) {
	router := Inject()

	object, buffer := map[string]interface{}{
		"number_one": 12.2,
		"number_two": 11.1,
		"operator":   "/",
	}, new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(object)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/execute", buffer)
	router.ServeHTTP(w, req)

	var response models.ServerSuccessResponse
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, models.NewServerSuccessResponse(fmt.Sprintf("%.5f", 12.2/11.1)), response)
}

func TestExecuteOperationRoute_Divide_By_Zero(t *testing.T) {
	router := Inject()

	object, buffer := map[string]interface{}{
		"number_one": 12.2,
		"number_two": 0,
		"operator":   "/",
	}, new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(object)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/execute", buffer)
	router.ServeHTTP(w, req)

	var response models.ServerErrorResponse
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, models.NewServerErrorResponse(error_messages.DivisionByZero), response)
}

func TestExecuteOperationRoute_Invalid_Operator(t *testing.T) {
	router := Inject()

	object, buffer := map[string]interface{}{
		"number_one": 1111,
		"number_two": 1111,
		"operator":   "ss",
	}, new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(object)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/execute", buffer)
	router.ServeHTTP(w, req)

	var response models.ServerErrorResponse
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, models.NewServerErrorResponse(error_messages.UnknownOperator), response)
}

package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	apperrors "github.com/emaldie/secret-api/pkg/errors"
)

type Response struct {
	StatusCode int             `json:"statusCode"`
	Success    bool            `json:"success"`
	Data       interface{}     `json:"data,omitempty"`
	Error      apperrors.Error `json:"error,omitempty"`
}

type ErrorInfo struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	Fields  []string `json:"fields,omitempty"`
}

func RespondJSON(w http.ResponseWriter, status int, data interface{}, error apperrors.Error) {
	response := Response{
		StatusCode: status,
		Success:    status >= 200 && status < 300,
		Error:      error,
		Data:       data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Response to JSON serialization failure", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func RespondError(w http.ResponseWriter, err error) {
	var appErr *apperrors.Error

	if !errors.As(err, &appErr) {
		appErr = apperrors.InternalError("Internal server error", err)
	}

	status := appErr.StatusCode()

	response := Response{
		StatusCode: status,
		Success:    false,
		Data:       nil,
		Error:      *appErr,
	}
	slog.Error(appErr.Message, "error", appErr, "type", string(appErr.Type))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("JSON serialization failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func DecodeJSON(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return apperrors.BadRequestError("Invalid JSON data", err)
	}
	return nil
}

func BindJSON(r *http.Request, v interface{}, validate func(interface{}) error) error {
	if err := DecodeJSON(r, v); err != nil {
		return err
	}

	if validate != nil {
		if err := validate(v); err != nil {
			return apperrors.ValidationError("Data validation error", err)
		}
	}

	return nil
}

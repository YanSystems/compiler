package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	slog.Debug("ReadJSON called")
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		slog.Error("Failed to decode JSON", "error", err)
		return err
	}
	slog.Debug("JSON decoded successfully", "data", data)

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		err := errors.New("body must have only a single JSON value")
		slog.Error("JSON body must have only a single value", "error", err)
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	slog.Debug("WriteJSON called", "status", status, "data", data)

	out, err := json.Marshal(data)
	if err != nil {
		slog.Error("Failed to marshal JSON", "error", err)
		return err
	}
	slog.Debug("JSON marshalled successfully")

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
		slog.Debug("Custom headers set", "headers", headers[0])
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		slog.Error("Failed to write JSON response", "error", err)
		return err
	}

	slog.Debug("JSON response written successfully")
	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	slog.Debug("ErrorJSON called", "error", err)
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
		slog.Debug("Custom status code set", "statusCode", statusCode)
	}

	var payload JsonResponse
	payload.Error = true
	payload.Message = err.Error()

	slog.Debug("Error payload prepared", "payload", payload)
	return WriteJSON(w, statusCode, payload)
}

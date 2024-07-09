package services

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/YanSystems/compiler/pkg/compiler"
	"github.com/YanSystems/compiler/pkg/utils"
)

func HandleExecutePython(w http.ResponseWriter, r *http.Request) {
	slog.Debug("HandleExecutePython called")

	var userSubmittedCode compiler.Code
	err := utils.ReadJSON(w, r, &userSubmittedCode)
	if err != nil {
		slog.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Debug("JSON request body read successfully", "code", userSubmittedCode)

	if userSubmittedCode.Lang == "" || userSubmittedCode.Src == "" {
		err := errors.New("missing fields in request payload")
		slog.Error("Validation error: missing fields in request payload", "error", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	slog.Debug("Request payload validation passed", "lang", userSubmittedCode.Lang, "src", userSubmittedCode.Src)

	slog.Info("Executing Python script", "script", userSubmittedCode.Src)
	output, err := userSubmittedCode.Execute()
	if err != nil {
		slog.Error("Failed to execute Python script", "error", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	slog.Info("Python script executed successfully", "output", output)

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Python script has been successfully processed",
		Data:    output,
	}

	slog.Info("Sending response payload", "responsePayload", responsePayload)
	utils.WriteJSON(w, http.StatusOK, responsePayload)
	slog.Info("Response sent successfully", "status", http.StatusOK)
}

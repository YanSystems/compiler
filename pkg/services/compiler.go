package services

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/YanSystems/compiler/pkg/compiler"
	"github.com/YanSystems/compiler/pkg/utils"
)

func HandleExecutePython(w http.ResponseWriter, r *http.Request) {
	var userSubmittedCode compiler.Code
	err := utils.ReadJSON(w, r, &userSubmittedCode)
	if err != nil {
		slog.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	if userSubmittedCode.Lang == "" || userSubmittedCode.Src == "" {
		err := errors.New("missing fields in request payload")
		slog.Error("Validation error", "error", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	output, err := userSubmittedCode.Execute()
	if err != nil {
		slog.Error("Failed to execute Python script", "error", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Python script has been successfully processed",
		Data:    output,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}

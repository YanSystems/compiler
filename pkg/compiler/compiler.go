package compiler

import (
	"fmt"
	"log/slog"
)

type Code struct {
	Lang string   `json:"lang"`
	Src  string   `json:"src"`
	Args []string `json:"args"`
}

type ExecutionResult struct {
	Error  bool   `json:"error"`
	Output string `json:"output"`
}

func (c *Code) Execute() (*ExecutionResult, error) {
	slog.Debug("Execute called", "language", c.Lang)
	var result *ExecutionResult
	var err error

	switch c.Lang {
	case "python":
		slog.Debug("Executing Python code", "source", c.Src)
		result, err = c.executePython()
		if err != nil {
			slog.Error("Failed to execute Python code", "error", err)
		} else {
			slog.Info("Python code executed successfully", "result", result)
		}
	default:
		slog.Error("Unsupported language", "language", c.Lang)
		result = &ExecutionResult{
			Error:  false,
			Output: fmt.Sprintf("unexpected error: language %s is not supported", c.Lang),
		}
		err = nil
	}

	slog.Debug("Execution result", "result", result)
	return result, err
}

package compiler

import (
	"bytes"
	"log/slog"
	"os/exec"
	"strings"
)

func (c *Code) executePython() (*ExecutionResult, error) {
	slog.Debug("Preparing to execute Python code", "source", c.Src, "args", c.Args)

	args := append([]string{"-c", c.Src}, c.Args...)
	cmd := exec.Command("python3", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	slog.Debug("Running Python command", "command", cmd.String())

	err := cmd.Run()
	if err != nil {
		errorMessage := cleanErrorMessage(stderr.String())
		slog.Error("Python script execution failed", "error", errorMessage)
		return &ExecutionResult{
			Error:  true,
			Output: errorMessage,
		}, nil
	}

	slog.Debug("Python script executed successfully", "output", out.String())
	return &ExecutionResult{
		Error:  false,
		Output: out.String(),
	}, nil
}

func cleanErrorMessage(msg string) string {
	lines := strings.Split(msg, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			return line
		}
	}
	return strings.TrimSpace(msg)
}

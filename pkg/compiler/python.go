package compiler

import (
	"bytes"
	"os/exec"
	"strings"
)

func (c *Code) executePython() (*ExecutionResult, error) {
	args := append([]string{"-c", c.Src}, c.Args...)
	cmd := exec.Command("python3", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return &ExecutionResult{
			Error:  true,
			Output: cleanErrorMessage(stderr.String()),
		}, nil
	}

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

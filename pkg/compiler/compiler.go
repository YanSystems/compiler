package compiler

import "fmt"

type Code struct {
	Lang string   `json:"language"`
	Src  string   `json:"src"`
	Args []string `json:"arguments"`
}

type ExecutionResult struct {
	Error  bool   `json:"error"`
	Output string `json:"output"`
}

func (c *Code) Execute() (*ExecutionResult, error) {
	var result *ExecutionResult
	var err error

	switch c.Lang {
	case "python":
		result, err = c.executePython()
	default:
		result = &ExecutionResult{
			Error:  false,
			Output: fmt.Sprintf("unexpected error: language %s is not supported", c.Lang),
		}
		err = nil
	}

	return result, err
}

package compiler

import "fmt"

type Code struct {
	Language  string
	Src       string
	Arguments []string
}

type ExecutionResult struct {
	Error  bool
	Output string
}

func (c *Code) Execute() (*ExecutionResult, error) {
	var result *ExecutionResult
	var err error

	switch c.Language {
	case "python":
		result, err = c.executePython()
	default:
		return nil, fmt.Errorf("unknown language: %s", c.Language)
	}

	return result, err
}

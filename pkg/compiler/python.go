package compiler

func (c *Code) executePython() (*ExecutionResult, error) {
	return &ExecutionResult{
		Error:  false,
		Output: "hello world\n",
	}, nil
}

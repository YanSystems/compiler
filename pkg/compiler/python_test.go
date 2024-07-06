package compiler

import "testing"

type TestCase struct {
	Label    string
	Code     Code
	Expected ExecutionResult
}

func TestValidPythonCode(t *testing.T) {
	tc := TestCase{
		Label: "Valid Python Code",
		Code: Code{
			Lang: "python",
			Src:  "print('hello world')",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "hello world\n",
		},
	}

	tc.Run(t)
}

func TestInvalidPythonCode(t *testing.T) {
	tc := TestCase{
		Label: "Invalid Python Code",
		Code: Code{
			Lang: "python",
			Src:  "print('hello world'",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  true,
			Output: "SyntaxError: '(' was never closed",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithArgs(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Args",
		Code: Code{
			Lang: "python",
			Src:  "import sys\nprint(sys.argv[1])",
			Args: []string{"test-arg"},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "test-arg\n",
		},
	}

	tc.Run(t)
}

func TestLongRunningPythonCode(t *testing.T) {
	tc := TestCase{
		Label: "Long-running Python Code",
		Code: Code{
			Lang: "python",
			Src:  "import time\ntime.sleep(2)\nprint('Done')",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "Done\n",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithSyntaxError(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Syntax Error",
		Code: Code{
			Lang: "python",
			Src:  "def func:\nprint('Hello')",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  true,
			Output: "SyntaxError: invalid syntax",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithRuntimeError(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Runtime Error",
		Code: Code{
			Lang: "python",
			Src:  "print(1 / 0)",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  true,
			Output: "ZeroDivisionError: division by zero",
		},
	}

	tc.Run(t)
}

func TestEmptyPythonCode(t *testing.T) {
	tc := TestCase{
		Label: "Empty Python Code",
		Code: Code{
			Lang: "python",
			Src:  "",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithUnicodeCharacters(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Unicode Characters",
		Code: Code{
			Lang: "python",
			Src:  "print('こんにちは世界')",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "こんにちは世界\n",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithMultipleArgs(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Multiple Args",
		Code: Code{
			Lang: "python",
			Src:  "import sys\nprint(' '.join(sys.argv[1:]))",
			Args: []string{"arg1", "arg2", "arg3"},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "arg1 arg2 arg3\n",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithEnvironmentVariable(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Environment Variable",
		Code: Code{
			Lang: "python",
			Src:  "import os\nprint(os.getenv('TEST_ENV'))",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "test_value\n",
		},
	}

	t.Setenv("TEST_ENV", "test_value")
	tc.Run(t)
}

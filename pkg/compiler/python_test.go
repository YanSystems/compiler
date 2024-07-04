package compiler

import "testing"

type TestCase struct {
	Label    string
	Code     Code
	Expected ExecutionResult
}

func (tc *TestCase) Run(t *testing.T) {
	t.Run(tc.Label, func(t *testing.T) {
		result, err := tc.Code.Execute()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if result.Error != tc.Expected.Error || result.Output != tc.Expected.Output {
			t.Errorf("executePython() = %v, Expected %v", result, tc.Expected)
		}
	})
}

func TestValidPythonCode(t *testing.T) {
	tc := TestCase{
		Label: "Valid Python Code",
		Code: Code{
			Language:  "python",
			Src:       "print('hello world')",
			Arguments: []string{},
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
			Language:  "python",
			Src:       "print('hello world'",
			Arguments: []string{},
		},
		Expected: ExecutionResult{
			Error:  true,
			Output: "SyntaxError: EOL while scanning string literal\n",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithArguments(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Arguments",
		Code: Code{
			Language:  "python",
			Src:       "import sys\nprint(sys.argv[1])",
			Arguments: []string{"test-arg"},
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
			Language:  "python",
			Src:       "import time\ntime.sleep(2)\nprint('Done')",
			Arguments: []string{},
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
			Language:  "python",
			Src:       "def func:\nprint('Hello')",
			Arguments: []string{},
		},
		Expected: ExecutionResult{
			Error:  true,
			Output: "SyntaxError: invalid syntax\n",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithRuntimeError(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Runtime Error",
		Code: Code{
			Language:  "python",
			Src:       "print(1 / 0)",
			Arguments: []string{},
		},
		Expected: ExecutionResult{
			Error:  true,
			Output: "ZeroDivisionError: division by zero\n",
		},
	}

	tc.Run(t)
}

func TestEmptyPythonCode(t *testing.T) {
	tc := TestCase{
		Label: "Empty Python Code",
		Code: Code{
			Language:  "python",
			Src:       "",
			Arguments: []string{},
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
			Language:  "python",
			Src:       "print('こんにちは世界')",
			Arguments: []string{},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "こんにちは世界\n",
		},
	}

	tc.Run(t)
}

func TestPythonCodeWithMultipleArguments(t *testing.T) {
	tc := TestCase{
		Label: "Python Code with Multiple Arguments",
		Code: Code{
			Language:  "python",
			Src:       "import sys\nprint(' '.join(sys.argv[1:]))",
			Arguments: []string{"arg1", "arg2", "arg3"},
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
			Language:  "python",
			Src:       "import os\nprint(os.getenv('TEST_ENV'))",
			Arguments: []string{},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: "test_value\n",
		},
	}

	t.Setenv("TEST_ENV", "test_value")
	tc.Run(t)
}

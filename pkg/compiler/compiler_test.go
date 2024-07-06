package compiler

import (
	"fmt"
	"testing"
)

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

func TestUnsupportedLanguage(t *testing.T) {
	lang := "unsupported_lang"
	tc := TestCase{
		Label: lang,
		Code: Code{
			Lang: "unsupported_lang",
			Src:  "print('hello world')",
			Args: []string{},
		},
		Expected: ExecutionResult{
			Error:  false,
			Output: fmt.Sprintf("unexpected error: language %s is not supported", lang),
		},
	}

	tc.Run(t)
}

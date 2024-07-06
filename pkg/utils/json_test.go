package utils

type ReadTestCase struct {
	Name           string `json:"name"`
	Input          string `json:"input"`
	ExpectedError  string `json:"expected_error"`
	ExpectedOutput string `json:"expected_output"`
}

package utils

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

type ReadTestCase struct {
	Name           string
	Input          string
	ExpectedError  string
	ExpectedOutput map[string]any
}

func (rtc *ReadTestCase) RunReadTest(t *testing.T) {
	t.Run(rtc.Name, func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", bytes.NewReader([]byte(rtc.Input)))
		rr := httptest.NewRecorder()

		var data map[string]any
		err := ReadJSON(rr, req, &data)

		actualError := ""
		if err != nil {
			actualError = err.Error()
		}

		if rtc.ExpectedError != "" {
			if actualError != rtc.ExpectedError {
				t.Errorf("expected error %q but got %q", rtc.ExpectedError, actualError)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error but got %q", actualError)
			}
			if !equalMaps(data, rtc.ExpectedOutput) {
				t.Errorf("expected output %v but got %v", rtc.ExpectedOutput, data)
			}
		}
	})
}

func TestReadValidJSON(t *testing.T) {
	rtc := ReadTestCase{
		Name:           "Valid JSON",
		Input:          `{"name":"John", "age":30}`,
		ExpectedError:  "",
		ExpectedOutput: map[string]any{"name": "John", "age": float64(30)},
	}

	rtc.RunReadTest(t)
}

func TestReadInvalidJSON(t *testing.T) {
	rtc := ReadTestCase{
		Name:           "Invalid JSON",
		Input:          `{"name":"John", "age":30`,
		ExpectedError:  "unexpected EOF",
		ExpectedOutput: nil,
	}

	rtc.RunReadTest(t)
}

func TestReadEmptyBody(t *testing.T) {
	rtc := ReadTestCase{
		Name:           "Empty Body",
		Input:          ``,
		ExpectedError:  "EOF",
		ExpectedOutput: nil,
	}

	rtc.RunReadTest(t)
}

func TestReadMultipleJSONObjects(t *testing.T) {
	rtc := ReadTestCase{
		Name:           "Multiple JSON Objects",
		Input:          `{"name":"John"}{"age":30}`,
		ExpectedError:  "body must have only a single JSON value",
		ExpectedOutput: nil,
	}

	rtc.RunReadTest(t)
}

func TestReadExceedMaxBytes(t *testing.T) {
	largeJSON := `{"name":"` + string(make([]byte, 1048560)) + `"}`

	rtc := ReadTestCase{
		Name:           "Exceed Max Bytes",
		Input:          largeJSON,
		ExpectedError:  "invalid character '\\x00' in string literal",
		ExpectedOutput: nil,
	}

	rtc.RunReadTest(t)
}

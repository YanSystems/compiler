package utils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ErrorTestCase struct {
	Name            string
	Err             error
	Status          []int
	ExpectedOutput  string
	ExpectedStatus  int
	ExpectedHeaders http.Header
	ExpectedError   string
}

func (etc *ErrorTestCase) RunErrorTest(t *testing.T) {
	t.Run(etc.Name, func(t *testing.T) {
		rr := httptest.NewRecorder()

		err := ErrorJSON(rr, etc.Err, etc.Status...)
		actualError := ""
		if err != nil {
			actualError = err.Error()
		}

		if etc.ExpectedError != "" {
			if actualError != etc.ExpectedError {
				t.Errorf("expected error %q but got %q", etc.ExpectedError, actualError)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error but got %q", actualError)
			}
			if rr.Code != etc.ExpectedStatus {
				t.Errorf("expected status %d but got %d", etc.ExpectedStatus, rr.Code)
			}
			if rr.Body.String() != etc.ExpectedOutput {
				t.Errorf("expected body %q but got %q", etc.ExpectedOutput, rr.Body.String())
			}
			for key, value := range etc.ExpectedHeaders {
				if rr.Header().Get(key) != value[0] {
					t.Errorf("expected header %q but got %q", value[0], rr.Header().Get(key))
				}
			}
		}
	})
}

func TestErrorDefaultStatus(t *testing.T) {
	etc := ErrorTestCase{
		Name:           "Default Status Code",
		Err:            errors.New("default error"),
		Status:         nil,
		ExpectedOutput: `{"error":true,"message":"default error"}`,
		ExpectedStatus: http.StatusBadRequest,
	}

	etc.RunErrorTest(t)
}

func TestErrorCustomStatus(t *testing.T) {
	etc := ErrorTestCase{
		Name:           "Custom Status Code",
		Err:            errors.New("custom error"),
		Status:         []int{http.StatusInternalServerError},
		ExpectedOutput: `{"error":true,"message":"custom error"}`,
		ExpectedStatus: http.StatusInternalServerError,
	}

	etc.RunErrorTest(t)
}

func TestErrorMessageFormatting(t *testing.T) {
	etc := ErrorTestCase{
		Name:           "Error Message Formatting",
		Err:            errors.New("formatted error"),
		Status:         nil,
		ExpectedOutput: `{"error":true,"message":"formatted error"}`,
		ExpectedStatus: http.StatusBadRequest,
	}

	etc.RunErrorTest(t)
}

func TestErrorWithHeaders(t *testing.T) {
	etc := ErrorTestCase{
		Name:           "Error With Headers",
		Err:            errors.New("error with headers"),
		Status:         []int{http.StatusForbidden},
		ExpectedOutput: `{"error":true,"message":"error with headers"}`,
		ExpectedStatus: http.StatusForbidden,
		ExpectedHeaders: http.Header{
			"Content-Type": {"application/json"},
		},
	}

	etc.RunErrorTest(t)
}

func TestErrorWriteFailure(t *testing.T) {
	etc := ErrorTestCase{
		Name:           "Error Write Failure",
		Err:            errors.New("write failure"),
		Status:         nil,
		ExpectedOutput: "",
		ExpectedStatus: http.StatusBadRequest,
		ExpectedError:  "write error",
	}

	t.Run(etc.Name, func(t *testing.T) {
		fw := &failWriter{
			ResponseWriter: httptest.NewRecorder(),
		}

		err := ErrorJSON(fw, etc.Err, etc.Status...)
		if err == nil || err.Error() != etc.ExpectedError {
			t.Errorf("expected error %q but got %v", etc.ExpectedError, err)
		}
	})
}

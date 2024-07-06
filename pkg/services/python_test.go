package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YanSystems/compiler/pkg/compiler"
	"github.com/YanSystems/compiler/pkg/utils"
)

type PythonHandlerTestCase struct {
	Name             string
	RequestPayload   any
	ExpectedStatus   int
	ExpectedResponse utils.JsonResponse
	ExpectedError    string
}

func (phtc *PythonHandlerTestCase) RunPythonHandlerTest(t *testing.T) {
	t.Run(phtc.Name, func(t *testing.T) {
		requestPayload, err := json.Marshal(phtc.RequestPayload)
		if err != nil {
			t.Fatalf("could not marshal request payload: %v", err)
		}
		server := httptest.NewServer(http.HandlerFunc(HandleExecutePython))
		defer server.Close()
		resp, err := http.Post(server.URL, "application/json", bytes.NewReader(requestPayload))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		var responsePayload utils.JsonResponse
		err = json.NewDecoder(resp.Body).Decode(&responsePayload)

		if phtc.ExpectedError != "" {
			if resp.StatusCode != phtc.ExpectedStatus {
				t.Fatalf("expected status %d but got %d", phtc.ExpectedStatus, resp.StatusCode)
			}

			if err != nil {
				t.Fatalf("could not unmarshal response: %v", err)
			}

			if responsePayload.Message != phtc.ExpectedError {
				t.Fatalf("expected error %q but got %q", phtc.ExpectedError, responsePayload.Message)
			}
		} else {
			if resp.StatusCode != phtc.ExpectedStatus {
				t.Fatalf("expected status %d but got %d", phtc.ExpectedStatus, resp.StatusCode)
			}

			if err != nil {
				t.Fatalf("could not unmarshal response: %v", err)
			}

			if responsePayload.Error != phtc.ExpectedResponse.Error {
				t.Fatalf("expected error %v but got %v", phtc.ExpectedResponse.Error, responsePayload.Error)
			}
			if responsePayload.Message != phtc.ExpectedResponse.Message {
				t.Fatalf("expected message %q but got %q", phtc.ExpectedResponse.Message, responsePayload.Message)
			}

			var actualData compiler.ExecutionResult
			dataBytes, err := json.Marshal(responsePayload.Data)
			if err != nil {
				t.Fatalf("could not marshal response data: %v", err)
			}
			err = json.Unmarshal(dataBytes, &actualData)
			if err != nil {
				t.Fatalf("could not unmarshal response data: %v", err)
			}

			expectedData, ok := phtc.ExpectedResponse.Data.(compiler.ExecutionResult)
			if !ok {
				t.Fatalf("expected response data to be of type compiler.ExecutionResult")
			}

			if actualData.Error != expectedData.Error {
				t.Fatalf("expected data error %v but got %v", expectedData.Error, actualData.Error)
			}
			if actualData.Output != expectedData.Output {
				t.Fatalf("expected data output %v but got %v", expectedData.Output, actualData.Output)
			}
		}
	})
}

func TestExecutePythonValidRequest(t *testing.T) {
	expectedResult := compiler.ExecutionResult{
		Error:  false,
		Output: "hello world\n",
	}
	htc := PythonHandlerTestCase{
		Name:           "Valid Request",
		RequestPayload: compiler.Code{Lang: "python", Src: "print('hello world')", Args: []string{}},
		ExpectedStatus: http.StatusOK,
		ExpectedResponse: utils.JsonResponse{
			Error:   false,
			Message: "Python script has been successfully processed",
			Data:    expectedResult,
		},
	}

	htc.RunPythonHandlerTest(t)
}

func TestExecutePythonInvalidJSON(t *testing.T) {
	htc := PythonHandlerTestCase{
		Name:           "Invalid JSON",
		RequestPayload: "invalid JSON",
		ExpectedStatus: http.StatusBadRequest,
		ExpectedResponse: utils.JsonResponse{
			Error:   true,
			Message: "json: cannot unmarshal string into Go value of type compiler.Code",
		},
		ExpectedError: "json: cannot unmarshal string into Go value of type compiler.Code",
	}

	htc.RunPythonHandlerTest(t)
}

func TestExecutePythonEmptyRequestBody(t *testing.T) {
	htc := PythonHandlerTestCase{
		Name:           "Empty Request Body",
		RequestPayload: nil,
		ExpectedStatus: http.StatusBadRequest,
		ExpectedResponse: utils.JsonResponse{
			Error:   true,
			Message: "missing fields in request payload",
		},
		ExpectedError: "missing fields in request payload",
	}

	htc.RunPythonHandlerTest(t)
}

func TestExecutePythonMissingFields(t *testing.T) {
	htc := PythonHandlerTestCase{
		Name:           "Missing Fields",
		RequestPayload: compiler.Code{Lang: "", Src: "", Args: []string{}},
		ExpectedStatus: http.StatusBadRequest,
		ExpectedResponse: utils.JsonResponse{
			Error:   true,
			Message: "missing fields in request payload",
		},
		ExpectedError: "missing fields in request payload",
	}

	htc.RunPythonHandlerTest(t)
}

func TestExecutePythonRuntimeError(t *testing.T) {
	expectedResult := compiler.ExecutionResult{
		Error:  true,
		Output: "ZeroDivisionError: division by zero",
	}
	htc := PythonHandlerTestCase{
		Name:           "Runtime Error",
		RequestPayload: compiler.Code{Lang: "python", Src: "print(1 / 0)", Args: []string{}},
		ExpectedStatus: http.StatusOK,
		ExpectedResponse: utils.JsonResponse{
			Error:   false,
			Message: "Python script has been successfully processed",
			Data:    expectedResult,
		},
	}

	htc.RunPythonHandlerTest(t)
}

package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type WriteTestCase struct {
	Name           string
	Status         int
	Data           any
	Headers        http.Header
	ExpectedOutput string
	ExpectedError  string
}

func (wtc *WriteTestCase) RunWriteTest(t *testing.T) {
	t.Run(wtc.Name, func(t *testing.T) {
		rr := httptest.NewRecorder()

		err := WriteJSON(rr, wtc.Status, wtc.Data, wtc.Headers)
		actualError := ""
		if err != nil {
			actualError = err.Error()
		}

		if wtc.ExpectedError != "" {
			if actualError != wtc.ExpectedError {
				t.Errorf("expected error %q but got %q", wtc.ExpectedError, actualError)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error but got %q", actualError)
			}
			if rr.Code != wtc.Status {
				t.Errorf("expected status %d but got %d", wtc.Status, rr.Code)
			}
			if !jsonEqual(rr.Body.String(), wtc.ExpectedOutput) {
				t.Errorf("expected body %q but got %q", wtc.ExpectedOutput, rr.Body.String())
			}
			if wtc.Headers != nil {
				for key, value := range wtc.Headers {
					if rr.Header().Get(key) != value[0] {
						t.Errorf("expected header %q but got %q", value[0], rr.Header().Get(key))
					}
				}
			}
		}
	})
}

func TestWriteValidJSON(t *testing.T) {
	wtc := WriteTestCase{
		Name:           "Valid JSON Response",
		Status:         http.StatusOK,
		Data:           map[string]any{"name": "John", "age": 30},
		Headers:        nil,
		ExpectedOutput: `{"name":"John","age":30}`,
		ExpectedError:  "",
	}

	wtc.RunWriteTest(t)
}

func TestWriteInvalidJSON(t *testing.T) {
	wtc := WriteTestCase{
		Name:           "Invalid JSON Data",
		Status:         http.StatusOK,
		Data:           make(chan int), // JSON marshaling error
		Headers:        nil,
		ExpectedOutput: "",
		ExpectedError:  "json: unsupported type: chan int",
	}

	wtc.RunWriteTest(t)
}

func TestWriteWithCustomHeaders(t *testing.T) {
	wtc := WriteTestCase{
		Name:           "With Custom Headers",
		Status:         http.StatusCreated,
		Data:           map[string]any{"message": "created"},
		Headers:        http.Header{"X-Custom-Header": {"CustomValue"}},
		ExpectedOutput: `{"message":"created"}`,
		ExpectedError:  "",
	}

	wtc.RunWriteTest(t)
}

func TestWriteEmptyData(t *testing.T) {
	wtc := WriteTestCase{
		Name:           "Empty Data",
		Status:         http.StatusNoContent,
		Data:           nil,
		Headers:        nil,
		ExpectedOutput: "null",
		ExpectedError:  "",
	}

	wtc.RunWriteTest(t)
}

func TestWriteMultipleHeaders(t *testing.T) {
	wtc := WriteTestCase{
		Name:           "Multiple Headers",
		Status:         http.StatusOK,
		Data:           map[string]any{"status": "ok"},
		Headers:        http.Header{"X-Custom-Header-1": {"Value1"}, "X-Custom-Header-2": {"Value2"}},
		ExpectedOutput: `{"status":"ok"}`,
		ExpectedError:  "",
	}

	wtc.RunWriteTest(t)
}

func TestWriteFailure(t *testing.T) {
	wtc := WriteTestCase{
		Name:           "Write Failure",
		Status:         http.StatusInternalServerError,
		Data:           map[string]any{"error": "something went wrong"},
		Headers:        nil,
		ExpectedOutput: "",
		ExpectedError:  "write error",
	}

	t.Run(wtc.Name, func(t *testing.T) {
		fw := &failWriter{}
		err := WriteJSON(fw, wtc.Status, wtc.Data, wtc.Headers)
		if err == nil || err.Error() != wtc.ExpectedError {
			t.Errorf("expected error %q but got %v", wtc.ExpectedError, err)
		}
	})
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestNewServer(t *testing.T) {
	api := Server{Port: "8000"}
	server := api.NewServer()

	// Check if the server address is correctly set
	expectedAddr := fmt.Sprintf(":%s", api.Port)
	if server.Addr != expectedAddr {
		t.Errorf("expected server address %s but got %s", expectedAddr, server.Addr)
	}

	// Check if the handler is not nil
	if server.Handler == nil {
		t.Error("expected server handler to be set but got nil")
	} else {
		// Check if the handler is a chi.Mux
		if _, ok := server.Handler.(*chi.Mux); !ok {
			t.Error("expected server handler to be of type chi.Mux")
		}
	}
}

func TestNewRouter(t *testing.T) {
	api := Server{Port: "8000"}
	router := api.NewRouter()

	t.Run("CORS Preflight", func(t *testing.T) {
		req, err := http.NewRequest("OPTIONS", "/python", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		req.Header.Set("Origin", "http://localhost:3000")
		req.Header.Set("Access-Control-Request-Method", "POST")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
		}

		if rr.Header().Get("Access-Control-Allow-Origin") != "http://localhost:3000" {
			t.Errorf("expected Access-Control-Allow-Origin to be %q but got %q", "http://localhost:3000", rr.Header().Get("Access-Control-Allow-Origin"))
		}
	})

	t.Run("Health Check", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/health", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
		}

		expectedBody := "OK"
		if rr.Body.String() != expectedBody {
			t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
		}
	})
}

type MockServer struct {
	ListenAndServeFunc func() error
}

func (m *MockServer) ListenAndServe() error {
	if m.ListenAndServeFunc != nil {
		return m.ListenAndServeFunc()
	}
	return nil
}

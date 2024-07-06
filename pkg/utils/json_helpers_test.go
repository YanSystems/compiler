package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func jsonEqual(a, b string) bool {
	var j1, j2 map[string]any
	err1 := json.Unmarshal([]byte(a), &j1)
	err2 := json.Unmarshal([]byte(b), &j2)
	if err1 != nil || err2 != nil {
		return false
	}
	return equalMaps(j1, j2)
}

func equalMaps(a, b map[string]any) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if vb, ok := b[k]; !ok || vb != v {
			return false
		}
	}
	return true
}

type failWriter struct {
	http.ResponseWriter
}

func (fw *failWriter) Header() http.Header {
	return http.Header{}
}

func (fw *failWriter) Write([]byte) (int, error) {
	return 0, errors.New("write error")
}

func (fw *failWriter) WriteHeader(statusCode int) {}

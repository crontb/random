package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/crontb/random/pkg/http/handler"
	"github.com/gorilla/mux"
)

func TestCard(t *testing.T) {
	tests := []struct {
		name       string
		route      string
		method     string
		wantStatus int
	}{
		{"Case 1", "/card", http.MethodOptions, http.StatusOK},
		{"Case 2", "/card", http.MethodGet, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name+": "+tt.route, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc(tt.route, handler.Card)
			router.ServeHTTP(recorder, req)

			if recorder.Code != tt.wantStatus {
				t.Errorf("failed on route %s: got %v, want %v",
					tt.route, recorder.Code, tt.wantStatus)
			}

			/*
				result := recorder.Result()
				defer result.Body.Close()
				body, err := ioutil.ReadAll(result.Body)
				if err != nil {
					t.Fatalf("could not read response: %v", err)
				}
			*/

			// @todo: validate body JSON (body.result, body.error) if requested in test case
		})
	}
}

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name               string
		queryParam         string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "missing value",
			queryParam:         "",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "value is empty",
		},
		{
			name:               "OK",
			queryParam:         "10",
			expectedStatusCode: http.StatusOK,
			expectedBody:       "20",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// ResponseRecorder is an implementation of [http.ResponseWriter] that
			// records its mutations for later inspection in tests.
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/double?value="+tc.queryParam, nil)
			DoubleHandler(w, r)
			require.Equal(t, tc.expectedStatusCode, w.Code)
			require.Equal(t, tc.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

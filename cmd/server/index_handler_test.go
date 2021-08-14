package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aifaniyi/sample/pkg/mock"
)

type indexTestCase struct {
	method       string
	path         string
	expectedCode int
	msg          string
}

func TestIndex(t *testing.T) {
	createRepo = mock.Repository
	svr := NewServer(
		createRepo(),
	)

	testCases := []indexTestCase{
		{http.MethodGet, "/", http.StatusOK, `{"msg":"Hello World!"}`},
		{http.MethodPost, "/", http.StatusMethodNotAllowed, ``},
	}

	for i, testCase := range testCases {
		request, _ := http.NewRequest(testCase.method, testCase.path, nil)
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		svr.router.ServeHTTP(response, request)

		if response.Code != testCase.expectedCode {
			t.Logf("test %d\nExpected: %v\nActual: %v",
				i+1, testCase.expectedCode, response.Code)
		}
	}
}

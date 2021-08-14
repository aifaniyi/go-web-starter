package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aifaniyi/sample/pkg/mock"
)

type authTestCase struct {
	method       string
	path         string
	body         string
	expectedCode int
	msg          string
}

func TestSignup(t *testing.T) {
	createRepo = mock.Repository
	svr := NewServer(
		createRepo(),
	)

	testCases := []authTestCase{
		{http.MethodGet, "/api/v1/signup", `{"email": "email", "password": "password", "confirmPassword": "password"}`, http.StatusMethodNotAllowed, ``},
		{http.MethodPost, "/api/v1/signup", `{"email": "email", "password": "password", "confirmPassword": "password"}`, http.StatusOK, ``},
		{http.MethodPost, "/api/v1/signup", `{"email": "email", "password": "password", "confirmPassword": "pass"}`, http.StatusBadRequest, ``},
	}

	for i, testCase := range testCases {
		request, _ := http.NewRequest(testCase.method, testCase.path,
			bytes.NewReader([]byte(testCase.body)))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		svr.router.ServeHTTP(response, request)

		if response.Code != testCase.expectedCode {
			t.Logf("test %d\nExpected: %v\nActual: %v",
				i+1, testCase.expectedCode, response.Code)
		}
	}
}

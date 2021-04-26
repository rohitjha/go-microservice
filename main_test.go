package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "index route",
			route:         "/",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"data\":null,\"message\":\"Hello, World!\",\"status\":200}",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	app := Setup()

	for _, test := range tests {
		req := httptest.NewRequest(
			"GET",
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := ioutil.ReadAll(res.Body)

		assert.Nilf(t, err, test.description)
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}

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
		method        string
		route         string
		body          string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "non existing route",
			method:        "GET",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
		{
			description:   "get all products",
			method:        "GET",
			route:         "/products",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  `[{"name":"Latte","description":"Frothy milky coffee","price":2.45,"sku":"abc323"},{"name":"Espresso","description":"Short and strong coffee without milk","price":1.99,"sku":"fjd34"}]`,
		},
		{
			description:   "put a product",
			method:        "PUT",
			route:         "/products",
			body:          `{"name":"Pie","description":"pi","price":3.1415,"sku":"123456"}`,
			expectedError: false,
			expectedCode:  201,
			expectedBody:  "",
		},
		{
			description:   "update a product",
			method:        "POST",
			route:         "/products/Pie",
			body:          `{"description":"pie"}`,
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "delete a product",
			method:        "DELETE",
			route:         "/products/Pie",
			body:          "",
			expectedError: false,
			expectedCode:  204,
			expectedBody:  "",
		},
	}

	app := Setup()

	for _, test := range tests {
		req := httptest.NewRequest(
			test.method,
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

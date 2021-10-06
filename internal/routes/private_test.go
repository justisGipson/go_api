package routes

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/CodeliciousProduct/bluebird/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPrivateRoutes(t *testing.T) {
	// load env.test
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}

	dataString := `{"id": "00000000-0000-0000-0000-000000000000"`

	// create acccess token
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		panic(err)
	}
	// define input and output data structure for single test case
	tests := []struct {
		// inputs
		description   string
		route         string
		method        string
		tokenString   string
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		// TODO: define test output
		{
			description:   "delete lesson without JWT or req.body",
			route:         "/api/v1/lesson",
			method:        "DELETE",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "delete lesson without valid credentials",
			route:         "/api/v1/lesson",
			method:        "DELETE",
			tokenString:   "Bearer" + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  401,
		},
		{
			description:   "delete lesson with valid credentials",
			route:         "/api/v1/lesson",
			method:        "DELETE",
			tokenString:   "Bearer" + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  401,
		},
		// TODO: write rest of tests
	}

	app := fiber.New()
	PrivateRoutes(app)

	for _, test := range tests {
		// create new http req with test case routes
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Authorization", test.tokenString)
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1) // -1 disables req latency
		assert.Equalf(t, test.expectedError, err != nil, test.description)
		// errors lead to broken responses
		// this lets us move forward
		if test.expectedError {
			continue
		}
		// verify expectedCode
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
	}
}

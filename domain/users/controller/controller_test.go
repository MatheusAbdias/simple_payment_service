package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MatheusAbdias/simple_payment_service/domain/users"
	domain "github.com/MatheusAbdias/simple_payment_service/domain/users"
	"github.com/MatheusAbdias/simple_payment_service/domain/users/adapters"
	"github.com/stretchr/testify/require"
)

func setup() *Controller {
	service := domain.NewService(adapters.NewTestRepository())
	return NewController(service)

}

func testRequest(
	method string,
	url string,
	body []byte,
) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(method, url, bytes.NewReader(body))
	recorder := httptest.NewRecorder()

	return recorder, request
}
func TestCreateUser(t *testing.T) {
	controller := setup()

	testCases := []struct {
		name               string
		userDTO            *domain.UserDTO
		expectedStatusCode int
	}{
		{
			"Should be can create user",
			&domain.UserDTO{
				FullName: "Jon Doe",
				Email:    "jon@email.com",
				Document: "68507344070",
			},
			http.StatusOK,
		},
		{
			"Should be cant create user with invalid email",
			&domain.UserDTO{
				FullName: "Jon Doe",
				Email:    "jon_email.com",
				Document: "68507344070",
			},
			http.StatusBadRequest,
		},
		{
			"Should be cant create user without email",
			&domain.UserDTO{
				FullName: "Jon Doe",
				Email:    "",
				Document: "68507344070",
			},
			http.StatusBadRequest,
		},
		{
			"Should be cant create user with invalid document",
			&domain.UserDTO{
				FullName: "Jon Doe",
				Email:    "jon@email.com",
				Document: "11111111111",
			},
			http.StatusBadRequest,
		},
		{
			"Should be cant create user without document",
			&domain.UserDTO{
				FullName: "Jon Doe",
				Email:    "jon@email.com",
				Document: "",
			},
			http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var requestBody []byte
			var err error
			if requestBody, err = json.Marshal(tc.userDTO); err != nil {
				t.Fatal(err)
			}

			recorder, request := testRequest("POST", "/users", requestBody)

			controller.Signup(recorder, request)

			if status := recorder.Code; status != tc.expectedStatusCode {
				t.Errorf(
					"handler returned wrong status code: got %v want %v",
					recorder.Code,
					tc.expectedStatusCode,
				)
			}

		})
	}
}

func TestShouldBeCantCreateUserWhenEmailIsAlreadyRegister(t *testing.T) {
	controller := setup()
	ctx := context.Background()

	firstUserDTO := &users.UserDTO{
		FullName: "Jon Dow",
		Email:    "jon@email.com",
		Document: "68507344070",
	}
	if _, err := controller.service.RegisterUser(ctx, firstUserDTO); err != nil {
		t.Fatal(err)
	}

	invalidUser := &users.UserDTO{
		FullName: "Jon Dow",
		Email:    "JON@email.com",
		Document: "57271936050",
	}

	var requestBody []byte
	var err error
	if requestBody, err = json.Marshal(invalidUser); err != nil {
		t.Fatal(err)
	}

	recorder, request := testRequest("POST", "/users", requestBody)

	controller.Signup(recorder, request)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestShouldBeCantCreateUserWhenDocumentIsAlreadyRegister(t *testing.T) {
	controller := setup()
	ctx := context.Background()

	firstUserDTO := &users.UserDTO{
		FullName: "Jon Dow",
		Email:    "jon@email.com",
		Document: "68507344070",
	}
	if _, err := controller.service.RegisterUser(ctx, firstUserDTO); err != nil {
		t.Fatal(err)
	}

	invalidUser := &users.UserDTO{
		FullName: "Mark Dow",
		Email:    "mark@email.com",
		Document: "68507344070",
	}

	var requestBody []byte
	var err error
	if requestBody, err = json.Marshal(invalidUser); err != nil {
		t.Fatal(err)
	}

	recorder, request := testRequest("POST", "/users", requestBody)

	controller.Signup(recorder, request)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

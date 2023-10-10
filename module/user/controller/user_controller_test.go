package controller_test

import (
	"belajar-go-echo/mocks"
	"belajar-go-echo/module/user/controller"
	"belajar-go-echo/module/user/controller/request"
	"belajar-go-echo/module/user/controller/response"
	"belajar-go-echo/module/user/entity"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUsers(t *testing.T) {

	e := echo.New()
	service := new(mocks.UserService)
	returnData := []entity.UserDTO{{ID: 1, Name: "Bryan", Email: "Bryan@gmail.com", Password: "qwerty"}}

	t.Run("Success get users", func(t *testing.T) {
		service.On("GetAllUser").Return(returnData, nil).Once()
		svc := controller.NewUserController(service)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/")
		var responseData response.GetUsersResponse

		if assert.NoError(t, svc.GetAllUserController(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.NoError(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0].Name, responseData.Data[0].Name)
		}
		service.AssertExpectations(t)
	})

	t.Run("Error get users", func(t *testing.T) {
		service.On("GetAllUser").Return(nil, errors.New("An error ocurred")).Once()
		svc := controller.NewUserController(service)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/")
		var responseData response.GetUsersResponse

		if assert.NoError(t, svc.GetAllUserController(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.NoError(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, []entity.UserDTO(nil), responseData.Data)
		}
		service.AssertExpectations(t)
	})

}

func TestCreateUser(t *testing.T) {

	e := echo.New()
	service := new(mocks.UserService)

	var testCases = []struct {
		name         string
		expectedCode int
	}{
		{
			name:         "Error create user invalid request",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Error create user in server error",
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:         "Success create user",
			expectedCode: http.StatusCreated,
		},
	}

	invalidRequest := struct {
		Email    int    `json:"email"`
		Name     string `json:"name"`
		Password bool   `json:"password"`
	}{
		Email:    123,
		Name:     "Bryan",
		Password: true,
	}
	validRequest := request.CreateUserRequest{
		Email:    "Bryan@gmail.com",
		Name:     "Bryan",
		Password: "qwerty",
	}
	jsonBody, _ := json.Marshal(invalidRequest)
	service.On("CreateUser", mock.Anything).Return(errors.New("An error ocurred")).Once()

	for idx, testCase := range testCases {
		if idx == 1 {
			jsonBody, _ = json.Marshal(validRequest)
		}
		if idx == 2 {
			service.On("CreateUser", mock.Anything).Return(nil).Once()
		}

		t.Run(testCase.name, func(t *testing.T) {
			svc := controller.NewUserController(service)

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(jsonBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			echoContext := e.NewContext(req, rec)
			echoContext.SetPath("/")
			var responseData response.CreateUserResponse

			if assert.NoError(t, svc.CreateUserController(echoContext)) {
				responseBody := rec.Body.String()
				err := json.Unmarshal([]byte(responseBody), &responseData)
				if err != nil {
					assert.NoError(t, err, "error")
				}
				assert.Equal(t, testCase.expectedCode, rec.Code)
			}
		})

		t.Cleanup(func() {
			service.AssertExpectations(t)
		})
	}

}

package service_test

import (
	"belajar-go-echo/mocks"
	"belajar-go-echo/module/user/entity"
	"belajar-go-echo/module/user/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertUser(t *testing.T) {

	repository := new(mocks.UserRepository)

	var testCases = []struct {
		name string
		err  error
	}{
		{
			name: "Error create user",
			err:  errors.New("An error ocurred"),
		},
		{
			name: "Success create user",
			err:  nil,
		},
	}
	repository.On("InsertUser", mock.Anything).Return(errors.New("An error ocurred")).Once()

	for idx, testCase := range testCases {
		if idx == 1 {
			repository.On("InsertUser", mock.Anything).Return(nil).Once()
		}

		t.Run(testCase.name, func(t *testing.T) {
			svc := service.NewUserServie(repository)

			err := svc.CreateUser(entity.UserDTO{})
			if idx == 0 {
				assert.Error(t, err)
				assert.Equal(t, testCase.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.err, err)
			}
		})

		t.Cleanup(func() {
			repository.AssertExpectations(t)
		})
	}

}

func TestGetAllUser(t *testing.T) {

	repository := new(mocks.UserRepository)
	returnData := []entity.UserDTO{{ID: 1, Name: "Bryan", Email: "Bryan@gmail.com", Password: "qwerty"}}

	var testCases = []struct {
		name string
		err  error
	}{
		{
			name: "Error create user",
			err:  errors.New("An error ocurred"),
		},
		{
			name: "Success create user",
			err:  nil,
		},
	}
	repository.On("GetAllUser", mock.Anything).Return(nil, errors.New("An error ocurred")).Once()

	for idx, testCase := range testCases {
		if idx == 1 {
			repository.On("GetAllUser", mock.Anything).Return(returnData, nil).Once()
		}

		t.Run(testCase.name, func(t *testing.T) {
			svc := service.NewUserServie(repository)

			data, err := svc.GetAllUser()
			if idx == 0 {
				assert.Error(t, err)
				assert.Equal(t, testCase.err, err)
				assert.Equal(t, []entity.UserDTO([]entity.UserDTO(nil)), data)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.err, err)
				assert.Equal(t, returnData, data)
			}
		})

		t.Cleanup(func() {
			repository.AssertExpectations(t)
		})
	}

}

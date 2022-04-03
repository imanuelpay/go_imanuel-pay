package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mvc/config"
	"mvc/controllers"
	"mvc/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type UserResponse struct {
	Message string              `json:"message" form:"message"`
	Data    models.UserResponse `json:"data" form:"data"`
}

func GenerateDataUsers() error {
	user := []models.User{
		{
			Name:     "Test1",
			Email:    "test1@test.com",
			Password: "test1",
		},
		{
			Name:     "Test2",
			Email:    "test2@test.com",
			Password: "test2",
		},
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
func TestGetAllUsersValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all user",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   2,
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.GetUsersController(c)) {
			assert.Equal(t, testCase.expectCode, response.Code)
			body := response.Body.String()

			type UserResponse struct {
				Message string        `json:"message" form:"message"`
				Data    []models.User `json:"data" form:"data"`
			}

			var user UserResponse
			err := json.Unmarshal([]byte(body), &user)

			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCase.sizeData, len(user.Data))
		}
	}
}

func TestGetUserValid(t *testing.T) {
	var testCases = []struct {
		name       string
		nameUser   string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "get user by id 2",
			nameUser:   "Test2",
			path:       "/users",
			paramName:  "id",
			paramValue: "2",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.GetUserController(c)
		if cek != nil {
			assert.Equal(t, testCase.expectCode, response.Code)
			t.Errorf("should not get error, get error: %s", cek)
			body := response.Body.String()

			var user = struct {
				message string
				data    models.User
			}{}

			err := json.Unmarshal([]byte(body), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			actual := user.data.Name
			expected := testCase.nameUser
			if actual != expected {
				t.Errorf("should return %s, get: %s", expected, actual)
			}
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestGetUserInvalidNoRecord(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "get user by id 3",
			path:       "/users",
			paramName:  "id",
			paramValue: "3",
			expectCode: http.StatusInternalServerError,
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.GetUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestGetUserInvalidParamId(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "get user by code 2",
			path:       "/users",
			paramName:  "code",
			paramValue: "2",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.GetUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestCreateUserValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusOK,
			data: map[string]interface{}{
				"name":     "Test3",
				"email":    "test3@test.com",
				"password": "test3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)

		cek := controllers.CreateUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}

		body := response.Body.String()

		type UserResponse struct {
			Message string      `json:"message" form:"message"`
			Data    models.User `json:"data" form:"data"`
		}

		var user UserResponse

		if err := json.Unmarshal([]byte(body), &user); err != nil {
			assert.Error(t, err, "error")
		}

		expectedName := testCase.data["name"]
		if user.Data.Name != expectedName {
			t.Errorf("user name should be %s, get: %s", expectedName, user.Data.Name)
		}
		expectedEmail := testCase.data["email"]
		if user.Data.Email != expectedEmail {
			t.Errorf("user email should be %s, get: %s", expectedEmail, user.Data.Email)
		}
		expectedPassword := testCase.data["password"]
		if user.Data.Password != expectedPassword {
			t.Errorf("user pasword should be %s, get: %s", expectedPassword, user.Data.Password)
		}
	}
}

func TestCreateUserInvalidBody(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			data:       map[string]interface{}{},
		},
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"name":     "Test3",
				"password": "test3",
			},
		},
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"name":  "Test3",
				"email": "test3@test.com",
			},
		},
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"password": "test3",
				"email":    "test3@test.com",
			},
		},
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"name": "Test3",
			},
		},
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"password": "test3",
			},
		},
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"email": "test3@test.com",
			},
		},
		{
			name:       "create user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"name":     "Test3",
				"email":    1,
				"password": "test3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)

		cek := controllers.CreateUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestUpdateUserValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "update user data",
			path:       "/users",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"name": "Test3",
			},
		},
		{
			name:       "update user data",
			path:       "/users",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"email": "test3@test.com",
			},
		},
		{
			name:       "update user data",
			path:       "/users",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"password": "test3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.UpdateUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestUpdateUserValidFullAttribute(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "update user data",
			path:       "/users",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"name":     "Test3",
				"email":    "test3@test.com",
				"password": "test3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.UpdateUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}

		body := response.Body.String()

		type UserResponse struct {
			Message string      `json:"message" form:"message"`
			Data    models.User `json:"data" form:"data"`
		}

		var user UserResponse

		if err := json.Unmarshal([]byte(body), &user); err != nil {
			assert.Error(t, err, "error")
		}

		expectedName := testCase.data["name"]
		if user.Data.Name != expectedName {
			t.Errorf("user name should be %s, get: %s", expectedName, user.Data.Name)
		}
		expectedEmail := testCase.data["email"]
		if user.Data.Email != expectedEmail {
			t.Errorf("user email should be %s, get: %s", expectedEmail, user.Data.Email)
		}
		expectedPassword := testCase.data["password"]
		if user.Data.Password != expectedPassword {
			t.Errorf("user pasword should be %s, get: %s", expectedPassword, user.Data.Password)
		}
	}
}

func TestUpdateUserInvalidParamID(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "update user data",
			path:       "/users",
			expectCode: http.StatusInternalServerError,
			paramName:  "id",
			paramValue: "3",
			data: map[string]interface{}{
				"name": "Test3",
			},
		},
		{
			name:       "update user data",
			path:       "/users",
			expectCode: http.StatusInternalServerError,
			paramName:  "id",
			paramValue: "4",
			data: map[string]interface{}{
				"email": "test3@test.com",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.UpdateUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestUpdateUserInvalidBody(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "update user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			paramName:  "id",
			paramValue: "1",
			data: map[string]interface{}{
				"name": 12344,
			},
		},
		{
			name:       "update user data",
			path:       "/users",
			expectCode: http.StatusUnprocessableEntity,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"email": 1010101,
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.UpdateUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestDeleteUserValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "delete user data",
			path:       "/users",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "1",
		},
		{
			name:       "delete user data",
			path:       "/users",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	for _, testCase := range testCases {
		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.DeleteUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestDeleteUserInvalidParamID(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "delete user data",
			path:       "/users",
			expectCode: http.StatusInternalServerError,
			paramName:  "id",
			paramValue: "3",
		},
		{
			name:       "delete user data",
			path:       "/users",
			expectCode: http.StatusInternalServerError,
			paramName:  "id",
			paramValue: "4",
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	for _, testCase := range testCases {
		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.DeleteUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestDeleteUserInvalidParamString(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "delete user data",
			path:       "/users",
			expectCode: http.StatusBadRequest,
			paramName:  "id",
			paramValue: "tiga",
		},
		{
			name:       "delete user data",
			path:       "/users",
			expectCode: http.StatusBadRequest,
			paramName:  "id",
			paramValue: "empat",
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	loginInfo, err := json.Marshal(map[string]interface{}{
		"email":    "test1@test.com",
		"password": "test1",
	})

	if err != nil {
		t.Errorf("marshalling new person failed")
	}

	loginRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	loginContext := e.NewContext(loginRequest, loginResponse)
	loginContext.SetPath("/login")

	// send request
	if err := controllers.LoginUserController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginResponse.Code != 200 {
		t.Errorf("should return 200, get: %d", loginResponse.Code)
	}

	// compare response
	user := UserResponse{}
	if err := json.Unmarshal(loginResponse.Body.Bytes(), &user); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if user.Data.Token == "" {
		t.Errorf("token expected")
	}

	token := user.Data.Token

	for _, testCase := range testCases {
		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.DeleteUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestLoginUserValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "create user data",
			path:       "/login",
			expectCode: http.StatusOK,
			data: map[string]interface{}{
				"email":    "test2@test.com",
				"password": "test2",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)

		cek := controllers.LoginUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}

		body := response.Body.String()

		user := UserResponse{}
		if err := json.Unmarshal([]byte(body), &user); err != nil {
			assert.Error(t, err, "error")
		}

		if user.Data.Token == "" {
			t.Errorf("token expected")
		}

		expectedEmail := testCase.data["email"]
		if user.Data.Email != expectedEmail {
			t.Errorf("user email should be %s, get: %s", expectedEmail, user.Data.Email)
		}
	}
}

func TestLoginUserInvalidData(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "create user data",
			path:       "/login",
			expectCode: http.StatusInternalServerError,
			data: map[string]interface{}{
				"email":    "test3@test.com",
				"password": "test3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)

		cek := controllers.LoginUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestLoginUserInvalidBody(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "create user data",
			path:       "/login",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"email":    "test3@test.com",
				"password": 123,
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataUsers()

	for _, testCase := range testCases {
		newUser, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new user failed")
		}

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newUser))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)

		cek := controllers.LoginUserController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

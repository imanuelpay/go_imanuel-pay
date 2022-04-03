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

	"github.com/stretchr/testify/assert"
)

func GenerateDataBooks() error {
	book := []models.Book{
		{
			Name:      "Buku1",
			Author:    "author1",
			Publisher: "publisher1",
		},
		{
			Name:      "Buku2",
			Author:    "author2",
			Publisher: "publisher2",
		},
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return err
	}

	user := []models.User{
		{
			Name:     "Test1",
			Email:    "test1@test.com",
			Password: "test1",
		},
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
func TestGetAllBooksValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all book",
			path:       "/books",
			expectCode: http.StatusOK,
			sizeData:   2,
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.GetBooksController(c)) {
			assert.Equal(t, testCase.expectCode, response.Code)
			body := response.Body.String()

			type BookResponse struct {
				Message string        `json:"message" form:"message"`
				Data    []models.Book `json:"data" form:"data"`
			}

			var book BookResponse
			err := json.Unmarshal([]byte(body), &book)

			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCase.sizeData, len(book.Data))
		}
	}
}

func TestGetBookValid(t *testing.T) {
	var testCases = []struct {
		name       string
		nameBook   string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "get book by id 2",
			nameBook:   "Buku2",
			path:       "/books",
			paramName:  "id",
			paramValue: "2",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.GetBookController(c)
		if cek != nil {
			assert.Equal(t, testCase.expectCode, response.Code)
			t.Errorf("should not get error, get error: %s", cek)
			body := response.Body.String()

			var book = struct {
				message string
				data    models.Book
			}{}

			err := json.Unmarshal([]byte(body), &book)
			if err != nil {
				assert.Error(t, err, "error")
			}

			actual := book.data.Name
			expected := testCase.nameBook
			if actual != expected {
				t.Errorf("should return %s, get: %s", expected, actual)
			}
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestGetBookInvalidNoRecord(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "get book by id 3",
			path:       "/books",
			paramName:  "id",
			paramValue: "3",
			expectCode: http.StatusInternalServerError,
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.GetBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestGetBookInvalidParamId(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "get book by code 2",
			path:       "/books",
			paramName:  "code",
			paramValue: "2",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.GetBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestCreateBookValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusOK,
			data: map[string]interface{}{
				"name":      "Book3",
				"author":    "author3",
				"publisher": "publisher3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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
		newBook, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new book failed")
		}

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newBook))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)

		cek := controllers.CreateBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}

		body := response.Body.String()

		type BookResponse struct {
			Message string      `json:"message" form:"message"`
			Data    models.Book `json:"data" form:"data"`
		}

		var book BookResponse

		if err := json.Unmarshal([]byte(body), &book); err != nil {
			assert.Error(t, err, "error")
		}

		expectedName := testCase.data["name"]
		if book.Data.Name != expectedName {
			t.Errorf("book name should be %s, get: %s", expectedName, book.Data.Name)
		}
		expectedAuthor := testCase.data["author"]
		if book.Data.Author != expectedAuthor {
			t.Errorf("book author should be %s, get: %s", expectedAuthor, book.Data.Author)
		}
		expectedPublisher := testCase.data["publisher"]
		if book.Data.Publisher != expectedPublisher {
			t.Errorf("book pasword should be %s, get: %s", expectedPublisher, book.Data.Publisher)
		}
	}
}

func TestCreateBookInvalidBody(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			data:       map[string]interface{}{},
		},
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"name":      "Book3",
				"publisher": "publisher3",
			},
		},
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"name":   "Book3",
				"author": "author3",
			},
		},
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"publisher": "publisher3",
				"author":    "author3",
			},
		},
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"name": "Book3",
			},
		},
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"publisher": "publisher3",
			},
		},
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"author": "author3",
			},
		},
		{
			name:       "create book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			data: map[string]interface{}{
				"name":      "Book3",
				"author":    1,
				"publisher": "publisher3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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
		newBook, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new book failed")
		}

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newBook))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)

		cek := controllers.CreateBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestUpdateBookValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "update book data",
			path:       "/books",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"name": "Book3",
			},
		},
		{
			name:       "update book data",
			path:       "/books",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"author": "author3",
			},
		},
		{
			name:       "update book data",
			path:       "/books",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"publisher": "publisher3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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
		newBook, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new book failed")
		}

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newBook))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.UpdateBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestUpdateBookValidFullAttribute(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "update book data",
			path:       "/books",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"name":      "Book3",
				"author":    "author3",
				"publisher": "publisher3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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
		newBook, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new book failed")
		}

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newBook))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.UpdateBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}

		body := response.Body.String()

		type BookResponse struct {
			Message string      `json:"message" form:"message"`
			Data    models.Book `json:"data" form:"data"`
		}

		var book BookResponse

		if err := json.Unmarshal([]byte(body), &book); err != nil {
			assert.Error(t, err, "error")
		}

		expectedName := testCase.data["name"]
		if book.Data.Name != expectedName {
			t.Errorf("book name should be %s, get: %s", expectedName, book.Data.Name)
		}
		expectedAuthor := testCase.data["author"]
		if book.Data.Author != expectedAuthor {
			t.Errorf("book author should be %s, get: %s", expectedAuthor, book.Data.Author)
		}
		expectedPublisher := testCase.data["publisher"]
		if book.Data.Publisher != expectedPublisher {
			t.Errorf("book pasword should be %s, get: %s", expectedPublisher, book.Data.Publisher)
		}
	}
}

func TestUpdateBookInvalidParamID(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "update book data",
			path:       "/books",
			expectCode: http.StatusInternalServerError,
			paramName:  "id",
			paramValue: "3",
			data: map[string]interface{}{
				"name": "Book3",
			},
		},
		{
			name:       "update book data",
			path:       "/books",
			expectCode: http.StatusInternalServerError,
			paramName:  "id",
			paramValue: "4",
			data: map[string]interface{}{
				"author": "author3",
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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
		newBook, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new book failed")
		}

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newBook))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.UpdateBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestUpdateBookInvalidBody(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
		data       map[string]interface{}
	}{
		{
			name:       "update book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			paramName:  "id",
			paramValue: "1",
			data: map[string]interface{}{
				"name": 12344,
			},
		},
		{
			name:       "update book data",
			path:       "/books",
			expectCode: http.StatusUnprocessableEntity,
			paramName:  "id",
			paramValue: "2",
			data: map[string]interface{}{
				"author": 1010101,
			},
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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
		newBook, err := json.Marshal(testCase.data)

		if err != nil {
			t.Errorf("marshalling new book failed")
		}

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newBook))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		c.SetPath(testCase.path)
		c.SetParamNames(testCase.paramName)
		c.SetParamValues(testCase.paramValue)

		cek := controllers.UpdateBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestDeleteBookValid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "delete book data",
			path:       "/books",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "1",
		},
		{
			name:       "delete book data",
			path:       "/books",
			expectCode: http.StatusOK,
			paramName:  "id",
			paramValue: "2",
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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

		cek := controllers.DeleteBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestDeleteBookInvalidParamID(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "delete book data",
			path:       "/books",
			expectCode: http.StatusInternalServerError,
			paramName:  "id",
			paramValue: "3",
		},
		{
			name:       "delete book data",
			path:       "/books",
			expectCode: http.StatusInternalServerError,
			paramName:  "id",
			paramValue: "4",
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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

		cek := controllers.DeleteBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

func TestDeleteBookInvalidParamString(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		paramName  string
		paramValue string
		expectCode int
	}{
		{
			name:       "delete book data",
			path:       "/books",
			expectCode: http.StatusBadRequest,
			paramName:  "id",
			paramValue: "tiga",
		},
		{
			name:       "delete book data",
			path:       "/books",
			expectCode: http.StatusBadRequest,
			paramName:  "id",
			paramValue: "empat",
		},
	}

	e := InitEchoTestAPI()
	GenerateDataBooks()

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

		cek := controllers.DeleteBookController(c)

		if cek != nil {
			t.Errorf("should not get error, get error: %s", cek)
			return
		}

		if response.Code != testCase.expectCode {
			t.Errorf("should return %d, get: %d", testCase.expectCode, response.Code)
		}
	}
}

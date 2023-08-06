package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

func Test_LoginSuccess(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	auth := entity.UserAuth{
		Login:    "user",
		Password: "password",
	}
	body, _ := json.Marshal(auth)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/login", bytes.NewReader(body))
	mockProcessor.On("LoginUser", mock.Anything, auth).Return(nil)
	h.Login(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	mockProcessor.AssertCalled(t, "LoginUser", mock.Anything, auth)

	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	bodyError := []byte("not a json")
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/login", bytes.NewReader(bodyError))
	h.Login(c)
	assert.Equal(t, http.StatusInternalServerError, r.Code)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	auth := entity.UserAuth{
		Login:    "user",
		Password: "password",
	}

	body, _ := json.Marshal(auth)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/login", bytes.NewReader(body))

	mockProcessor.On("LoginUser", mock.Anything, auth).Return(entity.ErrInvalidCredentials)
	h.Login(ctx)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockProcessor.AssertCalled(t, "LoginUser", mock.Anything, auth)
}

func TestRegister_Success(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	auth := entity.UserAuth{
		Login:    "user",
		Password: "password",
	}
	body, _ := json.Marshal(auth)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewReader(body))

	mockProcessor.On("RegisterUser", mock.Anything, auth).Return(nil)
	h.Register(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	mockProcessor.AssertCalled(t, "RegisterUser", mock.Anything, auth)
}

func TestRegister_LoginAlreadyInUse(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	auth := entity.UserAuth{
		Login:    "user",
		Password: "password",
	}
	body, _ := json.Marshal(auth)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewReader(body))
	mockProcessor.On("RegisterUser", mock.Anything, auth).Return(entity.ErrLoginAlreadyInUse)
	h.Register(ctx)
	assert.Equal(t, http.StatusConflict, w.Code)
	mockProcessor.AssertCalled(t, "RegisterUser", mock.Anything, auth)

	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	bodyError := []byte("not a json")
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewReader(bodyError))
	h.Register(c)
	assert.Equal(t, http.StatusInternalServerError, r.Code)
}

func TestRegister_InternalServerError(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	auth := entity.UserAuth{
		Login:    "user",
		Password: "password",
	}

	body, _ := json.Marshal(auth)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewReader(body))
	mockProcessor.On("RegisterUser", mock.Anything, auth).Return(errors.New("some error"))
	h.Register(ctx)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockProcessor.AssertCalled(t, "RegisterUser", mock.Anything, auth)
}

func TestList_Success(t *testing.T) {
	listTests := []struct {
		name         string
		data         string
		user         string
		id           uint32
		returnVal    interface{}
		serviceName  string
		expectedJson string
	}{
		{
			name: "List credentials: Succes",
			data: "credentials",
			user: "user1",
			id:   1,
			returnVal: []entity.CutCredentials{{
				Id:          1,
				ServiceName: "Credential 1",
			}},
			expectedJson: `[{"id":1, "meta":"", "service-name":"Credential 1"}]`,
		},
		{
			name: "List binary: Succes",
			data: "binary",
			user: "user1",
			id:   1,
			returnVal: []entity.CutBinary{{
				Id:    1,
				Title: "Title",
			}},
			expectedJson: `[{"id":1, "title":"Title","meta":""}]`,
		},
		{
			name: "List text: Succes",
			data: "text",
			user: "user1",
			id:   1,
			returnVal: []entity.CutText{{
				Id:    1,
				Title: "Title",
			}},
			expectedJson: `[{"id":1, "meta":"", "title":"Title"}]`,
		},
		{
			name: "List card: Succes",
			data: "card",
			user: "user1",
			id:   1,
			returnVal: []entity.CutCard{{
				Id:         1,
				CardNumber: "1234",
			}},
			expectedJson: `[{"id":1, "meta":"", "card-number":"1234"}]`,
		},
	}
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}

	for _, tt := range listTests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Params = []gin.Param{
				{Key: "data", Value: tt.data},
				{Key: "user", Value: tt.user},
			}

			mockProcessor.On("ListUserData", mock.Anything, tt.data, tt.user).
				Return(tt.returnVal, nil)

			h.List(ctx)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, tt.expectedJson, w.Body.String())

			mockProcessor.AssertCalled(t, "ListUserData", mock.Anything, tt.data, tt.user)
		})
	}

}

func TestList_InternalServerError(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = []gin.Param{
		{Key: "data", Value: "credentials"},
		{Key: "user", Value: "user1"},
	}
	mockProcessor.On("ListUserData", mock.Anything, "credentials", "user1").Return(nil, errors.New("some error"))
	h.List(ctx)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGet_Success(t *testing.T) {
	getTests := []struct {
		name         string
		data         string
		user         string
		id           string
		returnVal    interface{}
		expectedJson string
	}{
		{
			name:         "Get credentials: Succes",
			data:         "credentials",
			user:         "user1",
			id:           "1",
			returnVal:    entity.UserCredentials{},
			expectedJson: `{"Meta":"", "ServiceLogin":"", "ServiceName":"", "ServicePassword":""}`,
		},
		{
			name:         "Get binary: Succes",
			data:         "binary",
			user:         "user1",
			id:           "1",
			returnVal:    entity.UserBinary{},
			expectedJson: `{"B64Content":"", "Meta":"", "Title":""}`,
		},
		{
			name:         "Get text: Succes",
			data:         "text",
			user:         "user1",
			id:           "1",
			returnVal:    entity.UserText{},
			expectedJson: `{"Body":"", "Meta":"", "Title":""}`,
		},
		{
			name:         "Get card: Succes",
			data:         "card",
			user:         "user1",
			id:           "1",
			returnVal:    entity.UserCard{},
			expectedJson: `{"CardHolder":"", "CardNumber":"", "CardValidityPeriod":"", "CvcCode":"", "Meta":""}`,
		},
	}
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}

	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Params = []gin.Param{
				{Key: "data", Value: tt.data},
				{Key: "id", Value: tt.id},
				{Key: "user", Value: tt.user},
			}
			mockProcessor.On("GetUserData", mock.Anything, tt.data, tt.id, tt.user).
				Return(tt.returnVal, nil)
			h.Get(ctx)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, tt.expectedJson, w.Body.String())
			mockProcessor.AssertCalled(t, "GetUserData", mock.Anything, tt.data, tt.id, tt.user)
		})
	}

}

func TestCredentialsCreate(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	mockProcessor.On("CredentialsCreation", mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Once()
	body := `{"username":"testuser","password":"testpassword"}`
	req, _ := http.NewRequest("POST", "/credentials/create/testuser", strings.NewReader(body))
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	h.CredentialsCreate(ctx)
	mockProcessor.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)

	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	bodyError := []byte("not a json")
	c.Request, _ = http.NewRequest(http.MethodPost, "/credentials/create/testuser", bytes.NewReader(bodyError))
	h.CredentialsCreate(c)
	assert.Equal(t, http.StatusInternalServerError, r.Code)
}

func TestDelete(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	handler := &handler{
		processor: mockProcessor,
	}
	mockProcessor.On("DeleteUserData", mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Once()
	req, _ := http.NewRequest("DELETE", "/delete/data/123", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	handler.Delete(ctx)
	mockProcessor.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTextCreate(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	mockProcessor.On("TextCreation", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).
		Return(nil).Once()
	reqBody := []byte(`{"text":"This is a test text"}`)
	req, _ := http.NewRequest("POST", "/v1/secrets/text/user", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	h.TextCreate(ctx)
	mockProcessor.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)

	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	bodyError := []byte("not a json")
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/secrets/text/user", bytes.NewReader(bodyError))
	h.TextCreate(c)
	assert.Equal(t, http.StatusInternalServerError, r.Code)
}

func TestBinaryCreate(t *testing.T) {
	mockProcessor := newMockSecretsProcessor(t)
	h := &handler{
		processor: mockProcessor,
	}
	mockProcessor.On("BinaryCreation", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).
		Return(nil).Once()
	reqBody := []byte(`{"binary":"This is a test binary"}`)
	req, _ := http.NewRequest("POST", "/v1/secrets/binary/user", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	h.BinaryCreate(ctx)
	mockProcessor.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)

	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	bodyError := []byte("not a json")
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/secrets/binary/user", bytes.NewReader(bodyError))
	h.BinaryCreate(c)
	assert.Equal(t, http.StatusInternalServerError, r.Code)
}

func TestTextEdit(t *testing.T) {
	textEditTests := []struct {
		name         string
		body         string
		url          string
		returnErr    error
		expectedCode int
	}{
		{
			name:         "Edit text: Succes",
			body:         `{"text": "updated text"}`,
			url:          "/v1/secrets/text/1",
			returnErr:    nil,
			expectedCode: 200,
		},
		{
			name:         "Edit text: Fail",
			body:         ``,
			url:          "/v1/secrets/text/1",
			returnErr:    fmt.Errorf("some err"),
			expectedCode: 500,
		},
	}
	mockProcessor := newMockSecretsProcessor(t)
	handler := handler{processor: mockProcessor}
	for _, tt := range textEditTests {
		t.Run(tt.name, func(t *testing.T) {
			mockProcessor.On("TextEdit", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.returnErr)
			req, _ := http.NewRequest("PUT", tt.url, strings.NewReader(tt.body))
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			handler.TextEdit(ctx)
			mockProcessor.AssertExpectations(t)
			assert.Equal(t, tt.expectedCode, ctx.Writer.Status())
		})
	}
}

func TestCardsEdit(t *testing.T) {
	cardsEditTests := []struct {
		name         string
		body         string
		url          string
		returnErr    error
		expectedCode int
	}{
		{
			name:         "Edit cards: Succes",
			body:         `{"test": "test"}`,
			url:          "/v1/secrets/cards/:id",
			returnErr:    nil,
			expectedCode: 200,
		},
		{
			name:         "Edit cards: Fail",
			body:         ``,
			url:          "/v1/secrets/card/1",
			returnErr:    fmt.Errorf("some err"),
			expectedCode: 500,
		},
	}
	mockProcessor := newMockSecretsProcessor(t)
	handler := handler{processor: mockProcessor}
	for _, tt := range cardsEditTests {
		t.Run(tt.name, func(t *testing.T) {
			mockProcessor.On("CardEdit", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.returnErr)
			req, _ := http.NewRequest("PUT", tt.url, strings.NewReader(tt.body))
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			handler.CardEdit(ctx)
			mockProcessor.AssertExpectations(t)
			assert.Equal(t, tt.expectedCode, ctx.Writer.Status())
		})
	}

}

func TestBinaryEdit(t *testing.T) {
	binaryEditTests := []struct {
		name         string
		body         string
		url          string
		returnErr    error
		expectedCode int
	}{
		{
			name:         "Edit binary: Succes",
			body:         `{"test": "test"}`,
			url:          "/v1/secrets/binary/:id",
			returnErr:    nil,
			expectedCode: 200,
		},
		{
			name:         "Edit binary: Fail",
			body:         ``,
			url:          "/v1/secrets/binary/1",
			returnErr:    fmt.Errorf("some err"),
			expectedCode: 500,
		},
	}
	mockProcessor := newMockSecretsProcessor(t)
	handler := handler{processor: mockProcessor}
	for _, tt := range binaryEditTests {
		t.Run(tt.name, func(t *testing.T) {
			mockProcessor.On("BinaryEdit", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.returnErr)
			req, _ := http.NewRequest("PUT", tt.url, strings.NewReader(tt.body))
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			handler.BinaryEdit(ctx)
			mockProcessor.AssertExpectations(t)
			assert.Equal(t, tt.expectedCode, ctx.Writer.Status())
		})
	}
}

func TestCredentialsEdit(t *testing.T) {
	binaryEditTests := []struct {
		name         string
		body         string
		url          string
		returnErr    error
		expectedCode int
	}{
		{
			name:         "Edit binary: Succes",
			body:         `{"test": "test"}`,
			url:          "/v1/secrets/credentials/:id",
			returnErr:    nil,
			expectedCode: 200,
		},
		{
			name:         "Edit binary: Fail",
			body:         ``,
			url:          "/v1/secrets/credentials/1",
			returnErr:    fmt.Errorf("some err"),
			expectedCode: 500,
		},
	}
	mockProcessor := newMockSecretsProcessor(t)
	handler := handler{processor: mockProcessor}
	for _, tt := range binaryEditTests {
		t.Run(tt.name, func(t *testing.T) {
			mockProcessor.On("CredentialsEdit", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.returnErr)
			req, _ := http.NewRequest("PUT", tt.url, strings.NewReader(tt.body))
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			handler.CredentialsEdit(ctx)
			mockProcessor.AssertExpectations(t)
			assert.Equal(t, tt.expectedCode, ctx.Writer.Status())
		})
	}
}

func TestCardCreate(t *testing.T) {
	binaryEditTests := []struct {
		name         string
		body         string
		url          string
		returnErr    error
		expectedCode int
	}{
		{
			name:         "Edit binary: Succes",
			body:         `{"test": "test"}`,
			url:          "/v1/secrets/card/user",
			returnErr:    nil,
			expectedCode: 200,
		},
		{
			name:         "Edit binary: Fail",
			body:         `not a json`,
			url:          "/v1/secrets/card/user",
			returnErr:    fmt.Errorf("some err"),
			expectedCode: 500,
		},
	}
	mockProcessor := newMockSecretsProcessor(t)
	h := handler{processor: mockProcessor}
	for _, tt := range binaryEditTests {
		t.Run(tt.name, func(t *testing.T) {
			mockProcessor.On("CardCreation", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			req, _ := http.NewRequest("POST", tt.url, strings.NewReader(tt.body))
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			h.CardCreate(ctx)
			mockProcessor.AssertExpectations(t)
			assert.Equal(t, tt.expectedCode, ctx.Writer.Status())
		})
	}

	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)
	bodyError := []byte("not a json")
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/secrets/card/user", bytes.NewReader(bodyError))
	h.CardCreate(c)
	assert.Equal(t, http.StatusInternalServerError, r.Code)
}

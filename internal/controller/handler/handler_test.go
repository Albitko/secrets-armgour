package handler

import (
	"bytes"
	"encoding/json"
	"errors"
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
	handler := &handler{
		processor: mockProcessor,
	}
	mockProcessor.On("CredentialsCreation", mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Once()
	body := `{"username":"testuser","password":"testpassword"}`
	req, _ := http.NewRequest("POST", "/credentials/create/testuser", strings.NewReader(body))
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	handler.CredentialsCreate(ctx)
	mockProcessor.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
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
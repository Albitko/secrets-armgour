package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

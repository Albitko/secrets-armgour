package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

//go:generate mockery --name secretsProcessor
type secretsProcessor interface {
	CardCreation(ctx context.Context, card entity.UserCard, user string) error
	BinaryCreation(ctx context.Context, binary entity.UserBinary, user string) error
	TextCreation(ctx context.Context, text entity.UserText, user string) error
	CredentialsCreation(ctx context.Context, text entity.UserCredentials, user string) error

	ListUserData(ctx context.Context, data, user string) (interface{}, error)
	GetUserData(ctx context.Context, data, id, user string) (interface{}, error)
	DeleteUserData(ctx context.Context, data, id string) error

	CardEdit(ctx context.Context, index string, card entity.UserCard) error
	BinaryEdit(ctx context.Context, index string, binary entity.UserBinary) error
	TextEdit(ctx context.Context, index string, text entity.UserText) error
	CredentialsEdit(ctx context.Context, index string, text entity.UserCredentials) error

	RegisterUser(ctx context.Context, auth entity.UserAuth) error
	LoginUser(ctx context.Context, auth entity.UserAuth) error
}

type handler struct {
	processor secretsProcessor
}

// Login - handler for user login
func (h *handler) Login(ctx *gin.Context) {
	var auth entity.UserAuth
	if err := json.NewDecoder(ctx.Request.Body).Decode(&auth); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.LoginUser(ctx, auth)
	if errors.Is(err, entity.ErrInvalidCredentials) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// Register - handler for user register
func (h *handler) Register(ctx *gin.Context) {
	var auth entity.UserAuth
	if err := json.NewDecoder(ctx.Request.Body).Decode(&auth); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.RegisterUser(ctx, auth)
	if errors.Is(err, entity.ErrLoginAlreadyInUse) {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// List - handler for user login
func (h *handler) List(ctx *gin.Context) {
	data := ctx.Param("data")
	user := ctx.Param("user")
	res, err := h.processor.ListUserData(ctx, data, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	switch data {
	case entity.Credentials:
		res = res.([]entity.CutCredentials)
	case entity.Binary:
		res = res.([]entity.CutBinary)
	case entity.Text:
		res = res.([]entity.CutText)
	case entity.Card:
		res = res.([]entity.CutCard)
	}
	ctx.JSON(http.StatusOK, res)
}

// Get - handler for getting user secrets
func (h *handler) Get(ctx *gin.Context) {
	data := ctx.Param("data")
	id := ctx.Param("id")
	user := ctx.Param("user")

	res, err := h.processor.GetUserData(ctx, data, id, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	switch data {
	case entity.Credentials:
		res = res.(entity.UserCredentials)
	case entity.Binary:
		res = res.(entity.UserBinary)
	case entity.Text:
		res = res.(entity.UserText)
	case entity.Card:
		res = res.(entity.UserCard)
	}
	ctx.JSON(http.StatusOK, res)
}

// CredentialsCreate - handler for saving user credentials for services
func (h *handler) CredentialsCreate(ctx *gin.Context) {
	user := ctx.Param("user")

	var credentials entity.UserCredentials
	if err := json.NewDecoder(ctx.Request.Body).Decode(&credentials); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CredentialsCreation(ctx, credentials, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// Delete - handler for saving user secrets from service
func (h *handler) Delete(ctx *gin.Context) {
	data := ctx.Param("data")
	id := ctx.Param("id")
	err := h.processor.DeleteUserData(ctx, data, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// TextCreate - handler for saving user text secrets
func (h *handler) TextCreate(ctx *gin.Context) {
	user := ctx.Param("user")

	var text entity.UserText
	if err := json.NewDecoder(ctx.Request.Body).Decode(&text); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.TextCreation(ctx, text, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// BinaryCreate - handler for saving user binary secrets
func (h *handler) BinaryCreate(ctx *gin.Context) {
	user := ctx.Param("user")

	var binary entity.UserBinary
	if err := json.NewDecoder(ctx.Request.Body).Decode(&binary); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.BinaryCreation(ctx, binary, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// CardCreate - handler for saving user card secrets
func (h *handler) CardCreate(ctx *gin.Context) {
	user := ctx.Param("user")

	var card entity.UserCard
	if err := json.NewDecoder(ctx.Request.Body).Decode(&card); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CardCreation(ctx, card, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// CredentialsEdit - handler for edit user credentials
func (h *handler) CredentialsEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var credentials entity.UserCredentials
	if err := json.NewDecoder(ctx.Request.Body).Decode(&credentials); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CredentialsEdit(ctx, id, credentials)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// TextEdit - handler for edit user text secrets
func (h *handler) TextEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var text entity.UserText
	if err := json.NewDecoder(ctx.Request.Body).Decode(&text); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.TextEdit(ctx, id, text)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}

// BinaryEdit - handler for edit user binary secrets
func (h *handler) BinaryEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var binary entity.UserBinary
	if err := json.NewDecoder(ctx.Request.Body).Decode(&binary); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.BinaryEdit(ctx, id, binary)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}

// CardEdit - handler for edit user card secrets
func (h *handler) CardEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var card entity.UserCard
	if err := json.NewDecoder(ctx.Request.Body).Decode(&card); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CardEdit(ctx, id, card)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}

// New - create handlers instance
func New(processor secretsProcessor) *handler {
	return &handler{
		processor: processor,
	}
}

package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type secretsProcessor interface {
	CardCreation(card entity.UserCard, user string) error
	BinaryCreation(binary entity.UserBinary, user string) error
	TextCreation(text entity.UserText, user string) error
	CredentialsCreation(text entity.UserCredentials, user string) error

	ListUserData(data, user string) (interface{}, error)
	GetUserData(data, id string) (interface{}, error)
	DeleteUserData(data, id string) error

	CardEdit(index string, card entity.UserCard) error
	BinaryEdit(index string, binary entity.UserBinary) error
	TextEdit(index string, text entity.UserText) error
	CredentialsEdit(index string, text entity.UserCredentials) error

	RegisterUser(auth entity.UserAuth) error
	LoginUser(auth entity.UserAuth) error
}

type handler struct {
	processor secretsProcessor
}

func (h *handler) Login(ctx *gin.Context) {
	var auth entity.UserAuth
	if err := json.NewDecoder(ctx.Request.Body).Decode(&auth); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.LoginUser(auth)
	if errors.Is(err, entity.ErrInvalidCredentials) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) Logout(ctx *gin.Context) {

}

func (h *handler) Register(ctx *gin.Context) {
	var auth entity.UserAuth
	if err := json.NewDecoder(ctx.Request.Body).Decode(&auth); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.RegisterUser(auth)
	if errors.Is(err, entity.ErrLoginAlreadyInUse) {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) List(ctx *gin.Context) {
	data := ctx.Param("data")
	user := ctx.Param("user")
	res, err := h.processor.ListUserData(data, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	switch data {
	case "credentials":
		fmt.Println(res)
		res = res.([]entity.CutCredentials)
	case "binary":
		fmt.Println(res)
		res = res.([]entity.CutBinary)
	case "text":
		fmt.Println(res)
		res = res.([]entity.CutText)
	case "card":
		fmt.Println(res)
		res = res.([]entity.CutCard)
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *handler) Get(ctx *gin.Context) {
	data := ctx.Param("data")
	id := ctx.Param("id")
	res, err := h.processor.GetUserData(data, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	switch data {
	case "credentials":
		res = res.(entity.UserCredentials)
	case "binary":
		res = res.(entity.UserBinary)
	case "text":
		res = res.(entity.UserText)
	case "card":
		res = res.(entity.UserCard)
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *handler) GeneratePassword(ctx *gin.Context) {

}

func (h *handler) CredentialsCreate(ctx *gin.Context) {
	user := ctx.Param("user")

	var credentials entity.UserCredentials
	if err := json.NewDecoder(ctx.Request.Body).Decode(&credentials); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CredentialsCreation(credentials, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) Edit(ctx *gin.Context) {

}

func (h *handler) Delete(ctx *gin.Context) {
	data := ctx.Param("data")
	id := ctx.Param("id")
	err := h.processor.DeleteUserData(data, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) TextCreate(ctx *gin.Context) {
	user := ctx.Param("user")

	var text entity.UserText
	if err := json.NewDecoder(ctx.Request.Body).Decode(&text); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.TextCreation(text, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) BinaryCreate(ctx *gin.Context) {
	user := ctx.Param("user")

	var binary entity.UserBinary
	if err := json.NewDecoder(ctx.Request.Body).Decode(&binary); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.BinaryCreation(binary, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) CardCreate(ctx *gin.Context) {
	user := ctx.Param("user")

	var card entity.UserCard
	if err := json.NewDecoder(ctx.Request.Body).Decode(&card); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CardCreation(card, user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) CredentialsEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var credentials entity.UserCredentials
	if err := json.NewDecoder(ctx.Request.Body).Decode(&credentials); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CredentialsEdit(id, credentials)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) TextEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var text entity.UserText
	if err := json.NewDecoder(ctx.Request.Body).Decode(&text); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.TextEdit(id, text)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}

func (h *handler) BinaryEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var binary entity.UserBinary
	if err := json.NewDecoder(ctx.Request.Body).Decode(&binary); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.BinaryEdit(id, binary)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}

func (h *handler) CardEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	var card entity.UserCard
	if err := json.NewDecoder(ctx.Request.Body).Decode(&card); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CardEdit(id, card)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}

func New(processor secretsProcessor) *handler {
	return &handler{
		processor: processor,
	}
}

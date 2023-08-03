package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type secretsProcessor interface {
	CardCreation(card entity.UserCard) error
	BinaryCreation(binary entity.UserBinary) error
	TextCreation(text entity.UserText) error
	CredentialsCreation(text entity.UserCredentials) error

	ListUserData(data string) (interface{}, error)
	GetUserData(data, id string) (interface{}, error)
}

type handler struct {
	processor secretsProcessor
}

func (h *handler) Login(ctx *gin.Context) {

}

func (h *handler) Logout(ctx *gin.Context) {

}

func (h *handler) Register(ctx *gin.Context) {

}

func (h *handler) List(ctx *gin.Context) {
	data := ctx.Param("data")
	res, err := h.processor.ListUserData(data)
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
	var credentials entity.UserCredentials
	if err := json.NewDecoder(ctx.Request.Body).Decode(&credentials); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CredentialsCreation(credentials)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(credentials)

}

func (h *handler) Edit(ctx *gin.Context) {

}

func (h *handler) Delete(ctx *gin.Context) {

}

func (h *handler) TextCreate(ctx *gin.Context) {
	var text entity.UserText
	if err := json.NewDecoder(ctx.Request.Body).Decode(&text); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.TextCreation(text)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(text)
}

func (h *handler) BinaryCreate(ctx *gin.Context) {
	var binary entity.UserBinary
	if err := json.NewDecoder(ctx.Request.Body).Decode(&binary); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.BinaryCreation(binary)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(binary)
}

func (h *handler) CardCreate(ctx *gin.Context) {
	var card entity.UserCard
	if err := json.NewDecoder(ctx.Request.Body).Decode(&card); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := h.processor.CardCreation(card)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(card)
}

func New(processor secretsProcessor) *handler {
	return &handler{
		processor: processor,
	}
}

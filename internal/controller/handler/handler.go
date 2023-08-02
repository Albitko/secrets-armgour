package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type secretsProcessor interface {
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

}

func (h *handler) Get(ctx *gin.Context) {

}

func (h *handler) GeneratePassword(ctx *gin.Context) {

}

func (h *handler) CredentialsCreate(ctx *gin.Context) {
	var requestJSON entity.UserCredentials
	if err := json.NewDecoder(ctx.Request.Body).Decode(&requestJSON); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(requestJSON)

}

func (h *handler) Edit(ctx *gin.Context) {

}

func (h *handler) Delete(ctx *gin.Context) {

}

func New(processor secretsProcessor) *handler {
	return &handler{
		processor: processor,
	}
}

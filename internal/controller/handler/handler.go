package handler

import (
	"github.com/gin-gonic/gin"
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

func (h *handler) Create(ctx *gin.Context) {

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

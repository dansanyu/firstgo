package handlers

import (
	"firstgo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
func (h *UserHandler) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "用户创建成功"})
}

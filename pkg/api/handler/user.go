package handler

import (
	useCaseInterface "usermanagement/pkg/usecase/interface"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase useCaseInterface.UserUserCase
}

type Response struct {
	ID    uint
	Name  string
	Image string
}

func NewUserHandler(useCase useCaseInterface.UserUserCase) *UserHandler {
	return &UserHandler{useCase}
}

func (cr *UserHandler) GetAllUsers(c *gin.Context) {
}

func (cr *UserHandler) GetSingleUser(c *gin.Context) {

}

func (cr *UserHandler) CreateUser(c *gin.Context) {

}

func (cr *UserHandler) DeleteUser(c *gin.Context) {

}

func (cr *UserHandler) UpdateUser(c *gin.Context) {

}

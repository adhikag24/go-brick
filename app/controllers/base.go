package controllers

import (
	"github.com/adhikag24/go-brick/app/usecase"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	AccountValidate(c *gin.Context)
}

func NewControllers(us usecase.UC) Controllers {
	return &AccountController{UC: us}
}

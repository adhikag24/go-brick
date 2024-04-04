package controllers

import (
	"github.com/adhikag24/go-brick/app/usecase"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	UC usecase.UC
}

func (u *AccountController) AccountValidate(c *gin.Context) {

}

// func (u *UserController) DisplayUser(c *gin.Context) {
// 	userId := c.PostForm("userid")
// 	if userId == "" {
// 		apiErr := &utils.ApplicationError{
// 			Message:    "userid is required",
// 			StatusCode: http.StatusBadRequest,
// 			Code:       "bad_request",
// 		}
// 		utils.RespondError(c, apiErr)
// 		return
// 	}

// 	user, apiErr := u.UC.GetUser(userId)
// 	if apiErr != nil {
// 		utils.RespondError(c, apiErr)
// 		return
// 	}

// 	utils.Respond(c, http.StatusOK, user)
// }

// func (u *UserController) DisplayAllUsers(c *gin.Context) {
// 	user, apiErr := u.UC.GetAllUsers()

// 	if apiErr != nil {
// 		utils.RespondError(c, apiErr)
// 		return
// 	}

// 	utils.Respond(c, http.StatusOK, user)
// }

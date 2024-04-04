package routes

import (
	"github.com/adhikag24/go-brick/app/controllers"
	"github.com/gin-gonic/gin"
)

type route struct {
	ctrl controllers.Controllers
}

func NewRouter(ctr controllers.Controllers) route {
	return route{ctrl: ctr}
}

func (r *route) Router() *gin.Engine {
	var router *gin.Engine = gin.Default()

	router.POST("/displayuser", r.ctrl.DisplayUser)
	router.GET("/displayallusers", r.ctrl.DisplayAllUsers)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, this is Golang API Test Case."})
	})

	return router

}

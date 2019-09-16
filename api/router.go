package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/controllers"
)

func Router(e *gin.Engine) {
	api := e.Group("/api")

	api.GET("/basics", controllers.GetBasics)
}

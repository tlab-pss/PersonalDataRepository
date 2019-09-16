package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yuuis/PersonalDataRepository/api/controllers"
)

func Router(e *gin.Engine, d *gorm.DB) {
	ds := controllers.NewRegistry(d)
	api := e.Group("/api")

	api.GET("/basic", ds.GetBasics)
}

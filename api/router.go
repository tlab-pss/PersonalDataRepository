package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yuuis/PersonalDataRepository/api/controllers"
)

func Router(e *gin.Engine, d *gorm.DB) {
	ds := controllers.NewRegistry(d)
	api := e.Group("/api")

	// basic
	api.GET("/basics", ds.GetBasic)
	api.POST("/basics", ds.CreateBasic)

	// location
	api.GET("/locations", ds.GetLocation)
	api.POST("/locations", ds.CreateLocation)

  // health
	api.GET("/healths", ds.GetHealth)
	api.POST("/healths", ds.CreateHealth)

  // registration information
	api.GET("/registration-informations", ds.GetRegistrationInformation)
	api.POST("/registration-informations", ds.CreateRegistrationInformation)
}

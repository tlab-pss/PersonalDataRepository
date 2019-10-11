package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(e *gin.Engine, c *mongo.Client) {
	ds := controllers.NewRegistry(c)
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

	// registered information
	api.GET("/registered-informations", ds.GetRegisteredInformation)
	api.POST("/registered-informations", ds.CreateRegisteredInformation)

	// plugin service
	api.GET("/plugin-services", ds.GetPluginService)
	api.POST("/plugin-services", ds.CreatePluginService)

	// user like
	api.GET("/user-likes", ds.GetUserLike)
	api.POST("/user-likes", ds.CreateUserLike)
}

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

	// basic location
	api.GET("/basic-locations", ds.GetBasicLocation)
	api.POST("/basic-locations", ds.CreateBasicLocation)

	// location
	api.GET("/locations", ds.GetLocations)
	api.GET("/locations/latest", ds.GetLatestLocation)
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

	// conversation
	api.GET("/conversations", ds.GetConversation)
	api.GET("/conversations/:transactionID", ds.FindConversation)
	api.POST("/conversations", ds.CreateConversation)

	// plugins

	// hotpepper
	api.GET("/hotpepper/intakes", ds.GetIntake)
	api.POST("/hotpepper/intakes", ds.CreateIntake)
}

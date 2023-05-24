package apis

import (
	"github.com/ini8labs/admin-service/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(server Server) error {

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// API end point
	r.GET("/api/v1/user", server.userInfo)

	r.GET("/api/v1/users", server.userInfoByEventId)

	r.GET("/api/v1/event", server.eventInfo)
	r.POST("/api/v1/event", server.addNewEvent)
	r.DELETE("/api/v1/event/:EventUID", server.deleteEvent)

	// r.GET("/api/v1/event/Winners", server.getEventWinners)
	// r.POST("api/v1/event/AddWinner", server.addNewWinner)
	r.GET("/api/v1/events", server.getEventInfo)

	return r.Run(server.Addr)
}

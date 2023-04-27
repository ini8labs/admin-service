package apis

import (
	"github.com/gin-gonic/gin"
)

func NewServer(server Server) error {

	r := gin.Default()

	// API end point
	r.GET("/api/v1/userinfo", server.UserInfo)

	r.GET("/api/v1/eventinfo", server.EventsInfo)
	r.GET("/api/v1/eventinfo/Date", server.GetEventsByDate)
	r.GET("/api/v1/eventinfo/Daterange", server.GetEventsByDateRange)
	//r.GET("/api/v1/eventinfo/Winners", server.GetEventWinners) // will not work

	r.POST("/api/v1/event/Add", server.AddNewEvent)
	r.DELETE("/api/v1/event/Delete", server.DeleteEvent)

	return r.Run(server.Addr)
}

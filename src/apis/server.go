package apis

import (
	"github.com/gin-gonic/gin"
)

func NewServer(server Server) error {

	r := gin.Default()

	// API end point
	//r.GET("/api/v1/user/info", GetAllUserData)
	r.GET("/api/v1/userinfo/Phone", server.GetUserInfoByPhone)
	r.GET("/api/v1/userinfo/Gov_Id", server.GetUserInfoByGovID)
	r.GET("/api/v1/userinfo/UID", server.GetUserInfoByID)
	r.GET("/api/v1/userinfo/EventID", server.GetParticipantsInfoByEventID)

	r.GET("/api/v1/eventinfo", server.GetAllEvents)
	r.GET("api/v1/eventinfo/Eventtype", server.GetEventsByType)
	// r.GET("/api/v1/eventinfo/Date", server.GetEventsByDate)                  need to convert string to primitive.Date
	// r.GET("/api/v1/eventinfo/Daterange", server.GetEventsByDateRange)        need to convert string to primitive.Date
	//r.GET("/api/v1/eventinfo/Winners", server.GetEventWinners) // will not work

	r.POST("/api/v1/event/Add", server.AddNewEvent)
	r.DELETE("/api/v1/event/Delete", server.DeleteEvent)

	return r.Run(server.Addr)
}

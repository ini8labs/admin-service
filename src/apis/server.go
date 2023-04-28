package apis

import (
	"github.com/gin-gonic/gin"
)

func NewServer(server Server) error {

	r := gin.Default()

	// API end point
	r.GET("/api/v1/user", server.userInfo)
	r.GET("/api/v1/users", server.userInfoByEventId)

	r.GET("/api/v1/event", server.eventInfo)
	//r.GET("/api/v1/eventinfo/Winners", server.GetEventWinners)   will not work

	r.POST("/api/v1/event/Add", server.addNewEvent)
	r.DELETE("/api/v1/event/Delete", server.deleteEvent)

	return r.Run(server.Addr)
}

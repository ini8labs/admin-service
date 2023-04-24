package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	*logrus.Logger
	*lsdb.Client
	Addr string
}

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
	//r.GET("/api/v1/eventinfo/Date", server.GetEventsByDate)
	// r.GET("/api/v1/eventinfo/Daterange", server.GetEventsByDateRange)
	r.GET("/api/v1/eventinfo/Winners", server.GetEventWinners) // will not work

	r.POST("/api/v1/event/Add", server.AddNewEvent)
	r.DELETE("/api/v1/event/Delete", server.DeleteEvent)

	return r.Run(server.Addr)
}

func (s Server) GetUserInfoByPhone(c *gin.Context) {

	phonenumber := c.Query("phone")
	userphonenumber, _ := strconv.Atoi(phonenumber)

	resp, err := s.Client.GetUserInfoByPhone(int64(userphonenumber))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s Server) GetUserInfoByGovID(c *gin.Context) {

	govid := c.Query("id")

	resp, err := s.Client.GetUserInfoByGovID(govid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s Server) GetUserInfoByID(c *gin.Context) {

	uid := c.Query("uid")
	objID, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		panic(err)
	}

	resp, err := s.Client.GetUserInfoByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s Server) GetAllEvents(c *gin.Context) {

	// var eventinfo EventInfo

	resp, err := s.Client.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (s Server) GetEventsByType(c *gin.Context) {

	eventtype := c.Query("eventtype")

	resp, err := s.Client.GetEventsByType(eventtype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)

}

// func (s Server) GetEventsByDate(c *gin.Context) {
// date := c.Query("date")
// eventdate, _ := strconv.Atoi(date)
// resp, err := s.Client.GetEventsByDate(eventdate)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// func (s Server) GetEventsByDateRange(c *gin.Context) {

// 	var eventdaterange EventDateRange

// 	if err := c.BindJSON(&eventdaterange); err != nil {
// 		c.JSON(http.StatusBadRequest, "Bad Format")
// 		return
// 	}

// 	resp, err := s.Client.GetEventByDateRange(eventdaterange.StartDate, eventdaterange.EndDate)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

func (s Server) GetParticipantsInfoByEventID(c *gin.Context) {

	eventid := c.Query("eventid")
	id, err := primitive.ObjectIDFromHex(eventid)
	if err != nil {
		panic(err)
	}

	resp, err := s.Client.GetParticipantsInfoByEventID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)

}

func (s Server) GetEventWinners(c *gin.Context) {
	eventid := c.Query("eventid")
	id, err := primitive.ObjectIDFromHex(eventid)
	if err != nil {
		panic(err)
	}

	resp, err := s.Client.GetEventWinners(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s Server) AddNewEvent(c *gin.Context) {

	var addevent lsdb.LotteryEventInfo

	if err := c.BindJSON(&addevent); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	if err := s.Client.AddNewEvent(addevent); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, "Event added successfully")
}

func (s Server) DeleteEvent(c *gin.Context) {

	var deleteevent lsdb.LotteryEventInfo

	if err := c.BindJSON(&deleteevent); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	if err := s.Client.DeleteEvent(deleteevent.EventUID); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, "Event deleted successfully")
}

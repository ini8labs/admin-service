package apis

import (
	"fmt"
	"net/http"

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
	//r.GET("/api/v1/user_info", GetAllUserData)
	r.GET("/api/v1/eventdata", server.GetAllEvents)
	r.POST("/api/v1/userinfo_phonenumber", server.GetUserInfoByPhone)
	r.POST("/api/v1/userinfo_govID", server.GetUserInfoByGovID)
	r.POST("/api/v1/userinfo_ID", server.GetUserInfoByID)
	r.POST("api/v1/eventdata_type", server.GetEventsByType)
	r.POST("/api/v1/eventdata_date", server.GetEventsByDate)
	r.POST("/api/v1/eventdata_daterange", server.GetEventsByDateRange)
	r.POST("/api/v1/userinfo_eventID", server.GetParticipantsInfoByEventID)
	r.POST("/api/v1/add_event", server.AddNewEvent)
	r.POST("/api/v1/delete_event", server.DeleteEvent)
	r.POST("/api/v1/eventwinners", server.GetEventWinners) // will not work

	return r.Run(server.Addr)
}

type UserPhoneNumber struct {
	PhoneNumber int64 `json:"phone"`
}

type UserGovtID struct {
	GovtID string `json:"govId"`
}

type UserID struct {
	UID primitive.ObjectID `json:"id"`
}

type EventType struct {
	Type string `json:"eventtype"`
}

type EventDate struct {
	EventDate primitive.DateTime `json:"date"`
}

type EventDateRange struct {
	StartDate primitive.DateTime `json:"startdate"`
	EndDate   primitive.DateTime `json:"enddate"`
}

type EventId struct {
	Id primitive.ObjectID `json:"eventid"`
}

type EventWinners struct {
	EventId primitive.ObjectID `json:"eventid"`
}

func (s Server) GetUserInfoByPhone(c *gin.Context) {

	var userphonenumber UserPhoneNumber

	if err := c.BindJSON(&userphonenumber); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := s.Client.GetUserInfoByPhone(userphonenumber.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
	fmt.Println(resp)
}

func (s Server) GetUserInfoByGovID(c *gin.Context) {

	var govid UserGovtID

	if err := c.BindJSON(&govid); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := s.Client.GetUserInfoByGovID(govid.GovtID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s Server) GetUserInfoByID(c *gin.Context) {

	var userid UserID

	if err := c.BindJSON(&userid); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := s.Client.GetUserInfoByID(userid.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s Server) GetAllEvents(c *gin.Context) {

	resp, err := s.Client.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (s Server) GetEventsByType(c *gin.Context) {
	var eventtype EventType

	if err := c.BindJSON(&eventtype); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := s.Client.GetEventsByType(eventtype.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)

}

func (s Server) GetEventsByDate(c *gin.Context) {

	var eventdate EventDate

	if err := c.BindJSON(&eventdate); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := s.Client.GetEventsByDate(eventdate.EventDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s Server) GetEventsByDateRange(c *gin.Context) {

	var eventdaterange EventDateRange

	if err := c.BindJSON(&eventdaterange); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := s.Client.GetEventByDateRange(eventdaterange.StartDate, eventdaterange.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s Server) GetParticipantsInfoByEventID(c *gin.Context) {

	var eventid lsdb.EventParticipantInfo

	if err := c.BindJSON(&eventid); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	fmt.Println(eventid.EventUID)
	resp, err := s.Client.GetParticipantsInfoByEventID(eventid.EventUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)

}

func (s Server) GetEventWinners(c *gin.Context) {

	var eventId EventWinners

	if err := c.BindJSON(&eventId); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := s.Client.GetEventWinners(eventId.EventId)
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

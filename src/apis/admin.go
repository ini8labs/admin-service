package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewServer(addr string, log *logrus.Logger) error {

	r := gin.Default()

	// API end point
	//r.GET("/api/v1/user_info", GetAllUserData)
	r.GET("/api/v1/eventdata", GetAllEvents)
	r.POST("/api/v1/userinfo_phonenumber", GetUserInfoByPhone)
	r.POST("/api/v1/userinfo_govID", GetUserInfoByGovID)
	r.POST("/api/v1/userinfo_ID", GetUserInfoByID)
	r.POST("api/v1/eventdata_type", GetEventsByType)
	r.POST("/api/v1/eventdata_date", GetEventsByDate)
	r.POST("/api/v1/eventdata_daterange", GetEventsByDateRange)
	r.POST("/api/v1/userinfo_eventID", GetParticipantsInfoByEventID) // not working
	r.POST("/api/v1/add_event", AddNewEvent)
	r.POST("/api/v1/delete_event", DeleteEvent)
	r.POST("/api/v1/eventwinners", GetEventWinners) // will not work
	//r.POST("/api/v1/addnewevent", AddNewEvent)
	return r.Run(addr)
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

func GetUserInfoByPhone(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var userphonenumber UserPhoneNumber

	if err := c.BindJSON(&userphonenumber); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := dbClient.GetUserInfoByPhone(userphonenumber.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
	fmt.Println(resp)
}

func GetUserInfoByGovID(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var govid UserGovtID

	if err := c.BindJSON(&govid); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := dbClient.GetUserInfoByGovID(govid.GovtID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetUserInfoByID(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var userid UserID

	if err := c.BindJSON(&userid); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := dbClient.GetUserInfoByID(userid.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetAllEvents(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	resp, err := dbClient.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}

	c.JSON(http.StatusOK, resp)

}

func GetEventsByType(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var eventtype EventType

	if err := c.BindJSON(&eventtype); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := dbClient.GetEventsByType(eventtype.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)

}

func GetEventsByDate(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var eventdate EventDate

	if err := c.BindJSON(&eventdate); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := dbClient.GetEventsByDate(eventdate.EventDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetEventsByDateRange(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var eventdaterange EventDateRange

	if err := c.BindJSON(&eventdaterange); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := dbClient.GetEventByDateRange(eventdaterange.StartDate, eventdaterange.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetParticipantsInfoByEventID(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var eventid EventId

	if err := c.BindJSON(&eventid); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := dbClient.GetParticipantsInfoByEventID(eventid.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetEventWinners(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var eventId EventWinners

	if err := c.BindJSON(&eventId); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	resp, err := dbClient.GetEventWinners(eventId.EventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AddNewEvent(c *gin.Context) {

	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var addevent lsdb.LotteryEventInfo

	if err := c.BindJSON(&addevent); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	if err := dbClient.AddNewEvent(addevent); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, "Event added successfully")
}

func DeleteEvent(c *gin.Context) {
	dbClient, err := lsdb.NewClient()

	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	var deleteevent lsdb.LotteryEventInfo

	if err := c.BindJSON(&deleteevent); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	if err := dbClient.DeleteEvent(deleteevent.EventUID); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}
	c.JSON(http.StatusOK, "Event deleted successfully")
}

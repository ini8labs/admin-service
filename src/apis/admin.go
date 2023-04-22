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
	//r.GET("/api/v1/user_info", UserData)
	r.GET("/api/v1/eventdata", GetAllEvents)
	r.POST("/api/v1/userinfo_phonenumber", GetUserInfoByPhone)
	r.POST("/api/v1/userinfo_govID", GetUserInfoByGovID)
	r.POST("/api/v1/userinfo_ID", GetUserInfoByID)
	r.POST("api/v1/eventdata_bytype", GetEventsByType)
	r.POST("/api/v1/eventdata_bydate", GetEventsByDate)
	r.POST("/api/v1/eventdata_byrange", GetEventsByDateRange)
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

type EventType struct {
	Type string `json:"eventtype"`
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

type EventDate struct {
	EventDate primitive.DateTime `json:"date"`
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

type EventDateRange struct {
	StartDate primitive.DateTime `json:"startdate"`
	EndDate   primitive.DateTime `json:"enddate"`
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

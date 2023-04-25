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
	r.GET("/api/v1/userinfo/UID", server.GetUserInfoByID) //not working
	r.GET("/api/v1/userinfo/EventID", server.GetParticipantsInfoByEventID)

	r.GET("/api/v1/eventinfo", server.GetAllEvents)
	r.GET("api/v1/eventinfo/Eventtype", server.GetEventsByType)
	// r.GET("/api/v1/eventinfo/Date", server.GetEventsByDate)                  need to convert string to primitive.Date
	// r.GET("/api/v1/eventinfo/Daterange", server.GetEventsByDateRange)        need to convert string to primitive.Date
	r.GET("/api/v1/eventinfo/Winners", server.GetEventWinners) // will not work

	r.POST("/api/v1/event/Add", server.AddNewEvent)
	r.DELETE("/api/v1/event/Delete", server.DeleteEvent)

	return r.Run(server.Addr)
}

type UserInfoByEventId struct {
	EventUID   primitive.ObjectID `bson:"_id,omitempty"`
	EventDate  primitive.DateTime `bson:"event_date,omitempty"`
	UserName   string             `bson:"name,omitempty"`
	EventName  string             `bson:"name,omitempty"`
	EventType  string             `bson:"event_type,omitempty"`
	BetUID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty"`
	BetNumbers []int              `bson:"bet_numbers,omitempty"`
	Amount     int                `bson:"amount,omitempty"`
}

type EventsInfo struct {
	EventUID      primitive.ObjectID `bson:"_id,omitempty"`
	EventDate     primitive.DateTime `bson:"event_date,omitempty"`
	EventName     string             `bson:"name,omitempty"`
	EventType     string             `bson:"event_type,omitempty"`
	WinningNumber int                `bson:"winning_number,omitempty"`
}

type UserInfo struct {
	UID   primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Phone int64              `bson:"phone,omitempty"`
	GovID string             `bson:"gov_id,omitempty"`
	EMail string             `bson:"e_mail,omitempty"`
}

func (s Server) GetUserInfoByPhone(c *gin.Context) {

	phonenumber := c.Query("phone")
	userphonenumber, _ := strconv.Atoi(phonenumber)

	resp, err := s.Client.GetUserInfoByPhone(int64(userphonenumber))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err)
		return
	}

	var userinfo UserInfo
	if resp.Phone == int64(userphonenumber) {
		userinfo.Name = resp.Name
		userinfo.UID = resp.UID
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	c.JSON(http.StatusOK, userinfo)
}

func (s Server) GetUserInfoByGovID(c *gin.Context) {

	govid := c.Query("id")

	resp, err := s.Client.GetUserInfoByGovID(govid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err)
		return
	}
	var userinfo UserInfo
	if resp.GovID == govid {
		userinfo.Name = resp.Name
		userinfo.UID = resp.UID
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	c.JSON(http.StatusOK, userinfo)
}

func (s Server) GetUserInfoByID(c *gin.Context) {

	uid := c.Query("uid")
	objID, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		logrus.Infoln("Provided hex string is not a valid ObjectID")
		return
	}

	resp, err1 := s.Client.GetUserInfoByID(objID)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err1)
		return
	}
	var userinfo UserInfo
	if resp.UID == objID {
		userinfo.Name = resp.Name
		userinfo.UID = resp.UID
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	c.JSON(http.StatusOK, userinfo)

}

func (s Server) GetAllEvents(c *gin.Context) {

	resp, err := s.Client.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err)
		return
	}

	var eventinfo EventsInfo

	for i := 0; i < len(resp); i++ {
		eventinfo.EventUID = resp[i].EventUID
		eventinfo.EventDate = resp[i].EventDate
		eventinfo.EventName = resp[i].Name
		eventinfo.EventType = resp[i].EventType
		eventinfo.WinningNumber = resp[i].WinningNumber

		c.JSON(http.StatusOK, eventinfo)
	}
}

func (s Server) GetEventsByType(c *gin.Context) {

	eventtype := c.Query("eventtype")

	resp, err := s.Client.GetEventsByType(eventtype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err)
		return
	}

	var eventinfo EventsInfo
	for i := 0; i < len(resp); i++ {
		eventinfo.EventUID = resp[i].EventUID
		eventinfo.EventDate = resp[i].EventDate
		eventinfo.EventName = resp[i].Name
		eventinfo.EventType = resp[i].EventType
		eventinfo.WinningNumber = resp[i].WinningNumber

		c.JSON(http.StatusOK, eventinfo)
	}

}

// func (s Server) GetEventsByDate(c *gin.Context) {
// 	date := c.Query("date")

// 	resp, err := s.Client.GetEventsByDate(date)
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
		logrus.Infoln(err)
		return
	}

	resp, err1 := s.Client.GetParticipantsInfoByEventID(id)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err1)
		return
	}
	for i := 0; i < len(resp); i++ {

		var userinfobyevent UserInfoByEventId
		if resp[i].EventUID == id {
			userinfobyevent.EventUID = resp[i].EventUID
			userinfobyevent.BetUID = resp[i].BetUID
			userinfobyevent.UserID = resp[i].UserID
			userinfobyevent.Amount = resp[i].Amount
			userinfobyevent.BetNumbers = resp[i].BetNumbers

			resp2, err2 := s.Client.GetUserInfoByID(userinfobyevent.UserID)
			if err2 != nil {
				c.JSON(http.StatusInternalServerError, "something is wrong with the server")
				logrus.Infoln(err2)
				return
			}
			if userinfobyevent.UserID == resp2.UID {
				userinfobyevent.UserName = resp2.Name
			}

			// resp3, err3 := s.Client.GetUserInfoByID(eventinfo.UserID)
			// if err3 != nil {
			// 	c.JSON(http.StatusInternalServerError, "something is wrong with the server")
			// 	logrus.Infoln(err3)
			// 	return
			// }

			c.JSON(http.StatusOK, userinfobyevent)
		}
	}

}

func (s Server) GetEventWinners(c *gin.Context) {
	eventid := c.Query("eventid")
	id, err := primitive.ObjectIDFromHex(eventid)
	if err != nil {
		logrus.Infoln(err)
	}

	resp, err := s.Client.GetEventWinners(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s Server) AddNewEvent(c *gin.Context) {

	var addevent lsdb.LotteryEventInfo

	if err := c.BindJSON(&addevent); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		logrus.Infoln(err)
		return
	}

	if err := s.Client.AddNewEvent(addevent); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, "Event added successfully")
}

func (s Server) DeleteEvent(c *gin.Context) {

	var deleteevent lsdb.LotteryEventInfo

	if err := c.BindJSON(&deleteevent); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Format")
		logrus.Infoln(err)
		return
	}

	if err := s.Client.DeleteEvent(deleteevent.EventUID); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, "Event deleted successfully")
}

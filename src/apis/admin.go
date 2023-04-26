package apis

import (
	"net/http"
	"strconv"
	"time"

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
	//r.GET("/api/v1/userinfo/EventID", server.GetParticipantsInfoByEventID)

	r.GET("/api/v1/eventinfo", server.GetAllEvents)
	r.GET("api/v1/eventinfo/Eventtype", server.GetEventsByType)
	// r.GET("/api/v1/eventinfo/Date", server.GetEventsByDate)                  need to convert string to primitive.Date
	// r.GET("/api/v1/eventinfo/Daterange", server.GetEventsByDateRange)        need to convert string to primitive.Date
	//r.GET("/api/v1/eventinfo/Winners", server.GetEventWinners) // will not work

	r.POST("/api/v1/event/Add", server.AddNewEvent)
	r.DELETE("/api/v1/event/Delete", server.DeleteEvent)

	return r.Run(server.Addr)
}

type UserInfoByEventId struct {
	EventUID   string    `json:"_id,omitempty"`
	EventDate  time.Time `json:"event_date,omitempty"`
	UserName   string    `json:"name,omitempty"`
	EventName  string    `json:"eventname,omitempty"`
	EventType  string    `json:"event_type,omitempty"`
	BetUID     string    `json:"betid,omitempty"`
	UserID     string    `json:"user_id,omitempty"`
	BetNumbers []int     `json:"bet_numbers,omitempty"`
	Amount     int       `json:"amount,omitempty"`
}

type EventsInfo struct {
	EventUID      string    `json:"_id,omitempty"`
	EventDate     time.Time `json:"event_date,omitempty"`
	EventName     string    `json:"name,omitempty"`
	EventType     string    `json:"event_type,omitempty"`
	WinningNumber int       `json:"winning_number,omitempty"`
}

type UserInfo struct {
	UID   string `json:"_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone int64  `json:"phone,omitempty"`
	GovID string `json:"gov_id,omitempty"`
	EMail string `json:"e_mail,omitempty"`
}

func (s Server) GetUserInfoByPhone(c *gin.Context) {

	phonenumber := c.Query("phone")
	userphonenumber, _ := strconv.Atoi(phonenumber)

	resp, err := s.Client.GetUserInfoByPhone(int64(userphonenumber))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}

	var userinfo UserInfo
	if resp.Phone == int64(userphonenumber) {

		mongouserId := resp.UID
		stringUserID := mongouserId.Hex()

		userinfo.Name = resp.Name
		userinfo.UID = stringUserID
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
		s.Logger.Error(err.Error())
		return
	}
	var userinfo UserInfo
	if resp.GovID == govid {
		mongouserId := resp.UID
		stringUserID := mongouserId.Hex()

		userinfo.Name = resp.Name
		userinfo.UID = stringUserID
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
		s.Logger.Error(err.Error())
		return
	}

	resp, err1 := s.Client.GetUserInfoByID(objID)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	var userinfo UserInfo
	if resp.UID == objID {

		mongouserId := resp.UID
		stringUserID := mongouserId.Hex()

		userinfo.Name = resp.Name
		userinfo.UID = stringUserID
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
		s.Logger.Error(err.Error())
		return
	}

	var arr []EventsInfo

	for i := 0; i < len(resp); i++ {
		var eventinfo EventsInfo

		mongouserId := resp[i].EventUID
		stringEventID := mongouserId.Hex()
		eventinfo.EventUID = stringEventID

		mongoEventDate := resp[i].EventDate
		stringEvenDate := mongoEventDate.Time()
		eventinfo.EventDate = stringEvenDate

		eventinfo.EventName = resp[i].Name
		eventinfo.EventType = resp[i].EventType
		eventinfo.WinningNumber = resp[i].WinningNumber

		arr = append(arr, eventinfo)
	}
	c.JSON(http.StatusOK, arr)
}

func (s Server) GetEventsByType(c *gin.Context) {

	eventtype := c.Query("eventtype")

	resp, err := s.Client.GetEventsByType(eventtype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}

	var arr []EventsInfo

	for i := 0; i < len(resp); i++ {
		var eventinfo EventsInfo

		mongoEventId := resp[i].EventUID
		stringEventID := mongoEventId.Hex()
		eventinfo.EventUID = stringEventID

		mongoEventDate := resp[i].EventDate
		stringEvenDate := mongoEventDate.Time()
		eventinfo.EventDate = stringEvenDate

		eventinfo.EventName = resp[i].Name
		eventinfo.EventType = resp[i].EventType
		eventinfo.WinningNumber = resp[i].WinningNumber

		arr = append(arr, eventinfo)
	}
	c.JSON(http.StatusOK, arr)

}

// // func (s Server) GetEventsByDate(c *gin.Context) {
// // 	date := c.Query("date")
// // 	eventdate := primitive.NewDateTimeFromTime(time.Date(2023, time.April, 24, 0, 0, 0, 0, time.Local))

// // 	resp, err := s.Client.GetEventsByDate(eventdate)
// // 	if err != nil {
// // 		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// // 		return
// // 	}

// // 	var eventinfo EventsInfo

// // 	for i := 0; i < len(resp); i++ {
// // 		eventinfo.EventUID = resp[i].EventUID
// // 		eventinfo.EventDate = resp[i].EventDate
// // 		eventinfo.EventName = resp[i].Name
// // 		eventinfo.EventType = resp[i].EventType
// // 		eventinfo.WinningNumber = resp[i].WinningNumber

// // 		c.JSON(http.StatusOK, eventinfo)
// // 	}
// // }

// // func (s Server) GetEventsByDateRange(c *gin.Context) {

// // 	var eventdaterange EventDateRange

// // 	if err := c.BindJSON(&eventdaterange); err != nil {
// // 		c.JSON(http.StatusBadRequest, "Bad Format")
// // 		return
// // 	}

// // 	resp, err := s.Client.GetEventByDateRange(eventdaterange.StartDate, eventdaterange.EndDate)
// // 	if err != nil {
// // 		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// // 		return
// // 	}
// // 	c.JSON(http.StatusOK, resp)
// // }

// func (s Server) GetParticipantsInfoByEventID(c *gin.Context) {

// 	eventid := c.Query("eventid")
// 	id, err := primitive.ObjectIDFromHex(eventid)
// 	if err != nil {
// 		logrus.Infoln(err)
// 		return
// 	}

// 	resp, err1 := s.Client.GetParticipantsInfoByEventID(id)
// 	if err1 != nil {
// 		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// 		logrus.Infoln(err1)
// 		return
// 	}
// 	var arr []UserInfoByEventId

// 	for i := 0; i < len(resp); i++ {

// 		var userinfobyevent UserInfoByEventId

// 		if resp[i].EventUID == id {
// 			mongouserId := resp[i].UserID
// 			stringUserID := mongouserId.String()
// 			userinfobyevent.UserID = stringUserID

// 			mongoeventId := resp[i].EventUID
// 			stringEventID := mongoeventId.String()
// 			userinfobyevent.EventUID = stringEventID

// 			mongoBetId := resp[i].BetUID
// 			stringBetID := mongoBetId.String()
// 			userinfobyevent.BetUID = stringBetID

// 			userinfobyevent.Amount = resp[i].Amount
// 			userinfobyevent.BetNumbers = resp[i].BetNumbers

// 			UserID, err2 := primitive.ObjectIDFromHex(userinfobyevent.UserID)
// 			if err != nil {
// 				panic(err2)
// 			}

// 			resp2, err3 := s.Client.GetUserInfoByID(UserID)
// 			if err3 != nil {
// 				c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// 				logrus.Infoln(err2)
// 				return
// 			}

// 			if userinfobyevent.UserID == resp2.UID.String() {
// 				userinfobyevent.UserName = resp2.Name
// 			}

// 			// resp3, err3 := s.Client.GetUserInfoByID(eventinfo.UserID)
// 			// if err3 != nil {
// 			// 	c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// 			// 	logrus.Infoln(err3)
// 			// 	return
// 			// }
// 			arr = append(arr, userinfobyevent)
// 		}
// 	}
// 	c.JSON(http.StatusOK, arr)

// }

// func (s Server) GetEventWinners(c *gin.Context) {
// 	eventid := c.Query("eventid")
// 	id, err := primitive.ObjectIDFromHex(eventid)
// 	if err != nil {
// 		logrus.Infoln(err)
// 	}

// 	resp, err := s.Client.GetEventWinners(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// 		logrus.Infoln(err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

func (s Server) AddNewEvent(c *gin.Context) {

	// events := []string{"Monday Special", "Lucky Tuesday", "Midweek", "Fortune Thursday", "Friday Bonanza", "National Weekly"}
	//eventtypes := []string{"MS", "LT", "MW", "FT", "FB"}

	events := map[string]string{
		"Monday Special":   "MS",
		"Lucky Tuesday":    "LT",
		"Midweek":          "MW",
		"Fortune Thursday": "FT",
		"Friday Bonanza":   "FB",
		"National Weekly":  "NW",
	}

	event_date, exists1 := c.GetQuery("date")
	name, exists2 := c.GetQuery("name")
	eventtype, exists3 := c.GetQuery("eventtype")
	winingnumber, exists4 := c.GetQuery("winningnumber")

	intWinNumber, _ := strconv.Atoi(winingnumber)

	date, error := time.Parse("2006-01-02", event_date)
	if error != nil {
		s.Logger.Error("Can't parse date format")
		c.JSON(http.StatusBadRequest, "Date not in correct format")
		return
	}

	if intWinNumber < 1 || intWinNumber > 90 {
		c.JSON(http.StatusBadRequest, "Win numbers should be greater than 1 and less than 90")
		return
	}

	if _, ok := events[name]; !ok {
		c.JSON(http.StatusBadRequest, "Invalid Event")
		return
	}

	var result bool = false
	for _, value := range events {
		if value == eventtype {
			result = true
		}
	}

	if !result {
		c.JSON(http.StatusBadRequest, "Event type does not exist")
		return
	}

	if events[name] != eventtype {
		c.JSON(http.StatusBadRequest, "Event does not match event type")
		return
	}

	eventinfo := lsdb.LotteryEventInfo{
		Name:          name,
		EventDate:     primitive.NewDateTimeFromTime(date),
		WinningNumber: intWinNumber,
		EventType:     eventtype,
	}
	if !exists1 || !exists2 || !exists3 || !exists4 {
		c.JSON(http.StatusBadRequest, "Bad Format")
		return
	}

	if err := s.Client.AddNewEvent(eventinfo); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, "Event added successfully")
}

func (s Server) DeleteEvent(c *gin.Context) {
	eventid, exists := c.GetQuery("EventUID")

	if !exists {
		s.Logger.Error("Field Missing")
		c.JSON(http.StatusBadRequest, "Bad Format")
	}

	objID, err := primitive.ObjectIDFromHex(eventid)
	if err != nil {
		logrus.Infoln("Provided hex string is not a valid ObjectID")
		return
	}

	if err := s.Client.DeleteEvent(objID); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		logrus.Infoln(err)
		return
	}
	c.JSON(http.StatusOK, "Event deleted successfully")
}

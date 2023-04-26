package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func (s Server) GetEventWinners(c *gin.Context) {
	eventid := c.Query("eventid")
	id, err := primitive.ObjectIDFromHex(eventid)
	if err != nil {
		s.Logger.Error(err.Error())
	}

	resp, err := s.Client.GetEventWinners(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
)

func initializeEventInfo(resp []lsdb.LotteryEventInfo) []EventsInfo {
	var arr []EventsInfo

	for i := 0; i < len(resp); i++ {
		var eventinfo EventsInfo

		eventinfo.EventUID = primitiveToString(resp[i].EventUID)
		eventinfo.EventDate = convertPrimitiveToTime(resp[i].EventDate)
		eventinfo.EventName = resp[i].Name
		eventinfo.EventType = resp[i].EventType
		eventinfo.WinningNumber = resp[i].WinningNumber

		arr = append(arr, eventinfo)
	}
	return arr
}

func (s Server) GetAllEvents(c *gin.Context) {

	resp, err := s.Client.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}

	result := initializeEventInfo(resp)
	c.JSON(http.StatusOK, result)
}

func (s Server) GetEventsByType(c *gin.Context) {

	eventtype := c.Query("eventtype")

	resp, err := s.Client.GetEventsByType(eventtype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	result := initializeEventInfo(resp)
	c.JSON(http.StatusOK, result)

}

func (s Server) GetEventsByDate(c *gin.Context) {

	var eventDate GetEventsByDate
	if err := c.BindJSON(&eventDate); err != nil {
		c.JSON(http.StatusBadRequest, "bad Format")
		s.Logger.Error(err.Error())
		return
	}

	resp, err := s.Client.GetEventsByDate(convertTimeToPrimitive(eventDate.EventDate))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	result := initializeEventInfo(resp)
	c.JSON(http.StatusOK, result)
}

func (s Server) GetEventsByDateRange(c *gin.Context) {

	var eventDate GetEventsByDate
	if err := c.BindJSON(&eventDate); err != nil {
		c.JSON(http.StatusBadRequest, "bad Format")
		s.Logger.Error(err.Error())
		return
	}

	resp, err := s.Client.GetEventByDateRange(convertTimeToPrimitive(eventDate.StartDate), convertTimeToPrimitive(eventDate.EndDate))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	result := initializeEventInfo(resp)
	c.JSON(http.StatusOK, result)
}

// func (s Server) GetEventWinners(c *gin.Context) {
// 	eventid := c.Query("eventid")
// 	id, err := primitive.ObjectIDFromHex(eventid)
// 	if err != nil {
// 		s.Logger.Error(err.Error())
// 	}

// 	resp, err := s.Client.GetEventWinners(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
// 		s.Logger.Error(err.Error())
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

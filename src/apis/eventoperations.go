package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
)

func (s Server) AddNewEvent(c *gin.Context) {

	events := map[string]string{
		"Monday Special":   "MS",
		"Lucky Tuesday":    "LT",
		"Midweek":          "MW",
		"Fortune Thursday": "FT",
		"Friday Bonanza":   "FB",
		"National Weekly":  "NW",
	}

	var newEvent AddNewEventReq
	if err := c.ShouldBind(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, "bad Format")
		s.Logger.Error(err.Error())
		return
	}

	date := convertTimeToPrimitive(newEvent.EventDate)

	if newEvent.WinningNumber < 1 || newEvent.WinningNumber > 90 {
		c.JSON(http.StatusBadRequest, "Win numbers should be greater than 1 and less than 90")
		return
	}

	if _, ok := events[newEvent.Name]; !ok {
		c.JSON(http.StatusBadRequest, "Invalid Event")
		return
	}

	var result bool = false
	for _, value := range events {
		if value == newEvent.EventType {
			result = true
		}
	}

	if !result {
		c.JSON(http.StatusBadRequest, "Event type does not exist")
		return
	}

	if events[newEvent.Name] != newEvent.EventType {
		c.JSON(http.StatusBadRequest, "Event does not match event type")
		return
	}

	eventinfo := lsdb.LotteryEventInfo{
		Name:          newEvent.Name,
		EventDate:     date,
		WinningNumber: newEvent.WinningNumber,
		EventType:     newEvent.EventType,
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
		c.JSON(http.StatusBadRequest, "Bad Format")
	}

	if err := s.Client.DeleteEvent(stringToPrimitive(eventid)); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, "Event deleted successfully")
}

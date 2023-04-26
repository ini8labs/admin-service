package apis

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	event_date := c.Query("date")
	name := c.Query("name")
	eventtype := c.Query("eventtype")
	winingnumber := c.Query("winningnumber")

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

	if err := s.Client.DeleteEvent(StringToPrimitive(eventid)); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, "Event deleted successfully")
}

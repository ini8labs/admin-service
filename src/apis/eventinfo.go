package apis

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
)

var events map[string]string = map[string]string{
	"Monday Special":   "MS",
	"Lucky Tuesday":    "LT",
	"Midweek":          "MW",
	"Fortune Thursday": "FT",
	"Friday Bonanza":   "FB",
	"National Weekly":  "NW",
}

func (s Server) addNewEvent(c *gin.Context) {

	var newEvent AddNewEventReq

	if err := c.ShouldBind(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, "bad Format")
		s.Logger.Error(err.Error())
		fmt.Println(newEvent)
		return
	}

	date := convertTimeToPrimitive(newEvent.EventDate)

	for i := 0; i < len(newEvent.WinningNumber); i++ {
		if newEvent.WinningNumber[i] < 1 || newEvent.WinningNumber[i] > 90 {
			c.JSON(http.StatusBadRequest, "Win numbers should be greater than 0 and less than 90")
			return
		}
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
	c.JSON(http.StatusCreated, "Event added successfully")
}

func validateEventId(str string, resp []EventsInfo) bool {
	eventIdExist := true

	for i := 0; i < len(resp); i++ {
		if resp[i].EventUID == str {
			eventIdExist = true
			break
		}
		if resp[i].EventUID != str {
			eventIdExist = false
		}
	}
	return eventIdExist
}

func (s Server) deleteEvent(c *gin.Context) {
	eventid := c.Param("EventUID")

	resp, err := s.getEventInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	validation := validateEventId(eventid, resp)
	if !validation {
		c.JSON(http.StatusBadRequest, "EventId does not exist")
		return
	}

	if err := s.Client.DeleteEvent(stringToPrimitive(eventid)); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, "Event deleted successfully")
}

func initializeEventInfo(resp []lsdb.LotteryEventInfo) []EventsInfo {
	var arr []EventsInfo

	for i := 0; i < len(resp); i++ {
		eventinfo := EventsInfo{
			EventUID:      primitiveToString(resp[i].EventUID),
			EventDate:     convertPrimitiveToTime(resp[i].EventDate),
			EventName:     resp[i].Name,
			EventType:     resp[i].EventType,
			WinningNumber: resp[i].WinningNumber,
		}

		arr = append(arr, eventinfo)
	}
	return arr
}

func (s Server) eventInfo(c *gin.Context) {

	eventType := c.Query("eventType")
	date := c.Query("date")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	eventInfo, err := s.getEventByQueryParams(eventType, date, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Problem with the server")
		return
	}

	c.JSON(http.StatusOK, eventInfo)

}

func (s Server) getEventByQueryParams(eventType, date, startDate, endDate string) ([]EventsInfo, error) {
	var eventInfo []EventsInfo
	var err error

	switch {
	case eventType != "":
		eventInfo, err = s.getEventInfoByType(eventType)
	case date != "":
		eventInfo, err = s.getEventInfoByDate(date)
	case startDate != "" && endDate != "":
		eventInfo, err = s.getEventInfoByDateRange(startDate, endDate)
	default:
		eventInfo, err = s.getEventInfo()
	}

	return eventInfo, err
}

func (s Server) getEventInfoByDateRange(startDate, endDate string) ([]EventsInfo, error) {
	initialDate := strings.Split(startDate, "-")

	intStartYear, _ := strconv.Atoi(initialDate[0])
	intStartMonth, _ := strconv.Atoi(initialDate[1])
	intStartDay, _ := strconv.Atoi(initialDate[2])

	startRangeDate := Date{
		Year:  intStartYear,
		Month: intStartMonth,
		Day:   intStartDay,
	}

	lastDate := strings.Split(endDate, "-")

	intEndYear, _ := strconv.Atoi(lastDate[0])
	intEndMonth, _ := strconv.Atoi(lastDate[1])
	intEndDay, _ := strconv.Atoi(lastDate[2])

	endRangeDate := Date{
		Year:  intEndYear,
		Month: intEndMonth,
		Day:   intEndDay,
	}
	resp, err := s.Client.GetEventByDateRange(convertTimeToPrimitive(startRangeDate), convertTimeToPrimitive(endRangeDate))

	if err != nil {
		return []EventsInfo{}, err
	}

	result := initializeEventInfo(resp)
	return result, nil
}

func (s Server) getEventInfo() ([]EventsInfo, error) {
	resp, err := s.Client.GetAllEvents()
	if err != nil {
		return []EventsInfo{}, err
	}

	result := initializeEventInfo(resp)
	return result, nil
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

func (s Server) getEventInfoByType(eventType string) ([]EventsInfo, error) {
	resp, err := s.Client.GetEventsByType(eventType)
	if err != nil {
		return []EventsInfo{}, err
	}
	eventInfo := initializeEventInfo(resp)
	return eventInfo, nil
}

func (s Server) getEventInfoByDate(date string) ([]EventsInfo, error) {

	eventDate := strings.Split(date, "-")
	intYear, _ := strconv.Atoi(eventDate[0])
	intMonth, _ := strconv.Atoi(eventDate[1])
	intDay, _ := strconv.Atoi(eventDate[2])

	eventdate := Date{
		Year:  intYear,
		Month: intMonth,
		Day:   intDay,
	}

	resp, err := s.Client.GetEventsByDate(convertTimeToPrimitive(eventdate))
	if err != nil {
		return []EventsInfo{}, err
	}

	result := initializeEventInfo(resp)
	return result, nil
}

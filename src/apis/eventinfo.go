package apis

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var events map[string]string = map[string]string{
	"Monday Special":   "MS",
	"Lucky Tuesday":    "LT",
	"Midweek":          "MW",
	"Fortune Thursday": "FT",
	"Friday Bonanza":   "FB",
	"National Weekly":  "NW",
}

func validateGenerateEventDate(newEvent AddNewEventReq) (primitive.DateTime, error) {
	date := convertTimeToPrimitive(newEvent.EventDate)
	fmt.Println(newEvent)
	var err error

	if daysInMonth(newEvent.EventDate.Month, newEvent.EventDate.Year) < newEvent.EventDate.Day || daysInMonth(int(time.Now().Month()), time.Now().Year()) < newEvent.EventDate.Day || newEvent.EventDate.Day < 1 || newEvent.EventDate.Month > 12 || newEvent.EventDate.Month < 1 {
		err = errors.New("invalid date")
		return date, err
	}

	if ((daysInMonth(int(time.Now().Month()), time.Now().Year()) < newEvent.EventDate.Day) && (newEvent.EventDate.Year == time.Now().Year()) && (newEvent.EventDate.Month == int(time.Now().Month()))) || newEvent.EventDate.Year < time.Now().Year() || (newEvent.EventDate.Year <= time.Now().Year() && newEvent.EventDate.Month < int(time.Now().Month()) || ((newEvent.EventDate.Year == time.Now().Year() && newEvent.EventDate.Month == int(time.Now().Month())) && newEvent.EventDate.Day <= time.Now().Day())) {
		err = errors.New("events can only be generated for future")
		return date, err
	}
	return date, nil
}

func validateGenerateEventWinningNumbers(newEvent AddNewEventReq) error {
	var err error
	if len(newEvent.WinningNumber) > 5 || len(newEvent.WinningNumber) < 5 {
		err = errors.New("there should be 5 winning numbers")
		return err
	}

	for i := 0; i < len(newEvent.WinningNumber); i++ {
		if newEvent.WinningNumber[i] < 1 || newEvent.WinningNumber[i] > 90 {
			err = errors.New("winning numbers should be greater than 0 and less than 90")
			return err
		}
		count := 0
		for j := 0; j < len(newEvent.WinningNumber); j++ {
			if newEvent.WinningNumber[i] == newEvent.WinningNumber[j] {
				count++
			}
		}
		if count > 1 {
			err = errors.New("winning numbers should not be same")
			return err
		}
	}
	return nil
}

func validateGenerateEventNameAndType(newEvent AddNewEventReq) error {
	var result bool
	var err error

	for _, value := range events {
		if value == newEvent.EventType {
			result = true
			break
		}
	}

	if !result {
		err = errors.New("event type does not exist")
		return err
	}

	if _, ok := events[newEvent.Name]; !ok {
		err = errors.New("invalid Event")
		return err
	}

	if events[newEvent.Name] != newEvent.EventType {
		err = errors.New("event does not match event type")
		return err
	}

	return nil
}

func initializeGenerateEventInfo(newEvent AddNewEventReq, date primitive.DateTime) lsdb.LotteryEventInfo {
	eventinfo := lsdb.LotteryEventInfo{
		Name:          newEvent.Name,
		EventDate:     date,
		WinningNumber: newEvent.WinningNumber,
		EventType:     newEvent.EventType,
	}
	return eventinfo
}

func validateAddEvent(newEvent AddNewEventReq) (lsdb.LotteryEventInfo, error) {

	date, err := validateGenerateEventDate(newEvent)
	if err != nil {
		return lsdb.LotteryEventInfo{}, err
	}

	err = validateGenerateEventWinningNumbers(newEvent)
	if err != nil {
		return lsdb.LotteryEventInfo{}, err
	}

	err = validateGenerateEventNameAndType(newEvent)
	if err != nil {
		return lsdb.LotteryEventInfo{}, err
	}

	eventInfo := initializeGenerateEventInfo(newEvent, date)
	return eventInfo, nil
}

func initializeEventInfo(lottteryEventInfo []lsdb.LotteryEventInfo) []EventsInfo {
	var eventsInfoArr []EventsInfo

	for i := 0; i < len(lottteryEventInfo); i++ {
		eventinfo := EventsInfo{
			EventUID:      primitiveToString(lottteryEventInfo[i].EventUID),
			EventDate:     convertPrimitiveToTime(lottteryEventInfo[i].EventDate),
			EventName:     lottteryEventInfo[i].Name,
			EventType:     lottteryEventInfo[i].EventType,
			WinningNumber: lottteryEventInfo[i].WinningNumber,
		}

		eventsInfoArr = append(eventsInfoArr, eventinfo)
	}
	return eventsInfoArr
}

func (s Server) addNewEvent(c *gin.Context) {

	var addEvent AddNewEventReq

	if err := c.ShouldBind(&addEvent); err != nil {
		c.JSON(http.StatusBadRequest, "bad Format")
		s.Logger.Error(err.Error())
		fmt.Println(addEvent)
		return
	}
	validation, err := validateAddEvent(addEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		s.Logger.Error(err.Error())
		return
	}

	if err := s.Client.AddNewEvent(validation); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusCreated, "Event added successfully")
}

func checkEventId(eventId string, eventInfo []lsdb.LotteryEventInfo) bool {
	eventIdExist := true

	for i := 0; i < len(eventInfo); i++ {
		if eventInfo[i].EventUID == stringToPrimitive(eventId) {
			eventIdExist = true
			break
		}
		if eventInfo[i].EventUID != stringToPrimitive(eventId) {
			eventIdExist = false
		}
	}
	return eventIdExist
}

func (s Server) validateEventId(eventId string) (bool, error) {
	resp, err := s.GetAllEvents()
	if err != nil {
		s.Logger.Error(err.Error())
		return false, err
	}
	eventIdExist := checkEventId(eventId, resp)
	return eventIdExist, nil
}

func (s Server) deleteEvent(c *gin.Context) {
	eventid := c.Param("EventUID")

	validation, _ := s.validateEventId(eventid)
	if !validation {
		c.JSON(http.StatusBadRequest, "EventId does not exist")
		s.Logger.Error("invalid event id")
		return
	}

	if err := s.Client.DeleteEvent(stringToPrimitive(eventid)); err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		return
	}

	c.JSON(http.StatusOK, "Event deleted successfully")
}

func (s Server) eventInfo(c *gin.Context) {

	eventType := c.Query("eventType")
	date := c.Query("date")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	eventInfo, err := s.getEventByQueryParams(eventType, date, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		s.Logger.Error(err.Error())
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
	case startDate != "" || endDate != "":
		if startDate == "" {
			firstDayOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
			fmt.Println(firstDayOfYear)
			fmt.Println("@@@@@")
			startDate = convertTimeToString(firstDayOfYear)
		}
		if endDate == "" {
			lastDayOfYear := time.Date(time.Now().Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
			endDate = convertTimeToString(lastDayOfYear)
		}
		eventInfo, err = s.getEventInfoByDateRange(startDate, endDate)
	}

	return eventInfo, err
}

func (s Server) getEventInfoByDateRange(startDate, endDate string) ([]EventsInfo, error) {
	eventStartDate, err := convertStringToDate(startDate)
	if err != nil {
		return []EventsInfo{}, err
	}

	eventEndDate, err := convertStringToDate(endDate)
	if err != nil {
		return []EventsInfo{}, err
	}

	resp, err := s.GetEventByDateRange(convertTimeToPrimitive(eventStartDate), convertTimeToPrimitive(eventEndDate))
	if err != nil {
		err := errors.New("invalid date")
		s.Logger.Error(err.Error())
		return []EventsInfo{}, err
	}

	result := initializeEventInfo(resp)
	return result, nil
}

func (s Server) getEventInfo(c *gin.Context) {
	resp, err := s.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Error")
	}

	result := initializeEventInfo(resp)
	c.JSON(http.StatusOK, result)
}

func validaEventNameAndType(eventType string) error {
	var result bool

	for _, value := range events {
		if value == eventType {
			result = true
			break
		}
	}

	if !result {
		err := errors.New("event type does not exist")
		return err
	}
	return nil
}

func (s Server) getEventInfoByType(eventType string) ([]EventsInfo, error) {
	resp, err := s.GetEventsByType(eventType)
	if err != nil {
		return []EventsInfo{}, err
	}

	err = validaEventNameAndType(eventType)
	if err != nil {
		return []EventsInfo{}, err
	}

	eventInfo := initializeEventInfo(resp)
	return eventInfo, nil
}

func (s Server) getEventInfoByDate(date string) ([]EventsInfo, error) {

	eventDate, err := convertStringToDate(date)
	if err != nil {
		return []EventsInfo{}, err
	}

	resp, err := s.Client.GetEventsByDate(convertTimeToPrimitive(eventDate))
	if err != nil {
		err := errors.New("invalid date")
		return []EventsInfo{}, err
	}

	if len(resp) == 0 {
		err := errors.New("invalid date")
		return []EventsInfo{}, err
	}

	result := initializeEventInfo(resp)
	return result, nil
}

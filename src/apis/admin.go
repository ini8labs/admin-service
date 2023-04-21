package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewServer(addr string, log *logrus.Logger) error {

	r := gin.Default()

	// API end point
	r.GET("/api/v1/get_user_data", GetUserData)
	r.GET("/api/v1/get_event_data", GetEventData)
	r.POST("api/v1/event_data_byName", GetEventDataByName)

	return r.Run(addr)

}

type User struct {
	Id            string `json:"userid"`
	Name          string `json:"name"`
	ContactNumber string `json:"number"`
	IdProofNumber string `json:"idproof"`
}

func GetUserData(c *gin.Context) {

	content, err := os.ReadFile(".\\src\\apis\\usersinfo.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return
	}

	var users []User

	err2 := json.Unmarshal(content, &users)
	if err2 != nil {
		log.Fatal("Error during Unmarshal: ", err2)
		return
	}
	c.JSON(http.StatusOK, users)
}

type Event struct {
	EventName    string `json:"eventname"`
	EventId      string `json:"eventid"`
	Date         string `json:"date"`
	WinNumber    []int  `json:"winnumber"`
	Participants []struct {
		Id           string `json:"userid"`
		StakeNumbers []int  `json:"stakenumber"`
	} `json:"participants,omitempty"`
}

func GetEventData(c *gin.Context) {
	content, err := os.ReadFile(".\\src\\apis\\eventinfo.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return
	}

	var events []Event

	err2 := json.Unmarshal(content, &events)
	if err2 != nil {
		log.Fatal("Error during Unmarshal: ", err2)
		return
	}
	c.JSON(http.StatusOK, events)
}

type EventDataByName struct {
	EventName    string `json:"eventname"`
	EventId      string `json:"eventid"`
	Date         string `json:"date"`
	WinNumber    []int  `json:"winnumber"`
	Participants []struct {
		Id           string `json:"userid"`
		StakeNumbers []int  `json:"stakenumber"`
	} `json:"participants,omitempty"`
}

type RequestEventDataByName struct {
	EventName string `json:"eventname"`
}

func GetEventDataByName(c *gin.Context) {
	content, err := os.ReadFile(".\\src\\apis\\eventinfo.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return
	}

	var eventdatabyname []EventDataByName
	var requesteventdatabyname RequestEventDataByName

	if err := c.ShouldBind(&requesteventdatabyname); err != nil {
		c.JSON(http.StatusBadRequest, "bad Format")
		return
	}

	err2 := json.Unmarshal(content, &eventdatabyname)
	if err2 != nil {
		log.Fatal("Error during Unmarshal: ", err2)
		return
	}

	c.JSON(http.StatusOK, eventdatabyname)

}

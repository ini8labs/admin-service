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
	r.GET("/api/v1/user_data", UserData)
	r.GET("/api/v1/event_data", EventData)
	r.POST("api/v1/event_data_byName", EventDataByName)

	return r.Run(addr)

}

type User struct {
	Id            string `json:"userid"`
	Name          string `json:"name"`
	ContactNumber string `json:"number"`
	IdProofNumber string `json:"idproof"`
}

func UserData(c *gin.Context) {

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

func EventData(c *gin.Context) {

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

type EventDatabyName struct {
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
	EventName string `json:"eventname,omitempty"`
}

func EventDataByName(c *gin.Context) {

	var requestdata RequestEventDataByName
	var eventdatabyname []EventDatabyName
	//var eventdatabyname map[string]interface{}

	content, err := os.ReadFile(".\\src\\apis\\eventinfo.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return
	}

	if err := c.BindJSON(&requestdata); err != nil {
		c.JSON(http.StatusBadRequest, "bad Format")
		return
	}

	err2 := json.Unmarshal(content, &eventdatabyname)
	if err2 != nil {
		log.Fatal("Error during Unmarshal: ", err2)
		return
	}

	var expectedData []EventDatabyName
	for i := 0; i < len(eventdatabyname); i++ {
		if eventdatabyname[i].EventName == requestdata.EventName {
			expectedData = append(expectedData, eventdatabyname[i])
		}
	}

	c.JSON(http.StatusOK, expectedData)
}

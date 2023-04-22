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
	r.POST("/api/v1/event_data_byId", EventDataById)

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

type RequestEventDataByName struct {
	EventName string `json:"eventname,omitempty"`
}

func EventDataByName(c *gin.Context) {

	var requestdata RequestEventDataByName
	var eventdatabyname []Event

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

	var expectedData []Event
	for i := 0; i < len(eventdatabyname); i++ {
		if eventdatabyname[i].EventName == requestdata.EventName {
			expectedData = append(expectedData, eventdatabyname[i])
		}
	}

	c.JSON(http.StatusOK, expectedData)
}

type RequestEventDataById struct {
	EventId string `json:"eventid,omitempty"`
}

func EventDataById(c *gin.Context) {

	var requestdata RequestEventDataById
	var eventdatabyid []Event

	content, err := os.ReadFile(".\\src\\apis\\eventinfo.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return
	}

	if err := c.BindJSON(&requestdata); err != nil {
		c.JSON(http.StatusBadRequest, "bad Format")
		return
	}

	err2 := json.Unmarshal(content, &eventdatabyid)
	if err2 != nil {
		log.Fatal("Error during Unmarshal: ", err2)
		return
	}

	var expectedData []Event
	for i := 0; i < len(eventdatabyid); i++ {
		if eventdatabyid[i].EventId == requestdata.EventId {
			expectedData = append(expectedData, eventdatabyid[i])
		}
	}

	c.JSON(http.StatusOK, expectedData)
}

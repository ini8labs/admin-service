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
	//r.POST("api/v1/winning_number" WinningNumber)
	r.GET("/api/v1/get_event_data", GetEventData)

	return r.Run(addr)

}

type User struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	ContactNumber string `json:"number"`
	IdProofNumber string `json:"idproof"`
}

func GetUserData(c *gin.Context) {

	content, err := os.ReadFile(".\\src\\apis\\users.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return
	}

	var users []User

	err2 := json.Unmarshal(content, &users)
	if err2 != nil {
		log.Fatal("Error during Unmarshal(): ", err2)
		return
	}
	c.JSON(http.StatusOK, users)
}

type Event struct {
	EventName string `json:"eventname"`
	EventId   string `json:"eventid"`
	Date      string `json:"date"`
	WinNumber []int  `json:"winnumber"`
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
		log.Fatal("Error during Unmarshal(): ", err2)
		return
	}
	c.JSON(http.StatusOK, events)
}

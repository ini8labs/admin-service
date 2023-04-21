package apis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewServer(addr string, log *logrus.Logger) error {

	r := gin.Default()

	// API end point
	r.GET("/api/v1/get_data", GetUserData)
	//r.POST("api/v1/winning_number" WinningNumber)
	//r.POST("/api/v1/get_data_eventId", GetDataByEventId)

	return r.Run(addr)

}

type User struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	ContactNumber string `json:"number"`
	IdProofNumber string `json:"idproof"`
}

func GetUserData(c *gin.Context) {

	content, err := ioutil.ReadFile(".\\src\\apis\\users.json")
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

// func GetDataByEventId(c *gin.Context) {

// }

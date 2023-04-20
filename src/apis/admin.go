package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewServer(addr string, log *logrus.Logger) error {

	r := gin.Default()

	// API end point
	r.GET("/api/v1/get_data", GetUserData)
	//r.POST("api/v1/winning_number" WinningNumber)

	return r.Run(addr)

}

type winningNumber struct {
	EventId string `json:"eventId"`
	Numbers []int  `json:"numbers"`
}

type User struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	ContactNumber string `json:"number"`
	IdProofNumber string `json:"idproof"`
}

// type Users struct {
// 	Users []User `json:"users"`
// }

func GetUserData(c *gin.Context) {

	content, err := os.ReadFile(".\\src\\apis\\users.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var user User
	err = json.Unmarshal(content, &user)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	fmt.Println(user.ContactNumber)
	fmt.Println(user.Name)
	c.JSON(http.StatusOK, user)
}

func WinningNumber(c *gin.Context) {
}

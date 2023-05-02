package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
)

func (s Server) addWinner(c *gin.Context) {
	allEventsInfoResp, err := s.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error1")
		s.Logger.Error(err.Error())
	}
	fmt.Println(allEventsInfoResp)
	for i := 0; i < len(allEventsInfoResp); i++ {
		userInfoByEventIdResp, err := s.GetParticipantsInfoByEventID(allEventsInfoResp[i].EventUID)
		fmt.Println(userInfoByEventIdResp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal Server Error2")
			s.Logger.Error(err.Error())
		}
		fmt.Println("#######################")

		count := 0
		for j := 0; j < len(userInfoByEventIdResp); i++ {
			for k := 0; k < len(userInfoByEventIdResp[j].BetNumbers); k++ {
				for m := 0; m < len(allEventsInfoResp[i].WinningNumber); m++ {
					if userInfoByEventIdResp[j].BetNumbers[k] == allEventsInfoResp[i].WinningNumber[m] {
						count++
					}
				}
			}
			winAmount := 0
			switch {
			case count == 1:
				winAmount = userInfoByEventIdResp[j].Amount * 40
			case count == 2:
				winAmount = userInfoByEventIdResp[j].Amount * 240
			case count == 3:
				winAmount = userInfoByEventIdResp[j].Amount * 2100
			case count == 4:
				winAmount = userInfoByEventIdResp[j].Amount * 6000
			case count == 5:
				winAmount = userInfoByEventIdResp[j].Amount * 44000
			}
			addWinner := lsdb.WinnerInfo{
				UserID:    userInfoByEventIdResp[j].UserID,
				EventID:   userInfoByEventIdResp[j].EventUID,
				WinType:   allEventsInfoResp[i].EventType,
				AmountWon: winAmount,
			}

			if err := s.Client.AddNewWinner(addWinner); err != nil {
				c.JSON(http.StatusInternalServerError, "something is wrong with the server")
				s.Logger.Error(err.Error())
				return
			}
			c.JSON(http.StatusOK, "Winner added successfully")
		}
	}
}

func (s Server) getEventWinners(c *gin.Context) {

}

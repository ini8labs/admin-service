package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s Server) GetUserInfoByPhone(c *gin.Context) {

	phonenumber := c.Query("phone")
	userphonenumber, _ := strconv.Atoi(phonenumber)

	resp, err := s.Client.GetUserInfoByPhone(int64(userphonenumber))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}

	var userinfo UserInfo
	if resp.Phone == int64(userphonenumber) {
		userinfo.Name = resp.Name
		userinfo.UID = PrimitiveToString(resp.UID)
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	c.JSON(http.StatusOK, userinfo)
}

func (s Server) GetUserInfoByGovID(c *gin.Context) {

	govid := c.Query("id")

	resp, err := s.Client.GetUserInfoByGovID(govid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	var userinfo UserInfo
	if resp.GovID == govid {
		stringUserID := PrimitiveToString(resp.UID)

		userinfo.Name = resp.Name
		userinfo.UID = stringUserID
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	c.JSON(http.StatusOK, userinfo)
}

func (s Server) GetUserInfoByID(c *gin.Context) {

	uid := c.Query("uid")

	resp, err := s.Client.GetUserInfoByID(StringToPrimitive(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}
	var userinfo UserInfo
	if resp.UID == StringToPrimitive(uid) {

		mongouserId := resp.UID
		stringUserID := mongouserId.Hex()

		userinfo.Name = resp.Name
		userinfo.UID = stringUserID
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	c.JSON(http.StatusOK, userinfo)

}

func (s Server) GetParticipantsInfoByEventID(c *gin.Context) {
	eventid := c.Query("eventid")

	resp, err := s.Client.GetParticipantsInfoByEventID(StringToPrimitive(eventid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}

	var arr []UserInfoByEventId

	for i := 0; i < len(resp); i++ {
		var userinfobyevent UserInfoByEventId

		if resp[i].EventUID == StringToPrimitive(eventid) {
			userinfobyevent.UserID = PrimitiveToString(resp[i].UserID)
			userinfobyevent.EventUID = PrimitiveToString(resp[i].EventUID)
			userinfobyevent.BetUID = PrimitiveToString(resp[i].BetUID)
			userinfobyevent.Amount = resp[i].Amount
			userinfobyevent.BetNumbers = resp[i].BetNumbers

			resp2, err := s.Client.GetUserInfoByID(resp[i].UserID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "something is wrong with the server")
				s.Logger.Error(err.Error())
				return
			}

			if userinfobyevent.UserID == PrimitiveToString(resp2.UID) {
				userinfobyevent.UserName = resp2.Name
			}
			//stringUserID := mongouserId.Hex()
			//userinfobyevent.UserID = stringUserID
			// resp3, err3 := s.Client.GetUserInfoByID(eventinfo.UserID)
			// if err3 != nil {
			// 	c.JSON(http.StatusInternalServerError, "something is wrong with the server")
			// 	logrus.Infoln(err3)
			// 	return
			// }
			arr = append(arr, userinfobyevent)
		}
	}
	c.JSON(http.StatusOK, arr)

}

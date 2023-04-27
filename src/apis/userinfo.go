package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
)

func initializeUserInfo(resp *lsdb.UserInfo, str string) UserInfo {
	var userinfo UserInfo
	if resp.GovID == str {

		userinfo.Name = resp.Name
		userinfo.UID = primitiveToString(resp.UID)
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	return userinfo
}

func initializeUserInfoByPhone(resp *lsdb.UserInfo, num int) UserInfo {
	var userinfo UserInfo
	if resp.Phone == int64(num) {

		userinfo.Name = resp.Name
		userinfo.UID = primitiveToString(resp.UID)
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	return userinfo
}
func initializeUserInfoById(resp *lsdb.UserInfo, str string) UserInfo {
	var userinfo UserInfo
	id := stringToPrimitive(str)
	if resp.UID == id {

		userinfo.Name = resp.Name
		userinfo.UID = primitiveToString(resp.UID)
		userinfo.Phone = resp.Phone
		userinfo.GovID = resp.GovID
		userinfo.EMail = resp.EMail
	}
	return userinfo
}

func (s Server) initializeUserInfobByEventId(resp []lsdb.EventParticipantInfo, str string, c *gin.Context) []UserInfoByEventId {
	var arr []UserInfoByEventId

	for i := 0; i < len(resp); i++ {
		var userinfobyevent UserInfoByEventId

		if resp[i].EventUID == stringToPrimitive(str) {
			userinfobyevent.UserID = primitiveToString(resp[i].UserID)
			userinfobyevent.EventUID = primitiveToString(resp[i].EventUID)
			userinfobyevent.BetUID = primitiveToString(resp[i].BetUID)
			userinfobyevent.Amount = resp[i].Amount
			userinfobyevent.BetNumbers = resp[i].BetNumbers

			resp2, err := s.Client.GetUserInfoByID(resp[i].UserID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "something is wrong with the server")
				s.Logger.Error(err.Error())
			}
			if userinfobyevent.UserID == primitiveToString(resp2.UID) {
				userinfobyevent.UserName = resp2.Name
				userinfobyevent.PhoneNumber = resp2.Phone
			}
			//stringUserID := mongouserId.Hex()
			//userinfobyevent.UserID = stringUserID
			// resp3, err := s.Client.GetUserInfoByID(eventinfo.UserID)
			// if err3 != nil {
			// 	c.JSON(http.StatusInternalServerError, "something is wrong with the server")
			// 	s.Server.Error(err.Error())
			// 	return
			// }
			arr = append(arr, userinfobyevent)
		}
	}
	return arr
}

func (s Server) UserInfo(c *gin.Context) {
	phonenumber, exists1 := c.GetQuery("phone")
	uid, exists2 := c.GetQuery("uid")
	govid, exists3 := c.GetQuery("govid")
	eventid, exists4 := c.GetQuery("eventid")

	if exists1 {
		userphonenumber, _ := strconv.Atoi(phonenumber)

		resp, err := s.Client.GetUserInfoByPhone(int64(userphonenumber))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "something is wrong with the server")
			s.Logger.Error(err.Error())
			return
		}

		userInfo := initializeUserInfoByPhone(resp, userphonenumber)

		c.JSON(http.StatusOK, userInfo)
	}

	if exists2 {
		resp, err := s.Client.GetUserInfoByID(stringToPrimitive(uid))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "something is wrong with the server")
			s.Logger.Error(err.Error())
			return
		}

		userInfo := initializeUserInfoById(resp, uid)
		c.JSON(http.StatusOK, userInfo)
	}

	if exists3 {
		resp, err := s.Client.GetUserInfoByGovID(govid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "something is wrong with the server")
			s.Logger.Error(err.Error())
			return
		}
		userInfo := initializeUserInfo(resp, govid)
		c.JSON(http.StatusOK, userInfo)
	}

	if exists4 {
		resp, err := s.Client.GetParticipantsInfoByEventID(stringToPrimitive(eventid))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "something is wrong with the server")
			s.Logger.Error(err.Error())
			return
		}

		result := s.initializeUserInfobByEventId(resp, eventid, c)
		c.JSON(http.StatusOK, result)
	}
}

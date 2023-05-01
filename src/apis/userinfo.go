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
		userinfo = UserInfo{
			Name:  resp.Name,
			UID:   primitiveToString(resp.UID),
			Phone: resp.Phone,
			GovID: resp.GovID,
			EMail: resp.EMail,
		}
	}

	return userinfo
}

func (s Server) userInfo(c *gin.Context) {

	phonenumber := c.Query("phone")
	uid := c.Query("uid")
	govid := c.Query("govid")

	userInfo, err := s.getUserByQueryParams(phonenumber, uid, govid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Probelm with the server")
		return
	}

	c.JSON(http.StatusOK, userInfo)

}

func (s Server) getUserInfoByPhone(phoneNumber string) (UserInfo, error) {
	userPhoneNumber, _ := strconv.Atoi(phoneNumber)

	resp, err := s.Client.GetUserInfoByPhone(int64(userPhoneNumber))
	if err != nil {
		return UserInfo{}, err
	}

	userInfo := initializeUserInfoByPhone(resp, userPhoneNumber)
	return userInfo, nil

}

func (s Server) getUserInfoByGovID(govId string) (UserInfo, error) {
	resp, err := s.Client.GetUserInfoByGovID(govId)
	if err != nil {
		return UserInfo{}, err
	}
	userInfo := initializeUserInfo(resp, govId)
	return userInfo, nil
}

func (s Server) getUserInfoByUID(uid string) (UserInfo, error) {
	resp, err := s.Client.GetUserInfoByID(stringToPrimitive(uid))
	if err != nil {
		return UserInfo{}, err
	}

	userInfo := initializeUserInfoById(resp, uid)
	return userInfo, nil
}

func (s Server) initializeUserInfobByEventId(resp []lsdb.EventParticipantInfo, str string) ([]UserInfoByEventId, error) {
	var arr []UserInfoByEventId

	for i := 0; i < len(resp); i++ {
		var userinfobyevent UserInfoByEventId

		if resp[i].EventUID == stringToPrimitive(str) {
			userinfobyevent = UserInfoByEventId{
				UserID:     primitiveToString(resp[i].UserID),
				EventUID:   primitiveToString(resp[i].EventUID),
				BetUID:     primitiveToString(resp[i].BetUID),
				Amount:     resp[i].Amount,
				BetNumbers: resp[i].BetNumbers,
			}

			resp2, err := s.Client.GetUserInfoByID(resp[i].UserID)
			if err != nil {
				return []UserInfoByEventId{}, err
			}

			if userinfobyevent.UserID == primitiveToString(resp2.UID) {
				userinfobyevent.UserName = resp2.Name
				userinfobyevent.PhoneNumber = resp2.Phone
			}

			arr = append(arr, userinfobyevent)
		}
	}
	return arr, nil
}

func (s Server) getUserByQueryParams(phonenumber, uid, govid string) (UserInfo, error) {
	var userInfo UserInfo
	var err error

	switch {
	case phonenumber != "":
		userInfo, err = s.getUserInfoByPhone(phonenumber)
	case uid != "":
		userInfo, err = s.getUserInfoByUID(uid)
	case govid != "":
		userInfo, err = s.getUserInfoByGovID(govid)
	}

	return userInfo, err
}

func (s Server) userInfoByEventId(c *gin.Context) {

	eventId := c.Query("eventId")
	resp, err := s.Client.GetParticipantsInfoByEventID(stringToPrimitive(eventId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}

	result, _ := s.initializeUserInfobByEventId(resp, eventId)
	c.JSON(http.StatusOK, result)
}

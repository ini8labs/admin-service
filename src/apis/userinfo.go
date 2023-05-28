package apis

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ini8labs/lsdb"
)

func initializeUserInfoByGovId(userInfo *lsdb.UserInfo, govId string) UserInfo {
	var userinfo UserInfo

	if userInfo.GovID == govId {
		userinfo.Name = userInfo.Name
		userinfo.UID = primitiveToString(userInfo.UID)
		userinfo.Phone = userInfo.Phone
		userinfo.GovID = userInfo.GovID
		userinfo.EMail = userInfo.EMail
	}

	return userinfo
}

func initializeUserInfoByPhone(userInfo *lsdb.UserInfo, phoneNumber int) UserInfo {
	var userinfo UserInfo

	if userInfo.Phone == int64(phoneNumber) {
		userinfo.Name = userInfo.Name
		userinfo.UID = primitiveToString(userInfo.UID)
		userinfo.Phone = userInfo.Phone
		userinfo.GovID = userInfo.GovID
		userinfo.EMail = userInfo.EMail
	}

	return userinfo
}
func initializeUserInfoByUserId(userInfo *lsdb.UserInfo, userId string) UserInfo {
	var userinfo UserInfo
	id := stringToPrimitive(userId)

	if userInfo.UID == id {
		userinfo = UserInfo{
			Name:  userInfo.Name,
			UID:   primitiveToString(userInfo.UID),
			Phone: userInfo.Phone,
			GovID: userInfo.GovID,
			EMail: userInfo.EMail,
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
		c.JSON(http.StatusBadRequest, err.Error())
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
	if resp == nil {
		err = errors.New("invalid phone number")
		s.Logger.Error(err.Error())
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

	if resp == nil {
		err = errors.New("invalid govt id")
		s.Logger.Error(err.Error())
		return UserInfo{}, err
	}
	userInfo := initializeUserInfoByGovId(resp, govId)
	return userInfo, nil
}

func (s Server) getUserInfoByUID(uid string) (UserInfo, error) {
	resp, err := s.Client.GetUserInfoByID(stringToPrimitive(uid))
	if err != nil {
		return UserInfo{}, err
	}

	if resp == nil {
		err = errors.New("invalid user id")
		s.Logger.Error(err.Error())
		return UserInfo{}, err
	}

	userInfo := initializeUserInfoByUserId(resp, uid)
	return userInfo, nil
}

func (s Server) InitializeUserInfobByEventId(eventParticipantInfo []lsdb.EventParticipantInfo, eventId string) ([]UserInfoByEventId, error) {
	var userInfoArr []UserInfoByEventId

	for i := 0; i < len(eventParticipantInfo); i++ {
		var userinfobyevent UserInfoByEventId

		if eventParticipantInfo[i].EventUID == stringToPrimitive(eventId) {
			userinfobyevent = UserInfoByEventId{
				UserID:     primitiveToString(eventParticipantInfo[i].UserID),
				EventUID:   primitiveToString(eventParticipantInfo[i].EventUID),
				BetUID:     primitiveToString(eventParticipantInfo[i].BetUID),
				Amount:     eventParticipantInfo[i].Amount,
				BetNumbers: eventParticipantInfo[i].BetNumbers,
			}

			resp2, err := s.Client.GetUserInfoByID(eventParticipantInfo[i].UserID)
			if err != nil {
				return []UserInfoByEventId{}, err
			}

			if userinfobyevent.UserID == primitiveToString(resp2.UID) {
				userinfobyevent.UserName = resp2.Name
				userinfobyevent.PhoneNumber = resp2.Phone
			}

			userInfoArr = append(userInfoArr, userinfobyevent)
		}
	}
	return userInfoArr, nil
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
	validation, _ := s.validateEventId(eventId)
	if !validation {
		c.JSON(http.StatusBadRequest, "EventId does not exist")
		s.Logger.Error("invalid event id")
		return
	}

	resp, err := s.Client.GetParticipantsInfoByEventID(stringToPrimitive(eventId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err.Error())
		return
	}

	result, _ := s.InitializeUserInfobByEventId(resp, eventId)
	c.JSON(http.StatusOK, result)
}

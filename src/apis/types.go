package apis

import (
	"github.com/ini8labs/lsdb"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*logrus.Logger
	*lsdb.Client
	Addr string
}

type UserInfoByEventId struct {
	EventUID    string `json:"event_id,omitempty"`
	EventDate   Date   `json:"event_date,omitempty"`
	UserName    string `json:"username,omitempty"`
	EventName   string `json:"eventname,omitempty"`
	EventType   string `json:"event_type,omitempty"`
	BetUID      string `json:"betid,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	BetNumbers  []int  `json:"bet_numbers,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	PhoneNumber int64  `json:"phone,omitempty"`
}

type EventsInfo struct {
	EventUID      string `json:"event_id,omitempty"`
	EventDate     Date   `json:"event_date,omitempty"`
	EventName     string `json:"name,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	WinningNumber []int  `json:"winning_number,omitempty"`
}

type UserInfo struct {
	UID   string `json:"user_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone int64  `json:"phone,omitempty"`
	GovID string `json:"gov_id,omitempty"`
	EMail string `json:"e_mail,omitempty"`
}

type Date struct {
	Day   int `json:"day,omitempty"`
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}

type AddNewEventReq struct {
	EventDate     Date   `json:"event_date"`
	Name          string `json:"name,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	WinningNumber []int  `json:"win_number,omitempty"`
}

type GetEventsByDate struct {
	EventDate Date `json:"event_date"`
	StartDate Date `json:"start_date"`
	EndDate   Date `json:"end_date"`
}

type AddWinner struct {
	UserID    string `json:"user_id"`
	EventUID  string `json:"event_id"`
	AmountWon int    `json:"amountWon"`
	Name      string `json:"user_name"`
	Phone     int64  `json:"phone"`
	EMail     string `json:"e_mail"`
	EventDate Date   `json:"event_date"`
	WinType   string `json:"winType"`
}

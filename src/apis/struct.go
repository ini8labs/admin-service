package apis

import (
	"time"

	"github.com/ini8labs/lsdb"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*logrus.Logger
	*lsdb.Client
	Addr string
}

type UserInfoByEventId struct {
	EventUID   string    `json:"_id,omitempty"`
	EventDate  time.Time `json:"event_date,omitempty"`
	UserName   string    `json:"username,omitempty"`
	EventName  string    `json:"eventname,omitempty"`
	EventType  string    `json:"event_type,omitempty"`
	BetUID     string    `json:"betid,omitempty"`
	UserID     string    `json:"user_id,omitempty"`
	BetNumbers []int     `json:"bet_numbers,omitempty"`
	Amount     int       `json:"amount,omitempty"`
}

type EventsInfo struct {
	EventUID      string    `json:"_id,omitempty"`
	EventDate     time.Time `json:"event_date,omitempty"`
	EventName     string    `json:"name,omitempty"`
	EventType     string    `json:"event_type,omitempty"`
	WinningNumber int       `json:"winning_number,omitempty"`
}

type UserInfo struct {
	UID   string `json:"_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone int64  `json:"phone,omitempty"`
	GovID string `json:"gov_id,omitempty"`
	EMail string `json:"e_mail,omitempty"`
}

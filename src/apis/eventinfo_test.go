package apis

import (
	"reflect"
	"testing"

	"github.com/ini8labs/lsdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_initializeEventInfo(t *testing.T) {
	type args struct {
		lottteryEventInfo []lsdb.LotteryEventInfo
	}
	tests := []struct {
		name string
		args args
		want []EventsInfo
	}{
		{
			"VALID INPUT",
			args{
				[]lsdb.LotteryEventInfo{
					{
						EventUID:  stringToPrimitive("64521585398c724e121e7fa9"),
						EventDate: 1682899200000, Name: "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1683101061405,
						UpdatedAt:     1683101061405,
					},
					{
						EventUID:      stringToPrimitive("6452862df3a29fa62420130c"),
						EventDate:     1682899200000,
						Name:          "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1683129901783,
						UpdatedAt:     1683129901783,
					},
					{
						EventUID:      stringToPrimitive("645287379b9f5da89aa4720a"),
						EventDate:     1682899200000,
						Name:          "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1683130167325,
						UpdatedAt:     1683130167325,
					},
				},
			},
			[]EventsInfo{
				{
					EventUID:      "64521585398c724e121e7fa9",
					EventDate:     Date{Day: 1, Month: 5, Year: 2023},
					EventName:     "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
				{
					EventUID:      "6452862df3a29fa62420130c",
					EventDate:     Date{Day: 1, Month: 5, Year: 2023},
					EventName:     "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
				{
					EventUID:      "645287379b9f5da89aa4720a",
					EventDate:     Date{Day: 1, Month: 5, Year: 2023},
					EventName:     "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeEventInfo(tt.args.lottteryEventInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeEventInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateDate(t *testing.T) {
	type args struct {
		newEvent AddNewEventReq
	}
	tests := []struct {
		name    string
		args    args
		want    primitive.DateTime
		wantErr bool
	}{
		{
			"valid input",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   30,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			1685404800000,
			false,
		},
		{
			"negative date",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   -17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			1681344000000,
			true,
		},
		{
			"invalid date",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   34,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			1685750400000,
			true,
		},
		{
			"past date",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   7,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			1683417600000,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateGenerateEventDate(tt.args.newEvent)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateWinningNumbers(t *testing.T) {
	type args struct {
		newEvent AddNewEventReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid input",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			false,
		},
		{
			"negative winning number",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{-34, 65, 78, 3, 4},
				},
			},
			true,
		},
		{
			"winning number greater than 90",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{91, 65, 78, 3, 4},
				},
			},
			true,
		},
		{
			"repeating winning number",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{65, 65, 78, 3, 4},
				},
			},
			true,
		},
		{
			"more than 5 winning numbers",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{65, 89, 78, 3, 4, 68},
				},
			},
			true,
		},
		{
			"less than 5 winning numbers",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{65, 78, 3, 4},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateGenerateEventWinningNumbers(tt.args.newEvent); (err != nil) != tt.wantErr {
				t.Errorf("validateWinningNumbers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateEventNameAndType(t *testing.T) {
	type args struct {
		newEvent AddNewEventReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid input",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			false,
		},
		{
			"invalid event type",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MSS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			true,
		},
		{
			"event name and event type does not match",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "FB",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			true,
		},
		{
			"event name does not exist",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "MondaySpecial",
					EventType:     "MS",
					WinningNumber: []int{34, 65, 78, 3, 4},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateGenerateEventNameAndType(tt.args.newEvent); (err != nil) != tt.wantErr {
				t.Errorf("validateEventNameAndType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_initializeGenerateEventInfo(t *testing.T) {
	type args struct {
		newEvent AddNewEventReq
		date     primitive.DateTime
	}
	tests := []struct {
		name string
		args args
		want lsdb.LotteryEventInfo
	}{
		{
			"valid input",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{9, 65, 78, 3, 4},
				},
				1684281600000,
			},
			lsdb.LotteryEventInfo{
				EventUID:      stringToPrimitive("000000000000000000000000"),
				EventDate:     1684281600000,
				Name:          "Monday Special",
				EventType:     "MS",
				WinningNumber: []int{9, 65, 78, 3, 4},
				CreatedAt:     0,
				UpdatedAt:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeGenerateEventInfo(tt.args.newEvent, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeGenerateEventInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateAddEvent(t *testing.T) {
	type args struct {
		newEvent AddNewEventReq
	}
	tests := []struct {
		name    string
		args    args
		want    lsdb.LotteryEventInfo
		wantErr bool
	}{
		{
			"valid input",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   30,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{9, 65, 78, 3, 4},
				},
			},
			lsdb.LotteryEventInfo{
				EventUID:      stringToPrimitive("000000000000000000000000"),
				EventDate:     1685404800000,
				Name:          "Monday Special",
				EventType:     "MS",
				WinningNumber: []int{9, 65, 78, 3, 4},
				CreatedAt:     0,
				UpdatedAt:     0,
			},
			false,
		},
		{
			"invalid date",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   -17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{9, 65, 78, 3, 4},
				},
			},
			lsdb.LotteryEventInfo{},
			true,
		},
		{
			"invalid event name",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "MondaySpecial",
					EventType:     "MS",
					WinningNumber: []int{9, 65, 78, 3, 4},
				},
			},
			lsdb.LotteryEventInfo{},
			true,
		},
		{
			"invalid winning numbers",
			args{
				AddNewEventReq{
					EventDate: Date{
						Day:   17,
						Month: 5,
						Year:  2023,
					},
					Name:          "Monday Special",
					EventType:     "MS",
					WinningNumber: []int{-9, 65, 78, 3, 4},
				},
			},
			lsdb.LotteryEventInfo{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateAddEvent(tt.args.newEvent)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAddEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateAddEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkEventId(t *testing.T) {
	type args struct {
		eventId   string
		eventInfo []lsdb.LotteryEventInfo
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"valid input",
			args{
				eventId: "646329ec8522d4bc895325af",
				eventInfo: []lsdb.LotteryEventInfo{
					{
						EventUID:  stringToPrimitive("646329ec8522d4bc895325af"),
						EventDate: 1684281600000, Name: "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1684220396144,
						UpdatedAt:     1684220396144,
					},
					{
						EventUID:      stringToPrimitive("64632a6a7d5099ed5e04d6cd"),
						EventDate:     1684281600000,
						Name:          "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1684220522477,
						UpdatedAt:     1684220522477},
					{
						EventUID:      stringToPrimitive("64633927af0f4ac829285889"),
						EventDate:     1684281600000,
						Name:          "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1684224295875,
						UpdatedAt:     1684224295875,
					},
				},
			},
			true,
		},
		{
			"valid input",
			args{
				eventId: "2148584553543trtd",
				eventInfo: []lsdb.LotteryEventInfo{
					{
						EventUID:      stringToPrimitive("646329ec8522d4bc895325af"),
						EventDate:     1684281600000,
						Name:          "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1684220396144,
						UpdatedAt:     1684220396144,
					},
					{
						EventUID:      stringToPrimitive("64632a6a7d5099ed5e04d6cd"),
						EventDate:     1684281600000,
						Name:          "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1684220522477,
						UpdatedAt:     1684220522477,
					},
					{
						EventUID:      stringToPrimitive("64633927af0f4ac829285889"),
						EventDate:     1684281600000,
						Name:          "Monday Special",
						EventType:     "MS",
						WinningNumber: []int{34, 65, 78, 3, 4},
						CreatedAt:     1684224295875,
						UpdatedAt:     1684224295875,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEventId(tt.args.eventId, tt.args.eventInfo); got != tt.want {
				t.Errorf("checkEventId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validaEventNameAndType(t *testing.T) {
	type args struct {
		eventType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid input",
			args{eventType: "MS"},
			false,
		},
		{
			"invalid input",
			args{eventType: "MSS"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validaEventNameAndType(tt.args.eventType); (err != nil) != tt.wantErr {
				t.Errorf("validaEventNameAndType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

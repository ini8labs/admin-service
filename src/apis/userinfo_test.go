package apis

import (
	"reflect"
	"testing"

	"github.com/ini8labs/lsdb"
)

func Test_initializeUserInfoByGovId(t *testing.T) {
	type args struct {
		userInfo *lsdb.UserInfo
		govId    string
	}
	tests := []struct {
		name string
		args args
		want UserInfo
	}{
		{
			"valid input",
			args{
				&lsdb.UserInfo{
					UID:       stringToPrimitive("64457b80630b6fef47225939"),
					Name:      "Anand",
					Phone:     7506639417,
					GovID:     "ABCDEFG",
					EMail:     "anand@ini8labs.tech",
					CreatedAT: "2023-04-24T00:10:00+05:30",
					UpdatedAT: "2023-04-24T00:10:00+05:30",
				},
				"ABCDEFG",
			},
			UserInfo{
				UID:   "64457b80630b6fef47225939",
				Name:  "Anand",
				Phone: 7506639417,
				GovID: "ABCDEFG",
				EMail: "anand@ini8labs.tech",
			},
		},
		{
			"invalid input",
			args{
				&lsdb.UserInfo{},
				"InvalidGovId",
			},
			UserInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeUserInfoByGovId(tt.args.userInfo, tt.args.govId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeUserInfoByGovId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initializeUserInfoByPhone(t *testing.T) {
	type args struct {
		userInfo    *lsdb.UserInfo
		phoneNumber int
	}
	tests := []struct {
		name string
		args args
		want UserInfo
	}{
		{
			"valid input",
			args{
				&lsdb.UserInfo{
					UID:       stringToPrimitive("64457b80630b6fef47225939"),
					Name:      "Anand",
					Phone:     7506639417,
					GovID:     "ABCDEFG",
					EMail:     "anand@ini8labs.tech",
					CreatedAT: "2023-04-24T00:10:00+05:30",
					UpdatedAT: "2023-04-24T00:10:00+05:30",
				},
				7506639417,
			},
			UserInfo{
				UID:   "64457b80630b6fef47225939",
				Name:  "Anand",
				Phone: 7506639417,
				GovID: "ABCDEFG",
				EMail: "anand@ini8labs.tech",
			},
		},
		{
			"invalid input",
			args{
				&lsdb.UserInfo{},
				121212,
			},
			UserInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeUserInfoByPhone(tt.args.userInfo, tt.args.phoneNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeUserInfoByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initializeUserInfoByUserId(t *testing.T) {
	type args struct {
		userInfo *lsdb.UserInfo
		userId   string
	}
	tests := []struct {
		name string
		args args
		want UserInfo
	}{
		{
			"valid input",
			args{
				&lsdb.UserInfo{
					UID:       stringToPrimitive("64457b80630b6fef47225939"),
					Name:      "Anand",
					Phone:     7506639417,
					GovID:     "ABCDEFG",
					EMail:     "anand@ini8labs.tech",
					CreatedAT: "2023-04-24T00:10:00+05:30",
					UpdatedAT: "2023-04-24T00:10:00+05:30",
				},
				"64457b80630b6fef47225939",
			},
			UserInfo{
				UID:   "64457b80630b6fef47225939",
				Name:  "Anand",
				Phone: 7506639417,
				GovID: "ABCDEFG",
				EMail: "anand@ini8labs.tech",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeUserInfoByUserId(tt.args.userInfo, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeUserInfoByUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

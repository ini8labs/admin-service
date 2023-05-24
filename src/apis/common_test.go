package apis

import (
	"reflect"
	"testing"
)

func Test_daysInMonth(t *testing.T) {
	type args struct {
		month int
		year  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"days in may",
			args{
				month: 5,
				year:  2023,
			},
			31,
		},
		{
			"valid june",
			args{
				month: 6,
				year:  2023,
			},
			30,
		},
		{
			"days in feb in non leap year",
			args{
				month: 2,
				year:  2023,
			},
			28,
		},
		{
			"days in feb in leap year",
			args{
				month: 2,
				year:  2020,
			},
			29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := daysInMonth(tt.args.month, tt.args.year); got != tt.want {
				t.Errorf("daysInMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertStringToDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    Date
		wantErr bool
	}{
		{
			"valid date",
			args{"2023-05-17"},
			Date{Year: 2023, Month: 5, Day: 17},
			false,
		},
		{
			"invalid date",
			args{"02-2023-05-17"},
			Date{Year: 0, Month: 0, Day: 0},
			true,
		},
		{
			"invalid year",
			args{"ab2023-05-17"},
			Date{Year: 0, Month: 0, Day: 0},
			true,
		},
		{
			"invalid month",
			args{"2023-ab05-17"},
			Date{Year: 0, Month: 0, Day: 0},
			true,
		},
		{
			"invalid day",
			args{"2023-05-ab17"},
			Date{Year: 0, Month: 0, Day: 0},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertStringToDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertStringToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertStringToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

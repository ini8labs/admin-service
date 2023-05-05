package apis

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func primitiveToString(p primitive.ObjectID) string {
	return p.Hex()
}

func stringToPrimitive(s string) primitive.ObjectID {
	var a, _ = primitive.ObjectIDFromHex(s)
	return a
}

func convertTimeToPrimitive(date Date) primitive.DateTime {

	d := time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, time.UTC)

	return primitive.NewDateTimeFromTime(d)

}

func convertPrimitiveToTime(date primitive.DateTime) Date {
	t := date.Time()

	return Date{
		Day:   t.Day(),
		Month: int(t.Month()),
		Year:  t.Year(),
	}
}

func daysInMonth(month, year int) int {
	switch time.Month(month) {
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) { // leap year
			return 29
		}
		return 28
	default:
		return 31
	}
}

func convertTimeToString(date time.Time) string {
	str := date.Format("2006-01-02")

	return str
}

func convertStringToDate(date string) (Date, error) {
	var d Date

	dateArr := strings.Split(date, "-")

	if len(dateArr) != 3 {
		return d, fmt.Errorf("invalid date")
	}

	intYear, err := strconv.Atoi(dateArr[0])
	if err != nil {
		return d, err
	}
	intMonth, err := strconv.Atoi(dateArr[1])
	if err != nil {
		return d, err
	}
	intDay, err := strconv.Atoi(dateArr[2])
	if err != nil {
		return d, err
	}

	d = Date{
		Year:  intYear,
		Month: intMonth,
		Day:   intDay,
	}

	return d, nil
}

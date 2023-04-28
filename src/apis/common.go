package apis

import (
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

	d := time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, time.Local)

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

// func convertStringToInt(str string)int{
// 	a, err := strconv.Atoi(str)
// 	if err != nil{

// 	return a
// }

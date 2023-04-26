package apis

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PrimitiveToString(p primitive.ObjectID) string {
	return p.Hex()
}

func StringToPrimitive(s string) primitive.ObjectID {
	var a, _ = primitive.ObjectIDFromHex(s)
	return a
}

package model

import (
	"go.mongodb.org/mongo-driver/bson"
)

type GetOne struct {
	Base
	Data bson.M
}

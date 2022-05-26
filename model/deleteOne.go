package model

type DeleteOne struct {
	Base
	DeletedCount int64 `json:"deletedCount" bson:"deletedCount"`
}

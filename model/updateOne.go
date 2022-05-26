package model

type UpdateOne struct {
	Base
	MatchedCount  int64 `json:"matchedCount" bson:"matchedCount"`
	ModifiedCount int64 `json:"modifiedCount" bson:"modifiedCount"`
}

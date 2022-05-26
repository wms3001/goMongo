package model

type MCount struct {
	Base
	DocNum int64 `json:"docNum" bson:"docNum"`
}

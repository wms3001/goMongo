package model

type AddMany struct {
	Base
	InsertIds []string `json:"insertIds"`
}

package model

type MFilter struct {
	FilterMap map[string]interface{} `json:"filterMap" bson:"filterMap"`
	Type      string                 `json:"type" bson:"type"`
}

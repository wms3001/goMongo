package model

type MDistinct struct {
	Base
	Mfilds []interface{} `json:"mfilds" bson:"mfilds"`
}

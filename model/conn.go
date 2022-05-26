package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Conn struct {
	Base
	Client     *mongo.Client
	Collection *mongo.Collection
}

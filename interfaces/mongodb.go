package interfaces

import "go.mongodb.org/mongo-driver/mongo"

type Collection interface {
	GetCollection(name string) (*mongo.Collection, error)
}

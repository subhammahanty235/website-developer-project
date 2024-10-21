package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebsiteData struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	ProfileId        primitive.ObjectID `json:"profileId"`
	UserId           primitive.ObjectID `json:"userId"`
	Name             string             `json:"name"`
	DeploymentStatus bool               `json:"deploymentStatus"`
	CreatedAt        time.Time          `json:"created_at"`
}

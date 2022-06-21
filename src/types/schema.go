package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

type Image struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FlickrId     string             `bson:"flickId,omitempty" json:"flickId,omitempty"`
	Path         string             `bson:"path,omitempty" json:"path,omitempty"`
	Width        uint               `bson:"width,omitempty" json:"width,omitempty"`
	Height       uint               `bson:"height,omitempty" json:"height,omitempty"`
	Title        string             `bson:"title,omitempty" json:"title,omitempty"`
	Description  string             `bson:"description,omitempty" json:"description,omitempty"`
	License      string             `bson:"license,omitempty" json:"license,omitempty"`
	CreationDate *time.Time         `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	Tags         []Tag              `bson:"tags,omitempty" json:"tags,omitempty"`
}

type Tag struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	Origin       string             `bson:"origin,omitempty" json:"origin,omitempty"`
	CreationDate *time.Time         `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
}

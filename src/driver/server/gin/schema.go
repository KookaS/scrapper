package gin

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// elements here prevent cyclical imports in other folders

// BodyUpdateImage is the body for the update of an image.


type BodyUpdateImageTagsPull struct {
	Origin string             `bson:"origin,omitempty" json:"origin,omitempty"`
	ID     primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	Names  []string           `bson:"names,omitempty" json:"names,omitempty"`
}



type BodyImageCopy struct {
	Origin    string `bson:"origin,omitempty" json:"origin,omitempty"`
	OriginID  string `bson:"originID,omitempty" json:"originID,omitempty"`
	Name      string `bson:"name,omitempty" json:"name,omitempty"`
	Extension string `bson:"extension,omitempty" json:"extension,omitempty"`
}

type BodyTransferImage struct {
	OriginID string `bson:"originID,omitempty" json:"originID,omitempty"`
	From     string `bson:"from,omitempty" json:"from,omitempty"`
	To       string `bson:"to,omitempty" json:"to,omitempty"`
}

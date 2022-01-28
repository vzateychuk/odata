package meta

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Metadata struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	ValidFrom time.Time          `bson:"validFrom"`
	ValidTo   time.Time          `bson:"validTo"`
	Text      string             `bson:"text"`
	Meta      map[string]string  `bson:"meta"`
}

func NewMeta(text, countryCode string) *Metadata {

	meta := map[string]string{}
	meta["country"] = countryCode

	return &Metadata{
		ID:        primitive.NewObjectID(),
		ValidFrom: time.Now(),
		ValidTo:   time.Now().AddDate(100, 0, 0),
		Text:      text,
		Meta:      meta,
	}
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Type string

const (
	Url  Type = "url"
	Qr   Type = "qr"
	Both Type = "both "
)

type url struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Domain      string             `json:"domain" bson:"domain"`
	ShortCode   string             `json:"shortcode" bson:"shortcode"`
	UserId      string             `json:"user_id" bson:"user_id"`
	OriginalURL string             `json:"destination" bson:"destination"`
	Title       string             `json:"title" bson:"title"`
	Type        Type               `json:"type" bson:"type"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
	Expiration  time.Time          `bson:"expiration" json:"expiration"`
}

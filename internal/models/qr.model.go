package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QrStyle struct {
	Pattern     string `json:"pattern" bson:"pattern"`
	CornerStyle string `json:"cornerStyle" bson:"cornerStyle"`
	Foreground  string `json:"fgColor" bson:"fgColor" `
	Background  string `json:"bgColor" bson:"bgColor" `
}

type QR struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UrlId     primitive.ObjectID `json:"url_id" bson:"url_id"`
	UserId    string             `json:"user_id" bson:"user_id"`
	Style     QrStyle            `json:"style" bson:"style"`
	Image     string             `json:"qr_url" bson:"qr_url"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

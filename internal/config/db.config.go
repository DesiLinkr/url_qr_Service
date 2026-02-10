package config

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var urlCol *mongo.Collection
var qrCol *mongo.Collection

func init() {

	DB_Url := getEnv("DB_URL", "mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(DB_Url)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal("❌ db connection failed:", err)
	}

	log.Println("✅ db connected")
	urlCol = client.Database("Url_and_QR").Collection("url")
	qrCol = client.Database("Url_and_QR").Collection("qr")
}

func Url() *mongo.Collection {
	return urlCol
}

func Qr() *mongo.Collection {
	return qrCol
}

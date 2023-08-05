package models

import (
	"ipfs-pin-auth-server/connections"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Token struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Secret    string             `bson:"secret" json:"secret"`
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}

var TokensCollection *mongo.Collection = connections.OpenCollection(connections.Client, "tokens")

// func init() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	models := []mongo.IndexModel{
// 		{
// 			Keys:    bson.M{"createdAt": 1},
// 			Options: options.Index().SetExpireAfterSeconds(7 * 24 * 60 * 60),
// 		},
// 		{
// 			Keys:    bson.M{"accessToken": 1},
// 			Options: nil,
// 		},
// 	}
// 	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
// 	_, err := TokensCollection.Indexes().CreateMany(ctx, models, opts)

// 	// Check for the options errors
// 	if err != nil {
// 		fmt.Println("Indexes().CreateIndexes() ERROR:", err)
// 	} else {
// 		fmt.Println("CreateIndexes() opts:", opts)
// 	}
// 	defer cancel()
// }

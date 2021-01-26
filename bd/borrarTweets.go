package bd

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BorrarTweet .
func BorrarTweet(ID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("microblogging")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	// Pass these options to the Find method
	//findOptions := options.Find()
	//findOptions.SetLimit(2)
	_, err := col.DeleteOne(ctx, condicion)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

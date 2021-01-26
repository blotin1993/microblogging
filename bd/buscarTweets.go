package bd

import (
	"context"
	"log"
	"microblogging/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//BuscarTweets .
func BuscarTweets(ID string) ([]models.Tweet, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("microblogging")
	col := db.Collection("tweet")

	var tweets []models.Tweet
	//objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"userid": ID,
	}

	// Pass these options to the Find method
	//findOptions := options.Find()
	//findOptions.SetLimit(2)
	cur, err := col.Find(ctx, condicion)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem models.Tweet
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		tweets = append(tweets, elem)
	}
	// Close the cursor once finished
	cur.Close(ctx)

	return tweets, nil
}

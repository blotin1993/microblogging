package bd

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BorrarTweet .
func BorrarTweet(ID, userid string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("microblogging")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": userid,
	}

	// Pass these options to the Find method
	//findOptions := options.Find()
	//findOptions.SetLimit(2)

	result, err := col.DeleteOne(ctx, condicion)
	if err != nil || result.DeletedCount == 0 {
		err = errors.New("Ocurrió un error al intentar borrar el tweet")
	}

	return err
}

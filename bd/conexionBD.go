package bd

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:opeZRi2jdKIBXcmb@microblogging.p49ri.mongodb.net/test?authSource=admin&replicaSet=atlas-5u9fsz-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

/*ConectarBD es la función que me permite conectar a la base de datos
  Devuelve una conexión a la BD del tipo Mongo Client*/
func ConectarBD() *mongo.Client {
	//
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	//defer client.Disconnect(ctx)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la BD")
	return client
}

/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"microblogging/bd"
	"microblogging/models"
)

/*GraboTweet permite grabar el tweet en la BD */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet

	// decodificamos el body y armamos un registro
	err := json.NewDecoder(r.Body).Decode(&tweet)

	//
	tweet.UserID = IDUsuario
	tweet.Fecha = time.Now()
	// Para insertarlo en la base de datos necesitamos mapearlo a un bson
	_, status, err := bd.InsertoTweet(tweet)

	//si hay un error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, intentelo nuevamente.  "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

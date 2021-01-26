package routers

import (
	"microblogging/bd"
	"net/http"
)

//BorrarTweet .
func BorrarTweet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	err := bd.BorrarTweet(id)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet "+err.Error(), 400)
		return
	}
}

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
	err := bd.BorrarTweet(id, IDUsuario)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

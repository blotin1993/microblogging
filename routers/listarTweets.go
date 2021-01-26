package routers

import (
	"encoding/json"
	"microblogging/bd"
	"net/http"
)

//ListarTweets .
func ListarTweets(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	processedData, err := bd.BuscarTweets(id)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar listar los tweets "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "encoding/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(processedData)
}

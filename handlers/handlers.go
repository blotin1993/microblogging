package handlers

import (
	"log"
	"net/http"
	"os"

	"microblogging/middlew"

	"microblogging/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores : seteo mi puerto, el handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	//Endpoints ------------------------------------------------------------------------------------
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/listar_tweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListarTweets))).Methods("GET")
	router.HandleFunc("/borrar_tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.BorrarTweet))).Methods("DELETE")
	//-----------------------------------------------------------------------------------------------
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

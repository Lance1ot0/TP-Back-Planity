package main

import (
	"TP-Back-Planity/web/config"
	handlersMux "TP-Back-Planity/web/handlers"
	database "TP-Back-Planity/web/store"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("Erreur lors de la connexion à la base de données : %v", err)
	}

	defer db.Close()

	store := database.NewStore(db)

	mux := handlersMux.NewHandler(store)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"http://localhost:5173"}) // Remplacez cela par l'URL de votre application React en production
    methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

    err = http.ListenAndServe(":8081", handlers.CORS(headersOk, originsOk, methodsOk)(mux))
    if err != nil {
        log.Fatalf("Erreur lors du lancement du serveur : %v", err)
    }

}

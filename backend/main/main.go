package main

import (
	"TP-Back-Planity/web/config"
	handlers "TP-Back-Planity/web/handlers"
	database "TP-Back-Planity/web/store"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello from the backend!")

	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("Erreur lors de la connexion à la base de données : %v", err)
	}

	defer db.Close()

	store := database.NewStore(db)
	mux := handlers.NewHandler(store)

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
	}
}

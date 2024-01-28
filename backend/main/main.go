package main

import (
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	fmt.Println("test")

	conf := mysql.Config{
		User: 	"user",
		Passwd: "password",
		Net: 	"tcp",
		Addr: 	"localhost:3307",
		DBName: "mydb",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {

		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Test d'insertion
	db.QueryRow("INSERT INTO users (name) VALUES ('NomUtilisateur');")
	db.QueryRow("INSERT INTO Client (firstName, lastName, email, password) VALUES ('John', 'Doe', 'john.doe@example.com', 'motdepasse123');")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("on est co ww")
}
package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectToDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
	}

	conf := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_ADDRESS"),
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %v", err)
	}

	fmt.Println("Connecté à la base de données MySQL avec succès !")
	return db, nil
}

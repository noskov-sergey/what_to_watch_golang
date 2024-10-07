package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	http2 "github.ru/noskov-sergey/what_to_watch_golang/internal/delivery/http"
	oRep "github.ru/noskov-sergey/what_to_watch_golang/internal/repository/opinion"
	usecase "github.ru/noskov-sergey/what_to_watch_golang/internal/usecase/opinion"
)

func main() {
	cfg := flag.String("c", ".env", "config file path")
	flag.Parse()

	err := godotenv.Load(*cfg)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	defer db.Close()

	whatUC := usecase.New(oRep.NewOpinionRepository(db))

	service := http2.New(whatUC)

	log.Printf("server is running on %s", os.Getenv("BASE_URL"))
	err = http.ListenAndServe(os.Getenv("BASE_URL"), service)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	http2 "github.ru/noskov-sergey/what_to_watch_golang/internal/delivery/http"
	"github.ru/noskov-sergey/what_to_watch_golang/internal/metrics"
	oRep "github.ru/noskov-sergey/what_to_watch_golang/internal/repository/opinion"
	usecase "github.ru/noskov-sergey/what_to_watch_golang/internal/usecase/opinion"
)

func main() {
	cfg := flag.String("c", ".env", "config file path")
	flag.Parse()

	reg := prometheus.NewRegistry()
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

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

	m := metrics.NewMetrics(reg)

	service := http2.New(whatUC, m)

	go func() {
		log.Fatal(http.ListenAndServe(os.Getenv("MET_URL"), nil))
		log.Printf("prometheus is running on %s", os.Getenv("MET_URL"))
	}()

	log.Printf("server is running on %s", os.Getenv("BASE_URL"))
	err = http.ListenAndServe(os.Getenv("BASE_URL"), service)
	if err != nil {
		log.Fatal(err)
	}
}

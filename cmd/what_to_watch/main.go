package main

import (
	"database/sql"
	"flag"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"

	http2 "github.ru/noskov-sergey/what_to_watch_golang/internal/delivery/http"
	"github.ru/noskov-sergey/what_to_watch_golang/internal/metrics"
	oRep "github.ru/noskov-sergey/what_to_watch_golang/internal/repository/opinion"
	usecase "github.ru/noskov-sergey/what_to_watch_golang/internal/usecase/opinion"
)

func main() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	cfg := flag.String("c", ".env", "config file path")
	flag.Parse()

	reg := prometheus.NewRegistry()
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	err := godotenv.Load(*cfg)
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		logger.Fatal("failed to connect to database: %s", zap.Error(err))
	}
	defer db.Close()

	whatUC := usecase.New(oRep.NewOpinionRepository(db))

	m := metrics.NewMetrics(reg)

	service := http2.New(whatUC, m, logger)

	go func() {
		logger.Info("prometheus is running on %s", zap.String("status", os.Getenv("MET_URL")))
		logger.Fatal("prometheus fatal", zap.Error(http.ListenAndServe(os.Getenv("MET_URL"), nil)))
	}()

	logger.Info("server is running", zap.String("BASE_URL", os.Getenv("BASE_URL")))
	err = http.ListenAndServe(os.Getenv("BASE_URL"), service)
	if err != nil {
		logger.Fatal("server running error", zap.Error(err))
	}
}

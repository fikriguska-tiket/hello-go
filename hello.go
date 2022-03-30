package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name:        "http_requests_total",
		Help:        "Total number of HTTP requests",
		ConstLabels: prometheus.Labels{"server": "api"},
	},
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	httpRequestsTotal.Inc()
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	version := "v1.3"

	fmt.Fprintf(w, fmt.Sprintf("Hello World!! - from %s @ %s", hostname, version))
	log.Println(fmt.Sprintf("Hello World!! - from %s @ %s", hostname, version))
}

func main() {
	prometheus.MustRegister(httpRequestsTotal)

	r := mux.NewRouter()
	r.HandleFunc("/", helloServer)
	r.Handle("/metrics", promhttp.Handler())

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	addr := fmt.Sprintf(":%s", port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	log.Print("Starting server at ", addr)
	log.Fatal(srv.ListenAndServe())
}

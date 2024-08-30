package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Router() {
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("GET /api/healthz", handleReadiness)

	fmt.Printf("Server listening on %s\n", port)
	log.Fatal(server.ListenAndServe())
}

// A simple API for legendary animals.

package main

import (
	"context"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

type Legend struct {
	ID			string `json:"id"`
	Name		string `json:"name"`
	Type    string `json:"type"`
	Origin 	string `json:"origin"`
}

var legends []Legend

func getLegends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(legends)
}

func createLegend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newLegend Legend
	json.NewDecoder(r.Body).Decode(&newLegend)
	newLegend.ID = strconv.Itoa(len(legends) +1)
	legends = append(legends, newLegend)
	json.NewEncoder(w).Encode(newLegend)
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	ln, err := ngrok.Listen(ctx,
		config.LabeledTunnel(
			config.WithLabel("edge", os.Getenv("NGROK_LABEL")),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	router:=mux.NewRouter()
	router.HandleFunc("/legend", getLegends).Methods("GET")
	router.HandleFunc("/legend", createLegend).Methods("POST")

	log.Println("ngrok API gateway established:")
	log.Println("Tunnel:", ln.ID())
	log.Println("Edge(s):")
	for key, value := range ln.Labels() {
		log.Println(key, value)
	}

	return http.Serve(ln, router)
}

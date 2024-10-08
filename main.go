package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type kdrama struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var kdramas []kdrama

func getkdramas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kdramas)
}
func deletekdrama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range kdramas {
		if item.ID == params["id"] {
			kdramas = append(kdramas[:index], kdramas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(kdramas)
}
func getkdrama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range kdramas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// create
func createkdrama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var kdrama kdrama
	_ = json.NewDecoder(r.Body).Decode(&kdrama)
	kdrama.ID = strconv.Itoa(rand.Intn(100000000))
	kdramas = append(kdramas, kdrama)
	json.NewEncoder(w).Encode(kdrama)
}

// update
func updateKdrama(w http.ResponseWriter, r *http.Request) {

	//set son content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)
	//loop over the kdramas, range
	//delete the movie with bthe id that you have sent
	//add a new movie-the movie that we sent in the body of postman
	for index, item := range kdramas {
		if item.ID == params["id"] {
			kdramas = append(kdramas[:index], kdramas[index+1:]...)
			var kdrama kdrama
			_ = json.NewDecoder(r.Body).Decode(&kdrama)
			kdrama.ID = params["id"]
			kdramas = append(kdramas, kdrama)
			json.NewEncoder(w).Encode(kdrama)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	kdramas = append(kdramas, kdrama{ID: "1", Isbn: "438227", Title: "kdrama one", Director: &Director{Firstname: "ruvin", Lastname: "rae"}})
	kdramas = append(kdramas, kdrama{ID: "2", Isbn: "456777", Title: "kdrama two", Director: &Director{Firstname: "kaashvi", Lastname: "aggarwal"}})
	r.HandleFunc("/kdramas", getkdramas).Methods("GET")
	r.HandleFunc("/getkdramas/{id}", getkdrama).Methods("GET")
	r.HandleFunc("/createkdramas", createkdrama).Methods("POST")
	r.HandleFunc("/updatekdramas/{id}", updateKdrama).Methods("PUT")
	r.HandleFunc("/deletekdramas/{id}", deletekdrama).Methods("DELETE")
	fmt.Printf("starting server at port 8000\n ")
	log.Fatal(http.ListenAndServe(":8000", r))
}

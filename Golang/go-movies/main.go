package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getMovies")
}

func getMovieId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "getMovieId: %s", vars["id"])
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "createMovie")
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "deleteMovie: %s", vars["id"])
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "updateMovie: %s", vars["id"])
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "33", Isbn: "1244321124", Title: "Annen Evdemi 2", Director: &Director{Name: "Murtaza", Surname: "Candaş"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieId).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")

	fmt.Printf("Starting server at port 8080 \n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Year int    `json:"Year"`
}

type allMovies []movie

var movies = allMovies{
	{
		ID:   1,
		Name: "Shawshank Redemption",
		Year: 1994,
	},
}

var movieIdx = 2

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secure Api by Río")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID, e := strconv.Atoi(vars["id"])

	if e != nil {
		fmt.Fprintf(w, "Provide Valid movie ID")
		return
	}

	for _, movie := range movies {
		if movie.ID == movieID {
			w.WriteHeader(http.StatusFound)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Movie with ID:%v not found", movieID)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie movie
	reqBody, e := ioutil.ReadAll(r.Body)
	if e != nil {
		fmt.Fprintf(w, "Provide Valid task")
	}
	json.Unmarshal(reqBody, &newMovie)
	newMovie.ID = movieIdx
	movieIdx++
	movies = append(movies, newMovie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID, e := strconv.Atoi(vars["id"])

	if e != nil {
		fmt.Fprintf(w, "Provide Valid movie ID")
		return
	}

	for i, movie := range movies {
		if movie.ID == movieID {
			movies = append(movies[:i], movies[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Deleted movie with ID:%v", movieID)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Movie with ID:%v not found", movieID)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID, e := strconv.Atoi(vars["id"])
	var updatedMovie movie
	if e != nil {
		fmt.Fprintf(w, "Provide Valid movie ID")
		return
	}
	reqBody, e := ioutil.ReadAll(r.Body)

	if e != nil {
		fmt.Fprintf(w, "Provide Valid movie ID")
		return
	}
	json.Unmarshal(reqBody, &updatedMovie)

	for i, movie := range movies {
		if movie.ID == movieID {
			movies = append(movies[:i], movies[i+1:]...)
			updatedMovie.ID = movieID
			movies = append(movies, updatedMovie)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Updated movie with ID:%v", movieID)
			return
		}
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", indexRoute)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8080", r))
}

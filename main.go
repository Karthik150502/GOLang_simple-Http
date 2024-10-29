package main

import (
	// "fmt"
	"fmt"
	"net/http"
	"strconv" // To convert int to string

	"encoding/json" //For encoding the data while sending to the Client

	"github.com/gorilla/mux"

	"log"
	"math/rand"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{
		Id:       "1",
		Isbn:     "438227",
		Title:    "The Avatar",
		Director: &Director{Firstname: "James", Lastname: "Cameroon"},
	})
	movies = append(movies, Movie{
		Id:       "2",
		Isbn:     "549832",
		Title:    "Inception",
		Director: &Director{Firstname: "Christopher", Lastname: "Nolan"},
	})
	movies = append(movies, Movie{
		Id:       "3",
		Isbn:     "789654",
		Title:    "The Matrix",
		Director: &Director{Firstname: "Lana", Lastname: "Wachowski"},
	})
	movies = append(movies, Movie{
		Id:       "4",
		Isbn:     "345987",
		Title:    "Interstellar",
		Director: &Director{Firstname: "Christopher", Lastname: "Nolan"},
	})
	movies = append(movies, Movie{
		Id:       "5",
		Isbn:     "892374",
		Title:    "The Godfather",
		Director: &Director{Firstname: "Francis", Lastname: "Coppola"},
	})
	movies = append(movies, Movie{
		Id:       "6",
		Isbn:     "982374",
		Title:    "Pulp Fiction",
		Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"},
	})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movie/create", createMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting the port at 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json") //
	json.NewEncoder(res).Encode(movies)
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	res.Header().Set("Content-Type", "application/json")

	for index, item := range movies {
		if item.Id == params["id"] {
			//Getting the movies at the index to delete it, and at the movies[index] replacing it with all the other movies, then the movie we want to delete wont exist.
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)

}
func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(res).Encode(item) //Modifies/Encodes the response object with the item.
			break
		}
	}

}
func createMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movie)
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie) //Might provide a error, if the schema doesnt match.
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(res).Encode(movie)
			return
		}
	}
}

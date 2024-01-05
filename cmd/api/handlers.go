package main

import (
	"BE-Go-Movies/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// type Payload struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Version string `json:"version"`
// }

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "Active",
		Message: "Go Movie  up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	rd, _ := time.Parse("2006-01-02", "1986-01-02")
	highlander := models.Movie{
		ID:          1,
		Title:       "Higlander",
		ReleaseDate: rd,
		MPAARating:  "R",
		RunTime:     116,
		Description: "Very Nice Movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)
	rd, _ = time.Parse("2006-01-02", "1986-06-12")

	lotla := models.Movie{
		ID:          2,
		Title:       "Raider Lost Ark",
		ReleaseDate: rd,
		MPAARating:  "PG-13",
		RunTime:     115,
		Description: "Very Nice Movie Ark",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, lotla)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

package main

import(
	"fmt"
	"time"
	"net/http"

	"github.com/Praveen005/notesAPI/internal/data"
)

func(app *application) createMovieHandler(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w, "Create a Movie..")
}

func(app *application)showMovieHandler(w http.ResponseWriter, r *http.Request){
	

	id, err := app.readIDParam(r)

	if err != nil || id <1{
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:			id,
		CreatedAt:	time.Now(),
		Title:		"Casablanca",
		Runtime:	102,
		Genres:		[]string{"drama", "romance", "war"},
		Version:	1,
	}
	
	
		

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil{
		// app.logger.Println(err)
		// http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		
		app.serverErrorResponse(w, r, err)
	}

	// fmt.Fprintf(w, "Showing the details of the movie with id %d..", id)
}


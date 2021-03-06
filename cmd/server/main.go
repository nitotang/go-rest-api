package main

import (
	"fmt"
	"net/http"

	"github.com/nitotang/go-rest-api/internal/comment"
	"github.com/nitotang/go-rest-api/internal/database"
	transportHTTP "github.com/nitotang/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run - handles the startup of our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

// Our main entrypoint for the application
func main() {
	fmt.Println("REST API init")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}

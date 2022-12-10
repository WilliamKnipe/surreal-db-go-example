package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/surrealdb/surrealdb.go"
	"github.com/williamknipe/testing-go-library/handlers"
)

func main() {
	fmt.Println("Connecting to SurrealDB database")

	// Connect to Database via websocket
	db := InitDatabase()
	defer db.Close()

	fmt.Println("Starting server")
	// Create a new router
	router := mux.NewRouter()

	// Handlers import from handlers package
	HealthCheck := handlers.HealthCheck
	AddCafes := handlers.AddCafes
	GetCafes := handlers.GetCafes

	// Specify endpoints, handler functions and HTTP method
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/cafes", AddCafes(db)).Methods("POST")
	router.HandleFunc("/cafes", GetCafes(db)).Methods("GET")
	http.Handle("/", router)

	// Start server
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", router)
}

type Cafe struct {
	Name   string `json:"name"`
	Coffee string `json:"coffee"`
	Chairs string `json:"chairs"`
}

func InitDatabase() *surrealdb.DB {
	// Create new Surreal db connection.
	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		fmt.Println("InitDatabase new db failed")
		panic(err)
	}

	// Credentials set in surrealDB start command
	_, err = db.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	})
	if err != nil {
		fmt.Println("InitDatabase sign in falled")
		panic(err)
	}

	_, err = db.Use("test", "test")
	if err != nil {
		fmt.Println("InitDatabase db.Use failed")
		panic(err)
	}

	// Create cafes table with some example cafe data
	_, err = db.Create("cafes", Cafe{
		Name:   "hot comfy cafe",
		Coffee: "hot",
		Chairs: "comfy",
	})
	if err != nil {
		fmt.Println("InitDatabase db.Create failed")
		panic(err)
	}
	log.Println("InitDatabase Created cafes")

	userData, err := db.Select("cafes")
	if err != nil {
		fmt.Println("InitDatabase select cafes failed")
		panic(err)
	}

	if userData == nil {
		fmt.Println("InitDatabase Select cafes is nil")
	}

	fmt.Println("InitDatabase created DB")
	return db
}

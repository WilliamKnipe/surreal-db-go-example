package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/surrealdb/surrealdb.go"
	"github.com/williamknipe/testing-go-library/handlers"
)

func main() {
	fmt.Println("Connecting to SurrealDB database")

	c := make(chan *surrealdb.DB)

	// Connect to Database via websocket
	go InitDatabase(c)
	db := <-c
	fmt.Println(db)

	fmt.Println("Starting server")
	// Create a new router
	router := mux.NewRouter()

	// Handlers import from handler packages
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

func InitDatabase(c chan *surrealdb.DB) {
	var wg sync.WaitGroup
	wg.Add(1)

	// Create new Surreal db connection.
	db, err1 := surrealdb.New("ws://localhost:8000/rpc")

	if err1 != nil {
		fmt.Println("InitDatabase new db failed")
		panic(err1)
	}

	defer db.Close()
	fmt.Println(db)

	// Credentials set in surrealDB start command
	_, err2 := db.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	})

	if err2 != nil {
		fmt.Println("InitDatabase sign in falled")
		panic(err2)
	}

	_, err := db.Use("test", "test")

	if err != nil {
		fmt.Println("InitDatabase db.Use failed")
		panic(err)
	}

	// Create cafes table
	// Some example cafe data
	_, err = db.Create("cafes", Cafe{
		Name:   "hot comfy cafe",
		Coffee: "hot",
		Chairs: "comfy",
	})

	if err != nil {
		fmt.Println("InitDatabase db.Create failed")
		panic(err)
	}

	userData, err4 := db.Select("cafes")

	if err4 != nil {
		fmt.Println("InitDatabase select cafes failed")
		panic(err4)
	}

	log.Println("InitDatabase Created cafes")

	if userData == nil {
		fmt.Println("InitDatabase Select cafes is nil")
	}

	fmt.Println("InitDatabase created DB")

	c <- db

	fmt.Println("Sent pointer through channel")

	// Wait forever with this goroutine
	wg.Wait()
}

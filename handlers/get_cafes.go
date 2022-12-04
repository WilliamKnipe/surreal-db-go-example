package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/surrealdb/surrealdb.go"
)

func GetCafes(db *surrealdb.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get cafes")

		if db == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("db pointer reference is nil")
			return
		}

		_, err := db.Use("test", "test")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Use test test failed")
			return
		}

		userData, err2 := db.Select("cafes")

		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Select cafes failed")
			return
		}

		if userData == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "userData is nil")
			return
		}

		json, err3 := json.Marshal(userData)

		if err3 != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No Cafes Found")
			return
		}

		// Set content type and return JSON with status OK
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	})
}

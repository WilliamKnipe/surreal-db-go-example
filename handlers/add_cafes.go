package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/surrealdb/surrealdb.go"
)

func AddCafes(db *surrealdb.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Add cafes")

		if db == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "db pointer reference is nil")
			return
		}

		var p AddUsersRequest
		b := json.NewDecoder(r.Body)

		_, err := db.Use("test", "test")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Use test test failed")
			return
		}

		err = b.Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, s := range p.Cafes {
			_, err = db.Create("cafes", Cafe{
				Name:   s.Name,
				Coffee: s.Coffee,
				Chairs: s.Chairs,
			})
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Adding Cafes Failed")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Created Cafes")
	})
}

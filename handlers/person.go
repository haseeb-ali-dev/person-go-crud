package handlers

import (
	"encoding/json"
	"fmt"
	"go-crud/data"
	"log"
	"net/http"
	"strconv"
)

// Type Aliases for brevity
type Person = data.Person

// Function Aliases
var (
	createPerson = data.CreatePerson
	readPerson   = data.ReadPerson
	deletePerson = data.DeletePerson
	updatePerson = data.UpdatePerson
)

type Response struct {
	Message string `json:"message,omitempty"`
	Data    Person `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// API layer
func PersonHandler(w http.ResponseWriter, r *http.Request) {
	var res Response

	// send the response at the end
	defer func() {
		json.NewEncoder(w).Encode(res)
	}()

	// add headers
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case http.MethodPost:
		// read the json body and convert to struct
		var p data.Person

		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			res.Error = err.Error()
			return
		}

		id := createPerson(p)

		res.Message = fmt.Sprintf("Person created with ID: %d", id)
	case http.MethodGet:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("Invalid Id")
			res.Error = "Invalid Id"
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(strID)

		if err != nil {
			log.Println("Invalid Id", err)
			res.Error = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := readPerson(id)

		if err != nil {
			log.Println("Person not found", err)
			res.Error = err.Error()
			w.WriteHeader(http.StatusNotFound)
			return
		}

		res.Data = p
	case http.MethodPut:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("Invalid Id")
			res.Error = "Invalid Id"
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(strID)

		if err != nil {
			log.Println("Invalid Id", err)
			res.Error = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var p Person
		json.NewDecoder(r.Body).Decode(&p)

		err = updatePerson(id, p)

		if err != nil {
			log.Println("Person not updated", err)
			res.Error = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res.Message = fmt.Sprintf("Person updated with ID: %d", id)
	case http.MethodDelete:
		strID := r.URL.Query().Get("id")

		if strID == "" {
			log.Println("Invalid Id")
			res.Error = "Invalid Id"
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(strID)

		if err != nil {
			log.Println("Invalid Id", err)
			res.Error = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deletePerson(id)

		if err != nil {
			log.Println("Person not deleted", err)
			res.Error = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res.Message = fmt.Sprintf("Person deleted with ID: %d", id)
	}
}

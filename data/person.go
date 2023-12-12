package data

import (
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

// personID: Person
var DataLayer map[int]Person
var PersonID int

func init() {
	DataLayer = make(map[int]Person)
	PersonID = 0

	log.Println("datalayer initialized")
}

// Datalayer
func CreatePerson(p Person) int {
	pID := PersonID + 1
	DataLayer[pID] = p

	// update the personID
	PersonID = pID

	return pID
}

func ReadPerson(id int) (Person, error) {
	if person, ok := DataLayer[id]; ok {
		return person, nil
	}

	return Person{}, fmt.Errorf("Person not found with ID: %d", id)
}

func UpdatePerson(id int, p Person) error {
	if _, ok := DataLayer[id]; ok {
		DataLayer[id] = p

		return nil
	}

	return fmt.Errorf("Person for updation not found with ID: %d", id)
}

func DeletePerson(id int) error {
	if _, ok := DataLayer[id]; ok {
		delete(DataLayer, id)

		return nil
	}

	return fmt.Errorf("Person for deletion not found with ID: %d", id)
}

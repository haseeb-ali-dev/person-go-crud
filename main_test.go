package main

import (
	"go-crud/data"
	"testing"
)

// Type Aliases for brevity
type Person = data.Person

// Aliases
var (
	// Functions
	createPerson = data.CreatePerson
	readPerson   = data.ReadPerson
	deletePerson = data.DeletePerson
	updatePerson = data.UpdatePerson

	// Variables
	dataLayer = data.DataLayer
)

// Testing Layer
func TestCreatePerson(t *testing.T) {
	id := createPerson(Person{
		Name: "John",
		Age:  25,
	})

	p, ok := dataLayer[id]

	if id == 1 && ok && p.Name == "John" && p.Age == 25 {
		t.Log("Test passed: Person created")
		return
	}

	t.Error("Test failed: Person not created", id, ok, p, dataLayer)
}

func TestReadPerson(t *testing.T) {
	p, err := readPerson(1)

	if err != nil {
		t.Error("Test failed: Person not found", err)
		return
	}

	if p.Name != "John" || p.Age != 25 {
		t.Error("Test failed: Person not found", p)
		return
	}

	t.Log("Test passed: Person found", p)
}

func TestUpdatePerson(t *testing.T) {
	err := updatePerson(1, Person{
		Name: "John",
		Age:  30,
	})

	if err != nil {
		t.Error("Test failed: Person not updated", err)
		return
	}

	p := dataLayer[1]
	if p.Name != "John" || p.Age != 30 {
		t.Error("Test failed: Person not updated", p)
		return
	}

	t.Log("Test passed: Person updated", p)
}

func TestDeletePerson(t *testing.T) {
	err := deletePerson(1)

	if err != nil {
		t.Error("Test failed: Person not deleted", err)
		return
	}

	p, ok := dataLayer[1]
	if ok {
		t.Error("Test failed: Person not deleted expected nil, got: ", ok, p)
		return
	}

	t.Log("Test passed: Person deleted", ok, p)
}

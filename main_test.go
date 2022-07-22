package main

import (
	"testing"
)

func TestDatabaseInsertSuccessful(t *testing.T) {
	// Arrange
	InstantiateDBIfNotExists()
	AddRecord()

	// Act
	person := getRecord()
	
	// Assert
	if person.PersonId != 1 {
		t.Error("Invalid Id")
	}

	if person.LastName != "Simpson" {
		t.Error("Invalid last name")
	}

	if person.FirstName != "Homer" {
		t.Error("Invalid first name")
	}

	if person.Address != "742 Evergreen Terrace" {
		t.Error("Invalid address")
	}

	if person.City != "Springfield" {
		t.Error("Invalid city")
	}
}

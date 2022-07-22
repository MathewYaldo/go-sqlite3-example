package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	InstantiateDBIfNotExists()
	AddRecord()
	log.Println(getRecord())
}

func InstantiateDBIfNotExists() {
	if _, err := os.Stat("data.db"); err == nil {
		log.Println("Database exists.")
		return
	}

	os.Create("data.db")
	db, err := sql.Open("sqlite3", "data.db")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec("CREATE TABLE Persons ( PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255) );")
	if err != nil {
		log.Fatal(err)
	}
}

func AddRecord() {
	db, err := sql.Open("sqlite3", "data.db")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec("INSERT INTO Persons (PersonID, LastName, FirstName, Address, City) VALUES (1, 'Simpson', 'Homer', '742 Evergreen Terrace', 'Springfield');")
	if err != nil {
		log.Fatal(err)
	}
}

func getRecord() Person {
	db, err := sql.Open("sqlite3", "data.db")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return Person{}
	}

	var result = db.QueryRow("SELECT * FROM Persons WHERE LastName = 'Simpson'")
	var person = Person{}
	err = result.Scan(&person.PersonId, &person.LastName, &person.FirstName, &person.Address, &person.City)

	// handle error
	if err != nil {
		panic(err)
	}

	return person
}

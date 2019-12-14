package main

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type People struct {
	id int64
	name string
}

type Peoples []People

func GetPeoples() Peoples{
	db, err := sql.Open("mysql", "docker:docker@tcp(db)/test_database")
	if err != nil {
		log.Fatal("db error.")
	}
	
	rows, err := db.Query("select * from people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var peoples Peoples
	for rows.Next() {
		people := People{}
		if err := rows.Scan(&people.id, &people.name); err != nil {
			log.Fatal(err)
		}

		peoples = append(peoples, people)
	}
	defer db.Close()

	return peoples
}
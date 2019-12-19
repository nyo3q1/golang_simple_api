package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type People struct {
	id   int64
	name string
}

type Peoples []People

func GetPeoples() (*Peoples, error) {
	db, err := sql.Open("mysql", "docker:docker@tcp(db)/test_database")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var peoples Peoples
	for rows.Next() {
		people := People{}
		if err := rows.Scan(&people.id, &people.name); err != nil {
			return nil, err
		}

		peoples = append(peoples, people)
	}

	return &peoples, nil
}

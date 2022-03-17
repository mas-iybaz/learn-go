package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id      string
	nama    string
	email   string
	telepon string
	kelamin string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:aziz@tcp(127.0.0.1:3306)/99_iqbal")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getAllStudent() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM siswa_iqbal")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Student

	for rows.Next() {
		var each = Student{}
		var err = rows.Scan(&each.id, &each.nama, &each.email, &each.telepon, &each.kelamin)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.email)
	}
}

func getStudent(ID string) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = Student{}
	var id = ID
	err = db.
		QueryRow("SELECT * FROM siswa_iqbal WHERE id = ?", id).
		Scan(&result.id, &result.nama, &result.email, &result.telepon, &result.kelamin)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result.email)
}

func main() {
	getAllStudent()
	getStudent("1")
}

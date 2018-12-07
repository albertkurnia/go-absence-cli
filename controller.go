package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func CreateDivision(div Division) int {

	if rowExists("SELECT id FROM division WHERE id=$1", div.id) {
		fmt.Printf("ID %d already exist.\n", div.id)
		return 0
	}

	sqlStatement := `INSERT INTO division (id, name) VALUES ($1, $2) RETURNING id`

	id := 0

	err := DB.QueryRow(sqlStatement, div.id, div.name).Scan(&id)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	return id
}

func CreateEmployee(employee Employee) int {

	if rowExists("SELECT id FROM employee WHERE id=$1", employee.id) {
		fmt.Printf("ID %d already exists.\n", employee.id)
		return 0
	}

	sqlStatement := `SELECT id FROM division WHERE name=$1`

	idDivision := 0

	row := DB.QueryRow(sqlStatement, employee.division)

	switch err := row.Scan(&idDivision); err {
	case sql.ErrNoRows:
		fmt.Print("Division not found or mismatch name.\n")
		return 0
	case nil:
		break
	default:
		panic(err)
		return 0
	}

	sqlStatement = `INSERT INTO employee (id, "divisionId", name, status) VALUES ($1, $2, $3, $4) RETURNING id`

	id := 0

	err := DB.QueryRow(sqlStatement, employee.id, idDivision, employee.name, employee.status).Scan(&id)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	return id
}

func EmployeeIn(id int) *Output {

	// validate if employee is registered or not
	if rowExists("SELECT id FROM employee WHERE id=$1", id) == false {
		fmt.Printf("ID %d not found.\n", id)
		os.Exit(1)
	}

	// validate if employee status is ACTIVE, BANNED, INACTIVE
	sqlStatement := `SELECT status FROM employee WHERE id=$1`

	var status string

	err := DB.QueryRow(sqlStatement, id).Scan(&status)

	if err != nil {
		log.Fatal(err)
	}

	if status == "BANNED" {
		fmt.Println("Sorry, You are temporary banned.")
		os.Exit(1)
	}

	if status == "INACTIVE" {
		fmt.Println("You are not allowed.")
		os.Exit(1)
	}

	// get detail employee
	sqlStatement = `SELECT e.name, d.name FROM employee e JOIN division d ON e."divisionId"=d.id WHERE e.id=$1`

	var output Output

	err = DB.QueryRow(sqlStatement, id).Scan(&output.name, &output.division)

	if err != nil {
		log.Fatal(err)
	}

	// record absent history
	var currentTime = time.Now()

	sqlStatement = `INSERT INTO absent_history ("employeeId", "status", "absentTime") VALUES ($1,'IN', $2)`

	_, err = DB.Exec(sqlStatement, id, currentTime)

	if err != nil {
		log.Fatal(err)
	}

	output.currentTime = currentTime
	output.status = "IN"

	return &output
}

func EmployeeOut(id int) *Output {

	// validate if employee status is ACTIVE, BANNED, INACTIVE
	sqlStatement := `SELECT status FROM employee WHERE id=$1`

	var status string

	err := DB.QueryRow(sqlStatement, id).Scan(&status)

	if err != nil {
		log.Fatal(err)
	}

	if status == "BANNED" {
		fmt.Println("Sorry, You are temporary banned.")
		os.Exit(1)
	}

	if status == "INACTIVE" {
		fmt.Println("You are not allowed.")
		os.Exit(1)
	}

	// get detail employee
	sqlStatement = `SELECT e.name, d.name FROM employee e JOIN division d ON e."divisionId"=d.id WHERE e.id=$1`

	var output Output

	err = DB.QueryRow(sqlStatement, id).Scan(&output.name, &output.division)

	if err != nil {
		log.Fatal(err)
	}

	// record absent history
	var currentTime = time.Now()

	sqlStatement = `INSERT INTO absent_history ("employeeId", status, "absentTime") VALUES ($1, 'OUT', $2)`

	_, err = DB.Exec(sqlStatement, id, currentTime)

	if err != nil {
		log.Fatal(err)
	}

	output.currentTime = currentTime
	output.status = "OUT"

	return &output
}

func EmployeeLastStatus(id int) string {

	sqlStatement := `SELECT ah.status FROM absent_history ah JOIN employee e ON e.id=ah."employeeId" WHERE ah."employeeId"=$1 ORDER BY ah.id DESC LIMIT 1`

	var status sql.NullString

	row := DB.QueryRow(sqlStatement, id)

	switch err := row.Scan(&status); err {
	case sql.ErrNoRows:
		return "FIRST_TIME"
	case nil:
		return status.String
	default:
		panic(err)
	}
}

func rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := DB.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("error checking if row exists '%s' %v", args, err)
	}

	return exists
}

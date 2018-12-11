package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "albert"
	dbname = "absent"
)

var DB *sql.DB

func main() {

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)
	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer DB.Close()

	fmt.Println("-- Simple Absence Program --")
	fmt.Println("-------------------------------")
	fmt.Println("Type: `help -all` for list command.")

	reader := bufio.NewReader(os.Stdin)

	for {

		var input string
		fmt.Print("-> ")

		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSuffix(input, "\n")
		arrInput := strings.Fields(input)

		if len(arrInput) == 1 {

			if arrInput[0] == "help" {
				fmt.Println("Type: `help -all` for list command.")
			}

			i, err := strconv.Atoi(arrInput[0])

			if err != nil {
				log.Fatal(err)
			}

			if i == 0 {
				fmt.Println("System exit.")
				break
			}

			status := EmployeeLastStatus(i)

			if status == "FIRST_TIME" || status == "OUT" {
				output := EmployeeIn(i)
				ShowToConsole(output)
			} else if status == "IN" {
				output := EmployeeOut(i)
				ShowToConsole(output)
			}
		} else if len(arrInput) > 1 {
			if arrInput[0] == "help" {
				if arrInput[1] == "-all" {
					PrintUsage()
				}
			}
			if arrInput[0] == "INSERT-EMPLOYEE" {
				var employee Employee

				id, err := strconv.Atoi(arrInput[1])
				name := arrInput[2]
				division := arrInput[3]
				status := arrInput[4]

				if err != nil {
					log.Fatal(err)
				}

				employee.id = id
				employee.name = name
				employee.division = division
				employee.status = status

				resp := CreateEmployee(employee)

				if resp == 0 {
					continue
				} else {
					fmt.Printf("Employee with ID: %d created.\n", resp)
				}

			} else if arrInput[0] == "INSERT-DIVISION" {
				var division Division

				id, err := strconv.Atoi(arrInput[1])
				name := arrInput[2]

				if err != nil {
					log.Fatal(err)
				}

				division.id = id
				division.name = name

				resp := CreateDivision(division)

				if resp == 0 {
					continue
				} else {
					fmt.Printf("Division with ID: %d created.\n", resp)
				}

			}
		}

	}
}

func ShowToConsole(output *Output) {
	fmt.Printf("Welcome %s\n", output.name)
	fmt.Printf("Divison: %s\n", output.division)
	fmt.Printf("[%s] at %v\n", output.status, output.currentTime.Format("Mon Jan _2 15:04:05 2006"))
}

func PrintUsage() {
	fmt.Println("For Create New Division: ")
	fmt.Println("Type `INSERT-DIVISION [id] [name]`")
	fmt.Println("For Create New Employee: ")
	fmt.Println("Type `INSERT-Emplotee [id] [name] [division] [ACTIVE/BANNED/INACTIVE]`")
}

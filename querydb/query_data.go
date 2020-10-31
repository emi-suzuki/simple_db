package main

import (
	"database/sql"
	"dbutils"
	"fmt"

	"github.com/kisielk/sqlstruct"
)

type SeeAlso struct {
	Id           int
	Object_Id    int
	Reference_Id int
}

type BigResult struct {
	Id              int
	Object_Id       int
	Reference_Id    int
	AccessionNumber string
	Artist          string
	Continent       string
	Country         string
	Dated           string
	Department      string
	Description     string
	Dimension       string
	IdURL           string
	Image           string
	Height          int
	Width           int
	Medium          string
	Room            string
	Style           string
	Text            string
	Title           string
}

func SelectTable(db *sql.DB, tableName string) *sql.Rows {
	q := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error querying with exact match")
	}
	return rows
}

func ExactQuery(db *sql.DB, tableName string, field string, exactMatch interface{}) *sql.Rows {
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s = '?'", tableName, field)
	rows, err := db.Query(q, exactMatch)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error querying with exact match")
	}
	return rows
}

func GeneralQuery(db *sql.DB, tableName string, field string, partialMatch string) *sql.Rows {
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s LIKE '%s'", tableName, field, partialMatch)
	fmt.Println(q)
	rows, err := db.Query(q)

	// rows, err := db.Query(`SELECT * FROM ? WHERE ? LIKE '?';`, tableName, field, partialMatch)
	if err != nil {
		fmt.Println("Error querying with partial match")
	}
	return rows
}

func PrintQueryResults(rows *sql.Rows) {
	for rows.Next() {
		var object Object
		err := sqlstruct.Scan(&object, rows)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(object)
	}
}

func PrintQueryResultsSeeAlso(rows *sql.Rows) {
	for rows.Next() {
		var seeAlso SeeAlso
		err := sqlstruct.Scan(&seeAlso, rows)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(seeAlso)
	}
}

func JoinQuery(db *sql.DB, tableName1 string, tableName2 string, table1Col string, table2Col string) *sql.Rows {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s JOIN %s ON %s.%s = %s.%s; ", tableName1, tableName2, tableName1, table1Col, tableName2, table2Col))
	if err != nil {
		fmt.Println(err.Error)
	}
	return rows
}

func main() {
	db := dbutils.OpenDatabase("sqlite3", "/mnt/c/Users/Emi/Desktop/database/src/museumdb/artmuseum.db")

	// Select rows where department contains "Prints"
	rows1 := GeneralQuery(db, "objects", "Department", "Prints%")
	defer rows1.Close()
	PrintQueryResults(rows1)

	// Select the see_also table
	rows2 := SelectTable(db, "see_also")
	defer rows2.Close()
	PrintQueryResultsSeeAlso(rows2)

	// Retrieve all details for
	rows3 := JoinQuery(db, "see_also", "objects", "object_id", "id")
	for rows3.Next() {
		var results BigResult
		err := sqlstruct.Scan(&results, rows3)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(results)
	}
}

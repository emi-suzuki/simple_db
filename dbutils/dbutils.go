// Testing out go to try to make a database and query in the database

package dbutils

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3" // don't want to call by name, only through sql
)

/*
OpenDatabase opens and returns a reference to a database/sql defined database.
If the database does not exist, it is created.

Ex:
db := OpenDatabase("sqlite3", "/full/path/shop.db")
*/
func OpenDatabase(db_type string, db_name string) *sql.DB {
	db, err := sql.Open(db_type, db_name)
	if err != nil {
		fmt.Println("Error opening database/n", err.Error())
	}
	return db
}

/*
CreateTable creates a new table in the provided database with the given name
and schema. Schema should contain the field names as keys and field types as values.

Ex:
db := OpenDatabase("sqlite3", "/full/path/shop.db")
schema := map[string]string{"id": "INTEGER PRIMARY KEY", "name": "TEXT", "age": "INTEGER"}
CreateTable(db, "employees", schema)
*/
func CreateTable(db *sql.DB, tableName string, schema map[string]string) {
	// check that the table name contains no spaces or single quotes
	// var myRegex string = `^[\w]+$`
	// match, _ := regexp.MatchString(myRegex, tableName)
	// if !match {
	// 	errors.New(fmt.Sprintf("Table name must only contain alphanums and underscores < %s >", tableName))
	// }
	err := verifyName(tableName)
	if err != nil {
		fmt.Printf("Table name must only contain alphanums and underscores < %s >", tableName)
	}
	// convert the schema into a string
	var tableContents string
	for key, value := range schema {
		tableContents += fmt.Sprintf("%s %s, ", key, value)
	}
	tableContents = strings.TrimSuffix(tableContents, ", ")
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, tableContents))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Encountered an error during table creation <%s>", tableName)
	}
}

// func TableInsert(db *sql.DB, table_name string, fields []string)

/*
QueryString converts a list of strings into a string for queries

Ex:
var fields = []string{"id", "name", "age"}
stringFields = QueryString(fields)
fmt.Println(stringFields) // output: "id, name, age"
*/
func QueryString(fields []string) string {
	return strings.Join(fields, ", ")
}

/*
verifyName checks that the string/name given contains only alphanums and underscores
*/
func verifyName(name string) error {
	var myRegex string = `^[\w]+$`
	match, _ := regexp.MatchString(myRegex, name)
	if !match {
		return errors.New(fmt.Sprintf("Name must only contain alphanums and underscores < %s >", name))
	}
	return nil
}

/*
Input: The database, a string with the table_name, list of strings containing field names ###
Output: no output

Populate a table with data given the correct fields
*/
func PopulateTable(db *sql.DB, table_name string, fields []string, data [][]string, increment bool) {
	field_query := QueryString(fields)
	statement, _ := db.Prepare(fmt.Sprintf("INSERT INTO %s(%s) VALUES(?)", table_name, field_query))
	var id int = 1
	for _, entry := range data {
		entry_query := QueryString(entry)
		if increment {
			entry_query = fmt.Sprintf("%d, %s", id, entry_query)
			id++
		}
		statement.Exec(entry_query)
	}
}

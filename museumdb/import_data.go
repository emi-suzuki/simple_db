package main

import (
	"dbutils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var ObjTableName string = "objects"
var ExtTableName string = "see_also"

type Extra struct {
	SeeAlso []string `json:"see_also"`
}

var ExtSchema = map[string]string{
	"id":           "INTEGER PRIMARY KEY",
	"object_id":    "INTEGER",
	"reference_id": "INTEGER",
}
var ExtFields = []string{
	"object_id",
	"reference_id",
}

func main() {
	// initialize the database
	db := dbutils.OpenDatabase("sqlite3", "artmuseum.db")
	dbutils.CreateTable(db, ObjTableName, Schema)
	dbutils.CreateTable(db, ExtTableName, ExtSchema)

	// retrieve all json files from directory
	files, err := ioutil.ReadDir("./objects/0")
	if err != nil {
		fmt.Println("Error trying to access files in directory")
	}

	// create insert statement for objects
	fieldQuery := dbutils.QueryString(Fields)
	preparedQuery := fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		ObjTableName,
		fieldQuery,
	)
	// prepare string for secondary table
	extFieldQuery := dbutils.QueryString(ExtFields)

	// convert all json files into objects
	for _, file := range files {

		// open file
		jsonFile, err := os.Open(fmt.Sprintf("./objects/0/%s", file.Name()))
		if err != nil {
			fmt.Printf("Error trying to open json file <%s>", file.Name())
		}

		defer jsonFile.Close()

		// convert file to []byte
		byteVal, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Printf("Could not read file <%s>", file.Name())
		}

		var object Object
		var extra Extra

		json.Unmarshal(byteVal, &object)
		json.Unmarshal(byteVal, &extra)

		// create an id from the json file name, ex: "76.json" -> 76
		object.Id, err = strconv.Atoi(strings.Split(file.Name(), ".")[0])
		if err != nil {
			fmt.Println("Error creating object Id")
		}

		_, err = db.Exec(
			preparedQuery,
			object.Id,
			object.AccessionNumber,
			object.Artist,
			object.Continent,
			object.Country,
			object.Dated,
			object.Department,
			object.Description,
			object.Dimension,
			object.IdURL,
			object.Image,
			object.Height,
			object.Width,
			object.Medium,
			object.Room,
			object.Style,
			object.Text,
			object.Title,
		)

		if err != nil {
			fmt.Println(err.Error())
		}

		// insert all see_also references into a separate table
		for i := 0; i < len(extra.SeeAlso); i++ {
			// skip the lists which are populated with an empty string
			if extra.SeeAlso[i] == "" {
				break
			}
			db.Exec(fmt.Sprintf("INSERT INTO %s(%s) VALUES(?,?)", ExtTableName, extFieldQuery), object.Id, extra.SeeAlso[i])
		}

	}
}

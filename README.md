## Project Details

This project was created as a way to explore Golang and create a simple database. The resultant database was created using SQLite3, a Golang SQLite3 database driver from [https://github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) and a portion of the ArtsMIA collection's data. 

The *__dbutils__* package contains functions to create a new database and insert new tables. This package includes the functions which aren't specific to the dataset or the SQLite3, but may not be compatible with all database systems and all drivers.

This project uses data from the [ArtsMIA collection](https://github.com/artsmia/collection) labeled as objects. Each object is stored in an individual file labeled "<object_id>.json". Each object contains object specific fields such as "Title", "Height", "Width", etc. and a "See Also" list of other object ids.

The *__artmuseum.db__* file contains all objects sorted into two tables: an objects table and a see_also table. The objects table contains each object as a unique entry and uses the object id as the primary key. The see_also table matches an object's id to each object id in the object's "See Also" list. Each pair of ids is stored as a new entry. 

The *__museumdb__* main package creates the SQLite 3 database and contains functions, structs, and schema to populate the database with tables and entries. The package creates tables from the defined schema and loads all objects into two struct types: A single object struct and an extra struct containing the see_also list. The object struct is inserted once and the extra struct is iterated through to insert multiple entries.

The *__querydb__* main package contains functions to perform a small set of SQL select statements on the dataset and a set of print functions specfic to the returned results. Query results are returned as sql.Rows which are converted into structs to print results.

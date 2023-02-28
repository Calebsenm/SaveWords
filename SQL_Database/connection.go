
// The library that allows us to connect a MySQL


package SQL_Database

import (
  "fmt"
    "database/sql" // Interactuar con bases de datos
  _ "github.com/go-sql-driver/mysql"
  )


func  GetDatabase( ) (db *sql.DB, e error) {
  
  db, err :=  sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/Test")
	if err != nil {
		panic(err.Error())
  } else  {
    fmt.Println("Connect");
  } 
  return db,nil 
}


// The library that allows us to connect a MySQL


package SQL_Database

import (
  "fmt"
    "database/sql" // Interactuar con bases de datos
  _ "github.com/go-sql-driver/mysql"
  )


func  GetDatabase( ) (db *sql.DB, e error) {
  
  usuario := "root"
  pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "Words_App"
	// Debe tener la forma usuario:contraseña@protocolo(host:puerto)/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	


  if err != nil{
    return nil , err ;
  }else {
    fmt.Println("connect yes ")
  }
  
  return db,nil 
}

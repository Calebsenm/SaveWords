package Options

import (
  "fmt"
  "savewords/SQL_Database"
)

type Connection struct{
  User string
  Pass string

}

func  ( e * Connection)  Connection_Data() {

  db , err := SQL_Database.GetDatabase()
	defer db.Close()

  data , err := db.Query("SELECT username, password FROM users")
  
  if err != nil{
    fmt.Println("error ",err)
  }else {
    fmt.Println("connect yes !")
  }
  
  fmt.Println( e.Pass , e.User , db)
 
  fmt.Println("YEs ",   data)
  

}

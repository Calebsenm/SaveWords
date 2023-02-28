package Options

import (
  "fmt"
  "savewords/SQL_Database"
)

type User struct{
  ID        int 
  Name      string
  Password  string

}


type Connection struct{
  User string
  Pass string

}

func  ( e * Connection)  Connection_Data() {
  // call the Function for Connection to  database
  db , err := SQL_Database.GetDatabase()
	defer db.Close()
  
  // Make SQL Query 
  data , err := db.Query("SELECT * FROM users");
  
  if err != nil{
    fmt.Println("Error ",err)
  }

  newUsers := make(map[int]User)
  // clean Data and save  in the map  
  var iteraror int 
  for data.Next(){
    iteraror++
    var user User;
    err := data.Scan(&user.ID,&user.Name,&user.Password );
    
    newUsers[iteraror]= user
    if err != nil {
      fmt.Println(err );
    }
  }

  //check the data and the correct users  
  for _ , element := range newUsers{
    if element.Name  == e.User && element.Password == e.Pass {
      fmt.Println("Coneccion exitosa ");
    }
  } 
}



package Options

import (
  "fmt"
  "savewords/SQL_Database"
)

func Register_Data(){

  var username  string
  var password  string 
  var password2 string 

  for {
    fmt.Print("Put your Username ");
    fmt.Scan(&username);

    fmt.Print("Put your Password ");
    fmt.Scan(&password);

    fmt.Print("Repeat Your Password ");
    fmt.Scan(&password2);

    if password == password2{
      break;
    } else{
      fmt.Println("Passwords Do not Match ");
    }

  } 


  //Query Insert 
  db , err := SQL_Database.GetDatabase();
  data , err := db.Prepare("INSERT INTO users (username, password) VALUES(?, ?)");
 
  defer data.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = data.Exec(username , password );


  if err != nil{
    fmt.Println("Error ",err);
  } else{
    fmt.Println("User Saved Correct :) ");
  }
 

  defer db.Close();
}

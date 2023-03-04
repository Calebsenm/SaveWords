
package Options

import(
  "fmt"
  "savewords/SQL_Database"
)

type UsersData struct {

  Id     int 
  User   string
  Passw  string 
}

func PrintUsersData(){

  db , _:= SQL_Database.GetDatabase();
  defer db.Close()

  //Query 
  data , err2 := db.Query("SELECT * FROM users");
  if err2 != nil{
    fmt.Println("Eerror -> ", err2);
  }

  for data.Next(){
    
    var user UsersData 
    err := data.Scan(&user.Id, &user.User ,&user.Passw);
    if err != nil{
      fmt.Println("error ",err)
    }
   
    fmt.Println(user);

  }
  
  
}


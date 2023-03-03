
package main 

import (
  "fmt"
 op "savewords/Options"
)

var (
  OptionsMenu = map[int]string{
    1 : "Input  Words ",
    2 : "See Words ",
    3 : "Options ",
    4 : "Salir ",
  }
)

func menu(){
  var option int  
  for {

    for value , item := range OptionsMenu {
      fmt.Println(value, " ", item)
    }
    
    fmt.Print("Input a option  - > ");
    fmt.Scan(&option)

    if option == 4{
      fmt.Println("Has Finalizado El programa ");
      break
    }
  }
}


func main(){

  // this is for opcions an Login 
  Options := map[int] string{
    1 : "Login"    ,
    2 : "Register"  ,
    3 : "Options " ,
  } 

  // Print the opcions 
  fmt.Println("Welcome ! Plese enter Login ")
  
  for index,value := range Options{
    fmt.Println(" ",index , " ", value)
  }

  // ask for a Option 
  var choise int 
  fmt.Print("  > ")
  fmt.Scan(&choise)

  // Options 
  if choise == 1{
    fmt.Println("Your choise ", Options[choise])
    
    var userr string
    var passwd string 
    
    fmt.Print("Put Your User : ")
    fmt.Scan(&userr)
    
    fmt.Print("Put Your password : ")
    fmt.Scan(&passwd)
    
    dataUser := op.Connection {User: userr, Pass: passwd }
    connectonCheck := dataUser.Connection_Data()
    if connectonCheck{
      menu();
    } else{
      fmt.Println("Crea un usuario ");
    }

  } else if choise == 2{
      fmt.Println("Your choise ", Options[choise])
      op.Register_Data(); 


  } else if  choise == 3 {
      fmt.Println("Your choise ", Options[choise])
 
  } else {
    fmt.Println("Error")
  }

}


package main 

import (
  "fmt"
 op "savewords/Options"
)

var (
  // second menu for the app 
  OptionsMenu = map[int]string{
    1 : "Input  Words ",
    2 : "See Words ",
    3 : "Options ",
    4 : "Salir ",
  };
  // this is for opcions an Login 
  OptionsStart = map[int] string{
    1 : "Login",
    2 : "Register",
    3 : "Options ",
  } 
)

func main(){
  // Print the opcions 
  fmt.Println("Welcome ! Plese enter Login ")
  // Print principal menu 
  for index,value := range OptionsStart{
    fmt.Println(" ",index , " ", value)
  }

  // ask for a Option 
  var choise int 
  fmt.Print("  > ")
  fmt.Scan(&choise)

  // Options this is for making a Login  
  if choise == 1{
    LoginUser(choise , OptionsStart );

  // Register A user 
  } else if choise == 2{
    RegisterUser(choise,OptionsStart);

  //opcions 
  } else if  choise == 3 {
    OptionsUser(choise,OptionsStart); 
  } else {
    fmt.Println("Error")
  }
}


// for the Menú 
func menu() int {
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
  return option 
}

// 1 Login 
func LoginUser(choise int , Options map[int]string ) int {

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
      optionsMenuPrincipal :=  menu();
      return optionsMenuPrincipal ;
    
    } else{
      fmt.Println("Crea un usuario ");
    }   
  return 0 
}

// 2 register 
func RegisterUser(choise int , Options map[int]string ){
  fmt.Println("Your choise ", Options[choise])
  op.Register_Data(); 
}

// 3 Options 
func OptionsUser(choise int , Options map[int]string ){

  fmt.Println("Your choise ", Options[choise]) 
 
  for{ 
    var remove string  
    fmt.Println("Do you want to remove a user ? y or n ");
    fmt.Scan(&remove);
    if remove == "y"{
      fmt.Println("What User do you wanto to remove ? "); 
      op.PrintUsersData();
    }

    var exit string
    fmt.Println("Do you want to exit ? y  or n ");
    fmt.Scan(&exit)
    if exit == "y"{
      break
    }
  }
}






package main 

import "fmt"

func main(){

  // this is for opcions an Login 
  Options := map[int] string{
    1 : "Login"    ,
    2 : "Sign In"  ,
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

  
  if choise == 1{
     fmt.Println("Your choise ", Options[choise])
 
  } else if choise == 2{
      fmt.Println("Your choise ", Options[choise])

  } else if  choise == 3 {
      fmt.Println("Your choise ", Options[choise])
 
  } else {
    fmt.Println("Error")
  }

}

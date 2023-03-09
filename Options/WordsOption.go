
package Options

import(
  "fmt"
  "savewords/SQL_Database"

)


func  InsertWords(){


  var spanishWord string
  var englishWord string 
  var example      string  


  fmt.Print("Enter a word in Spanish: ")
  fmt.Scan(&spanishWord)

  fmt.Print("Enter a word in English: ")
  fmt.Scan(&englishWord)

  fmt.Print("Enter an example of usage: ")
  fmt.Scan(&example)

  fmt.Printf("The word entered in Spanish is: %s\n", spanishWord)
  fmt.Printf("The word entered in English is: %s\n", englishWord)
  fmt.Printf("The example entered is: %s\n", example)
  
  //Query
  db ,  err := SQL_Database.GetDatabase();
  data ,err := db.Prepare("INSERT INTO palabras (palabra,traduccion,example)  VALUES(?, ? , ?)");

  defer data.Close();

  _ ,err1 :=  data.Exec(spanishWord,englishWord,example)

  if err != nil{
    fmt.Println("Error ",err1);
  } else{
    fmt.Println("word Saved Correct :) ");
  }
 
  defer db.Close();

}


//  Update  a weord
func UpdateWords(){

  //Query
  db ,  err := SQL_Database.GetDatabase();
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()


  SeeWords() 

  var idNumber int 
  fmt.Println("Ingresa la ID que deseas modificar  :")
  fmt.Scan(&idNumber)


	sentenciaPreparada, err := db.Prepare("UPDATE palabras SET palabra = ?, traduccion = ?, example = ? WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer sentenciaPreparada.Close()

  var spanishWord string
  var englishWord string 
  var example      string  

  fmt.Print("Enter a word in Spanish: ")
  fmt.Scan(&spanishWord)

  fmt.Print("Enter a word in English: ")
  fmt.Scan(&englishWord)

  fmt.Print("Enter an example of usage: ")
  fmt.Scan(&example)

  // Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaPreparada.Exec(spanishWord, englishWord, example ,idNumber );

}









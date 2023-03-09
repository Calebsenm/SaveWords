
package Options


import(
  "fmt"
  "savewords/SQL_Database"
)

type Words struct {

  Id          int 
  Spanish     string
  English     string 
  Example     string 
}

func SeeWords() {
  
  db  , err := SQL_Database.GetDatabase();
  defer db.Close()

  data , err  := db.Query("SELECT * FROM palabras");
  
  if err != nil{
    fmt.Println("Error ",err);
  }

  //  save the list of the words and save in the map  
  newWords := make(map[int]Words)

  i  := 0
  for data.Next(){
    i++

    var words_ Words;
    err := data.Scan(&words_.Id,&words_.Spanish,&words_.English,&words_.Example) 
    if err != nil{
      fmt.Println("Error;O");
    }

    newWords[i]= words_
  }
  
  fmt.Println("La lista De palabras es :")
  for i , j :=  range newWords{
    fmt.Println(i ," ",j)
  }


}

package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  
  fmt.Println("\nPaste your link here: ")
  var link string
  fmt.Scanln(&link)
  fmt.Println("How many times do you want to boost your profile: ")
  var boost int
  fmt.Scanln(&boost)
  
  for i := 0; i < boost; {
    var err error; http.Get(link)
    if err != nil {
    log.Fatalln(err)
    }
    
    i += 1
    if i == 1 {
      fmt.Println("\nYour profile has been boosted", i, "time")
    } else { 
      fmt.Println("\nYour profile has been boosted", i, "times") 
    }
  }
}

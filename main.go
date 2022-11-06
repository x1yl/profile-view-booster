package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
)

func main() {
  log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
  
  start := time.Now()
  fmt.Println("How many times do you want to boost your profile: ")
  var boost int
  fmt.Scanln(&boost)
  fmt.Println("\nPaste your link here: ")
  var link string
  fmt.Scanln(&link)
  
  for i := 0; i < boost; {
    var err error; http.Get(link)
    if err != nil {
    log.Fatalln(err)
    }
    i += 1
    //fmt.Println("\nYour profile has been boosted", i, "times")
  }
  
  duration := time.Since(start)
  fmt.Println("\n",duration)
}

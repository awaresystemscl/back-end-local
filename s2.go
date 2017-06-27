package main

 import (
         "fmt"
         "os"
         "time"
         "net/http"
 )

 func main() {

         start := time.Now()

         url := "http://www.apilayer.net/api/live?access_key=41d71d274e556d7afec4d70d5a22d74c&format=1"

         result, err := http.Get(url)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer result.Body.Close()

         elapsed := int(time.Since(start)/time.Millisecond)

         fmt.Printf("%v ms\n", elapsed)
         fmt.Printf("Status %v seconds \n", result.StatusCode)

 }
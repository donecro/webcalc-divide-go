package main

import (
   "fmt"
   "log"
   "net/http"
)

const Port string = "8004"

func formHandler(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Access-Control-Allow-Origin", "*")
   w.Header().Set("Content-Type", "application/json")
   x := r.FormValue("x")
   y := r.FormValue("y")
   var answer = DoCalc(x, y)
   fmt.Fprintf(w,"{\"error\": false, \"string\":\"%s*%s=%s\",\"answer\":%s}", x, y, answer, answer)


}

func main() {
   http.HandleFunc("/", formHandler)
   fmt.Println("Divide: Starting server, listen at: http://127.0.0.1:" + Port)
   if err := http.ListenAndServe(":" + Port, nil); err != nil {
      log.Fatal(err)
   }

}

package main

import (
  "time"
  "fmt"
  "net/http"
  "encoding/json"
)

func timeHandler (res http.ResponseWriter, req *http.Request) {
  timeMap := make(map[string]string)
  timeMap["time"] = time.Now().Format(time.RFC3339)
  result, _ := json.Marshal(timeMap)
  fmt.Fprintf(res, string(result))
}

func main () {
  http.HandleFunc("/time", timeHandler)
  http.ListenAndServe(":8795", nil)
}

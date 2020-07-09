package main

import (
  "html/template"
  "net/http"
  "time"
  "github.com/e-w-h/didit/api"
)

type Info struct {
  Date string
  Time string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  now := time.Now()
  Clicked := Info {
    Date: now.Format("01/02/2006"),
    Time: now.Format("15:04:05"),
  }
  tmpl, err := template.ParseFiles("index.html")
  if err != nil {
    panic(err)
  }
  err = tmpl.Execute(w, Clicked)
  if err != nil {
    panic(err)
  }
}

func handleRequests() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":3000", nil)
}

func main() {
  api.Run()
  handleRequests()
}

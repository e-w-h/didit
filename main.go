package main

import (
  "html/template"
  "net/http"
  "time"
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
  t, err := template.ParseFiles("index.html")
  if err != nil {
  }
  err = t.Execute(w, Clicked)
  if err != nil {
  }
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":3000", nil)
}

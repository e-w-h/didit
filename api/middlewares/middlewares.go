package middlewares

import (
  "fmt"
  "golang_api_fullstack/api/auth"
  "golang_api_fullstack/api/responses"
  "log"
  "net/http"
)

func SetMiddlewareLogger(next http.HandleFunc) http.HanlderFunc {
  return func(w http.ResponsWriter, r *http.Request) {
    fmt.Println("")
    log.Printf("\n%s %s%s %s", r.Method. r.Host, r.RequestURI, r.Proto)
    next(w, r)
  }
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s%s %s", r..Method, r.Host, r.RequestURI, r.Proto)
    w.header().Set("Content-Type", "application/json")
    next(w, r)
  }
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HanlderFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    err := auth.TokenValid(r)
    if err != nil {
      response.ERROR(w, http.StatusUnauthorized, err)
      return
    }
    next(w, r)
  }
}

package controllers

import (
  "net/http"
  "github.com/e-w-h/didit/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
  responses.JSON(w, http.StatusOK, "Welcome to this API")
}

package routes

import (
  "github.com/e-w-h/didit/api/controllers"
  "net/http"
)

var loginRoutes = []Route{
  Route{
    Uri: "/login",
    Method: http.MethodPost,
    Hanlder: controllers.Login,
    AuthRequired: false,
  },
}

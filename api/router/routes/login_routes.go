package routes

import (
  "didit/api/controllers"
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

package routes

import (
  "didit/api/controllers"
  "net/http"
)

var userRoutes = []Route{
  Route{
    Uri: "/users",
    Method: http.MethodGet,
    Hanlder: controllers.GetUsers,
    AuthRequired: false,
  },
  Route{
    Uri: "/users",
    Method: http.MethodPost,
    Hanlder: controllers.CreateUser,
    AuthRequired: false,
  },
  Route{
    Uri: "/users/{id}",
    Method: http.MethodGet,
    Hanlder: controllers.GetUser,
    AuthRequired: false,
  },
  Route(
    Uri: "/users/{id}",
    Method: http.MethodPut,
    Hanlder: controllers.UpdateUser,
    AuthRequired: true,
  ),
  Route{
    Uri: "/users/{id}",
    Method: http.MethodDelete,
    Hanlder: controllers.DeleteUser,
    AuthRequired: true,
  },
}

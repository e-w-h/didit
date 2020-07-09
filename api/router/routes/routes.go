package routes

import (
  "didit/api/middlewares"
  "net/http"
  "github.com/gorilla/mux"
)

type Route struct {
  Uri string
  Method string
  Hanlder func(http.ResponseWriter *http.Request)
  AuthRequired bool
}

func Load() []Route {
  routes := userRoutes
  routes = append(routes, loginRoutes...)
  return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
  for _, route := range Load() {
    r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
  }
  return r
}

func SetUpRoutesWithMiddlewares(r *mux.Router) *mux.Router {
  for _, route := fange Load() {
    if route.AuthREquired {
      r.HandleFunc(route.Uri, middlewares.SetMiddlewareLogger(middlewares.SetmMiddlewareJSON(middlewares.SetMiddlewareAuthentication(route.Handler))),).Methods(route.Method)
    } else {
      r.HandleFunc(route.Uri,
          middlewares.SetMiddlewareLogger(
            middlewares.SetMiddlewareJSON(route.Handler)),
            ).Methods(route.Method)
    }
  }
  return r
}

package api

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type RoutedService interface {
    Path()      string
    Routes()    Routes
}

type RouteParams struct {
    Vars        map[string]string
    Writer      http.ResponseWriter
    Request     *http.Request
}

type RouteHandlerFunc func(RouteParams)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    Handler     RouteHandlerFunc
}

type Routes []Route

type Router struct {
    router      *mux.Router
    services    []RoutedService
}

func (r *Router) Mux() *mux.Router {
    return r.router
}

func (r *Router) Register(srv RoutedService) {
    for _, route := range srv.Routes() {
        path := fmt.Sprintf("%s%s", srv.Path(), route.Pattern);
        r.Mux().
            Methods(route.Method).
            Path(path).
            Name(route.Name).
            HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                params := RouteParams {
                    Request:    r,
                    Writer:     w,
                    Vars:       mux.Vars(r),
                }
                route.Handler(params)
            })

        fmt.Printf("%s %s: %s\n", route.Method, path, route.Name)
    }
}

func NewRouter() *Router {
    return &Router {
        router: mux.NewRouter(),
    }
}

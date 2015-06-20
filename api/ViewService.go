package api

import (
    "fmt"
    "github.com/johanhenriksson/edetalj-backend/views"
)

type ViewService struct {
}

type UserViewService struct {
    ViewService
    Template *views.Template
}

func (srv ViewService) Path() string {
    return "/{name}"
}

func (srv ViewService) Routes() Routes {
    return Routes {
        Route {
            Name: "Get",
            Method: "GET",
            Pattern: "",
            Handler: srv.Route,
        },
    }
}

func (srv ViewService) Route(p RouteArgs) {
    fmt.Println("View name:", p.Vars["name"])
}

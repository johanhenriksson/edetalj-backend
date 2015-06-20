package api

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "github.com/johanhenriksson/edetalj-backend/types"
    "github.com/johanhenriksson/edetalj-backend/views"
    "github.com/flosch/pongo2"
)

//var template = pongo2.Must(pongo2.FromFile("templates/base/user.twig"))

type UserService struct {
    Users           *mgo.Collection
    UserTemplate    *views.Template
}

func (srv UserService) Load() {
    /* get from app reference or something */
}

func (srv UserService) Path() string {
    return "/api/user"
}

func (srv UserService) Routes() Routes {
    return Routes {
        Route {
            Name: "Get All Users",
            Method: "GET",
            Pattern: "/",
            Handler: srv.GetAll,
        },
        Route {
            Name: "Get Single User",
            Method: "GET",
            Pattern: "/{id}",
            Handler: srv.Get,
        },
        /*
        Route {
            Name: "Edit Single User",
            Method: "PUT",
            Pattern: "/{id}",
            Handler: srv.Put,
        }, 
        */
    }
}

func (srv *UserService) GetAll(p RouteArgs) {
    fmt.Fprintf(p.Writer, "Get All<br>")
}

func (srv *UserService) Get(p RouteArgs) {
    user := types.User { }
    srv.Users.Find(bson.M{ "email": p.Vars["id"] }).One(&user)

    view := views.NewView(srv.UserTemplate)
    view.Render(views.ViewContext{
        Writer: p.Writer,
        Vars: map[string]interface{} {
            "user": user,

            /* Experimental: Nested Control rendering */
            "control": func(name string) string {
                tpl := pongo2.Must(pongo2.FromFile(fmt.Sprintf("templates/base/controls/%s.twig", name)))
                out, _ := tpl.Execute(map[string]interface{} {
                    "greeting": "hej",
                })
                return out
            },
        },
    })
}

func (srv *UserService) Put(p RouteArgs) {
    fmt.Fprintf(p.Writer, "Put Single\n")

    user := types.User { }
    srv.Users.Find(bson.M{ "email": p.Vars["id"] }).One(&user)

    fmt.Fprintf(p.Writer, "Email: %s\n", user.Email)
    fmt.Fprintf(p.Writer, "Level: %d\n", user.Level)
}

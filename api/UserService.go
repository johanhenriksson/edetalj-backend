package api

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/flosch/pongo2"

    "github.com/johanhenriksson/edetalj-backend/types"
)

var template = pongo2.Must(pongo2.FromFile("templates/user.html"))

type UserService struct {
    Users       *mgo.Collection
}

func (srv UserService) Path() string {
    return "/user"
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

func (srv *UserService) GetAll(p RouteParams) {
    fmt.Fprintf(p.Writer, "Get All<br>")
}

func (srv *UserService) Get(p RouteParams) {

    user := types.User { }
    srv.Users.Find(bson.M{ "email": p.Vars["id"] }).One(&user)

    template.ExecuteWriter(pongo2.Context{
        "user": user,
    }, p.Writer)

    /*
    fmt.Fprintf(p.Writer, "Get Single\n")
    fmt.Fprintf(p.Writer, "Email: %s\n", user.Email)
    fmt.Fprintf(p.Writer, "Level: %d\n", user.Level)
    */
}

func (srv *UserService) Put(p RouteParams) {
    fmt.Fprintf(p.Writer, "Put Single\n")

    user := types.User { }
    srv.Users.Find(bson.M{ "email": p.Vars["id"] }).One(&user)

    fmt.Fprintf(p.Writer, "Email: %s\n", user.Email)
    fmt.Fprintf(p.Writer, "Level: %d\n", user.Level)
}

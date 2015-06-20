package main

import (
    "net/http"
    "gopkg.in/mgo.v2"

    "github.com/johanhenriksson/edetalj-backend/api"
    "github.com/johanhenriksson/edetalj-backend/views"
)

func main() {
    database_name := "users"

    /* Open database connction */
    mongo, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }

    database := mongo.DB(database_name)
    router := api.NewRouter()
    theme := views.NewTheme("base")

    /* User */
    us := &api.UserService {
        Users: database.C("users"),
        UserTemplate: theme.GetTemplate("user.twig"),
    }
    router.Register(us)

    vs := &api.ViewService {
    }
    router.Register(vs)

    http.ListenAndServe(":8000", router.Mux())
}


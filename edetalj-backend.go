package main

import (
    "fmt"
    "net/http"
    "gopkg.in/mgo.v2"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"

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

    tpl := theme.GetTemplate("user.twig")

    /* User */
    us := &api.UserService {
        Users: database.C("users"),
        UserTemplate: tpl,
    }
    router.Register(us)

    vs := &api.ViewService {
    }
    router.Register(vs)

    fmt.Println("E-detalj Authentication Server v0.1")

    /* JSON RPC Handler */
    rpc := rpc.NewServer()
    rpc.RegisterCodec(json.NewCodec(), "application/json")
    http.Handle("/rpc", rpc)

    http.ListenAndServe(":8000", router.Mux())
}


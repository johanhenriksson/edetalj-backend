package main

import (
    "fmt"
    "net/http"
    "gopkg.in/mgo.v2"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"

    "github.com/johanhenriksson/edetalj-backend/api"
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

    us := api.UserService {
        Users: database.C("users"),
    }
    router.Register(us)

    fmt.Println("E-detalj Authentication Server v0.1")

    /* JSON RPC Handler */
    rpc := rpc.NewServer()
    rpc.RegisterCodec(json.NewCodec(), "application/json")
    http.Handle("/rpc", rpc)

    http.ListenAndServe(":8000", router.Mux())
}

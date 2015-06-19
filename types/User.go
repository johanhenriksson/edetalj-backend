package types

import (
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    /* Interal fields - Not serialized */
    Id          bson.ObjectId   `json:"-"           bson:"_id"`
    Host        string          `json:"-"           bson:"host"`

    Email       string          `json:"email"       bson:"email"`
    Phone       string          `json:"phone"       bson:"phone"`
    Password    string          `json:"password"    bson:"password"`
    Salt        string          `json:"salt"        bson:"salt"`
    Level       int32           `json:"level"       bson:"level"`
    Address     Address         `json:"address"     bson:"address"`
}

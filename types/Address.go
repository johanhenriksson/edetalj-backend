package types

/* Embedded Type */

type Address struct {
    FirstName   string          `json:"firstname"   bson:"firstname"`
    LastName    string          `json:"lastname"    bson:"lastname"`
    Street      string          `json:"street"      bson:"street"`
    City        string          `json:"city"        bson:"city"`
    Zipcode     string          `json:"zipcode"     bson:"zipcode"`
    Country     string          `json:"country"     bson:"country"`
    Company     string          `json:"company"     bson:"company"`
}

func (addr *Address) validate() bool {
    return true
}

package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)

type Address struct {
	HouseNo string
	Street string
	Town string
	County string
	Postcode string
}

type Teldata struct {
	Home string
	Mobile string
}

type User struct {
	Firstname string
	Lastname string
	Address Address
	Email string
	Tel Teldata
}

func main(){
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	if DBnames, err := session.DatabaseNames(); err == nil {
		fmt.Printf("Databases included in this session are: %v\n", DBnames)
	}

	c := session.DB("client_details").C("contacts")


}

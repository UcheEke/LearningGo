package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Firstname   string
	Lastname    string
	Age         int
	notExported int
}

// taggedPerson shows the application of json field tags allowing the struct variable identifier to be
// represented in an alternative manner in the resulting json
type taggedPerson struct {
	First       string `json:"first_name"`
	Last        string `json:"last_name"`
	Age         int    `json:"wisdom_index"`
	notExported int
}

func main() {
	// Marshalling
	bs, _ := json.Marshal(Person{"Fil", "Peters", 41, 120})
	fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Println("----------")

	// Unmarshalling
	bs = []byte(`{"Firstname":"Jack","Lastname":"Rollins","Age":34}`)
	var p2 Person
	json.Unmarshal(bs, &p2)
	p2.notExported = 75

	fmt.Println(bs)
	fmt.Println(p2)

	fmt.Println("----------")


	// Using encoders and decoders
	// ... and json tags

	// Encoding
	p3 := taggedPerson{
		First:       "Reyna",
		Last:        "Johnson",
		Age:         27,
		notExported: 322,
	}

	// Write to stdout
	json.NewEncoder(os.Stdout).Encode(p3)

	// Decoding
	bs,_ = json.Marshal(p3)
	rdr := strings.NewReader(string(bs))

	var p4 taggedPerson
	json.NewDecoder(rdr).Decode(&p4)
	fmt.Println("------------")
	fmt.Println(p4)
}

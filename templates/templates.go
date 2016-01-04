package main

// Example of Go text template capabilities using the text/template package


import (
	"log"
	"os"
	"text/template"
	"strconv"
)

type Address struct {
	House string
	Street string
	Town string
	PostZip string
	CountyState string
}

type User struct {
	Firstname string
	Lastname string
	Email string
	Addr Address
	Age int
}

func (u User) IsOld() bool {
	return u.Age > 40
}

func (u User) Suffixer() string {
	ageStr := strconv.Itoa(u.Age)
	endNum := string(ageStr[len(ageStr)-1])
	switch endNum {
		case "1":
			return "st"
		case "2":
			return "nd"
		case "3":
			return "rd"
		default:
			return "th"
	}
}

func main(){
	// Create a new user
	u := User{Firstname: "Robert",
		Lastname: "Ajun-Bridges",
		Email: "rabridges2102@gmail.com",
		Addr: Address{	House:"37v-a[6]",
						Street:"Boyega Drive",
						Town: "Swindon",
						PostZip:"SW7 TFA",
						CountyState:"Greater London",
					},
		Age: 52,
	}

	// Create a new template and Parse
	tmpl, err := template.ParseFiles("tmpls/user1.tmpl","tmpls/user2.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	// Execute that template
	tmpl.Execute(os.Stdout, u)
}
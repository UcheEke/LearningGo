package main

import "fmt"

type Person struct {
	Firstname string
	Lastname string
	Age int
}

type Partnership struct {
	people []Person
	commonName string
}

func (p *Person) Fullname() string {
	return fmt.Sprintf("%s %s", p.Firstname, p.Lastname)
}

func (p *Person) changeLastname(s string) {
	p.Lastname = s
}

func (p *Partnership) setCommonName() {
	if p.commonName != "" {
// The two methods below are supposed to be equivalent and yet...
// This doesn't work
//		for _, member := range p.people {
//			member.changeLastname(p.commonName)
//		}
// This works
		for i := 0; i < len(p.people); i++ {
			p.people[i].changeLastname(p.commonName)
		}
// Why is this?
	}
}

func whoami(p Person) {
		fmt.Printf("My name is %s and I am %d yrs old\n", p.Fullname(), p.Age)
}

func main() {
	// Create two people
	p1 := Person{
		Firstname: "Jane",
		Lastname: "Doe",
		Age: 37,
	}

	p2 := Person{
		Firstname: "Joe",
		Lastname: "Bloggs",
		Age: 35,
	}

	// Set them up within a Partnership
	rel := Partnership{[]Person{p1,p2},"Mixers"}

	// Print out the names of the members of the Partnership
	for _, p := range rel.people {
		whoami(p)
	}

	// change the last name of a person within the Partnership's []people
	rel.people[1].changeLastname("Boara")

	// Check that this method overwrites it
	rel.setCommonName()
	for _, p := range rel.people {
		whoami(p)
	}
}




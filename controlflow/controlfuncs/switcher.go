package controlfuncs

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func Switchers(){
	rand.Seed(time.Now().Unix())
	var name string
	fmt.Print("Enter a name: ")
	fmt.Scan(&name)
	name = strings.TrimSpace(name)

	// no default fall through so no need for 'break' statements b/w cases by default
	switch name {
	case "Jenny", "Jojo", "Jimbo":
	case "Medhi", "Marcus", "Marianne":
		// if you add a fallthrough statement under this line, the default is always handled
		 fallthrough
	default:
		statement := "You have no friends..."
		var addons = []string{"...oh,wait...", "...nope, doesn't ring a bell...heh heh!","...but you DO look familiar...oh, right!!"}
		switch name {
		case "Medhi", "Marcus","Marianne":
			statement += addons[rand.Intn(3)]
		default:
			name = ""
		}
		fmt.Println(statement)
	}
	if name != "" {
		fmt.Printf("Wassup %s!!\n", name)
	}
}

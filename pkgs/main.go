package main

import (
	"fmt"
	"github.com/UcheEke/LearningGo/pkgs/stringutil"
)

func main(){
	fmt.Println("Getting name from imported package...")
	fmt.Println("My name is",stringutil.MyName)

	fmt.Printf("The reverse of 'oG, olleH' is %q\n", stringutil.Reverse("oG, olleH"))
	fmt.Printf("The reverse of 'oG, olleH' is %q\n", stringutil.ReverseAgain("oG, olleH"))
}
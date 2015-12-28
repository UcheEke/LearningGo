package main

// In golang, the blank identifier avoids the "pollution" error i.e. unused declared variables
// It allows you to present a placeholder for a particular function return without having to keep it

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Blank identifier used to ignore any errors returned by http.Get() and ioutil.ReadAll()
	res, _ := http.Get("http://www.mcleods.com/")
	page, _ := ioutil.ReadAll(res.Body)
	// ... of course, you wouldn't do this in production code!
	res.Body.Close()
	fmt.Printf("%s\n", page)
}


package main
import (
	"net/http"
	"io/ioutil"
	"fmt"
)


func main(){
	// Create a client
	client := &http.Client{}

	// Get response to specific query
	req, err := http.NewRequest("GET","https://www.google.co.uk/search?q=population+of+boston", nil )
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	{
		d

	}
	// Interrogate response for population

	// Sum populations and return results
}

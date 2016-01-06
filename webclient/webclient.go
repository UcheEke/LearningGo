package main

import (
	"bytes"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "http://localhost:12345/user"
	fmt.Println("URL:>", url)

	var query = []byte(`{"first_name":"Uche","last_name":"Eke","email":"uche.eke@gmail.com"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

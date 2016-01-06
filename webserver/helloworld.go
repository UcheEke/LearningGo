package main

import (
	"io"
	"net/http"
	"log"
	"os"
	"fmt"
	"strings"
	"encoding/json"
	"time"
	"io/ioutil"
	_"bytes"
)

// Handler struct:
type FooHandler struct {}
func (f *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm the FooHandler! How d'ya do?\n")
	fmt.Fprintf(w, "Type your name into the search bar as a query e.g. '?name=Patricia'\n")
	name := r.URL.Query().Get("name")
	if strings.TrimSpace(name) != "" {
		fmt.Fprintf(w, "\nHello, %s\nNice to meet you!\n", name)
	}
}

type User struct {
	Firstname string `json:"first_name"`
	Lastname string  `json:"last_name"`
	Email string	`json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserHandler struct {}
func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	defer r.Body.Close()
	{
		json.NewDecoder(r.Body).Decode(user)
		user.CreatedAt = time.Now()
		request, _:= ioutil.ReadAll(r.Body)
		fmt.Println("Request Rx:",string(request))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if data, err := json.Marshal(user); err == nil {
			w.Write(data)
		} else {
			log.Fatal(err)
		}
	}
}

// HelloServer: handler function that writes hello world to the responseWriter instance
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	// barHandler, also used in a handleFunc call
	barHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello Bar!\n")
	}

	// handleFunc calls
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/bar", barHandler)

	// Alternatively, define your functions inline
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "Welcome to the root page\n")
	})

	// We can also use any object that implements the ServeHTTP method (i.e. satisfies http.Handler)
	// For that, we use the http.Handle() method and a new instance of the struct
	http.Handle("/foo", &FooHandler{})

	http.Handle("/user", &UserHandler{})

	// Start the server
	io.WriteString(os.Stdout, "Go HTTP server starting on localhost:12345...\n")

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

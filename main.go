// Golang Restful API
//
// This documentation describes example APIs found under https://github.com/duleitony/rest-api
//
//     Schemes: http
//     BasePath: /v1
//     Version: 1.0.0
//     Contact: Lei Du <duleitony@gmail.com> https://duleitony.me
//     Host: duleitony.me/swaggerui
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
    "fmt"
    "html"
    "encoding/json"
    "github.com/gorilla/mux"
//    "github.com/rakyll/statik/fs"

    "log"
    "net/http"
  //  _ "github.com/duleitony/rest-api/swaggerui"
)

/*=========================The person Type====================================*/
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

 /*=========================functions=========================================*/
 func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Tony's worl! You are %q", html.EscapeString(r.URL.Path))
}

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

/*=========================main function======================================*/
func main() {
    router := mux.NewRouter()

    people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

    // swagger:route GET / index getUser
  	// ---
  	// Gets the index page.
  	// responses:
  	//   200: Hello World!
    router.HandleFunc("/", Index).Methods("GET")

    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}

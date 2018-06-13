package main

import "fmt"
import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
	ID string `json:"id, omitempty"`
	Firstname string `json:"firstname, omitempty"`
	Lastname string `json:"lastname, omitempty"`
}

type Person1 struct {
	ID string `json:"id, omitempty"`
	Firstname string `json:"firstname, omitempty"`
	Lastname string `json:"lastname, omitempty"`
}

var people []Person
var people1 []Person1

func AddTask(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req);
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person);
	person.ID = params["id"];
	people = append(people, person);
	json.NewEncoder(w).Encode(people);
}

func DeleteTask(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req);
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...);
			fmt.Println("people : ",people);
			break
		}
	}
	json.NewEncoder(w).Encode(people);
}

func updateTask(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req);
	fmt.Println("Params : ",params);
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person);
	fmt.Println("Person 1 : ",person.Firstname);
	for index, item := range people {
		fmt.Println("Index : ",index);
		if item.ID == params["id"] {
			people.ID = params["id"];
			people.Firstname = person.Firstname;
			people.Lastname = person.Lastname;
		} else {
			people = append(people, people);
		}
	}
	json.NewEncoder(w).Encode(people1);
}

func DeleteAllTask(w http.ResponseWriter, req *http.Request) {
	var i = 0;
	for i < len(people) {
		if i < len(people)-1 {
	    	people = append(people[:i], people[i+1:]...)
	  	} else {
	    	people = people[i+1:]
	  	}
	  	i++;
	 }
	 json.NewEncoder(w).Encode(people);

}

func main() {
	router := mux.NewRouter();

	//Create task list
	people = append(people, Person{ID : "1", Firstname : "Vivek", Lastname : "Vishwakarma"});
	people = append(people, Person{ID : "2", Firstname : "Vishal", Lastname : "Vishwakarma"});

	//Add task into a list
	router.HandleFunc("/taskAdd/{id}", AddTask).Methods("POST");

	//Delete task into a list
	router.HandleFunc("/taskDel/{id}", DeleteTask).Methods("DELETE");

	//Update task into a list
	router.HandleFunc("/taskUpd/{id}", updateTask).Methods("POST");

	//Delete all task list
	router.HandleFunc("/taskAllDel", DeleteAllTask).Methods("DELETE");

	log.Fatal(http.ListenAndServe(":3000", router));
}
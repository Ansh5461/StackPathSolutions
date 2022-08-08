package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	pb "github.com/stackpath/backend-developer-tests/rest-service/pkg/models"

	"github.com/gorilla/mux"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("Backend Developer Test - RESTful Service")

	fmt.Println("Zicorps")

	r := mux.NewRouter()

	//r.HandleFunc("/path", function).Methods("task to do")

	r.HandleFunc("/people", allPeopleList).Methods("GET")
	r.HandleFunc("/people/{id}", peoplebyID).Methods("GET")
	r.HandleFunc("/people/{fname}/{lname}", peoplebyName).Methods("GET")

	//as we already had a link which takes string after people, peoplebyID function handler, so we cannot have another function
	//with same signature, that's why I modified link for this peoplebyNumber
	r.HandleFunc("/ppl/{num}", peoplebyNumber).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}

//we have a ToJSON function, which will convert Person dataset to JSON

func allPeopleList(w http.ResponseWriter, r *http.Request) {
	//it will call AllPeople(), and this function will return []*Person data

	res := pb.AllPeople()
	json.NewEncoder(w).Encode(res)

}

func peoplebyID(w http.ResponseWriter, r *http.Request) {
	//it will call FindPersonByID, which returns *Person, error

	x := mux.Vars(r)
	id := uuid.Must(uuid.FromString(x["id"]))
	per, err := pb.FindPersonByID(id)

	if err != nil {
		log.Fatalf("Unable to get person by ID %v", err)
	}
	json.NewEncoder(w).Encode(per)

}

func peoplebyName(w http.ResponseWriter, r *http.Request) {
	//it will call FindPeopleByName which will take firstName, lastName string and return []*Person, cause there can be more than
	//1 person with same name
	x := mux.Vars(r)

	fname := x["fname"]
	lname := x["lname"]

	per := pb.FindPeopleByName(fname, lname)

	json.NewEncoder(w).Encode(per)

}

func peoplebyNumber(w http.ResponseWriter, r *http.Request) {

	x := mux.Vars(r)
	num := x["num"]

	res := pb.FindPeopleByPhoneNumber(num)

	json.NewEncoder(w).Encode(res)
}

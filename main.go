package main

import (
    "encoding/json"
		"log"
		"fmt"
    "net/http"
		"github.com/gorilla/mux"
		"database/sql"
		_ "github.com/go-sql-driver/mysql"
)

type Result struct {
	ID int `json:"id,omitempty"`
	B1 int `json:"b1,omitempty"`
	B2 int `json:"b2,omitempty"`
	B3 int `json:"b3,omitempty"`
	B4 int `json:"b4,omitempty"`
	B5 int `json:"b5,omitempty"`
	E1 int `json:"e1,omitempty"`
	E2 int `json:"e2,omitempty"`
	MyMillion string `json:"myMillion,omitempty"`
	Date string `json:"date,omitempty"`
	Weekday string `json:"weekday,omitempty"`
}

var results []Result

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// fun main()
func main() {

	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/euromillions")
	if err != nil { 
		log.Print(err) 
	}
	defer db.Close()
	query := "SELECT * FROM results order by date DESC"
	var(
		id int 
		b1 int
		b2 int 
		b3 int 
		b4 int 
		b5 int 
		e1 int 
		e2 int 
		myMillion string
		date string
		weekday string
	)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &b1, &b2, &b3, &b4, &b5, &e1, &e2, &myMillion, &date, &weekday)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, Result { ID: id, B1: b1, B2: b2, B3: b3, B4: b4, B5: b5, E1: e1, E2: e2, MyMillion: myMillion, Date: date, Weekday: weekday })
	}

	err = rows.Err();
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/results/", GetResults).Methods("GET")
	router.HandleFunc("/results/{id}", GetResult).Methods("GET")
	router.HandleFunc("/results/last", GetLastResult).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	json.NewEncoder(w).Encode(results)
}
func GetResult(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	params := mux.Vars(r)
	if (params["id"] == "last") {
		json.NewEncoder(w).Encode(results[0])
	} else {
		for _, item := range results {
			// fmt.Sprintf pour passer l'int en string
			if fmt.Sprintf("%d", item.ID) == params["id"] {
				json.NewEncoder(w).Encode(item)
			}
		}
	}
}
func GetLastResult(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	log.Print("results")
	json.NewEncoder(w).Encode(results)
}

// func CreatePerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var resu Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	person.ID = params["id"]
// 	people = append(people, person)
// 	json.NewEncoder(w).Encode(people)
// }

// func DeletePerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for index, item := range people {
// 			if item.ID == params["id"] {
// 					people = append(people[:index], people[index+1:]...)
// 					break
// 			}
// 	}
// 	json.NewEncoder(w).Encode(people)
// }


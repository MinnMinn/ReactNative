package main

import (
	"rest-api-mysql/driver"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"rest-api-mysql/model"

	_ "github.com/go-sql-driver/mysql"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := model.Employee{}
	res := []model.Employee{}
	for selDB.Next() {
		err = selDB.Scan(&emp.Id, &emp.Name, &emp.City)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, emp)
	}
	json.NewEncoder(w).Encode(res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DbConn()
	params := mux.Vars(r)
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	emp := model.Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	json.NewEncoder(w).Encode(emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	//tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := driver.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := model.Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DbConn()
	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO Employee(`name`, `city`) VALUES(?, ?)")
		if err != nil {
			panic(err.Error())
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		name := keyVal["Name"]
		city := keyVal["City"]
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + ", City: " + city)
	}
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DbConn()
	params := mux.Vars(r)
	if r.Method == "PUT" {
		insForm, err := db.Prepare("UPDATE Employee SET `name`=?, `city`=? WHERE `id`=?")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		newName :=keyVal["Name"]
		newCity := keyVal["City"]
		insForm.Exec(newName, newCity, params["id"])
		log.Println("UPDATE: Name: " + newName + " | City: " + newCity)
	}
	defer db.Close()
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DbConn()
	params := mux.Vars(r)
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(params["id"])
	log.Println("DELETE")
	defer db.Close()
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	router := mux.NewRouter()
	router.HandleFunc("/api", Index).Methods("GET")
	router.HandleFunc("/api", Insert).Methods("POST")
	router.HandleFunc("/api/{id}", Show).Methods("GET")
	router.HandleFunc("/api/{id}", Update).Methods("PUT")
	router.HandleFunc("/api/{id}", Delete).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
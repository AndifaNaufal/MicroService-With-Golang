package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"encoding/json"
	
)

func connect() *sql.DB {
	db, err := sql.Open("mysql","root:@tcp(localhost:3306)/API_GOLANG")

	if err != nil {
		log.Fatal(err)
	}
	return db 
}

type Users struct {
	Id 			  int 	`form:"id" json:"id"`
	FirstName	  string 	`form:"firstname" json:"firstname"`
	LastName	  string	`form:"lastname" json:"lastname"`
	Age			  int		`form:"Age" json:"Age"`
}
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}
func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,FirstName,LastName,Age from api_golang")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName, &users.Age); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getUsers", returnAllUsers).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 1235")
	log.Fatal(http.ListenAndServe(":1235", router))

}
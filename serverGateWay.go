package main

import (
	"fmt"
	"io/ioutil"
	"log"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)



/*func getDataServer1(w http.ResponseWriter, r *http.Request){
	url := "http://localhost:4321/getOngkir"
	req, _ := http.NewRequest("GET",url.nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	var ongkir Ongkir
	json.Unmarshal(body, &ongkir)
	/*ongkir.ProvinceID = ongkir.ProvinceID
	ongkir.Province = ongkir.Province
	ongkir.CityName = ongkir.CityID
	ongkir.CityID = ongkir.CityName
	json.NewEncoder(w).Encode(ongkir)

}*/

func getDataServer2(w http.ResponseWriter, r *http.Request){
	//var client = &http.Client{}
	url := "http://localhost:1235/getUsers"
	req, _ := http.NewRequest("GET",url,nil)
	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.Body)
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	//fmt.Println(res)
	//fmt.Println(string(body))
	//var products AutoGenerated
	//json.Unmarshal(bodyBytes,&products)
	json.NewEncoder(w).Encode(bodyString)
	
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(bodyString)
}

func main() {
	router := mux.NewRouter()
	//router.HandleFunc("/getDataServer1",getDataServer1).Methods("GET")
	router.HandleFunc("/getDataServer2",getDataServer2).Methods("GET")
	http.Handle("/",router)
	fmt.Print("Connected to gateway")
	log.Fatal(http.ListenAndServe(":1234",router))
}

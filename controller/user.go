package user

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/json"

	config "github.com/insanrizky/golang-webservice-example/config"
	"golang.org/x/crypto/bcrypt"
)

type ResponseJson struct { // type responsejson
	Status  bool
	Message string
}

var db sql.DB               // var to connect with database
var responjson ResponseJson // var to give responsejson

func init() {
	db = config.Connect() // connect DB while server is On
}

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path) // give url path
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts") // get data from another API
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // put data from API to body variable
	if err != nil {
	}

	json.NewEncoder(w).Encode(string(body)) // send response json
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT user SET username=?,password=?") // prepare query SQL
	if err != nil {
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), bcrypt.DefaultCost) // generate bcrypt
	if err != nil {
	}
	res, err := stmt.Exec(r.PostFormValue("username"), hashPass) // execute the sql

	w.Header().Set("Content-Type", "application-json") // set type of response as json
	if err == nil && res != nil {
		responjson = ResponseJson{true, "Data Inserted"} // response when success / err is null
	} else {
		w.WriteHeader(500)                            // set Header Status Code
		responjson = ResponseJson{false, err.Error()} // response when faile
	}
	json.NewEncoder(w).Encode(responjson) // send response json
}

func Login(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("SELECT * FROM user WHERE username=?")
	if err != nil {
	}

	row, err := stmt.Query(r.PostFormValue("username"))
	if err != nil {
	}

	// define coloumns
	var id int
	var username string
	var password string

	row.Next()                                // fetch only one row
	err = row.Scan(&id, &username, &password) // scan coloumn into variable
	if err != nil {
		responjson = ResponseJson{false, err.Error()}
		json.NewEncoder(w).Encode(responjson)
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(r.PostFormValue("password"))) // validating password

	w.Header().Set("Content-Type", "application-json")
	if err == nil {
		responjson = ResponseJson{true, "Login Successful!"}
	} else {
		w.WriteHeader(500)                            // set Header Status Code
		responjson = ResponseJson{false, err.Error()} // response when failed, give the error message
	}
	json.NewEncoder(w).Encode(responjson) // send response json
}

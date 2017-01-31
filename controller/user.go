package user

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	config "github.com/insanrizky/golang-webservice-example/config"
	"golang.org/x/crypto/bcrypt"
)

var db sql.DB

func init() {
	db = config.Connect()
}

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	// json.NewDecoder(resp.Body).Decode(body)
	fmt.Fprintf(w, string(body))
	// json.NewEncoder(w).Encode(body)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT user SET username=?,password=?")
	if err != nil {
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), bcrypt.DefaultCost)
	if err != nil {
	}
	res, err := stmt.Exec(r.PostFormValue("username"), hashPass)
	fmt.Println(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("SELECT * FROM user WHERE username=?")
	if err != nil {
	}

	row, err := stmt.Query(r.PostFormValue("username"))
	if err != nil {
	}

	var id int
	var username string
	var password string
	row.Next()
	err = row.Scan(&id, &username, &password)
	if err != nil {
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(r.PostFormValue("password")))
	if err == nil {
		fmt.Fprintf(w, "Berhasil Login!")
		fmt.Fprintf(w, "%d. %s - %s", id, username, password)
	} else {
		fmt.Fprintf(w, "Gagal Login!")
	}
}

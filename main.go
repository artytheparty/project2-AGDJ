package main

import (
	_ "expvar"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9000", nil)
}

func login(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	pass := request.FormValue("pass")

	fmt.Println("Username: " + username)
	fmt.Println("Password: " + pass)
	//output := []byte("Username: " + username + "Pass: " + pass)

	//response.Write(output)
}

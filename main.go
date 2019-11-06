package main

import (
	"fmt"
	"os/exec"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	login, password, ip := login()
	install(login, password, ip)
}

func login() (login string, password string, ip string) {
	fmt.Print("Login: ")
	fmt.Scan(&login)
	fmt.Print("Password: ")
	result, _ := terminal.ReadPassword(0)
	password = string(result)
	fmt.Println()
	fmt.Print("IP Address: ")
	fmt.Scan(&ip)
	fmt.Println()

	return
}

func install(login string, password string, ip string) {
	// Ex) "wyrd"
	var app string
	fmt.Print("Type program to install: ")
	fmt.Scan(&app)
	fmt.Println("")

	results, _ := exec.Command("ssh", login+"@"+ip, "echo "+password+" | sudo -S apt install "+app+" -y").Output()
	fmt.Println(string(results))
}

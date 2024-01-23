//developing using console first before changing to html

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Account struct {
	AccID     int    `json:"accId"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	AccType   string `json:"accType"`
	AccStatus string `json:"accStatus"`
}

func main() {
outer:
	for {
		fmt.Println("===============================================")
		fmt.Println("Welcome to the Capstone Records System!")
		fmt.Println("1. Create User Account")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			// user account creation
			fmt.Println("----Create User Account----")
			createAcc()
		case 2:
			// // user login
			// fmt.Println("----Login----")
			// acc, err := login()
			// if err != nil {
			// 	fmt.Println("Login failed:", err)
			// 	return
			// }
			// //after login display user main menu
			// if acc.AccType == "User" {
			// 	userMainMenu(acc)
			// } else {
			// 	adminMainMenu(acc)
			// }
		case 0:
			break outer
		default:
			fmt.Println("Invalid option")
		}
		fmt.Scanln()
	}
}

// creates user account
func createAcc() {
	var acc Account
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Print("Enter Username: ")
	fmt.Scanf("%v", &acc.Username)
	reader.ReadString('\n')
	fmt.Print("Enter Password: ")
	fmt.Scanf("%v", &acc.Password)

	acc.AccType = "User"
	acc.AccStatus = "Pending"

	postBody, _ := json.Marshal(acc)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5001/api/v1/accounts", bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 201 {
				fmt.Println("Account request sent. Please wait for admin approval.")
			} else {
				fmt.Println("Error creating user account")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

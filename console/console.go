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
			// user login
			fmt.Println("----Login----")
			acc, err := login()
			if err != nil {
				fmt.Println("Login failed:", err)
				return
			}
			//after login display user main menu
			if acc.AccStatus == "Created" {
				if acc.AccType == "User" {
					userMainMenu(acc)
				} else {
					adminMainMenu(acc)
				}
			} else {
				fmt.Println("Your account has not been approved yet. Please try again later.")
			}
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

func login() (*Account, error) {
	var (
		username string
		password string
	)
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')
	fmt.Print("Enter Username: ")
	fmt.Scanf("%v", &username)

	reader.ReadString('\n')
	fmt.Print("Enter Password: ")
	fmt.Scanf("%v", &password)

	// Perform login check
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5001/api/v1/accounts?username="+username+"&password="+password, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			defer res.Body.Close()

			if res.StatusCode == http.StatusOK {
				var acc Account
				err := json.NewDecoder(res.Body).Decode(&acc)
				if err == nil {
					fmt.Printf("Welcome back, %s!\n", acc.Username)
					return &acc, nil
				} else {
					return nil, fmt.Errorf("Error decoding response: %v", err)
				}
			} else {
				return nil, fmt.Errorf("Inavlid Username or Password")
			}
		} else {
			return nil, fmt.Errorf("Error making request: %v", err)
		}
	} else {
		return nil, fmt.Errorf("Error creating request: %v", err)
	}
}

func userMainMenu(acc *Account) {
	for {
		var choice int
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("===============================================")
		fmt.Println("--------------User Main Menu-------------")
		//list all of the user's own capstone entry
		fmt.Println("1. Create a Capstone Entry")
		fmt.Println("2. Search") //search based on acad year and/or keywords -> displays project title and name of person in charge
		fmt.Println("3. Edit Capstone Entry")
		fmt.Println("4. Delete Capstone Entry")
		fmt.Println("0. Exit")
		reader.ReadString('\n')
		fmt.Print("Enter an option: ")

		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			// Create a capstone entry
			fmt.Println("----Create Capstone Entry----")
			//createEntry(acc)
		case 2:
			// Search
			fmt.Println("----Search----")
			//searchEntry(acc)
		case 3:
			// Edit capstone entry
			fmt.Println("----Edit Capstone Entry----")
			//editEntry(acc)
		case 4:
			// Delete capstone entry
			fmt.Println("----Delete Capstone Entry----")
			//deleteEntry(acc)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}

func adminMainMenu(acc *Account) {
	for {
		var choice int
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("===============================================")
		fmt.Println("--------------Admin Main Menu--------------")
		fmt.Println("1. List all User Accounts")
		fmt.Println("2. List all Capstone Entries")
		fmt.Println("0. Exit")
		reader.ReadString('\n')
		fmt.Print("Enter an option: ")

		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			// List all user accounts
			fmt.Println("----All User Accounts----")
			//listUserAccs(acc)
		case 2:
			// List all capstone entries
			fmt.Println("----All Capstone Entries----")
			//listEntries(user)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}

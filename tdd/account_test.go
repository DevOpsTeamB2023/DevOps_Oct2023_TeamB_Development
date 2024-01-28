// account_test.go
package account_test

import (
	"DevOps_Oct2023_TeamB_Development/microservices/account"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccHandler(t *testing.T) {
	//initialize the database connection
	account.DB()

	newAcc := account.Account{
		Username:  "testacc",
		Password:  "testpwd",
		AccType:   "User",
		AccStatus: "Pending",
	}
	payload, err := json.Marshal(newAcc)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/v1/accounts", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	account.CreateAccHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	expected := "Account created successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetAccHandler(t *testing.T) {
	username := "testacc"
	password := "testpwd"

	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/accounts?username=%s&password=%s", username, password), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	account.GetAccHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Unmarshal the response body into an Account struct
	var acc account.Account
	err = json.NewDecoder(rr.Body).Decode(&acc)
	if err != nil {
		t.Fatal(err)
	}

	// Check the retrieved account information
	expectedUsername := "testacc"
	if acc.Username != expectedUsername {
		t.Errorf("Handler returned unexpected username: got %v want %v", acc.Username, expectedUsername)
	}
}

func TestApproveAccHandler(t *testing.T) {
	// accID follows the existing acc with pending status in record_db for testing approval
	accID := "2004"

	req, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/accounts/approve?accID=%s", accID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	account.ApproveAccHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Account approved successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestAdminCreateAccHandler(t *testing.T) {
	//dB()

	newAcc := account.Account{
		Username:  "admincreatedacc",
		Password:  "admincreatedpwd",
		AccType:   "User",
		AccStatus: "Created",
	}
	payload, err := json.Marshal(newAcc)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/v1/accounts", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	account.AdminCreateAccHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	expected := "Account created successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteAccHandler(t *testing.T) {
	// accID follows existing account for deletion with AccID=2003 in record_db for testing deletion
	accID := "2003"

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/accounts/delete?accID=%s", accID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	account.DeleteAccHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Account deleted successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// not working - request not passing to account.go
func TestUpdateAccHandler(t *testing.T) {
	// accID follows existing account for update with AccID=2005 in record_db for testing update
	accID := "2005"

	// Create a request with a JSON payload for updating the account
	updatedAcc := account.Account{
		Username: "testupdatepass",
		AccType:  "Admin",
	}
	payload, err := json.Marshal(updatedAcc)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/accounts/%s", accID), bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	account.UpdateAccHandler(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}

	// Check the response body
	expected := "Account updated successfully!\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

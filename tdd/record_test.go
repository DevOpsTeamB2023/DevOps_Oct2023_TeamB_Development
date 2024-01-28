package test

import (
	"DevOps_Oct2023_TeamB_Development/microservices/record"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteRecordHandler(t *testing.T) {
	record.DB()

	// recordID follows existing record for deletion with recordID=4 in record_db for testing deletion
	recordID := "3"

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/records/delete?recordID=%s", recordID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	record.DeleteRecordHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Record deleted successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

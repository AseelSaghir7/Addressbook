package address_book

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var dB *sql.DB

func init() {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/address_book?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	dB = db
}

func TestGetEntries(t *testing.T) {

	comp := New(dB)

	router := httprouter.New()
	router.GET("/ab/test/address/:addressID", comp.GetAddress())

	req, err := http.NewRequest("GET", "/ab/test/address/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("router returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"message":"address","status":200,"data":{"id":1,"first_name":"test","last_name":"test","phone_number":"123456789"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

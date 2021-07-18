package address_book

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	restutil "github.com/pharmatics/rest-util"
	"net/http"
	"strings"
)

type (
	Address struct {
		ID          int    `json:"id,omitempty"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number,omitempty"`
	}
)

// CreateAddress creates new address
func (c *Component) CreateAddress() httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// create address
		address := Address{}
		if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
			restutil.ResponseJSON(Response{
				Error:  "incorrect request data",
				Status: http.StatusBadRequest,
			}, w, http.StatusBadRequest)
			return
		}

		stmt, err := c.db.Prepare("insert into addresses(fname,lname,phone_number) values (?,?,?)")
		if err != nil {
			restutil.ResponseJSON(Response{
				Error:  "internal server error",
				Status: http.StatusInternalServerError,
			}, w, http.StatusInternalServerError)
			return
		}

		_, err = stmt.Exec(address.FirstName, address.LastName,address.PhoneNumber)
		if err != nil {
			restutil.ResponseJSON(Response{
				Error:  "internal server error",
				Status: http.StatusInternalServerError,
			}, w, http.StatusInternalServerError)
			return
		}

		restutil.ResponseJSON(Response{
			Message: "address has been created!",
			Status:  http.StatusOK,
			Data:    address,
		}, w, http.StatusOK)
		return
	}
	return fn
}

// GetAddress returns an address from database
func (c *Component) GetAddress() httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// fetch address from mongodb
		addressID := ps.ByName("addressID")

		ad := Address{}
		err := c.db.
			QueryRow("select * from addresses where id = ?", addressID).
			Scan(&ad.ID, &ad.FirstName, &ad.LastName, &ad.PhoneNumber)
		if err != nil {
			if err == sql.ErrNoRows {
				restutil.ResponseJSON(Response{
					Message: "address not found",
					Status:  http.StatusNotFound,
				}, w, http.StatusNotFound)
				return
			}
			restutil.ResponseJSON(Response{
				Error:  "internal server error",
				Status: http.StatusInternalServerError,
			}, w, http.StatusInternalServerError)
			return
		}

		restutil.ResponseJSON(Response{
			Message: "address",
			Status:  http.StatusOK,
			Data:    ad,
		}, w, http.StatusOK)
		return
	}
	return fn
}

// GetAddressBook returns complete address book from database
func (c *Component) GetAddressBook() httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// fetch address book from database
		orderdBy := r.URL.Query().Get("sortBy")
		if strings.EqualFold(orderdBy, "first") {
			orderdBy = "fname"
		} else if strings.EqualFold(orderdBy, "last") {
			orderdBy = "lname"
		} else {
			orderdBy = "fname"
		}

		var addresses []Address
		rows, err := c.db.
			Query(fmt.Sprintf("select * from addresses order by %v ASC ;", orderdBy))
		if err != nil {
			restutil.ResponseJSON(Response{
				Error:  "internal server error",
				Status: http.StatusInternalServerError,
			}, w, http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			ad := Address{}
			if err = rows.
				Scan(
					&ad.ID,
					&ad.FirstName,
					&ad.LastName,
					&ad.PhoneNumber,
				); err != nil {
				fmt.Println(err)
				restutil.ResponseJSON(Response{
					Error:  "internal server error",
					Status: http.StatusInternalServerError,
				}, w, http.StatusInternalServerError)
				return
			}

			addresses = append(addresses, ad)
		}

		restutil.ResponseJSON(Response{
			Message: "address book",
			Status:  http.StatusOK,
			Data:    addresses,
		}, w, http.StatusOK)
		return
	}
	return fn
}

// RemoveAddress deletes an address from database
func (c *Component) RemoveAddress() httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// remove address from db
		addressID := ps.ByName("addressID")

		_, err := c.db.Exec("delete from addresses where id = ?", addressID)
		if err != nil {
			fmt.Println(err)
			restutil.ResponseJSON(Response{
				Error:  "internal server error",
				Status: http.StatusInternalServerError,
			}, w, http.StatusInternalServerError)
			return
		}

		restutil.ResponseJSON(Response{
			Message: "address has been removed!",
			Status:  http.StatusOK,
			Data:    addressID,
		}, w, http.StatusOK)
		return
	}
	return fn
}

func (c *Component) GetSearchedAddressBook() httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// fetch address book from database
		searchStr := r.URL.Query().Get("search")
		searchStr = "%" + searchStr + "%"
		var addresses []Address
		rows, err := c.db.
			Query(fmt.Sprintf("select * from addresses where fname like '%v' or lname like '%v'", searchStr, searchStr))
		if err != nil {
			fmt.Println(err)
			restutil.ResponseJSON(Response{
				Error:  "internal server error",
				Status: http.StatusInternalServerError,
			}, w, http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			ad := Address{}
			if err = rows.
				Scan(
					&ad.ID,
					&ad.FirstName,
					&ad.LastName,
					&ad.PhoneNumber,
				); err != nil {
				fmt.Println(err)
				restutil.ResponseJSON(Response{
					Error:  "internal server error",
					Status: http.StatusInternalServerError,
				}, w, http.StatusInternalServerError)
				return
			}

			addresses = append(addresses, ad)
		}

		restutil.ResponseJSON(Response{
			Message: "address book",
			Status:  http.StatusOK,
			Data:    addresses,
		}, w, http.StatusOK)
		return
	}
	return fn
}

package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itsfhz/flutter-golang-server/models"
	"github.com/itsfhz/flutter-golang-server/utils"
)

// func HomePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	log.Println("Endpoint Hit: homePage")
// }
func TestLink(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Connection to server successful..")
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpense function is invoked")

	vars := mux.Vars(r)
	id := vars["id"]

	db := utils.DbHandle()
	defer db.Close()

	var expense models.Expense

	err := db.QueryRow("SELECT id, type, title, date, rate FROM expenses WHERE id = $1", id).Scan(&expense.Id, &expense.Type, &expense.Title, &expense.Date, &expense.Rate)
	log.Println("Before switch")
	log.Println(expense)
	switch {
	case err == sql.ErrNoRows:
		//log.Fatalf("no user with id %d", id)
		json.NewEncoder(w).Encode(fmt.Sprint("Activity with id:", id, " does not exist"))

	case err != nil:
		log.Fatal(err)
		panic(err)
	default:
		//log.Printf("name is %s\n", partner.Name)
		json.NewEncoder(w).Encode(expense)
	}

}

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllExpenses function is invoked")

	db := utils.DbHandle()
	defer db.Close()

	log.Println("GetAllExpenses: db handle created")

	rows, err := db.Query("SELECT id, type, title, date, rate FROM expenses ")
	if err != nil {
		// handle this error better tharoduct
		panic(err)
	}
	defer rows.Close()

	log.Println("GetAllActivities: rows retrieved")

	retrievedExpense := []*models.Expense{}

	for rows.Next() {
		var expense models.Expense

		err1 := rows.Scan(&expense.Id, &expense.Type, &expense.Title, &expense.Date, &expense.Rate)

		if err1 != nil {
			panic(err1)
		}

		retrievedExpense = append(retrievedExpense, &expense)

	}

	json.NewEncoder(w).Encode(retrievedExpense)

}

func AddExpenses(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Add Product function is invoked")
	log.Println("inside AddExpenses function")
	var bodyBytes []byte
	var err1 error
	if r.Body != nil {
		bodyBytes, err1 = ioutil.ReadAll(r.Body)
		if err1 != nil {
			log.Printf("Body reading error: %v", err1)
			return
		}
		defer r.Body.Close()
	}

	// Below function is only for pretty printing of request contents
	utils.PrintRequest(r, bodyBytes)

	var expense models.Expense

	json.Unmarshal(bodyBytes, &expense)
	log.Println(expense.Type)
	log.Println(expense.Title)
	log.Println(expense.Date)
	log.Println(expense.Rate)
	db := utils.DbHandle()
	defer db.Close()

	sqlStatement := "INSERT INTO expenses (type, title, date, rate) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int32 = 0

	err := db.QueryRow(sqlStatement, expense.Type, expense.Title, expense.Date, expense.Rate).Scan(&id)
	if err != nil {
		panic(err)
	}
	log.Println("New Expenses ID is: ", id)

	//fmt.Println("Parnter received is :")
	//fmt.Println(partner)
	//json.NewEncoder(w).Encode(partner)
	message := fmt.Sprint("New Expenses is created with id: ", id)
	json.NewEncoder(w).Encode(message)

	//fmt.Fprintf(w, "%+v", string(reqBody))
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteExpense function is invoked")
	vars := mux.Vars(r)
	id := vars["id"]

	db := utils.DbHandle()
	defer db.Close()

	sqlStatement := "DELETE FROM expenses WHERE id = $1;"

	res, err1 := db.Exec(sqlStatement, id)
	if err1 != nil {
		panic(err1)
	}

	count, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	var message string
	if count > 0 {
		message = fmt.Sprint("Expenses with id:", id, " is successfully deleted")
	} else {
		message = fmt.Sprint("Expenses with id:", id, " does not exist")
	}

	json.NewEncoder(w).Encode(message)

}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateExpense function is invoked")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var expense models.Expense

	json.Unmarshal(reqBody, &expense)

	vars := mux.Vars(r)
	id := vars["id"]

	db := utils.DbHandle()
	defer db.Close()

	sqlStatement := "UPDATE expenses SET type = $1, title = $2, date = $3, rate = $4 WHERE id = $5;"

	res, err1 := db.Exec(sqlStatement, expense.Type, expense.Title, expense.Date, expense.Rate)
	if err1 != nil {
		panic(err1)
	}

	count, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	var message string
	if count > 0 {
		message = fmt.Sprint("Expense with id:", id, " is successfully updated")
	} else {
		message = fmt.Sprint("Expense with id:", id, " does not exist")
	}

	json.NewEncoder(w).Encode(message)

}

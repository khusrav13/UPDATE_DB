package main

import (
	"Example5/DBX"
	"Example5/moduls"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	data, err := sql.Open("sqlite3", "Bank")
	if err != nil {
		log.Fatal("Error", err)
	}
	err = DBX.DbInit(data)
	err = DBX.DbInitCurrency(data)
	err = DBX.DbAccounts(data)
	if err != nil {
		log.Fatal("error ", err)
	}
	fmt.Println("OKAY")


	///If you want you can add new user
//1)
//		NewUser := moduls.Users{
//					Name:    "Rai",
//					Surname: "Arbobova",
//					Age:     29,
//					Sex:     "Female",}
//
//				err = DBX.AddClient(data, NewUser)
//				if err != nil {
//					log.Fatal("Error")
//				}

	/////////////////Imagine you have ID 2 and you want to change information. Below We can all information about ID 2
	NewUser := moduls.Users{
		Name:    "Sino",
		Surname: "Ahamadzoda",
		Age:     29,
		Sex:     "Male",
	}
	err = DBX.UpdateClient(data, NewUser)
	if err != nil {
		log.Fatal("Error")
	}

	err = DBX.UpdateRemove(data)
	if err != nil{
		log.Fatal("Error. Please try again", err)
	}


}



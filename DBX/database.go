package DBX

import (
	"Example5/moduls"
	"database/sql"
	"fmt"
	"log"
)

func init()  {
	fmt.Println("hello, I am test")
}

func DbInit(db *sql.DB) (err error) {
	const usersDB = `CREATE TABLE IF NOT EXISTS users (
    ID integer PRIMARY KEY AUTOINCREMENT UNIQUE,
    Name TEXT NOT NULL,
    Surname TEXT NOT NULL,
    Sex TEXT,
    Age INTEGER NOT NULL,
    remove BOOLEAN NOT NULL DEFAULT FALSE
)`
	_, err = db.Exec(usersDB)

	if err != nil {
		return err
	}
	return nil
}

func DbInitCurrency(db *sql.DB) (err error) {
	const currency = `CREATE TABLE IF NOT EXISTS currency (
    ID integer PRIMARY KEY AUTOINCREMENT,
    Name text
)`
	_, err = db.Exec(currency)

	if err != nil {
		return err
	}
	return nil
}


func DbAccounts(db *sql.DB) (err error){
	const accounts  = `CREATE TABLE IF NOT EXISTS accounts (
    ID integer PRIMARY KEY AUTOINCREMENT UNIQUE,
    UserId integer references users(ID) not null ,
    Number integer,
    Amount integer,
    Currency integer references currency(ID),
    Remove BOOLEAN NOT NULL DEFAULT FALSE
)`
	_, err = db.Exec(accounts)

	if err != nil {
		return err
	}
	return nil
}


func AddClient(db *sql.DB, client moduls.Users) (err error){
	_, err = db.Exec(`INSERT INTO users(Name, Surname, Sex, Age) VALUES (($1), ($2), ($3), ($4))`, client.Name, client.Surname, client.Sex, client.Age)
	if err != nil {
		log.Fatal("Error try again")
		return err
	}
	return err
}

/////////////////Imagine you have ID 2 and you want to change information. Below we can change information about ID 2
func UpdateClient(db *sql.DB, client moduls.Users) (err error) {
	_, err = db.Exec(`UPDATE users SET (Name, Surname,Age, Sex) = (($1), ($2),($3), ($4)) WHERE ID = 2`, client.Name, client.Surname, client.Age, client.Sex)
	if err != nil {
		log.Fatal("Cannot Update. Please try again!", err)
		return err
	}
	return err
}


///////////Second Implementation by Remove /////////////////
func UpdateRemove(db *sql.DB) (err error) {
	_, err = db.Exec(`UPDATE users SET remove = TRUE WHERE Age <= 20 `)
	if err != nil {
		fmt.Println("ERROR.CANNOT UPDATE REMOVE", err)
		return err
	}
	return err
}




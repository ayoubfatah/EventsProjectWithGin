package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB  *sql.DB

func InitDB(){
var err error
 DB, err = sql.Open("sqlite3","api.db")
 if(err !=nil ){
	panic("Couldn't connect to database")
 }
 DB.SetMaxOpenConns(10)
 DB.SetMaxIdleConns(5)
 createTables()
}

func createTables(){
 createUsersTable := `
 CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  firstName TEXT NOT NULL  ,
  secondName TEXT NOT NULL  ,
  userName TEXT NOT NULL UNIQUE ,
  email TEXT NOT NULL  UNIQUE,
  password TEXT NOT NULL
 )
 `
 _ , err := DB.Exec(createUsersTable)
if err != nil {
    panic(err)
}

 createEventsTable := `
 CREATE TABLE IF NOT EXISTS events (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  location TEXT NOT NULL, 
  dateTime DATETIME NOT NULL,
  userId  INTEGER  REFERENCES users(id)
 )
 `	

_ , err = DB.Exec(createEventsTable)
if err != nil {
    panic(err)
}
 createRegisterTable := `
 CREATE TABLE IF NOT EXISTS registrations (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  eventId INTEGER,
  userId INTEGER
  FOREIGNER KEY(eventId) REFERENCES events(id)
  FOREIGNER KEY(userId) REFERENCES users(id)
 )
 `	

_ , err = DB.Exec(createRegisterTable)
if err != nil {
    panic(err)
}
}
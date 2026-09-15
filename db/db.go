package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB  *sql.DB

func InitDB(){
var err error
DB, err = sql.Open("sqlite", "file:events.db?cache=shared&mode=memory")
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
    slug TEXT NOT NULL UNIQUE,
    city TEXT NOT NULL,
    location TEXT NOT NULL,
    date DATETIME NOT NULL,
    organizerName TEXT NOT NULL,
    imageUrl TEXT NOT NULL,
    description TEXT NOT NULL,
    userId INTEGER REFERENCES users(id)
    )
 `	

_ , err = DB.Exec(createEventsTable)
if err != nil {
    panic(err)
}
 createRegisterTable := `
    CREATE TABLE IF NOT EXISTS registrations (
    eventId INTEGER NOT NULL,
    userId INTEGER NOT NULL,
    PRIMARY KEY (eventId, userId),
    FOREIGN KEY (eventId) REFERENCES events(id),
    FOREIGN KEY (userId) REFERENCES users(id)
    )
 `	

_ , err = DB.Exec(createRegisterTable)
if err != nil {
    panic(err)
}
}
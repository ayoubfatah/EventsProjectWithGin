package models

import (
	"gin-quickstart/db"
	"time"
)

type Event struct {
	Id int64
	Name string `binding:"required"` 
	Description string `binding:"required"` 
	Location string `binding:"required"` 
	DateTime time.Time `binding:"required"` 
	UserId int64 
}


var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, userId) 
			  VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.DateTime,
		e.UserId,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.Id = id
	return nil
}

func GetAllEvents() ([]Event, error){
	query := `SELECT * FROM events`
	rows , err := db.DB.Query(query)
	 if(err != nil){
		return  nil,err 
	 }
	defer rows.Close()


	var events []Event
	for rows.Next(){
		var event Event 
		err := rows.Scan(&event.Id, &event.Name , &event.Description , &event.Location , &event.DateTime , &event.UserId )
		 if(err != nil){
		return  nil,err 
		 }
		events = append(events, event)
	}
	return events , err 
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE Id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return  nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var event Event

	err = row.Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId,
	)
	if err != nil {
		return  nil, err
	}

	return  &event, nil
}

// update 


func (e *Event) Update() error {
	query := `
	UPDATE events 
	SET name = ? , description = ? , location = ? , dateTime = ?
	WHERE id = ? 
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		&e.Name,
		&e.Description,
		&e.Location,
		&e.DateTime,
		&e.Id,
	)
	if err != nil {
		return err
	}

	return nil
}



func (e *Event) Delete() error {
	query := `
	Delete FROM events 
	WHERE id = ? 
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		&e.Id,
	)
	if err != nil {
		return err
	}

	return nil
}


func (e *Event) Register(userId int64 ) error{
	query := `INSERT INTO registrations(eventId userId)
			  VALUES`
	stmt , err := db.DB.Prepare(query)

	if err != nil {
		return err 
	}
	defer stmt.Close()

	_ , err = stmt.Exec(e.Id , userId)
		if err != nil {
		return err 
	}

	return  nil
}	


func (e *Event) CancelRegistration(userId int64 ) error{
	query := `DELETE  FROM  registrations WHERE eventId = ? AND userId = ?  `
	stmt , err := db.DB.Prepare(query)

	if err != nil {
		return err 
	}
	defer stmt.Close()

	_ , err = stmt.Exec(e.Id , userId)
		if err != nil {
		return err 
	}

	return  nil
}	
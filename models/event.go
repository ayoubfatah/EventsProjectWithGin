package models

import (
	"fmt"
	"gin-quickstart/db"
	"time"
)

type Event struct {
    Id            int64     `json:"id"`
    Name          string    `json:"name" binding:"required"`
    Slug          string    `json:"slug" binding:"required"`
    City          string    `json:"city" binding:"required"`
    Location      string    `json:"location" binding:"required"`
    Date          time.Time `json:"date" binding:"required"`
    OrganizerName string    `json:"organizerName" binding:"required"`
    ImageUrl      string    `json:"imageUrl" binding:"required"`
    Description   string    `json:"description" binding:"required"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
    UserId        int64     `json:"userId"`
}

type PaginatedEvents struct {
	Events  []Event `json:"events"`
	Page    int     `json:"page"`
	Limit   int     `json:"limit"`
	HasMore bool    `json:"hasMore"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(
			name,
			slug,
			city,
			location,
			date,
			organizerName,
			imageUrl,
			description,
			userId
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		e.Name,
		e.Slug,
		e.City,
		e.Location,
		e.Date,
		e.OrganizerName,
		e.ImageUrl,
		e.Description,
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

func GetAllEvents(page int, limit int) (PaginatedEvents, error) {
	offset := (page - 1) * limit

	query := `
		SELECT
			id,
			name,
			slug,
			city,
			location,
			date,
			organizerName,
			imageUrl,
			description,
			userId
		FROM events
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := db.DB.Query(query, limit+1, offset)
	if err != nil {
		return PaginatedEvents{}, err
	}
	defer rows.Close()

	events := make([]Event, 0, limit+1)

	for rows.Next() {

		var event Event

		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Slug,
			&event.City,
			&event.Location,
			&event.Date,
			&event.OrganizerName,
			&event.ImageUrl,
			&event.Description,
			&event.UserId,
		)

		if err != nil {
			return PaginatedEvents{}, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return PaginatedEvents{}, err
	}

	hasMore := len(events) > limit

	// We don't send the extra event to the frontend.
	if hasMore {
		events = events[:limit]
	}

	return PaginatedEvents{
		Events:  events,
		Page:    page,
		Limit:   limit,
		HasMore: hasMore,
	}, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var event Event

	err = row.Scan(
		&event.Id,
		&event.Name,
		&event.Slug,
		&event.City,
		&event.Location,
		&event.Date,
		&event.OrganizerName,
		&event.ImageUrl,
		&event.Description,
		&event.UserId,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
func GetEventsByCity(city string) ([]Event, error) {
	query := `SELECT * FROM events WHERE city = ? 
			  ORDER BY id DESC
	`

	rows, err := db.DB.Query(query, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Slug,
			&event.City,
			&event.Location,
			&event.Date,
			&event.OrganizerName,
			&event.ImageUrl,
			&event.Description,
			&event.UserId,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
















func GetEventBySlug(slug string) (*Event, error) {
	query := `SELECT * FROM events WHERE slug = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(slug)

	var event Event

	err = row.Scan(
		&event.Id,
		&event.Name,
		&event.Slug,
		&event.City,
		&event.Location,
		&event.Date,
		&event.OrganizerName,
		&event.ImageUrl,
		&event.Description,
		&event.UserId,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
// Update

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET
			name = ?,
			slug = ?,
			city = ?,
			location = ?,
			date = ?,
			organizerName = ?,
			imageUrl = ?,
			description = ?
		WHERE id = ?
	`





	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		e.Name,
		e.Slug,
		e.City,
		e.Location,
		e.Date,
		e.OrganizerName,
		e.ImageUrl,
		e.Description,
		e.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete() error {
	query := `
		DELETE FROM events
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id)

	if err != nil {
		return err
	}

	return nil
}

// Register

func (e *Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations(eventId, userId)
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	if err != nil {
		return err
	}

	return nil
}

// Cancel Registration

func (e *Event) CancelRegistration(userId int64) error {
	query := `
		DELETE FROM registrations
		WHERE eventId = ? AND userId = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	if err != nil {
		return err
	}

	return nil
}



func (e *Event) IsRegistered(userId int64) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM registrations
			WHERE userId = ? AND eventId = ?
		)
	`

	var registered bool

	err := db.DB.QueryRow(query, userId, e.Id).Scan(&registered)
	if err != nil {
		return false, err
	}
	fmt.Println("waaaa")
	fmt.Println(registered)
	return registered, nil
}



func GetEventsByUserID(userID int64) ([]Event , error){


query := `
    SELECT
        id,
        name,
        slug,
        city,
        location,
        date,
        organizerName,
        imageUrl,
        description,
        userId
    FROM events
    WHERE userId = ?
    ORDER BY id DESC

`

	rows , err := db.DB.Query(query, userID)
	if err != nil {
		return nil , err
	}
	defer rows.Close()

	events := make([]Event, 0)

	for rows.Next(){
		var event Event
		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Slug,
			&event.City,
			&event.Location,
			&event.Date,
			&event.OrganizerName,
			&event.ImageUrl,
			&event.Description,
			&event.UserId,
		)
		
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
    return nil, err
}
	return events , nil
}
 






func GetRegisteredEventsByUser(userId int64) ([]Event, error) {
query := `
    SELECT 
        e.id,
        e.name,
        e.slug,
        e.city,
        e.location,
        e.date,
        e.organizerName,
        e.imageUrl,
        e.description,
        e.userId
    FROM events e
    INNER JOIN registrations r ON e.id = r.eventId
    WHERE r.userId = ?
    ORDER BY e.id DESC
`

	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Slug,
			&event.City,
			&event.Location,
			&event.Date,
			&event.OrganizerName,
			&event.ImageUrl,
			&event.Description,
			&event.UserId,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
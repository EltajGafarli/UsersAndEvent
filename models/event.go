package models

import (
	"BookingApp/db"
	"database/sql"
	"errors"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	Date        time.Time `binding:"required" json:"date"`
	UserID      int       `json:"user_id"`
}

func (e Event) Save() error {

	insertQuery := `
	INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)
`

	stmt, err := db.DB.Prepare(insertQuery)

	if err != nil {
		return errors.New("Error preparing statement: " + err.Error())
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserID)

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	id, err := result.LastInsertId()

	e.ID = int(id)

	return err
}

func GetAllEvents() ([]Event, error) {
	selectQuery := "select * from events"

	rows, err := db.DB.Query(selectQuery)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	if err != nil {
		return nil, errors.New("Error preparing statement: " + err.Error())
	}

	events := make([]Event, 0)

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
		if err != nil {
			return nil, errors.New("Error scanning row: " + err.Error())
		}
		events = append(events, event)
	}

	return events, nil

}

func GetEventByID(id int) (*Event, error) {
	selectQuery := "select * from events where id=?"
	row := db.DB.QueryRow(selectQuery, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
	if err != nil {
		return nil, errors.New("event Not Found")
	}

	return &event, nil
}

func (event *Event) UpdateEvent() error {
	query := `
	update events set name=?, description=?, location=?, datetime=?, user_id=?
	where id=?
`
	stmt, err := db.DB.Prepare(query)

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Date, event.UserID, event.ID)

	return err

}

func (event *Event) DeleteEvent() error {
	deleteQuery := "delete from events where id=?"
	stmt, err := db.DB.Prepare(deleteQuery)

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(event.ID)
	return err
}

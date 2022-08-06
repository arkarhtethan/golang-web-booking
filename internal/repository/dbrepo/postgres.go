package dbrepo

import (
	"context"
	"time"

	"github.com/arkarhtethan/golang-web-booking/internal/models"
)

func (m *postgresDbRepo) AllUsers() bool {
	return true
}

func (m *postgresDbRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var newId int
	stmt := `insert into reservations (first_name,last_name,email,phone,start_date,end_date,room_id,created_at,updated_at) values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`
	err := m.DB.QueryRowContext(ctx, stmt, res.FirstName, res.LastName, res.Email, res.Phone, res.StartDate, res.EndDate, res.RoomID, time.Now(), time.Now()).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (m *postgresDbRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `insert into room_restrictions(
			start_date,
			end_date,
			room_id,
			reservation_id,
			created_at,
			updated_at,
			restriction_id
		 )
		  values($1,$2,$3,$4,$5,$6,$7)`
	now := time.Now()
	_, err := m.DB.ExecContext(
		ctx,
		stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		now,
		now,
		r.RestrictionID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDbRepo) SearchAvailabilityByDatesRoomID(start, end time.Time, roomId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select
			count(id)
		from
			 room_restrictions
		where
			room_id = $1 and
			$2 < end_date and $3 > start_date;
	`
	var numRow int
	row := m.DB.QueryRowContext(ctx, query, roomId, start, end)
	err := row.Scan(&numRow)
	if err != nil {
		return false, err
	}
	if numRow == 0 {
		return true, nil
	}
	return false, nil
}
func (m *postgresDbRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select
			r.id, r.room_name
		from
			rooms r
		where r.id not in (
			select
				room_id
			from
				 room_restrictions rr
			where
				$1 < rr.end_date and $2 > rr.start_date
		)
	`
	var rooms []models.Room
	rows, err := m.DB.QueryContext(ctx, query, start, end)

	if err != nil {
		return rooms, err
	}
	for rows.Next() {
		var room models.Room
		err = rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)
	}
	if err != nil {
		return rooms, err
	}
	return rooms, nil
}

func (m *postgresDbRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var room models.Room

	query := `
		select id, room_name, created_at, updated_at from rooms where id = $1
	`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}
	return room, nil
}

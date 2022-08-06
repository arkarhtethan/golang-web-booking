package dbrepo

import (
	"errors"
	"time"

	"github.com/arkarhtethan/golang-web-booking/internal/models"
)

func (m *testDbRepo) AllUsers() bool {
	return true
}

func (m *testDbRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

func (m *testDbRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

func (m *testDbRepo) SearchAvailabilityByDatesRoomID(start, end time.Time, roomId int) (bool, error) {
	return false, nil
}
func (m *testDbRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (m *testDbRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 3 {
		return room, errors.New("some error")
	}
	return room, nil
}

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

func (m *testDbRepo) GetUserByID(id int) (models.User, error) {
	var user models.User
	return user, nil
}

func (m *testDbRepo) Authenticate(email, password string) (int, string, error) {
	return 0, "", nil
}

func (m *testDbRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDbRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

func (m *testDbRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

func (m *testDbRepo) GetReservationByID(id int) (models.Reservation, error) {
	var reservations models.Reservation
	return reservations, nil
}

func (m *testDbRepo) UpdateReservation(res models.Reservation) error {
	return nil
}

func (m *testDbRepo) DeleteReservation(id int) error {
	return nil
}

func (m *testDbRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (m *testDbRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (m *testDbRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	var rooms []models.RoomRestriction
	return rooms, nil
}

func (m *testDbRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	return nil
}

func (m *testDbRepo) DeleteBlockByID(id int) error {
	return nil
}

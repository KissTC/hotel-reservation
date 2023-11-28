package db

// esta constante antes estaba en user_store
// la movimos porque no tiene nada que ver con user
const (
	DBNAME     = "hotel_reservation"
	TestDBNAME = "hotel-reservation-test"
	DBURI      = "mongodb://localhost:27017"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}

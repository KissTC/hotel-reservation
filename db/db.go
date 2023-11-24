package db

import "go.mongodb.org/mongo-driver/bson/primitive"

// esta constante antes estaba en user_store
// la movimos porque no tiene nada que ver con user
const DBNAME = "hotel_reservation"

// la helper function
func ToObjectId(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id) //oid object id
	if err != nil {
		panic(err)
	}

	return oid
}

package db

import (
	"context"

	"github.com/kisstc/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

// tipo user store interface para implementar lo que queremos
// puede ser mongodb, postgress etc
type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
}

// this will be the implementation of the interface,
// we can have PostgressUserStore for other implementation eg
type MongoUserStroe struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(c *mongo.Client) *MongoUserStroe {
	coll := c.Database(DBNAME).Collection(userColl)
	return &MongoUserStroe{
		client: c,
		coll:   coll,
	}
}

func (ms *MongoUserStroe) GetUserById(ctx context.Context, id string) (*types.User, error) {
	// validate the correctnes of the id
	// converting this
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// para que funcione necesitamos collection y todo eso
	// así que a MongoUserStore le agregamos más campos, hasta el momento solo tenía el atributo Client
	var user types.User
	if err := ms.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// bson.M no podemos pasar el id tal cual porque no es de tipo string lo que acepta es un object id
// asi que creamos una helper function en db.go
// la helper function se cambio a una validacion en el metodo get user by id
// debido a la prueba de obtener un usuario erroneo

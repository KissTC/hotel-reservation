package db

import (
	"context"
	"fmt"

	"github.com/kisstc/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type Dropper interface {
	Drop(context.Context) error
}

// tipo user store interface para implementar lo que queremos
// puede ser mongodb, postgress etc
type UserStore interface {
	Dropper
	GetUserById(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, bson.M, types.UpdateUserParams) error
}

// this will be the implementation of the interface,
// we can have PostgressUserStore for other implementation eg
type MongoUserStroe struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(c *mongo.Client, dbname string) *MongoUserStroe {
	coll := c.Database(dbname).Collection(userColl)
	return &MongoUserStroe{
		client: c,
		coll:   coll,
	}
}

func (ms *MongoUserStroe) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := ms.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
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

func (ms *MongoUserStroe) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := ms.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (ms *MongoUserStroe) DeleteUser(ctx context.Context, id string) error {
	// esto se ocupa porque _id de mongo es de tipo object id, no string
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	//TODO: check if we need the res, res only have the deletecount for how many documents deleted
	_, err = ms.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	return nil
}

func (ms *MongoUserStroe) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {

	update := bson.D{
		{
			Key: "$set", Value: params.ToBSON(),
		},
	}

	_, err := ms.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoUserStroe) Drop(ctx context.Context) error {
	fmt.Println("---- dropping user collection")
	return s.coll.Drop(ctx)
}

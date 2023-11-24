package types

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// const (
// 	minFirstNameLen = 2
// 	minLastNameLen  = 2
// 	minPasswordLen  = 7
// )

type CreateUserParams struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=100"`
	LastName  string `json:"lastName"  validate:"required,min=2,max=100"`
	Email     string `json:"email"     validate:"required,email"`
	Password  string `json:"password"  validate:"required,min=7"`
}

func (params CreateUserParams) Validate(ctx context.Context) error {
	validate := validator.New()
	if err := validate.StructCtx(ctx, params); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "FirstName":
					return fmt.Errorf("error en el campo 'FirstName': %s", fieldError.Tag())
				case "LastName":
					return fmt.Errorf("error en el campo 'LastName': %s", fieldError.Tag())
				case "Email":
					return fmt.Errorf("error en el campo 'Email': %s", fieldError.Tag())
				case "Password":
					return fmt.Errorf("error en el campo 'Password': %s", fieldError.Tag())
				}
			}
		}
		return fmt.Errorf("error en la validación: %v", err)
	}
	return nil
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lasttName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}

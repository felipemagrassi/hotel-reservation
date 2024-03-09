package types

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string `bson:"firstName" json:"firstName"`
	LastName          string `bson:"lastName" json:"lastName"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"-"`
}

type UserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewUserCreateDTO(dto UserDTO) (*User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         dto.FirstName,
		LastName:          dto.LastName,
		Email:             dto.Email,
		EncryptedPassword: string(encryptedPassword),
	}, nil
}

func NewUserUpdateDTO(dto UserDTO) (*User, error) {
	return &User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}, nil
}

func (dto UserDTO) Validate() (map[string]string, bool) {
	errors := make(map[string]string)
	if len(dto.FirstName) == 0 {
		errors["firstName"] = "missing first name"
	}
	if len(dto.LastName) == 0 {
		errors["lastName"] = "missing last name"
	}
	if err := isEmailValid(dto.Email); err != nil {
		errors["email"] = err.Error()
	}
	if len(dto.Password) == 0 {
		errors["password"] = "missing password"
	}
	return errors, len(errors) == 0
}

func isEmailValid(email string) error {
	if len(email) == 0 {
		return fmt.Errorf("missing email")
	}

	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

	if !regexp.MustCompile(regex).MatchString(email) {
		return fmt.Errorf("invalid email")
	}

	return nil
}

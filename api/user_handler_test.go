package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/felipemagrassi/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func TestCanCreateAndGetUser(t *testing.T) {
	handler := NewUserHandler(testDatabase.Store)

	app := fiber.New()
	app.Get("/user/:id", handler.HandleGetUser)
	app.Post("/user", handler.HandleCreateUser)

	createDto := types.UserDTO{
		Email:     "felipe.magrassi@gmail.com",
		FirstName: "Felipe",
		LastName:  "Magrassi",
		Password:  "123456",
	}

	b, err := json.Marshal(createDto)
	if err != nil {
		t.Fatalf("Could not marshal user dto: %s", err)
	}

	req := httptest.NewRequest("POST", "/user", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		log.Fatalf("Could not make request: %s", err)
	}

	var createdUser types.User
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	if err != nil {
		log.Fatalf("Could not decode response: %s", err)
	}

	if createdUser.ID == "" {
		t.Errorf("Expected user id to be set, got empty string")
	}

	if len(createdUser.EncryptedPassword) > 0 {
		t.Errorf("Expected user password to be empty, got %s", createdUser.EncryptedPassword)
	}

	if createdUser.Email != createDto.Email {
		t.Errorf("Expected user email to be %s, got %s", createDto.Email, createdUser.Email)
	}

	if createdUser.FirstName != createDto.FirstName {
		t.Errorf("Expected user first name to be %s, got %s", createDto.FirstName, createdUser.FirstName)
	}

	if createdUser.LastName != createDto.LastName {
		t.Errorf("Expected user last name to be %s, got %s", createDto.LastName, createdUser.LastName)
	}

	req = httptest.NewRequest("GET", "/user/"+createdUser.ID, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req)
	if err != nil {
		log.Fatalf("Could not make request: %s", err)
	}
	defer resp.Body.Close()
	var foundUser types.User

	err = json.NewDecoder(resp.Body).Decode(&foundUser)
	if err != nil {
		log.Fatalf("Could not decode response: %s", err)
	}

	if foundUser.ID != createdUser.ID {
		t.Errorf("Expected user id to be %s, got %s", createdUser.ID, foundUser.ID)
	}

	if len(foundUser.EncryptedPassword) > 0 {
		t.Errorf("Expected user password to be empty, got %s", foundUser.EncryptedPassword)
	}

	if foundUser.Email != createdUser.Email {
		t.Errorf("Expected user email to be %s, got %s", createdUser.Email, foundUser.Email)
	}

	if foundUser.FirstName != createdUser.FirstName {
		t.Errorf("Expected user first name to be %s, got %s", createdUser.FirstName, foundUser.FirstName)
	}

	if foundUser.LastName != createdUser.LastName {
		t.Errorf("Expected user last name to be %s, got %s", createdUser.LastName, foundUser.LastName)
	}
}

func TestCanCreateAndDeleteUser(t *testing.T) {
	handler := NewUserHandler(testDatabase.Store)

	createDto := types.UserDTO{
		Email:     "felipe.magrassi@gmail.com",
		FirstName: "Felipe",
		LastName:  "Magrassi",
		Password:  "123456",
	}

	app := fiber.New()
	app.Get("/user/:id", handler.HandleGetUser)
	app.Post("/user", handler.HandleCreateUser)
	app.Delete("/user/:id", handler.HandleDeleteUser)
	b, _ := json.Marshal(createDto)

	req := httptest.NewRequest("POST", "/user", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		log.Fatalf("Could not make request: %s", err)
	}

	var createdUser types.User
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	if err != nil {
		log.Fatalf("Could not decode response: %s", err)
	}

	if createdUser.ID == "" {
		t.Errorf("Expected user id to be set, got empty string")
	}

	if len(createdUser.EncryptedPassword) > 0 {
		t.Errorf("Expected user password to be empty, got %s", createdUser.EncryptedPassword)
	}

	if createdUser.Email != createDto.Email {
		t.Errorf("Expected user email to be %s, got %s", createDto.Email, createdUser.Email)
	}

	if createdUser.FirstName != createDto.FirstName {
		t.Errorf("Expected user first name to be %s, got %s", createDto.FirstName, createdUser.FirstName)
	}

	if createdUser.LastName != createDto.LastName {
		t.Errorf("Expected user last name to be %s, got %s", createDto.LastName, createdUser.LastName)
	}

	req = httptest.NewRequest("DELETE", "/user/"+createdUser.ID, nil)
	deleteResp, err := app.Test(req)
	if err != nil {
		log.Fatalf("Could not make request: %s", err)
	}

	if deleteResp.StatusCode != 200 {
		log.Fatalf("Expected status code to be 200, got %d", deleteResp.StatusCode)
	}

	req = httptest.NewRequest("GET", "/user/"+createdUser.ID, nil)
	resp, err = app.Test(req)
	if err != nil {
		log.Fatalf("Could not make request: %s", err)
	}

	defer resp.Body.Close()
	var foundUser types.User

	err = json.NewDecoder(resp.Body).Decode(&foundUser)
	if err != nil {
		log.Fatalf("Could not decode response: %s", err)
	}

	if foundUser.ID != "" {
		t.Errorf("Expected user id to be empty, got %s", foundUser.ID)
	}
}

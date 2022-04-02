package Tests

import (
	"testing"
	"CAW/Backend/signupauth/controllers"
	"CAW/Backend/signupauth/database"
	"golang.org/x/crypto/bcrypt"
)

var testusername = "sivapraneeth"
var testpassword = "1234"
var testaddress = "root:Praneeth11@/users"
var testdbName = "users"
var data map[string]string

func TestHashPassword(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data["abcdefg"]),14)
	if err != nil {
		t.Error("Error when hashing password")
	}
	if len(hash) != 60 {
		t.Error("The length of hash is wrong.")
	}
}

func TestCheckPasswordHash(t *testing.T) {
  verbose = false
  err = database.DB.connectDB(testusername, testpassword, testaddress, testdbName)
  
	if !bcrypt.CompareHashAndPassword("sivapraneeth", "1234") {
		t.Error("Correct password did not pass the test")
	}
  if bcrypt.CompareHashAndPassword("sivapraneeth", "123") {
		t.Error(" wrong password passed the test.")
	}
}

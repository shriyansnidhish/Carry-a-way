package tests

import (
	"CAW/Backend/signupauth/database"
	"CAW/Backend/signupauth/password_recovery"
	"bytes"
	"net/http"
	"net/http/httptest"
	"server/utils"
	"strings"
	"testing"
)

func TestForgotPasswordHandler(t *testing.T) {
	database.Connect()

	var testdata = []byte(`{
		                     "token":"5XcwTHjK89Lo4"
		                     "email":"dev@forgot.com",
	                         "password":"newsecret"
						    }`)

	req, err := http.NewRequest("POST", "/Login", bytes.NewBuffer(testdata))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	nr := httptest.NewRecorder()
	handler := http.HandlerFunc(password_recovery.forgotPWverHandler())
	handler.ServeHTTP(nr, req)
	if status := nr.Code; status != http.StatusOK {
		t.Errorf("handler gave incorrect status code %v while we want %v",	status, http.StatusOK)
	}

	expected := `{"message":"Password reset successfully"}`
	if strings.Contains(nr.Body.String(), expected) {
	} else {
		t.Errorf("handler gave unexpected message %v while we expect %v", nr.Body.String(), expected)
	}
}

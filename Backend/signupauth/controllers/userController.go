package controllers

import (
	"CAW/Backend/signupauth/database"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetUsers() ([]User, error) {

	rows, err := database.DB.Query("SELECT Id, FirstName, LastName, Email from users")

	if err != nil {
		return nil, err
	}
	//closes rows after action is completed
	defer rows.Close()

	listofcustomers := make([]User, 0, 20)

	for rows.Next() {
		eachuser := User{}
		err = rows.Scan(&eachuser.Id, &eachuser.FirstName, &eachuser.LastName, &eachuser.Email)

		if err != nil {
			return nil, err
		}
		listofcustomers = append(listofcustomers, eachuser)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return listofcustomers, err
}
func GetUserswithId() ([]User, error) {

	rows, err := database.DB.Query("SELECT Id, FirstName, LastName, Email from users WHERE Id=? ")
	if err != nil {
		return nil, err
	}
	//closes rows after action is completed
	defer rows.Close()

	listofcustomers := make([]User, 0, 20)

	for rows.Next() {
		eachcustomerbyid := User{}
		err = rows.Scan(&eachcustomerbyid.StudentId, &eachcustomerbyid.FirstName, &eachcustomerbyid.LastName, &eachcustomerbyid.Email)

		if err != nil {
			return nil, err
		}
		listofcustomers = append(listofcustomers, eachcustomerbyid)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return listofcustomers, err
}

func isemailExists(email string) bool {
	row := database.DB.QueryRow("select email from users where email= ?", email)
	retrievedemail := ""
	row.Scan(&retrievedemail)
	return retrievedemail != ""
}
func AddUsers(userInstance User) (bool, error) {

	ctx, err := database.DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := ctx.Prepare("INSERT INTO users (firstname, lastname,email) VALUES (?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	newmail := isemailExists(userInstance.Email)
	if newmail {
		return false, fmt.Errorf("Customer with this Email already exists. Please login to continue")
	} else {

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userInstance.Password), bcrypt.DefaultCost)

		if err != nil {
			return false,

				fmt.Errorf("Unexpected error in encryption of user password")
		}

		newmail.Password = string(encryptedPassword)
		_, err = stmt.Exec(userInstance.FirstName, userInstance.LastName, userInstance.Email, userInstance.Password)

		if err != nil {
			return false, err
		}

		ctx.Commit()

		return true, nil
	}
}

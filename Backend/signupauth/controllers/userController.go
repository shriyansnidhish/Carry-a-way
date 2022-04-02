package controllers

import (
	"CAW/Backend/signupauth/database"
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

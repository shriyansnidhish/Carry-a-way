package models
type User struct{
	Id uint `json:"id"`
	FirstName string `json:"fname"`
	LastName string `json:"lname"`
	Email string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
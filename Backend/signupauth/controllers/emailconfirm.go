package controllers

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"
)

func EmailConfirmation(mailid string, orderid uint) {
	// sender data
	from := os.Getenv("FromEmailAddr")
	password := os.Getenv("SMTPpwd")
	// receiver address
	//toEmail := os.Getenv("ToEmailAddr") //
	toEmail := mailid
	fmt.Println(toEmail)
	var s string = strconv.FormatUint(uint64(orderid), 10)

	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message

	subject := "Carry-a-way: Order Confirmation!!!"
	body := "Dear Customer\n,Thank you for choosing carry-a-way to deliver your luggage\nYour order id is:\n"
	body = body + s
	body = body + "\nRegards,\nLugless team\n"

	message := []byte(subject + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	error := smtp.SendMail(address, auth, from, to, message)
	if error != nil {
		fmt.Println("error:", error)
	}
}

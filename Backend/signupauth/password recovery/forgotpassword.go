package main
import(
	"database/sql"
	"html/template"
	"errors"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
	"unicode"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
// 	"CAW/Backend/signupauth/models"
// 	"CAW/Backend/signupauth/database"
// 	"CAW/Backend/signupauth/controllers"
// 	"CAW/Backend/signupauth/usersessions"
 )
 //template to access FE html pages
var tmp*template.Template 
var db *sql.DB
type UserData struct { //UserData to store carry a way users details...works in conjuction with models.User
	Username   string
	Email      string
	AuthInfo   string
	ErrMessage string
	Message    string
}
//variable to store cookies 
var cookiedb = sessions.NewCookieStore([]byte(os.Getenv("Cookiedbstore")))
//Handler to take username and email verification password for account recovery
func forgotpwChangeHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("u") //takes the username for password recovery
	emailVerPassword := r.FormValue("evpw")// evpw provides the account recovery password
	fmt.Println("username:", username)
	fmt.Println("emailVerPassword:", emailVerPassword)
	var ud UserData
	ud.AuthInfo = "?u=" + username + "&evpw=" + emailVerPassword
	tmp.ExecuteTemplate(w, "login.html", ud)//access login page for password reset
}

//function to capture the password value
func forgotPasswordValue(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	fmt.Println("email:", email)
	var ud UserData
	//if entered incorrect email value
	ud.ErrMessage = "Sorry, your account seems to be absent from the database, please try with a correct email"
	//Begin the transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error occured before beginning the transaction:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("Error occured during the rollback:", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	//rollback changes if no action is taken
	defer tx.Rollback()
	var username string
	//Query to select email for sending password recovery info
	row := db.QueryRow("SELECT email, username FROM users WHERE email = ?", email)
	//read the emailid
	err = row.Scan(&email, &username)
	//check if the email id valid or not
	if err != nil {
		fmt.Println("email not found in db")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error in rolling back changes", rollbackErr)
		}
	
		tmp.ExecuteTemplate(w, "login.html", nil)
		return
	}
	//clock starts ticking ,the user has to reset password before it becomes invalid
	now := time.Now()
	timeout := now.Add(time.Minute * 45)
	rand.Seed(time.Now().UnixNano())
	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")//string of characters allowed
	emailVerRandRune := make([]rune, 64)
	for i := 0; i < 64; i++ {
		emailVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}
	fmt.Println("change password emailVerRandRune:", emailVerRandRune)
	emailVerPassword := string(emailVerRandRune)
	// fmt.Println("emailVerPassword:", emailVerPassword)
	// fmt.Println("emailVerPassword len:", len(emailVerPassword))
	var emailVerPWhash []byte
	emailVerPWhash, err = bcrypt.GenerateFromPassword([]byte(emailVerPassword), bcrypt.DefaultCost) //hasing the password
	//if error is present...rollback the changes
	if err != nil {
		fmt.Println("bcrypt err:", err)
		//if error in rollback...print the appropriate message on console
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error in rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	var updateEmailVerStmt *sql.Stmt
	//updating the new data from user to database
	updateEmailVerStmt, err = tx.Prepare("UPDATE email_ver_hash SET ver_hash = ?, timeout = ? WHERE email = ?;")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			//error in rollback
			fmt.Println("there was an error in rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	defer updateEmailVerStmt.Close()
	emailVerPWhashStr := string(emailVerPWhash)
	var result sql.Result
	result, err = updateEmailVerStmt.Exec(emailVerPWhashStr, timeout, email)
	fmt.Println("err:", err)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("err:", err)
	if err != nil || rowsAff != 1 {
		fmt.Println("error has occured while inserting new user", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error in rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	from := os.Getenv("FromEmailAddr") 
	password := os.Getenv("SMTPpwd")   
	to := []string{email}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	subject := "Subject: Carry-A-Way account recovery\n"
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://localhost:4200/login?u=" + username + "&evpw=" + emailVerPassword + "\">Change Password</a></body>"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", from, password, host)
	fmt.Println("message:", string(message))
	err = smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("error sending reset password email, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error in rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error in commiting the changes", commitErr)
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
	}
	tmp.ExecuteTemplate(w, "login.html", nil)
}

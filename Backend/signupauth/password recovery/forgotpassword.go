package main
import(
	"database/sql"
	"html/template"
	"errors"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
//	"strings"
	"time"
	"unicode"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
//	"github.com/gorilla/context"
//	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
 	"CAW/Backend/signupauth/models"
 	"CAW/Backend/signupauth/database"
 	"CAW/Backend/signupauth/controllers"
	"CAW/Backend/signupauth/usersessions"
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
	//updating the new user data will be deferred if there is a delay in transaction commit
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
	//block of code to send email for account recovery info
	from := os.Getenv("FromEmailAddr") 
	password := os.Getenv("SMTPpwd")   
	to := []string{email}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	//Header of the email
	subject := "Subject: Carry-A-Way account recovery\n"
	//body containing link to frontend password reset page
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://localhost:4200/login?u=" + username + "&evpw=" + emailVerPassword + "\">Change Password</a></body>"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", from, password, host)
	fmt.Println("message:", string(message))
	err = smtp.SendMail(address, auth, from, to, message)
	//if error in sending account recovery info
	if err != nil {
		fmt.Println("error sending reset password email, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error in rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	//committing the changes 
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error in commiting the changes", commitErr)
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
	}
	tmp.ExecuteTemplate(w, "login.html", nil)
}


func forgotPWverHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("u")
	emailVerPassword := r.FormValue("evpw")
	userPassword := r.FormValue("password")
	confirmPassword := r.FormValue("confirmpassword")
	fmt.Println("username:", username)
	fmt.Println("emailVerPassword:", emailVerPassword)
	fmt.Println("userPassword:", userPassword)
	fmt.Println("confirmPassword:", confirmPassword)
	var ud UserData
	ud.ErrMessage = "Sorry, there was an issue recovering account, please try again"
	ud.AuthInfo = "?u=" + username + "&evpw=" + emailVerPassword
	// check if userPassword and confirmpassword are the same
	if userPassword != confirmPassword {
		fmt.Println("passwords are not matching")
		ud.ErrMessage = "passwords must match"
		tmp.ExecuteTemplate(w, "login.html", ud)
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "forgotpassword.html", ud.ErrMessage)
		return
	}
	// rollback will be ignored if the tx has been committed later in the function
	defer tx.Rollback()
	// retrieving ver_hash and timeout from email_ver_hash table
	var dbEmailVerHash string
	var timeout time.Time
	row := db.QueryRow("SELECT ver_hash, timeout FROM email_ver_hash WHERE username = ?", username)
	err = row.Scan(&dbEmailVerHash, &timeout)
	if err != nil {
		fmt.Println("ver_hash not found in db")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	// check if within timelimit
	currentTime := time.Now()
	// func (t Time) After(u Time) bool, After reports whether the time instant t is after u.
	if currentTime.After(timeout) {
		fmt.Println("users:", username, "didn't verify account within 24 hours")
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes", rollbackErr)
		}
		return
	}
	fmt.Println("dbEmailVerHash:", dbEmailVerHash)
	// check if db ver_hash is the same as the hash of emailVerPassword from email
	err = bcrypt.CompareHashAndPassword([]byte(dbEmailVerHash), []byte(emailVerPassword))
	if err != nil {
		fmt.Println("dbEmailVerHash and hash of emailVerPassword are not the same")
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	fmt.Println("dbEmailVerHash and hash of emailVerPassword are the same")
	// check userPassword criteria
	err = checkPasswordCriteria(userPassword)
	if err != nil {
		ud.AuthInfo = "?u=" + username + "&evpw=" + emailVerPassword
		// saving password criteria error to inform user
		ud.ErrMessage = err.Error()
		tmp.ExecuteTemplate(w, "login.html", ud)
		return
	}
	// generate hash for new userPassword
	var hash []byte
	// generate emailVerPassword hash for db
	hash, err = bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "signup.html", ud.ErrMessage)
		return
	}
	// update db with new userPasswordHash
	stmt := "UPDATE users SET hash = ? WHERE username = ?"
	updateHashStmt, err := tx.Prepare(stmt)
	if err != nil {
		fmt.Println("error preparing updateHashStmt err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes", rollbackErr)
		}
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	defer updateHashStmt.Close()
	var result sql.Result
	result, err = updateHashStmt.Exec(hash, username)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	// check for successfull insert
	if err != nil || rowsAff != 1 {
		fmt.Println("error inserting new user, err:", err)
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes", rollbackErr)
		}
		return
	}
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error in commiting changes", commitErr)
		tmp.ExecuteTemplate(w, "login.html", ud.ErrMessage)
		return
	}
	fmt.Println("forgotten password has been reset")
	ud.Message = "Password has been successfully Updated"
	tmp.ExecuteTemplate(w, "login.html", ud)
}

func checkPasswordCriteria(password string) error {
	var err error
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		//func to check if password contains lower case characters
		case unicode.IsLower(char):
			pswdLowercase = true
		//func to check if password contains upper case characters
		case unicode.IsUpper(char):
			pswdUppercase = true
			err = errors.New("Pa")
		//func to check if password contains numericals
		case unicode.IsNumber(char):
			pswdNumber = true
	//func to check if password contains special characters
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		//func to check if password contains empty characters
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	// check password length
	if 11 < len(password) && len(password) < 60 {
		pswdLength = true
	}
	// create error message for any criteria not passed
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces {
		switch false {
		case pswdLowercase:
			err = errors.New("Password must contain atleast one lower case letter")
		case pswdUppercase:
			err = errors.New("Password must contain atleast one uppercase letter")
		case pswdNumber:
			err = errors.New("Password must contain atleast one number")
		case pswdSpecial:
			err = errors.New("Password must contain atleast one special character")
		case pswdLength:
			err = errors.New("Passward length must atleast 12 characters and less than 60")
		case pswdNoSpaces:
			err = errors.New("Password cannot have any spaces")
		}
		return err
	}
	return nil
}


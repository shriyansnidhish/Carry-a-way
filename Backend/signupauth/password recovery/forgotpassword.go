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
// 	"CAW/Backend/signupauth/models"
// 	"CAW/Backend/signupauth/database"
// 	"CAW/Backend/signupauth/controllers"
// 	"CAW/Backend/signupauth/usersessions"
 )
var tmp*template.Template
var db *sql.DB
type UserData struct {
	Username   string
	Email      string
	AuthInfo   string
	ErrMessage string
	Message    string
}
var cookiedb = sessions.NewCookieStore([]byte(os.Getenv("Cookiedbstore")))
func forgotpwChangeHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("u")
	emailVerPassword := r.FormValue("evpw")
	fmt.Println("username:", username)
	fmt.Println("emailVerPassword:", emailVerPassword)
	var ud UserData
	ud.AuthInfo = "?u=" + username + "&evpw=" + emailVerPassword
	tmp.ExecuteTemplate(w, "login.html", ud)
}
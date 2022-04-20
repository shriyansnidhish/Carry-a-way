package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var tmp *template.Template //struct to access html file in FE
var cookiedb = sessions.NewCookieStore([]byte("secret key"))//variable to store cookies
func createSessionHandler(w http.ResponseWriter, r *http.Request) {
	
	session, err := cookiedb.Get(r, "session-name")//declare a session name
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
		HttpOnly: true,
	}
	r.ParseForm()
	name := r.FormValue("name")
	if name != "" {
		session.Values["name"] = name//giving session values
	}
	fmt.Println("session:", session)
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.ExecuteTemplate(w, "login.html", name)//redirects to login page
}



func deleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := cookiedb.Get(r, "session-name")
	session.Options.MaxAge = -1
	fmt.Println("session:", session)
	session.Save(r, w)//saving the current session
	tmp.ExecuteTemplate(w, "deleteloginsession.html", nil)
}


func main() {
	tmp, _ = template.ParseGlob("Froentend/src/app/*.html")//access FE templates
	http.HandleFunc("/loginsession", createSessionHandler)//calling to create sessions
	http.HandleFunc("/delete", deleteSessionHandler)//delete the user session
	http.ListenAndServe("localhost:8000", context.ClearHandler(http.DefaultServeMux))//wrapping in case mux is not used
}






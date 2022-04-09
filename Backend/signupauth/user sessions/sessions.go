package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var tmp *template.Template
var store = sessions.NewCookieStore([]byte("secret key"))
func createSessionHandler(w http.ResponseWriter, r *http.Request) {
	
	session, err := store.Get(r, "session-name")
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
		session.Values["name"] = name
	}
	fmt.Println("session:", session)
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.ExecuteTemplate(w, "login.html", name)
}



func deleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Options.MaxAge = -1
	fmt.Println("session:", session)
	session.Save(r, w)
	tmp.ExecuteTemplate(w, "deleteloginsession.html", nil)
}


func main() {
	tmp, _ = template.ParseGlob("Froentend/src/app/*.html")
	http.HandleFunc("/login", createSessionHandler)
	http.HandleFunc("/delete", deleteSessionHandler)
	http.ListenAndServe("localhost:8000", context.ClearHandler(http.DefaultServeMux))
}






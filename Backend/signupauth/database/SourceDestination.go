package database

import (
	"CAW/Backend/signupauth/database"
	"CAW/Backend/signupauth/models"
	"encoding/json"
	"log"
	"net/http"

	// "database/sql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SourceDest struct {
	Source  string `json:"Source"`
	Dest    string `json:"Dest"`
	Mincost string `json:"MinCost`
}

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Praneeth11@/users"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	DB = connection
}

//variable to disclose all possible source,destination info and base cost
var data []SourceDest

func getSourceDestInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func main() {
	r := mux.NewRouter()
	data = append(data, models.Distance)
	database.DB.Commit().Statement.ReflectValue.Close()
	r.HandleFunc("/Sourcedestinfo", getSourceDestInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

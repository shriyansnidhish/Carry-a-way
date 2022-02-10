package database
import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"CAW/Backend/signupauth/models"
)
var DB *gorm.DB
func Connect(){
connection, err:=gorm.Open(mysql.Open("root:Praneeth11@/users"), &gorm.Config{})
if err!=nil{
	panic("could not connect to the database")
}
DB = connection
connection.AutoMigrate(&models.User{})
}
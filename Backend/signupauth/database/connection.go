package database
import(
	"gorm.io/driver/mysql"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
	"CAW/Backend/signupauth/models"
)

var DB *gorm.DB
//connecting to MYSQL database
func Connect(){
connection, err:=gorm.Open(mysql.Open("root:Praneeth11@/users"), &gorm.Config{})
if err!=nil{
	panic("could not connect to the database")
}
DB = connection
 connection.AutoMigrate(&models.User{},&models.Bookingtable{},&models.Dispute{},&models.Orderstable{},&models.Pricetable{},&models.Distance{})

}

// var DB *gorm.DB

// func Connect() {
// 	connection, err := gorm.Open("sqlite3", "users")

// 	if err != nil {
// 		panic("could not connect to the database")
// 	}
// 	DB = connection
// }
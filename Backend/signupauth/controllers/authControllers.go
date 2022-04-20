package controllers
import(
"github.com/gofiber/fiber/v2"
"CAW/Backend/signupauth/models"
"golang.org/x/crypto/bcrypt"//password encryption package
"CAW/Backend/signupauth/database"
"github.com/dgrijalva/jwt-go"
"strconv"
"time"
"fmt"
"database/sql"
)
const SecretKey="secret"
//User SignUp
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err:=c.BodyParser(&data); err!=nil{
		return err
	}
	password, _:=bcrypt.GenerateFromPassword([]byte(data["password"]),14)
	user:= models.User{
		FirstName: data["fname"],
		LastName: data["lname"],
		Email: data["email"],
		Password:password,
	}
	 stmt := "SELECT Id FROM users WHERE Email = ?"
	 row := database.DB.Query(stmt,user.Email)
	 var uID string
	err := row.Scan(&uID)
	 if err != sql.ErrNoRows {
	 	fmt.Println("Email already exists, err:", err)
		
		
	 }
	database.DB.Create(&user)
	return c.JSON(user)
}//User login
	func Login(c *fiber.Ctx) error{
		var data map[string]string
		if err:=c.BodyParser(&data); err!=nil{
			return err
		}
		var user models.User
		database.DB.Where("email=?",data["email"]).First(&user)
		if user.Id==0{
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message":"user not found",
			})
		}
		if err:=bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]));err!=nil{
c.Status(fiber.StatusBadRequest)
return c.JSON(fiber.Map{
	"message":"incorrect password",
})
		}
claims:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.StandardClaims{
	Issuer:strconv.Itoa(int(user.Id)),
	ExpiresAt: time.Now().Add(time.Hour*24).Unix(),
})
token, err:=claims.SignedString([]byte(SecretKey))
if err!=nil{
	c.Status(fiber.StatusInternalServerError)
	return c.JSON(fiber.Map{
		"message":"Not able to log in",
	})
}
cookie:=fiber.Cookie{
	Name:"jwt",
	Value:token,
	Expires:time.Now().Add(time.Hour*24),
	HTTPOnly:true,
}
c.Cookie(&cookie)
return c.JSON(fiber.Map{
	"message":"success",
})
		
		
	}//retrieving logged in user
	func User(c *fiber.Ctx) error{
		cookie:=c.Cookies("jwt")
		token,err:=jwt.ParseWithClaims(cookie,&jwt.StandardClaims{},func(token *jwt.Token)(interface{},error){
			return []byte(SecretKey),nil
		})
		if err!=nil{
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message":"unauthenticated",
			})
		}
		claims:=token.Claims.(*jwt.StandardClaims)
		var user models.User
		database.DB.Where("id=?",claims.Issuer).First(&user)
		return c.JSON(user)
	}//User Logout
	func Logout(c *fiber.Ctx) error{
		cookie:=fiber.Cookie{
			Name:"jwt",
			Value:"",
			Expires:time.Now().Add(-time.Hour),
			HTTPOnly:true,
		}
		c.Cookie(&cookie)
		return c.JSON(fiber.Map{
			"message":"success",
		})
	}
	//User booking
	func Booking(c *fiber.Ctx) error {
		var data map[string]string
		//var data1 map[string]uint
		if err:=c.BodyParser(&data); err!=nil{
			return err
		}
		
		booking:= models.Bookingtable{
			 Orderid:data["Orderid"],
			Source: data["source"],
			Destination: data["dest"],
			Arrivaldate :data["ad"],
	         Numberofbags:data["nb"],
	        Orderstatus:data["os"],
			Cost:data["cost"],
		}
		database.DB.Create(&booking)
		return c.JSON(booking)
	}
	func Orderstable(c *fiber.Ctx) error {
		var data map[string]string
		//var data1 map[string]uint
		if err:=c.BodyParser(&data); err!=nil{
			return err
		}
		
		booking:= models.Orderstable{
		    Orid:data["Orderid"],
			Status: data["status"],
			Description: data["desc"],
			Disputeeligibility :data["de"],
	        
		}
		database.DB.Create(&booking)
		return c.JSON(booking)
	}

	  
	

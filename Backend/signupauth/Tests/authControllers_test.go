package Tests
import(
	"CAW/Backend/signupauth/controllers"
	"CAW/Backend/signupauth/database"
	"CAW/Backend/signupauth/models"
	
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)
//Signup test
func TestSignup(t *testing.T){
	err:=c.Register()
	userinf:=User{
		
		"testfirstname",
		"testlastname",
		"test@a.com",
		"password",
	}
	body, err := json.Marshal(userinf)
	check(err)
	req, err := http.NewRequest("POST", "localhost:8000/api/register", bytes.NewReader(body))
    check(err)
	rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.Register(c*fiber.Ctx))
	handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK && status != http.StatusBadRequest {
    t.Errorf("handler returned wrong status code: got %v want %v or %v",
      status, http.StatusOK, http.StatusBadRequest)
  }

}
//Signin test
// func TestLogin(t *testing.T){
// 	// setup
// 	err = database.DB.ConnectDB(testusername, testpassword, testaddress, testdbName)
// 	defer database.DB.disconnectDB()
  
// 	signinInfo := SigninInfo{
// 	  "testUsername",
// 	  "testPassword",
// 	}
  
// 	body, err := json.Marshal(signinInfo)
// 	check(err)
  
// 	req, err := http.NewRequest("POST", "localhost:8000/login", bytes.NewReader(body))
// 	check(err)
  
  
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controllers.Login(c*fiber.Ctx))
  
// 	handler.ServeHTTP(rr, req)
  
// 	if status := rr.Code; status != http.StatusOK {
// 	  t.Errorf("handler returned wrong status code: got %v want %v",
// 		status, http.StatusOK)
// 	}
  
	
  
//   }
func TestLogin(t *testing.T) {
	var data = []byte(`{
		"username": "siva000@gmail.com",
		"password": "1234"
	}`)

	app := fiber.New()

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("400")
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}
  
  func TestLogout(t *testing.T) {
	var data = []byte(`{}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/logout", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler returned expected status code %v",http.StatusOK)
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}
func TestLogoutWhenFailure(t *testing.T) {
	var data = []byte(`{}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/logout", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler returned a wrong status code")
	}

	assert.Equal(t, fiber.StatusInternalServerError, response.StatusCode)
}

func check(err error) {
	if err != nil {
	  log.Fatal(err)
	}
  }
  

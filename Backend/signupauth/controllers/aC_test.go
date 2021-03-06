package controllers
import(
	"fmt"
	"strings"
	"time"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/gofiber/fiber/v2"
)

  
  func TestHelloRoute(t *testing.T) {
	
	tests := []struct{
		Firstname string
		Lastname string
		Email string
		Password string
	}
	{
		{
			Firstname: "testfirstname"
			Lastname:"testlastname"
			Email:"test@a.com"
			Password:"password"
		}
	}
  
	// Define Fiber app.
	app := fiber.New()
  
	// Create route with POST method for test
	app.Post("/register", func(c *fiber.Ctx) error {
	  // Return simple string as response
	  return c.SendString("Hello, World!")
	})
  
	// Iterate through test single test cases
	for _, test := range tests {
	  // Create a new http request with the route from the test case
	  req := httptest.NewRequest("POST", test.route, nil)
  
	  // Perform the request plain with the app,
	  // the second argument is a request latency
	  // (set to -1 for no latency)
	  resp, _ := app.Test(req, 1)
  
	  // Verify, if the status code is as expected
	  assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
  }

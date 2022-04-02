package Tests

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"CAW/Backend/signupauth/models"
	"CAW/Backend/signupauth/controllers"
	"github.com/gofiber/fiber/v2"
	
)
var testusername = "sivapraneeth"
var testpassword = "1234"
var testaddress = "root:Praneeth11@/users"
var testdbName = "users"

func TestGetOrders(t *testing.T) {
	err := models.connectDB(testusername, testpassword, testaddress, testdbName)
	if err != nil {
		log.Fatal(err)
	}
	

	//Setting the fiber router
	
app.Get("/", func(c *fiber.Ctx) error{
	c.Accepts("application/json"),
})

	fiber.App.GET("/api/orders", GetOrders)

	t.Run("GetOrderstest", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/orders", nil)
		if err != nil {
			t.Fatalf("Request could not be created: %v\n", err)
		}

		
		r := httptest.NewRecorder()

		
		
		c.ServeHTTP(r, req)
		fmt.Println(r.Body)
		
		// Getting and compare the response
		if r.Code == http.StatusOK {
			t.Logf("Test case returns status %d, which is the same code as the expected code %d\n", http.StatusOK, r.Code)
		} else {
			t.Logf("Test case returns status %d, which is not same as the expected code %d\n", http.StatusOK, r.Code)
		}
	})


}
func TestGetOrderById(t *testing.T) {
	err := models.ConnectDB(testusername, testpassword, testaddress, testdbName)
	if err != nil {
		log.Fatal(err)
	}
	

	
	app.Get("/", func(c *fiber.Ctx) error{
		c.Accepts("application/json"),
	})
	
		fiber.App.GET("/api/orders/id", GetOrders)
	

	t.Run("Get order by id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/orders/14", nil)
		if err != nil {
			t.Fatalf("Specified request could not be created: %v\n", err)
		}

		
		q := httptest.NewRecorder()

	
		
		c.ServeHTTP(w, req)
		fmt.Print("\n")
		fmt.Println(q.Body)
		
		
		if q.Code == http.StatusOK {
			t.Logf("Test case returns status %d, which is the same code as the expected code %d\n", http.StatusOK, q.Code)
		} else {
			t.Logf("Test case returns status %d, which is not same as the expected code %d\n", http.StatusOK, q.Code)
		}
	})
}
func TestUpdateTransitStatus(t *testing.T) {
	err := models.ConnectDB(testusername, testpassword, testaddress, testdbName)
	if err != nil {
		log.Fatal(err)
	}
	

	
	app.PUT("/", func(c *fiber.Ctx) error{
		c.Accepts("application/json"),
	})
	
		fiber.App.PUT("/api/orders/id", UpdateTransitStatus)
		t.Run("Test to update transit status", func(t *testing.T) {

			req, err := http.NewRequest(http.MethodPut, "/api/orders/65", bytes.NewBuffer(emptyData))
			if err != nil {
				t.Fatalf("Specified request cannot be created: %v\n", err)
			}
	
			
			q := httptest.NewRecorder()
	
			
			
			c.ServeHTTP(q, req)
			fmt.Println(q.Body)
			
		
			if w.Code == http.StatusOK {
				t.Logf("Test case returns status %d, which is the same code as the expected code %d\n", http.StatusOK, q.Code)
			} else {
				t.Logf("Test case returns status %d, which is not same as the expected code %d\n", http.StatusOK, q.Code)
			}
		})
	}
}

func TestCancelLuggageOrder(t *testing.T) {
	err := models.ConnectDB(testusername, testpassword, testaddress, testdbName)
	if err != nil {
		log.Fatal(err)
	}
	

	
	app.DELETE("/", func(c *fiber.Ctx) error{
		c.Accepts("application/json"),
	})
	
		fiber.App.DELETE("/api/cancelluggage/id", CancelLuggageOrder)
		t.Run("Test to update transit status", func(t *testing.T) {

			req, err := http.NewRequest(http.MethodDelete, "/api/cancelluggageorder/65", bytes.NewBuffer(emptyData))
			if err != nil {
				t.Fatalf("Specified request cannot be created: %v\n", err)
			}
	
			
			q := httptest.NewRecorder()
	
			
			
			c.ServeHTTP(q, req)
			fmt.Println(q.Body)
			
		
			if w.Code == http.StatusOK {
				t.Logf("Test case returns status %d, which is the same code as the expected code %d\n", http.StatusOK, q.Code)
			} else {
				t.Logf("Test case returns status %d, which is not same as the expected code %d\n", http.StatusOK, q.Code)
			}
		})
	}
}







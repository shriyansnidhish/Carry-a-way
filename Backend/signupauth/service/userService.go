package service

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func GetCustomers(c *fiber.Ctx) {
	customers, err := controllers.GetUsers()
	check(err)

	if customers == nil {
		c.JSON(http.StatusBadRequest, fiber.Map{"error": "No Customers are Found"})
		return
	} else {
		c.JSON(http.StatusOK, fiber.Map{"data": customers})
	}
}
func GetCustomerswithId(c *fiber.Ctx) {
	customerid := c.Params("customerid")

	customer, err := controllers.GetCustomerswithId(customerid)
	check(err)

	if Customer.FirstName == "" {
		c.JSON(http.StatusBadRequest, fiber.Map{"error": "Could not find the specified customer with requested id"})
		return
	} else {
		c.JSON(http.StatusOK, fiber.Map{"info": customer})
	}
}

package service

import (
	"log"
	"net/http"
	"strconv"

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
func AddCustomers(c *fiber.Ctx) {
	var json controllers.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, fiber.Map{"error": err.Error()})
		return
	}

	success, err := controllers.AddCustomers(json)

	if success {
		c.JSON(http.StatusOK, fiber.Map{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, fiber.Map{"error": err.Error()})
	}
}

func UpdateCustomerdetails(c *fiber.Ctx) {
	customerid := c.Params("customerid")

	customer, err := controllers.GetCustomerswithId(customerid)
	check(err)
	// if the Firstname is blank we can assume nothing is found and no need to perform Update task
	if customer.FirstName == "" {
		c.JSON(http.StatusBadRequest, fiber.Map{"error": "Customer details with requested id is not found"})
		return
	} else {
		var json controllers.User

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, fiber.Map{"error": err.Error()})
			return
		}

		custId, err := strconv.Atoi(customerid)

		if err != nil {
			c.JSON(http.StatusBadRequest, fiber.Map{"error": "Invalid customerid"})
		}

		success, err := controllers.UpdateCustomerdetails(json, custId)

		if success {
			c.JSON(http.StatusOK, fiber.Map{"message": "Success"})
		} else {
			c.JSON(http.StatusBadRequest, fiber.Map{"error": err.Error()})
		}
	}
}
func DeleteCustomer(c *fiber.Ctx) {
	customerid := c.Params("customerid")

	customer, err := controllers.GetUserById(id)
	check(err)

	if customer.FirstName == "" {
		c.JSON(http.StatusBadRequest, fiber.Map{"error": "The customer id requested to delete is not found"})
		return
	} else {
		custId, err := strconv.Atoi(customerid)

		if err != nil {
			c.JSON(http.StatusBadRequest, fiber.Map{"error": "Invalid customerid"})
		}
		success, err := controllers.DeleteUser(custId)

		if success {
			c.JSON(http.StatusOK, fiber.Map{"message": "Success"})
		} else {
			c.JSON(http.StatusBadRequest, fiber.Map{"error": err.Error()})
		}
	}
}

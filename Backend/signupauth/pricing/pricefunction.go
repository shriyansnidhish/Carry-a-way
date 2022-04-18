package main

import (
	"CAW/Backend/signupauth/database"
	"CAW/Backend/signupauth/models"
	"CAW/Backend/signupauth/controllers"
	"github.com/gofiber/fiber/v2"
	"log"
	"fmt"
	"net/http"
	"database/sql"
)
type MyDB struct {
	*sql.DB
  }

func main() {
	//connects to database
	database.Connect()
	//api
	http.HandleFunc("/price", PriceFunc(&models.Distance.SourceDest,&models.Distance.Dist))
	http.HandleFunc("/dispute", DisputeFuncPrice(&models.Dispute.Disputeeligibility))

	http.ListenAndServe(":8080", nil)
}
//function to charge min cost to customers based on the shipping distance
func (db *MyDB ) PriceFunc(SourceDest string,Dist int ) {
	//retrieves and scans base cost
	q1 := fmt.Sprintf(`SELECT %d FROM Distance.result WHERE SourceDest =%s`, models.Distance.MinCost, SourceDest)
	var MinCost int8
	err := db.QueryRow(q1).Scan(&MinCost)
	if err != nil {
		log.Fatal(err)
	}
	//adds the base cost to shipping service cost and gives final price to customer
	q2 := fmt.Sprintf(`UPDATE Bookingtable.result SET %d = %d WHERE orderid =%d';`, models.Bookingtable.Cost + models.Distance.MinCost, models.Bookingtable.Cost, orderid)
	r, err2 := db.Query(q2)
	if r == nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err)
	}
	//returns data to postman api
	if controllers.Booking(c* fiber.Ctx) {
		c.JSON(
			http.StatusOK,
			models.user.ApiResponse{
				Status:  http.StatusOK,
				Message: Success 200 OK,
				Data: map[string]interface{}{
					"SourceDest":models.Distance.SourceDest,
					"Dist": models.Distance.Dist,
					"Final Cost":models.Distance.MinCost
				}},
		)
		return
	}
	//api to be returned in case of failure
	else{
		c.JSON(
			http.StatusBadRequest,
			models.user.ApiResponse{
				Status:http.StatusBadRequest,
				Message: FAILURE 400,
				Data: map[string]interface{}{
					"SourceDest":models.Distance.SourceDest,//cannot retrieve the data as error occured
					"Dist": models.Distance.Dist,
					"Final Cost":models.Distance.MinCost
				}
			}
		)
	}
}

func (db *MyDB)DisputeFuncPrice(Disputeeligibility string) {
	//scans dispute eligibility schema for "YES" to process dispute claim
	q1:=fmt.Sprintf(`SELECT %s FROM Dispute.result`,models.Dispute.Disputeeligibility)
	err:=db.Query(q1)
	if err != nil{
		log.Fatal(err)
	}
    if(q1=="YES"){
		//scans the order id of disputed order
		squery:=fmt.Sprintf("SELECT %s FROM Dispute.result",models.Dispute.Orderid)
		var oidverif string
		err := db.QueryRow(squery).Scan(&oidverif)
		//verifies the order id from booking table and sets cost to 0 effectively initiating a full refund
		q2:=fmt.Sprintf(`UPDATE Bookingtable.result SET %d = %d WHERE Bookingtable.Orderid = Dispute.Orderid;`, models.Bookingtable.Cost, 0)
		//returns data in postman api
		c.JSON(
			http.StatusOK,
			models.user.ApiResponse{
				Status:http.StatusOK,
				Message: Success 200 OK,
				Data: map[string]interface{}{
					"Disputeeligibility":models.Dispute.Disputeeligibility,
					"Description":models.Dispute.Description
				}
			}
		)
	}
}
package database
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"CAW/Backend/signupauth/models"
	"database/sql"
)
type MyDB struct {
	*sql.DB
  }
  var db MyDB
func (db *MyDB) getShipmentResult(orid int) ([]models.Orderstable, error) {
	q := fmt.Sprintf(`SELECT status,orid,description,disputeeligibility FROM Orderstable.result WHERE orid = %d ;`, orid)
	rows, err := db.Query(q)
	var results []models.Orderstable

	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var result models.Orderstable
		err = rows.Scan(&result.Status, &result.Orid, &result.Description, &result.Disputeeligibility)
		results = append(results, result)
	}
	return results, nil
}
func (db *MyDB) getDisputestatus() ([]models.Dispute, error) {
	q := fmt.Sprintf(`SELECT disputeeligibility,description FROM Dispute.result WHERE disputeeligibility = "YES" ;`)
	rows, err := db.Query(q)
	var results []models.Dispute

	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var result models.Dispute
		err = rows.Scan( &result.Disputeeligibility,&result.Description)
		results = append(results, result)
	}
	return results, nil
}
func (db *MyDB) getBookingTable(orderid int) ([]models.Bookingtable, error) {
	q := fmt.Sprintf(`SELECT orderid,source,destination,arrivaldate,numberofbags,orderstatus FROM Bookingtable.result WHERE orderid = %d ;`, orderid)
	rows, err := db.Query(q)
	var results []models.Bookingtable

	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var result models.Bookingtable
		err = rows.Scan(&result.Orderid, &result.Source, &result.Destination, &result.Arrivaldate,&result.Numberofbags,&result.Orderstatus)
		results = append(results, result)
	}
	return results, nil
}
func (db *MyDB) addBagsNumber(orderid int,numberofbags int) {
	q1 := fmt.Sprintf(`SELECT %d FROM Bookingtable.result WHERE orderid ='%d';`,numberofbags , orderid)
	var number int8
	err := db.QueryRow(q1).Scan(&number)
	
	if err != nil {
		log.Fatal(err)
	}

	q2 := fmt.Sprintf(`UPDATE Bookingtable.result SET %d = %d WHERE orderid = %d;`, numberofbags, number+1, orderid)
	r, err3 := db.Query(q2)
	if r == nil {
		log.Fatal(err)

	}
	if err3 != nil {
		log.Fatal(err)
	}
}



package models

//various structs automigrated to database as tables

type User struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}
type Pricetable struct {
	Category string `json:"cat" gorm:"primary_key"`
	Price    string `json:"price"`
}
type Bookingtable struct {
	Orderid     string   `json:"Orderid" gorm:"primary_key"`
	Source      string `json:"source"`
	Destination string `json:"dest"`
	Arrivaldate string `json:"ad"`
	Numberofbags string `json:"nb"`
	Orderstatus string `json:"os"`
	Cost string `json:"cost"`

}
type Orderstable struct{
	Status string `json:"status" gorm:"primary_key"`
	Orid uint `gorm:"foreignKey:Orderid"`
	Description string `json:"desc"`
	Disputeeligibility string `json:"de"`
}
type Dispute struct{
	Disputeeligibility string `json:"de"`
	Description string `json:"desc"`
} 
type SigninInfo struct{
	Username string
	Password string
}
type Distance struct{
	SourceDest string `json:"sd"`
	Dist int `json:"dist"`
}





package model

import "gorm.io/gorm"

//User registration model
type Customers struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Gender    string `json:"gender"`
	Email	  string `json:"email"`
	Phone	  string `json:"phone"`
	Password  string `json:"password"`
}
type Seller struct{
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Gender    string `json:"gender"`
}
//user Login model
type Login struct{
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}
//Products available model
type AvailableProducts struct{
	gorm.Model
	ProductName string	`json:"name"`
	ProductType string	`json:"type"`
	Quantity	int		`json:"quantity"`
	Price		float64	`json:"price"`
}
//Out of stock products
type Orders struct{
	ID	uint64 			`json:"-" gorm:"primaryKey;autoincrement:true"`
	ProductType string	`json:"type"`
	ProductName string	`json:"name"`
	Quantity	int		`json:"quantity"`
	BuyingPrice float64	`json:"price"`
}
//sold products
type SoldProducts struct{
	ID 	uint64			`json:"-" gorm:"primaryKey;autoIncrement:true"`
	ProductType		string	`json:"type"`
	SellingPrice	string	`json:"price"`
	QuantitySold	float64	`json:"quantity"`
	TotalAmount		float64	`json:"-"`
}
type Admin struct{
	ID 	uint64			`json:"-" gorm:"primaryKey;autoIncrement:true"`
	Email string		`json:"email"`
	Password string		`json:"password"`
}


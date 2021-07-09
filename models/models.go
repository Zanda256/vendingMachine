package models

type ProductModel struct {
	ID              int `json:"id"`
	AmountAvailable int
	Cost            int
	ProductName     string
	SellerID        int
}

type UserModel struct {
	ID       int `json:"id"`
	UserName string
	Password string
	Deposit  int
	role     interface{}
}

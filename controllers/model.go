package controllers

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
}

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
type Transactions struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type ProductResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Products `json:"data"`
}

type ProductsResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Products `json:"data"`
}

type TransactionResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    Transactions `json:"data"`
}

type TransactionsResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    []Transactions `json:"data"`
}

type DetailTrannsaction struct {
	ID       int      `json:"id transaction"`
	User     User     `json:"user"`
	Products Products `json:"products"`
	Quantity int      `json:"quantity"`
}

type DetailTransactionsResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    []DetailTrannsaction `json:"data"`
}

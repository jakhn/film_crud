package models

type OrderPrimarKey struct {
	Id string `json:"order_id"` 
}

type CreateOrder struct {
	BooksId string `json:"books_id"`
	UsersId string `json:"users_id"`
}

type Order struct {
	Id       	string `json:"order_id"`
	BooksId    	string `json:"books_id"`
	UsersId    	string `json:"users_id"` 
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateOrder struct {
	Id          string `json:"order_id"`
	BooksId    	string `json:"books_id"`
	UsersId    	string `json:"users_id"`  
}


type OrderBook struct {
	Id		 			string		`json:"order_id"`
	FullName			string 		`json:"full_name"`
	Price				int64 		`json:"price"`
	CreatedAt	 		string  	`json:"created_at"`
	UpdatedAt 			string  	`json:"updated_at"`	
}

type GetListOrderRequest struct {
	Limit  int32
	Offset int32
}

type GetListOrderResponse struct {
	Count int32   	`json:"count"`
	Orders []*OrderBook `json:"orders"`
}

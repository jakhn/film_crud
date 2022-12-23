package models

type UserPrimarKey struct {
	Id string `json:"user_id"`
	Login string `json:"login"`
}

type CreateUser struct {
	FirstName 		string `json:"first_name"`
	LastName 		string `json:"last_name"`
	Password		string	`json:"password`
	Login 			string	`json:"login"` 
	PhoneNumber 	string `json:"phone_number"`
	Balance 		int64	`json:"balance"` 
}
	

type User struct {
	Id        		string `json:"user_id"`
	FirstName    	string `json:"first_name"`
	LastName 		string `json:"last_name"`
	Login 			string	`json:"login"` 
	Password		string	`json:"password`
	PhoneNumber 	string `json:"phone_number"`
	Balance 		int64    `json:"balance"`
	CreatedAt   	string `json:"created_at"`
	UpdatedAt   	string `json:"updated_at"`
}

type UpdateUser struct {
	Id          	string `json:"user_id"`
	FirstName    	string `json:"first_name"`
	LastName 		string `json:"last_name"`
	Login 			string	`json:"login"` 
	Password		string	`json:"password`
	PhoneNumber 	string `json:"phone_number"`
	Balance 		int64  `json:"balance"`  
}

type GetListUserRequest struct {
	Limit  int32
	Offset int32
}

type GetListUserResponse struct {
	Count int32   `json:"count"`
	Users []*User `json:"users"`
}

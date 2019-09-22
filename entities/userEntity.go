package entities

type User struct{
	Email string `json:"email"`
	Password string `json:"password"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	City  string `json:"city"`
	Country string `json:"country"`
	Phoneno int64 `json:"phone_no"`
	Salary  int64 `json:"salary"`

}
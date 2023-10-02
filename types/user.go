package types

type User struct{
	ID int `json:"id"`
	NAME string `json:"name"`
}

// validate user entries
func validateUser(u *User) bool {return true}
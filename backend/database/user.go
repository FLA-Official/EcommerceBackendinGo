package database

type User struct {
	ID          int    `json:"id"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

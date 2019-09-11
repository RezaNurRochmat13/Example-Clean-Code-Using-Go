package model

type User struct {
	IDUser      string `json:"id_user"`
	Username    string `json:"user_name"`
	UserAddress string `json:"user_address"`
	UserPhone   string `json:"user_phone"`
	UserAge     int    `json:"user_age"`
}

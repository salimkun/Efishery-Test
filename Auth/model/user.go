package model

type User struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Role       int32  `json:"role_id"`
	Password   string `json:"password"`
	Registered string `json:"registered_At"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type UserClaims struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Role       int32  `json:"role_id"`
	Registered string `json:"registered_At"`
}

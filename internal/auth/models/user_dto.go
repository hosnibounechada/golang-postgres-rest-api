package models

type RegisterUserReq struct {
	FirstName string `json:"first_name" binding:"required,min=1,max=50"`
	LastName  string `json:"last_name"  binding:"required,min=1,max=50"`
	Email     string `json:"email"  binding:"required,email"`
	Password  string `json:"password"  binding:"required,min=8"`
}

type RegisterUserRes struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// type LoginUserRes struct {
// 	AccessToken string `json:"access_token"`
// 	ID          int64  `json:"id"`
// 	Username    string `json:"username"`
// }

type LoginUserRes struct {
	Jwt  TokenRes `json:"jwt"`
	User User     `json:"user"`
}

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

package models

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Password  string   `json:"-"`
	Verified  bool     `json:"-"`
	Devices   []Device `json:"-" gorm:"foreignKey:UserID"`
}

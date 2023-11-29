package models

type Device struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	OS      string `json:"os"`
	Browser string `json:"browser"`
	Token   string `json:"token"`
	UserID  int64  `json:"user_id"`
}

type DeviceRes struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	OS      string `json:"os"`
	Browser string `json:"browser"`
}

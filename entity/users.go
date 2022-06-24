package entity

type User struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Age         uint8  `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}

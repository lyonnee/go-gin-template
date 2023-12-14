package iomodel

type LoginIM struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginOM struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}

type RegisterIM struct {
	Name        string `json:"name"`
	Age         uint8  `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

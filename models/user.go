package models

type Usre struct {
	User string `json:"name"`
	Birthday string `json:"birthday"`
	Address string	`json:"address"`
	Nick string	`json:"nick"`
	Password string `json:"password"`
}

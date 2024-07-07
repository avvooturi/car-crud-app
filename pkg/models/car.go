package models

type Car struct {
	Id    string `json:"id"`
	Year  int    `json:"year"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Color string `json:"color"`
	Price int    `json:"price"`
	Miles int    `json:"miles"`
}

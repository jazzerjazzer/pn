package api

type Found struct {
	ID             string `json:"id"`
	Price          string `json:"price"`
	ExpirationDate string `json:"expiration_date"`
}

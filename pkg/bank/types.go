package types

// Money представляет денежную сумму в минимальных единицах валюты
type Money int64

// Currency представляет код валюты
type Currency string

// Коды валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN представляет номер карты
type PAN string

// представляет
type Card struct {
	Id       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

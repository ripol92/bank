package types

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID 		 int
	PAN   	 PAN
	Balance  Money
	Currency Currency
	Color 	 string
	Name     string
	Active 	 bool
	MinBalance Money
}

type PaymentSource struct {
	Type string
	Number string
	Balance Money
}

type Category string

type Status string

const (
	StatusOk = "OK"
	StatusFail = "FAIL"
	StatusInProgress = "INPROGRESS"
)

type Payment struct {
	ID int
	Amount Money
	Category Category
	Status Status
}
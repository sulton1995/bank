package types

type Money int

type Category string

type Payment struct {
	ID int
	Amount Money
	Category Category
}
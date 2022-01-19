package types

// Money представляет собой денежную сумму в минимальных единицах(центы, копейки, дирамы и.т.д).
type Money int64

// Category представляет собой категорию, в котоой был совершен платеж(авто, аптеки, рестораны и.т.д.).
type Category string

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

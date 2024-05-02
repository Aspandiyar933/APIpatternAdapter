package pattern

type Todo struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Do          bool   `json:"do"`
}

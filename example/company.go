package example

//@Table(company)
type Company struct {
	ID   int64  `db:"id, pk, autoincrement"`
	Name string `db:"name, length:255"`
}
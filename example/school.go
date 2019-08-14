package example

//@Table(school)
type School struct {
	ID   int64  `db:"id, pk, autoincrement"`
	Name string `db:"name, length:255"`
}

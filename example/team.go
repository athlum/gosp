package example

//@Table(team)
type Team struct {
	ID   int64  `db:"id, pk, autoincrement"`
	Name string `db:"name, length:255"`
}

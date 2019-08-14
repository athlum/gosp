package example

//@Table(guild)
type Guild struct {
	ID   int64  `db:"id, pk, autoincrement"`
	Name string `db:"name, length:255"`
}
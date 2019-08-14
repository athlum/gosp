package query

const (
	QueryModel_Query = iota
	QueryModel_Model
)

type FieldSet interface {
	Array() []string
}

type Model interface {
	TableName() string
	FieldArray() []string
	PK() (*Field, interface{})
}

type QueryModel interface {
	TableName() string
	FieldArray() []string
	PK() (*Field, interface{})
	Query() *Query
}
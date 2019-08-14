package example

import (
	q "github.com/athlum/gosp/query"
)

type STC struct {
	ID      int64 `db:"id, pk, autoincrement"`
	Person  int64 `db:"person"`
	Company int64 `db:"company"`
}

type STCModel struct {
	query *q.Query
}

const _stcTableName = "stc"

func (s *STCModel) TableName() string {
	return _stcTableName
}

func (s *STC) TableName() string {
	return _stcTableName
}

type stcFields struct {
	ID      *q.Field
	Person  *q.Field
	Company *q.Field
}

func (fs *stcFields) Array() []string {
	return []string{
		fs.ID.String(),
		fs.Person.String(),
		fs.Company.String(),
	}
}

var _stcFields = &stcFields{
	ID:      q.TableField(_stcTableName, "id"),
	Person:  q.TableField(_stcTableName, "person"),
	Company: q.TableField(_stcTableName, "company"),
}

func (s *STCModel) Fields() *stcFields {
	return _stcFields
}

func (s *STC) Fields() *stcFields {
	return _stcFields
}

func (s *STCModel) FieldArray() []string {
	return _stcFields.Array()
}

func (s *STC) FieldArray() []string {
	return _stcFields.Array()
}

func (s *STCModel) PK() (*q.Field, interface{}) {
	m, _ := s.Queryset().GetMain()
	return m.Model.PK()
}

func (s *STC) PK() (*q.Field, interface{}) {
	return s.Fields().ID, s.ID
}

func (s *STCModel) Queryset() *q.Query {
	if s.query == nil {
		s.query = q.NewQuery(&STC{})
	}
	return s.query
}
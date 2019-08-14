package example

import (
	q "github.com/athlum/gosp/query"
)

type STG struct {
	ID        int64 `db:"id, pk, autoincrement"`
	StudentID int64 `db:"student_id"`
	GuildID   int64 `db:"guild_id"`
}

type STGModel struct {
	query *q.Query
}

const _stgTableName = "stg"

func (s *STGModel) TableName() string {
	return _stgTableName
}

func (s *STG) TableName() string {
	return _stgTableName
}

type stgFields struct {
	ID        *q.Field
	StudentID *q.Field
	GuildID   *q.Field
}

func (fs *stgFields) Array() []string {
	return []string{
		fs.ID.String(),
		fs.StudentID.String(),
		fs.GuildID.String(),
	}
}

var _stgFields = &stgFields{
	ID:        q.TableField(_stgTableName, "id"),
	StudentID: q.TableField(_stgTableName, "student_id"),
	GuildID:   q.TableField(_stgTableName, "guild_id"),
}

func (s *STGModel) Fields() *stgFields {
	return _stgFields
}

func (s *STG) Fields() *stgFields {
	return _stgFields
}

func (s *STGModel) FieldArray() []string {
	return _stgFields.Array()
}

func (s *STG) FieldArray() []string {
	return _stgFields.Array()
}

func (s *STGModel) PK() (*q.Field, interface{}) {
	m, _ := s.Query().GetMain()
	return m.Model.PK()
}

func (s *STG) PK() (*q.Field, interface{}) {
	return s.Fields().ID, s.ID
}

func (s *STGModel) Query() *q.Query {
	if s.query == nil {
		s.query = q.NewQuery(&STG{})
	}
	return s.query
}